package measurement

import (
	"fmt"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/device"
	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func LtbPow(operation bool) {
	i2c_bus := util.DetectMCP2221()

	// MAX7320
	device.PCA9548A(i2c_bus, 0x70, 7)
	time.Sleep(750 * time.Millisecond)

	if operation {
		// Set "ON" for LTB Power Outputs
		device.MAX7320(i2c_bus, 0x59, true)
		fmt.Println("LTB is ON now.")
	} else {
		// Set "OFF" for LTB Power Outputs
		device.MAX7320(i2c_bus, 0x59, false)
		fmt.Println("LTB is OFF now.")
	}

	// Reset I2C Mux
	device.PCA9548A(i2c_bus, 0x70, -1)
	time.Sleep(750 * time.Millisecond)
}
