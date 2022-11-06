package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ucla-gaps-tof/pb-software-ssl/pkg/measurement"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No argument is given.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ts":
		measurement.PbTemperature()
	case "cs":
		measurement.PbCurrent()
	case "preamp_ts":
		measurement.PreampTemperature()
	case "bias_read":
		measurement.PreampBiasRead()
	case "bias_set":
		measurement.PreampBiasSet(3545)
		time.Sleep(10 * time.Millisecond)
		measurement.PreampBiasRead()
	case "bias_off":
		measurement.PreampBiasSet(0)
		time.Sleep(10 * time.Millisecond)
		measurement.PreampBiasRead()
	case "ltb_on":
		measurement.LtbPow(true)
	case "ltb_off":
		measurement.LtbPow(false)
	default:
		fmt.Println("Wrong Argument")
	}
}
