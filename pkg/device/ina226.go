package device

import (
	"log"
	"math"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

const (
	// Register
	CONFIG_INA226  = 0x00
	SHUNT_INA226   = 0x01
	BUS_INA226     = 0x02
	POWER_INA226   = 0x03
	CURRENT_INA226 = 0x04
	CALIB_INA226   = 0x05
	// Configuration Parameters
	CONFIG_RST_INA226         = 0xC000
	CONFIG_AVG_1_INA226       = 0x4000
	CONFIG_AVG_4_INA226       = 0x4200
	CONFIG_AVG_16_INA226      = 0x4400
	CONFIG_AVG_64_INA226      = 0x4600
	CONFIG_AVG_128_INA226     = 0x4800
	CONFIG_AVG_256_INA226     = 0x4A00
	CONFIG_AVG_512_INA226     = 0x4C00
	CONFIG_AVG_1024_INA226    = 0x4E00
	CONFIG_VBUSCT_140_INA226  = 0x4000
	CONFIG_VBUSCT_204_INA226  = 0x4040
	CONFIG_VBUSCT_332_INA226  = 0x4080
	CONFIG_VBUSCT_588_INA226  = 0x40C0
	CONFIG_VBUSCT_1100_INA226 = 0x4100
	CONFIG_VBUSCT_2116_INA226 = 0x4140
	CONFIG_VBUSCT_4156_INA226 = 0x4180
	CONFIG_VBUSCT_8244_INA226 = 0x41C0
	CONFIG_VSHCT_140_INA226   = 0x4000
	CONFIG_VSHCT_204_INA226   = 0x4008
	CONFIG_VSHCT_332_INA226   = 0x4010
	CONFIG_VSHCT_588_INA226   = 0x4018
	CONFIG_VSHCT_1100_INA226  = 0x4020
	CONFIG_VSHCT_2116_INA226  = 0x4028
	CONFIG_VSHCT_4156_INA226  = 0x4030
	CONFIG_VSHCT_8244_INA226  = 0x4038
	CONFIG_MODE_PDS_INA226    = 0x4000 // Power-Down (or Shutdown)
	CONFIG_MODE_SVT_INA226    = 0x4001 // Shunt Voltage, Triggered
	CONFIG_MODE_BVT_INA226    = 0x4002 // Bus Voltage, Triggered
	CONFIG_MODE_SBT_INA226    = 0x4003 // Shunt and Bus, Triggered
	CONFIG_MODE_PDS_2_INA226  = 0x4004 // Power-Down (or Shutdown) 2
	CONFIG_MODE_SVC_INA226    = 0x4005 // Shunt Voltage, Continuous
	CONFIG_MODE_BVC_INA226    = 0x4006 // Bus Voltage, Continuous
	CONFIG_MODE_SBC_INA226    = 0x4007 // Shunt and Bus, Continuous
)

type INA226DATA struct {
	ShuntVoltage float32
	BusVoltage   float32
	Current      float32
	Power        float32
}

func INA226(bus int, address uint8, rshunt, mec float32) INA226DATA {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open INA226: %+v", err)
	}
	defer c.CloseI2C()

	// Configure INA226 Chip
	config := uint16(CONFIG_AVG_16_INA226 | CONFIG_VBUSCT_332_INA226 | CONFIG_VSHCT_332_INA226 | CONFIG_MODE_SBC_INA226)
	c.WriteBlockDataI2C(address, CONFIG_INA226, wordToByteArray_INA226(config))

	// Calibrate INA226 Chip
	cLSB := mec / float32(math.Pow(2, 15))
	pLSB := 25 * cLSB
	cal := 0.00512 / (cLSB * rshunt)
	calib_buf := make([]byte, 2)
	calib_buf[0] = byte(uint16(cal) >> 8)
	calib_buf[1] = byte(uint16(cal))
	c.WriteBlockDataI2C(address, CALIB_INA226, calib_buf)

	// Wait for 10 ms
	time.Sleep(10 * time.Millisecond)

	// Measure Shunt Voltage
	shuntVoltage_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, SHUNT_INA226, shuntVoltage_buf)
	if err != nil {
		log.Fatalf("Could not read shunt voltage: %+v", err)
	}
	shuntVoltage_adc := uint16(shuntVoltage_buf[0])<<8 | uint16(shuntVoltage_buf[1])
	shutVoltage := measShuntVoltage_INA226(shuntVoltage_adc)

	// Measure Bus Voltage
	busVoltage_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, BUS_INA226, busVoltage_buf)
	if err != nil {
		log.Fatalf("Could not read bus voltage: %+v", err)
	}
	busVoltage_adc := uint16(busVoltage_buf[0])<<8 | uint16(busVoltage_buf[1])
	busVoltage := measBusVoltage_INA226(busVoltage_adc)

	// Measure Current
	current_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, CURRENT_INA226, current_buf)
	if err != nil {
		log.Fatalf("Could not read current: %+v", err)
	}
	current_adc := uint16(current_buf[0])<<8 | uint16(current_buf[1])
	current := measCurrent_INA226(current_adc, cLSB)

	// Measure Power
	power_buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, POWER_INA226, power_buf)
	if err != nil {
		log.Fatalf("Could not read power: %+v", err)
	}
	power_adc := uint16(power_buf[0])<<8 | uint16(power_buf[1])
	power := measPower_INA226(power_adc, pLSB)

	ina226_data := INA226DATA{
		ShuntVoltage: shutVoltage,
		BusVoltage:   busVoltage,
		Current:      current,
		Power:        power,
	}

	return ina226_data

}

func measShuntVoltage_INA226(adc uint16) float32 {
	sign := +1.0
	if adc >= 0x8000 {
		sign = -1
		adc = (adc & 0x7FFF) + 1
	}
	voltage := float32(adc) * float32(sign) * 0.0000025

	return voltage
}

func measBusVoltage_INA226(adc uint16) float32 {
	voltage := float32(adc) * 0.00125

	return voltage
}

func measCurrent_INA226(adc uint16, cLSB float32) float32 {
	sign := +1.0
	if adc >= 0x8000 {
		sign = -1
		adc = (adc & 0x7FFF) + 1
	}
	current := float32(adc) * float32(sign) * cLSB

	return current
}

func measPower_INA226(adc uint16, pLSB float32) float32 {
	power := float32(adc) * pLSB

	return power
}

func wordToByteArray_INA226(w uint16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(w >> 8)
	buf[1] = byte(w)

	return buf
}
