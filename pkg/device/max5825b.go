package device

import (
	"log"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

const (
	// Configuration Register
	WDOG_MAX5825B        = 0x10
	REF_MAX5825B         = 0x20
	SW_GATE_CLR_MAX5825B = 0x30
	SW_GATE_SET_MAX5825B = 0x31
	WD_REFRESH_MAX5825B  = 0x32
	WD_RESET_MAX5825B    = 0x33
	SW_CLEAR_MAX5825B    = 0x34
	SW_RESET_MAX5825B    = 0x35
	POWER_MAX5825B       = 0x40
	CONFIG_MAX5825B      = 0x50
	DEFAULT_MAX5825B     = 0x60

	// DAC Register
	RETURNn_MAX5825B           = 0x70
	CODEn_MAX5825B             = 0x80
	LOADn_MAX5825B             = 0x90
	CODEn_LOAD_ALL_MAX5825B    = 0xA0
	CODEn_LOAD_n_MAX5825B      = 0xB0
	CODE_ALL_MAX5825B          = 0xC0
	LOAD_ALL_MAX5825B          = 0xC1
	CODE_ALL_LOAD_ALL_MAX5825B = 0xC2
	RETURN_ALL_MAX5825B        = 0xC3

	// No Operation Register
	NO_0_MAX5825B = 0xC4
	NO_1_MAX5825B = 0xC8
	NO_2_MAX5825B = 0xCC

	// Configuration Parametersz
	REF_PWR_MAX5825B      = 0x00
	REF_DAC_MAX5825B      = 0x04
	REF_MODE_EXT_MAX5825B = 0x00
	REF_MODE_2V5_MAX5825B = 0x01
	REF_MODE_2V0_MAX5825B = 0x02
	REF_MODE_4V1_MAX5825B = 0x03
)

func MAX5825B(bus int, address, channel uint8, adc_vol uint16) {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open MAX5825B: %+v", err)
	}
	defer c.CloseI2C()

	// Configure MAX5825B Reference
	ref_reg := uint8(REF_MAX5825B | REF_MODE_EXT_MAX5825B)
	c.WriteBlockDataI2C(address, ref_reg, []byte{0, 0})

	// Wait for 10 ms
	time.Sleep(10 * time.Millisecond)

	// Set Voltage from Input
	code_reg := uint8(CODEn_MAX5825B | channel)
	load_reg := uint8(LOADn_MAX5825B | channel)
	c.WriteBlockDataI2C(address, code_reg, wordToByteArray_MAX5825B(adc_vol))
	c.WriteBlockDataI2C(address, load_reg, []byte{0, 0})
}

func wordToByteArray_MAX5825B(w uint16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(w>>4) & 0xFF
	buf[1] = byte(w&15) << 4

	return buf
}
