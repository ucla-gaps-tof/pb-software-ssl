package device

import (
	"log"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

// var err error

const (
	// Register
	TEMP_TMP1075 = 0x00 // Temperature Result Register
	CFGR_TMP1075 = 0x01 // Configuration Register
	LLIM_TMP1075 = 0x02 // Low Limit Register
	HLIM_TMP1075 = 0x03 // High Limit Register
	// Configuration Parameters
	CONFIG_OS_TMP1075    = 0x80A0 // One-shot conversion mode
	CONFIG_R_35_TMP1075  = 0x60A0 // 35 ms conversion rate (Read-only)
	CONFIG_F_1_TMP1075   = 0x00A0 // 1 fault
	CONFIG_F_2_TMP1075   = 0x08A0 // 2 fault
	CONFIG_F_4_TMP1075   = 0x10A0 // 4 fault
	CONFIG_F_6_TMP1075   = 0x18A0 // 6 fault
	CONFIG_POL_L_TMP1075 = 0x00A0 // Active low ALERT pin
	CONFIG_POL_H_TMP1075 = 0x04A0 // Active high ALERT pin
	CONFIG_TM_CM_TMP1075 = 0x00A0 // ALERT pin functions in comparator mode
	CONFIG_TM_IM_TMP1075 = 0x02A0 // ALERT pin functions in interrupt mode
	CONFIG_SD_CC_TMP1075 = 0x00A0 // Device is in continuos conversion
	CONFIG_SD_SM_TMP1075 = 0x01A0 // Device is in shutdown conversion

)

func TMP1075(bus int, address uint8) float32 {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open TMP1075: %+v", err)
	}
	defer c.CloseI2C()

	// Configure TMP1075 Chip
	config := uint16(CONFIG_R_35_TMP1075 | CONFIG_F_1_TMP1075 | CONFIG_POL_L_TMP1075 | CONFIG_TM_CM_TMP1075 | CONFIG_SD_CC_TMP1075)
	c.WriteBlockDataI2C(address, CFGR_TMP1075, wordToByteArray_TMP1075(config))

	// Wait for 10 ms
	time.Sleep(10 * time.Millisecond)

	// Read Temperature
	buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, TEMP_TMP1075, buf)
	if err != nil {
		log.Fatalf("Could not read temperature register: %+v", err)
	}

	// 12-bit ADC Value
	adc_value := ((uint16(buf[0]) << 4) | (uint16(buf[1]) >> 4)) & 0xFFF

	temperature := convTemp_TMP1075(adc_value)

	return temperature
}

func wordToByteArray_TMP1075(w uint16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(w >> 8)
	buf[1] = byte(w)

	return buf
}

func convTemp_TMP1075(adc uint16) float32 {
	sign := +1.0
	if adc >= 0x800 {
		sign = -1
		adc -= 0x1000
	}
	temp := float32(adc) * float32(sign) * 0.0625

	return temp
}
