package measurement

import (
	"fmt"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/device"
	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func PreampBiasRead() {
	i2c_bus := util.DetectMCP2221()

	// Set Mux Channel for ADC 1
	device.PCA9548A(i2c_bus, 0x70, 1)
	time.Sleep(750 * time.Millisecond)
	preamp_sen_0 := device.MAX11615(i2c_bus, 0x33, 3, 3.0)
	preamp_sen_1 := device.MAX11615(i2c_bus, 0x33, 2, 3.0)
	preamp_sen_2 := device.MAX11615(i2c_bus, 0x33, 1, 3.0)
	preamp_sen_3 := device.MAX11615(i2c_bus, 0x33, 0, 3.0)
	preamp_sen_4 := device.MAX11617(i2c_bus, 0x35, 10, 3.0)
	preamp_sen_5 := device.MAX11617(i2c_bus, 0x35, 9, 3.0)
	preamp_sen_6 := device.MAX11617(i2c_bus, 0x35, 8, 3.0)
	preamp_sen_7 := device.MAX11617(i2c_bus, 0x35, 0, 3.0)

	// Set Mux Channel for ADC 2
	device.PCA9548A(i2c_bus, 0x70, 3)
	time.Sleep(750 * time.Millisecond)
	preamp_sen_8 := device.MAX11615(i2c_bus, 0x33, 3, 3.0)
	preamp_sen_9 := device.MAX11615(i2c_bus, 0x33, 2, 3.0)
	preamp_sen_10 := device.MAX11615(i2c_bus, 0x33, 1, 3.0)
	preamp_sen_11 := device.MAX11615(i2c_bus, 0x33, 0, 3.0)
	preamp_sen_12 := device.MAX11617(i2c_bus, 0x35, 10, 3.0)
	preamp_sen_13 := device.MAX11617(i2c_bus, 0x35, 9, 3.0)
	preamp_sen_14 := device.MAX11617(i2c_bus, 0x35, 8, 3.0)
	preamp_sen_15 := device.MAX11617(i2c_bus, 0x35, 0, 3.0)

	// Print Result
	fmt.Printf("Ch0	:	%.3fV\n", preampBiasConv(preamp_sen_0))
	fmt.Printf("Ch1	:	%.3fV\n", preampBiasConv(preamp_sen_1))
	fmt.Printf("Ch2	:	%.3fV\n", preampBiasConv(preamp_sen_2))
	fmt.Printf("Ch3	:	%.3fV\n", preampBiasConv(preamp_sen_3))
	fmt.Printf("Ch4	:	%.3fV\n", preampBiasConv(preamp_sen_4))
	fmt.Printf("Ch5	:	%.3fV\n", preampBiasConv(preamp_sen_5))
	fmt.Printf("Ch6	:	%.3fV\n", preampBiasConv(preamp_sen_6))
	fmt.Printf("Ch7	:	%.3fV\n", preampBiasConv(preamp_sen_7))
	fmt.Printf("Ch8	:	%.3fV\n", preampBiasConv(preamp_sen_8))
	fmt.Printf("Ch9	:	%.3fV\n", preampBiasConv(preamp_sen_9))
	fmt.Printf("Ch10	:	%.3fV\n", preampBiasConv(preamp_sen_10))
	fmt.Printf("Ch11	:	%.3fV\n", preampBiasConv(preamp_sen_11))
	fmt.Printf("Ch12	:	%.3fV\n", preampBiasConv(preamp_sen_12))
	fmt.Printf("Ch13	:	%.3fV\n", preampBiasConv(preamp_sen_13))
	fmt.Printf("Ch14	:	%.3fV\n", preampBiasConv(preamp_sen_14))
	fmt.Printf("Ch15	:	%.3fV\n", preampBiasConv(preamp_sen_15))

	// Reset I2C Mux
	device.PCA9548A(i2c_bus, 0x70, -1)
	time.Sleep(750 * time.Millisecond)
}

func preampBiasConv(voltage float32) float32 {
	bias_voltage := 22.27659574468085 * voltage

	return bias_voltage
}
