package measurement

import (
	"fmt"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/device"
	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func PreampTemperature() {
	i2c_bus := util.DetectMCP2221()

	// Set Mux Channel for ADC 1
	device.PCA9548A(i2c_bus, 0x70, 1)
	time.Sleep(750 * time.Millisecond)
	preamp_tmp_0 := preampTmp(device.MAX11615(i2c_bus, 0x33, 7, 3.0))
	preamp_tmp_1 := preampTmp(device.MAX11615(i2c_bus, 0x33, 6, 3.0))
	preamp_tmp_2 := preampTmp(device.MAX11615(i2c_bus, 0x33, 5, 3.0))
	preamp_tmp_3 := preampTmp(device.MAX11615(i2c_bus, 0x33, 4, 3.0))
	preamp_tmp_4 := preampTmp(device.MAX11617(i2c_bus, 0x35, 4, 3.0))
	preamp_tmp_5 := preampTmp(device.MAX11617(i2c_bus, 0x35, 5, 3.0))
	preamp_tmp_6 := preampTmp(device.MAX11617(i2c_bus, 0x35, 6, 3.0))
	preamp_tmp_7 := preampTmp(device.MAX11617(i2c_bus, 0x35, 7, 3.0))

	// Set Mux Channel for ADC 2
	device.PCA9548A(i2c_bus, 0x70, 3)
	time.Sleep(750 * time.Millisecond)
	preamp_tmp_8 := preampTmp(device.MAX11615(i2c_bus, 0x33, 7, 3.0))
	preamp_tmp_9 := preampTmp(device.MAX11615(i2c_bus, 0x33, 6, 3.0))
	preamp_tmp_10 := preampTmp(device.MAX11615(i2c_bus, 0x33, 5, 3.0))
	preamp_tmp_11 := preampTmp(device.MAX11615(i2c_bus, 0x33, 4, 3.0))
	preamp_tmp_12 := preampTmp(device.MAX11617(i2c_bus, 0x35, 4, 3.0))
	preamp_tmp_13 := preampTmp(device.MAX11617(i2c_bus, 0x35, 5, 3.0))
	preamp_tmp_14 := preampTmp(device.MAX11617(i2c_bus, 0x35, 6, 3.0))
	preamp_tmp_15 := preampTmp(device.MAX11617(i2c_bus, 0x35, 7, 3.0))

	// Print Result
	fmt.Printf("Preamp Board 0 Temperature	:	%.3f°C\n", preamp_tmp_0)
	fmt.Printf("Preamp Board 1 Temperature	:	%.3f°C\n", preamp_tmp_1)
	fmt.Printf("Preamp Board 2 Temperature	:	%.3f°C\n", preamp_tmp_2)
	fmt.Printf("Preamp Board 3 Temperature	:	%.3f°C\n", preamp_tmp_3)
	fmt.Printf("Preamp Board 4 Temperature	:	%.3f°C\n", preamp_tmp_4)
	fmt.Printf("Preamp Board 5 Temperature	:	%.3f°C\n", preamp_tmp_5)
	fmt.Printf("Preamp Board 6 Temperature	:	%.3f°C\n", preamp_tmp_6)
	fmt.Printf("Preamp Board 7 Temperature	:	%.3f°C\n", preamp_tmp_7)
	fmt.Printf("Preamp Board 8 Temperature	:	%.3f°C\n", preamp_tmp_8)
	fmt.Printf("Preamp Board 9 Temperature	:	%.3f°C\n", preamp_tmp_9)
	fmt.Printf("Preamp Board 10 Temperature	:	%.3f°C\n", preamp_tmp_10)
	fmt.Printf("Preamp Board 11 Temperature	:	%.3f°C\n", preamp_tmp_11)
	fmt.Printf("Preamp Board 12 Temperature	:	%.3f°C\n", preamp_tmp_12)
	fmt.Printf("Preamp Board 13 Temperature	:	%.3f°C\n", preamp_tmp_13)
	fmt.Printf("Preamp Board 14 Temperature	:	%.3f°C\n", preamp_tmp_14)
	fmt.Printf("Preamp Board 15 Temperature	:	%.3f°C\n", preamp_tmp_15)

	// Reset I2C Mux
	device.PCA9548A(i2c_bus, 0x70, -1)
	time.Sleep(750 * time.Millisecond)
}

func preampTmp(voltage float32) float32 {
	temperature := (voltage - 0.5) * 100.0
	if -40 > temperature || temperature > 150 {
		temperature = 500.000
	}

	return temperature
}
