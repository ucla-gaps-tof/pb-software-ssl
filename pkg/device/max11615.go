package device

import (
	"log"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

const (
	// Register
	SETUP_MAX11615  = 0x80
	CONFIG_MAX11615 = 0x00
	// Setup Parameters
	SETUP_Vdd_AI_NC_OFF_MAX11615 = 0x00 // Reference Voltage = Vdd, AIN_/REF = Analog Input, REF = Not Connected, Internal Reference State = Always Off
	SETUP_ER_RI_RI_OFF_MAX11615  = 0x20 // Reference Voltage = External Reference, AIN_/REF = Reference Input, REF = Reference Input, Internal Reference State = Always Off
	SETUP_IR_AI_NC_OFF_MAX11615  = 0x40 // Reference Voltage = Internal Reference, AIN_/REF = Analog Input, REF = Not Connected, Internal Reference State = Always Off
	SETUP_IR_AI_NC_ON_MAX11615   = 0x50 // Reference Voltage = Internal Reference, AIN_/REF = Analog Input, REF = Not Connected, Internal Reference State = Always On
	SETUP_IR_RO_RO_OFF_MAX11615  = 0x60 // Reference Voltage = Internal Reference, AIN_/REF = Reference Output, REF = Reference Output, Internal Reference State = Always Off
	SETUP_IR_RO_RO_ON_MAX11615   = 0x70 // Reference Voltage = Internal Reference, AIN_/REF = Reference Output, REF = Reference Output, Internal Reference State = Always On
	SETUP_INT_CLK_MAX11615       = 0x00 // Internal Clock
	SETUP_EXT_CLK_MAX11615       = 0x08 // External Clock
	SETUP_UNI_MAX11615           = 0x00 // Unipolar
	SETUP_BIP_MAX11615           = 0x04 // Bipolar
	SETUP_RST_MAX11615           = 0x00 // Reset
	SETUP_NA_MAX11615            = 0x01 // No Action
	// Configuration Parameters
	CONFIG_SCAN_0_MAX11615 = 0x00 // Scans up from AIN0 to the input selected by CS3-CS0.
	CONFIG_SCAN_1_MAX11615 = 0x20 // Converts the input selected by CS3-CS0 eight times.
	CONFIG_SCAN_2_MAX11615 = 0x40 // Scans up from AIN6 to the input selected by CS3-CS0.
	CONFIG_SCAN_3_MAX11615 = 0x60 // Converts channel selected by CS3-CS0.
	CONFIG_DIF_MAX11615    = 0x00 // Differential
	CONFIG_SGL_MAX11615    = 0x01 // Single-ended
)

func MAX11615(bus int, address, channel uint8, ref_voltage float32) float32 {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open MAX11615: %+v", err)
	}
	defer c.CloseI2C()

	// Setup MAX11615 Chip
	setup := uint16(SETUP_MAX11615 | SETUP_ER_RI_RI_OFF_MAX11615 | SETUP_INT_CLK_MAX11615 | SETUP_UNI_MAX11615 | SETUP_RST_MAX11615)
	c.WriteWordI2C(address, 0, setup)

	// Configure MAX11615 Chip
	config := uint16(CONFIG_MAX11615 | CONFIG_SCAN_3_MAX11615 | channelSelector_MAX11615(channel) | CONFIG_SGL_MAX11615)
	c.WriteWordI2C(address, 0, config)

	// Wait for 10 ms
	time.Sleep(10 * time.Millisecond)

	// Read ADC
	buf := make([]byte, 2)
	err = c.ReadBlockDataI2C(address, uint8(config), buf)
	if err != nil {
		log.Fatalf("Could not read ADC register: %+v", err)
	}

	// 12-bit ADC Value
	adc_value := uint16(buf[0]&0x0f)<<8 | uint16(buf[1]&0xFF)

	// Convert ADC to Voltage
	voltage := float32(adc_value) * ref_voltage / 4096.0

	return voltage

}

func channelSelector_MAX11615(channel uint8) uint8 {
	var chn uint8
	switch channel {
	case 0:
		chn = 0x00
	case 1:
		chn = 0x02
	case 2:
		chn = 0x04
	case 3:
		chn = 0x06
	case 4:
		chn = 0x08
	case 5:
		chn = 0x0A
	case 6:
		chn = 0x0C
	case 7:
		chn = 0x0E
	default:
		chn = 0x00
	}

	return chn
}
