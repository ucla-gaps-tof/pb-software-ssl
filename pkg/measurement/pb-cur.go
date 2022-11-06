package measurement

import (
	"fmt"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/device"
	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func PbCurrent() {
	i2c_bus := util.DetectMCP2221()

	// Preamp Board Sensors
	// +3.6V
	device.PCA9548A(i2c_bus, 0x70, 6)
	time.Sleep(750 * time.Millisecond)
	p3v6a_preamp_ina226 := device.INA226(i2c_bus, 0x46, 0.1, 0.42)
	p3v6a_preamp_vol := p3v6a_preamp_ina226.BusVoltage
	p3v6a_preamp_cur := p3v6a_preamp_ina226.Current
	p3v6a_preamp_pow := p3v6a_preamp_ina226.Power
	// -1.6V
	device.PCA9548A(i2c_bus, 0x70, 1)
	time.Sleep(750 * time.Millisecond)
	n1v6_preamp_vol := device.MAX11617(i2c_bus, 0x35, 3, 3.0)
	n1v6_preamp_cur := device.MAX11617(i2c_bus, 0x35, 2, 3.0) / 50 / 0.1
	n1v6_preamp_pow := n1v6_preamp_vol * n1v6_preamp_cur

	// Local Trigger Board Sensors
	// +3.4V Trenz
	device.PCA9548A(i2c_bus, 0x70, 5)
	time.Sleep(750 * time.Millisecond)
	p3v4a_ltb_ina219 := device.INA219(i2c_bus, 0x46, 0.1, 0.04)
	p3v4a_ltb_vol := p3v4a_ltb_ina219.BusVoltage
	p3v4a_ltb_cur := p3v4a_ltb_ina219.Current
	p3v4a_ltb_pow := p3v4a_ltb_ina219.Power
	// +3.4V DAC
	p3v4b_ltb_ina219 := device.INA219(i2c_bus, 0x47, 0.1, 0.1)
	p3v4b_ltb_vol := p3v4b_ltb_ina219.BusVoltage
	p3v4b_ltb_cur := p3v4b_ltb_ina219.Current
	p3v4b_ltb_pow := p3v4b_ltb_ina219.Power
	// +3.6V
	p3v6a_ltb_ina219 := device.INA219(i2c_bus, 0x4C, 0.1, 0.23)
	p3v6a_ltb_vol := p3v6a_ltb_ina219.BusVoltage
	p3v6a_ltb_cur := p3v6a_ltb_ina219.Current
	p3v6a_ltb_pow := p3v6a_ltb_ina219.Power
	// -1.6V
	device.PCA9548A(i2c_bus, 0x70, 3)
	time.Sleep(750 * time.Millisecond)
	n1v6_ltb_vol := device.MAX11617(i2c_bus, 0x35, 3, 3.0)
	n1v6_ltb_cur := device.MAX11617(i2c_bus, 0x35, 2, 3.0) / 100 / 0.1
	n1v6_ltb_pow := n1v6_ltb_vol * n1v6_ltb_cur

	// Print Result
	fmt.Printf("########## Preamp Board ##########\n")
	fmt.Printf("Preamp 3.6V Voltage	:	%.3fV\n", p3v6a_preamp_vol)
	fmt.Printf("Preamp 3.6V Current	:	%.3fA\n", p3v6a_preamp_cur)
	fmt.Printf("Preamp 3.6V Power	:	%.3f\nW", p3v6a_preamp_pow)
	fmt.Printf("Preamp -1.6V Voltage	:	%.3fV\n", -1*n1v6_preamp_vol)
	fmt.Printf("Preamp -1.6V Current	:	%.3fA\n", n1v6_preamp_cur)
	fmt.Printf("Preamp -1.6V Power	:	%.3fW\n", n1v6_preamp_pow)
	fmt.Printf("########## Local Trigger Board ##########\n")
	fmt.Printf("LTB 3.4V Trenz Voltage	:	%.3fV\n", p3v4a_ltb_vol)
	fmt.Printf("LTB 3.4V Trenz Current	:	%.3fA\n", p3v4a_ltb_cur)
	fmt.Printf("LTB 3.4V Trenz Power	:	%.3fW\n", p3v4a_ltb_pow)
	fmt.Printf("LTB 3.4V DAC Voltage	:	%.3fV\n", p3v4b_ltb_vol)
	fmt.Printf("LTB 3.4V DAC Current	:	%.3fA\n", p3v4b_ltb_cur)
	fmt.Printf("LTB 3.4V DAC Power	:	%.3fW\n", p3v4b_ltb_pow)
	fmt.Printf("LTB 3.6V Voltage	:	%.3fV\n", p3v6a_ltb_vol)
	fmt.Printf("LTB 3.6V Current	:	%.3fA\n", p3v6a_ltb_cur)
	fmt.Printf("LTB 3.6V Power		:	%.3fW\n", p3v6a_ltb_pow)
	fmt.Printf("LTB -1.6V Voltage	:	%.3fV\n", -1*n1v6_ltb_vol)
	fmt.Printf("LTB -1.6V Current	:	%.3fA\n", n1v6_ltb_cur)
	fmt.Printf("LTB -1.6V Power		:	%.3fW\n", n1v6_ltb_pow)

	// Reset I2C Mux
	device.PCA9548A(i2c_bus, 0x70, -1)
	time.Sleep(750 * time.Millisecond)

}
