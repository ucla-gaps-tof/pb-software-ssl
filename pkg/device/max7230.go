package device

import (
	"log"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/util"
)

func MAX7320(bus int, address uint8, operation bool) {
	c, err := util.OpenI2C(bus, address)
	if err != nil {
		log.Fatalf("Could not open MAX7320: %+v", err)
	}
	defer c.CloseI2C()

	if operation {
		c.WriteRegI2C(address, 0, 0x0F)
	} else {
		c.WriteRegI2C(address, 0, 0xF0)
	}
}
