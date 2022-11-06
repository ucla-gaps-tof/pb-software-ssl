package measurement

import (
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/device"
	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func PreampBiasSet(adc uint16) {
	i2c_bus := util.DetectMCP2221()

	// Set Mux Channel for ADC 1
	device.PCA9548A(i2c_bus, 0x70, 0)
	time.Sleep(750 * time.Millisecond)
	device.MAX5825B(i2c_bus, 0x1F, 0, adc)
	device.MAX5825B(i2c_bus, 0x1F, 1, adc)
	device.MAX5825B(i2c_bus, 0x1F, 2, adc)
	device.MAX5825B(i2c_bus, 0x1F, 3, adc)
	device.MAX5825B(i2c_bus, 0x1F, 4, adc)
	device.MAX5825B(i2c_bus, 0x1F, 5, adc)
	device.MAX5825B(i2c_bus, 0x1F, 6, adc)
	device.MAX5825B(i2c_bus, 0x1F, 7, adc)

	// Set Mux Channel for ADC 2
	device.PCA9548A(i2c_bus, 0x70, 2)
	time.Sleep(750 * time.Millisecond)
	device.MAX5825B(i2c_bus, 0x1F, 0, adc)
	device.MAX5825B(i2c_bus, 0x1F, 1, adc)
	device.MAX5825B(i2c_bus, 0x1F, 2, adc)
	device.MAX5825B(i2c_bus, 0x1F, 3, adc)
	device.MAX5825B(i2c_bus, 0x1F, 4, adc)
	device.MAX5825B(i2c_bus, 0x1F, 5, adc)
	device.MAX5825B(i2c_bus, 0x1F, 6, adc)
	device.MAX5825B(i2c_bus, 0x1F, 7, adc)

	// Reset I2C Mux
	device.PCA9548A(i2c_bus, 0x70, -1)
	time.Sleep(750 * time.Millisecond)
}
