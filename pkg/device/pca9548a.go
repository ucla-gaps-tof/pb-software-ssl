package device

import (
	"log"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func PCA9548A(bus int, address uint8, channel int8) {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open PCA9548A: %+v", err)
	}
	defer c.CloseI2C()

	c.WriteRegI2C(address, 0x00, 0x00)

	var chn uint8
	switch channel {
	case 0:
		chn = 0x01
	case 1:
		chn = 0x02
	case 2:
		chn = 0x04
	case 3:
		chn = 0x08
	case 4:
		chn = 0x10
	case 5:
		chn = 0x20
	case 6:
		chn = 0x40
	case 7:
		chn = 0x80
	case -1:
		chn = 0x00
	default:
		chn = 0x00
	}

	err = c.WriteRegI2C(address, 0x00, chn)
	if err != nil {
		log.Fatalf("Could not write byte: %+v", err)
	}
}
