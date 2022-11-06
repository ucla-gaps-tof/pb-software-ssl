package device

import (
	"log"
	"math"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

const (
	// Register
	CONFIG_INA219  = 0x00
	SHUNT_INA219   = 0x01
	BUS_INA219     = 0x02
	POWER_INA219   = 0x03
	CURRENT_INA219 = 0x04
	CALIB_INA219   = 0x05
	// Configuration Parameters
	CONFIG_RST_INA219        = 0x8000
	CONFIG_BRNG_16_INA219    = 0x0000
	CONFIG_BRNG_32_INA219    = 0x2000
	CONFIG_PG_40_INA219      = 0x0000
	CONFIG_PG_80_INA219      = 0x0800
	CONFIG_PG_160_INA219     = 0x1000
	CONFIG_PG_320_INA219     = 0x1800
	CONFIG_BADC_9B_INA219    = 0x0000
	CONFIG_BADC_10B_INA219   = 0x0080
	CONFIG_BADC_11B_INA219   = 0x0100
	CONFIG_BADC_12B_D_INA219 = 0x0180
	CONFIG_BADC_12B_INA219   = 0x0400
	CONFIG_BADC_2S_INA219    = 0x0480
	CONFIG_BADC_4S_INA219    = 0x0500
	CONFIG_BADC_8S_INA219    = 0x0580
	CONFIG_BADC_16S_INA219   = 0x0600
	CONFIG_BADC_32S_INA219   = 0x0680
	CONFIG_BADC_64S_INA219   = 0x0700
	CONFIG_BADC_128S_INA219  = 0x0780
	CONFIG_SADC_9B_INA219    = 0x0000
	CONFIG_SADC_10B_INA219   = 0x0008
	CONFIG_SADC_11B_INA219   = 0x0010
	CONFIG_SADC_12B_D_INA219 = 0x0018
	CONFIG_SADC_12B_INA219   = 0x0040
	CONFIG_SADC_2S_INA219    = 0x0048
	CONFIG_SADC_4S_INA219    = 0x0050
	CONFIG_SADC_8S_INA219    = 0x0058
	CONFIG_SADC_16S_INA219   = 0x0060
	CONFIG_SADC_32S_INA219   = 0x0068
	CONFIG_SADC_64S_INA219   = 0x0070
	CONFIG_SADC_128S_INA219  = 0x0078
	CONFIG_MODE_PD_INA219    = 0x0000 // Power-Down
	CONFIG_MODE_SVT_INA219   = 0x001  // Shunt Voltage, Triggered
	CONFIG_MODE_BVT_INA219   = 0x002  // Bus Voltage, Triggered
	CONFIG_MODE_SBT_INA219   = 0x003  // Shunt and Bus, Triggered
	CONFIG_MODE_ADO_INA219   = 0x004  // ADC off (disabled)
	CONFIG_MODE_SVC_INA219   = 0x005  // Shunt Voltage, Continuous
	CONFIG_MODE_BVC_INA219   = 0x006  // Bus Voltage, Continuous
	CONFIG_MODE_SBC_INA219   = 0x007  // Shunt and Bus, Continuous
)

type INA219DATA struct {
	ShuntVoltage float32
	BusVoltage   float32
	Current      float32
	Power        float32
}

func INA219(bus int, address uint8, rshunt, mec float32) INA219DATA {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open INA219: %+v", err)
	}
	defer c.CloseI2C()

	// Configure INA219 Chip
	config := uint16(CONFIG_BRNG_32_INA219 | CONFIG_PG_320_INA219 | CONFIG_BADC_16S_INA219 | CONFIG_SADC_16S_INA219 | CONFIG_MODE_SBC_INA219)
	c.WriteBlockDataI2C(address, CONFIG_INA219, wordToByteArray_INA219(config))

	// Calibrate INA219 Chip
	cLSB := mec / float32(math.Pow(2, 15))
	pLSB := 20 * cLSB
	cal := 0.04096 / (cLSB * rshunt)
	calib_buf := make([]byte, 2)
	calib_buf[0] = byte(uint16(cal) >> 8)
	calib_buf[1] = byte(uint16(cal))
	c.WriteBlockDataI2C(address, CALIB_INA219, calib_buf)

	// Wait for 10 ms
	time.Sleep(10 * time.Millisecond)

	// Measure Shunt Voltage
	shuntVoltage_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, SHUNT_INA219, shuntVoltage_buf)
	if err != nil {
		log.Fatalf("Could not read shunt voltage: %+v", err)
	}
	shuntVoltage_adc := uint16(shuntVoltage_buf[0])<<8 | uint16(shuntVoltage_buf[1])
	shutVoltage := measShuntVoltage_INA219(shuntVoltage_adc)

	// Measure Bus Voltage
	busVoltage_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, BUS_INA219, busVoltage_buf)
	if err != nil {
		log.Fatalf("Could not read bus voltage: %+v", err)
	}
	busVoltage_adc := ((uint16(busVoltage_buf[0]) << 8) | uint16(busVoltage_buf[1])) >> 3 & 0x1FFF
	busVoltage := measBusVoltage_INA219(busVoltage_adc)

	// Measure Current
	current_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, CURRENT_INA219, current_buf)
	if err != nil {
		log.Fatalf("Could not read current: %+v", err)
	}
	current_adc := uint16(current_buf[0])<<8 | uint16(current_buf[1])
	current := measCurrent_INA219(current_adc, cLSB)

	// Measure Power
	power_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, POWER_INA219, power_buf)
	if err != nil {
		log.Fatalf("Could not read current: %+v", err)
	}
	power_adc := uint16(power_buf[0])<<8 | uint16(power_buf[1])
	power := measPower_INA219(power_adc, pLSB)

	ina219_data := INA219DATA{
		ShuntVoltage: shutVoltage,
		BusVoltage:   busVoltage,
		Current:      current,
		Power:        power,
	}

	return ina219_data

}

func wordToByteArray_INA219(w uint16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(w >> 8)
	buf[1] = byte(w)

	return buf
}

func measShuntVoltage_INA219(adc uint16) float32 {
	sign := +1.0
	if adc >= 0x8000 {
		sign = -1
		adc = (adc & 0x7FFF) + 1
	}
	voltage := float32(adc) * float32(sign) * 0.00001

	return voltage
}

func measBusVoltage_INA219(adc uint16) float32 {
	voltage := float32(adc) * 0.004

	return voltage
}

func measCurrent_INA219(adc uint16, cLSB float32) float32 {
	sign := +1.0
	if adc >= 0x8000 {
		sign = -1
		adc = (adc & 0x7FFF) + 1
	}
	current := float32(adc) * float32(sign) * cLSB

	return current
}

func measPower_INA219(adc uint16, pLSB float32) float32 {
	power := float32(adc) * pLSB

	return power
}
