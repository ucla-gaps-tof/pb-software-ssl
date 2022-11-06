package util

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func DetectMCP2221() int {
	ps, err := exec.Command("bash", "-c", "i2cdetect -l | grep -i i2c-mcp2221").Output()
	if err != nil {
		log.Fatalf("Could not find MCP2221 Device on your computer: %+v\n", err)
		os.Exit(1)
	}
	i2c_bus_str := strings.Split(strings.Split(string(ps), "\t")[0], "-")[1]
	i2c_bus, err := strconv.Atoi(i2c_bus_str)
	if err != nil {
		log.Fatal(err)
	}

	return i2c_bus
}
