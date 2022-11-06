package measurement

import (
	"fmt"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/device"
	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func PbTemperature() {
	i2c_bus := util.DetectMCP2221()

	device.PCA9548A(i2c_bus, 0x70, 4)
	time.Sleep(750 * time.Millisecond)

	pds_temp := device.TMP1075(i2c_bus, 0x48)
	pas_temp := device.TMP1075(i2c_bus, 0x49)
	nas_temp := device.TMP1075(i2c_bus, 0x4A)
	shv_temp := device.TMP1075(i2c_bus, 0x4B)

	// Print Result
	fmt.Printf("PDS(U8) Temperature	:	%.3f°C\n", pds_temp)
	fmt.Printf("PAS(U10) Temperature	:	%.3f°C\n", pas_temp)
	fmt.Printf("NAS(U11) Temperature	:	%.3f°C\n", nas_temp)
	fmt.Printf("SHV(U12) Temperature	:	%.3f°C\n", shv_temp)

	// Reset I2C Mux
	device.PCA9548A(i2c_bus, 0x70, -1)
	time.Sleep(750 * time.Millisecond)

}
