package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	WebUi   bool
	WebPort int

	SerialPort           string
	SerialPortAutoDetect bool
	SerialBaudRate       int

	ParcourDirection int

	UnixSocketAddress string

	MaxSpeed float64

	UltrasonicOutPin int
	UltrasonicInPin  int
}

func ParseConfiguration(configFile string) (configuration Configuration, err error) {
	// Create a default configuration.
	config := Configuration{true, 1234, "/dev/null", true, 9600, 1, "/tmp/team9.sock", 50.0, 17, 18}

	// Open the configuration file.
	fmt.Println("try to open config file: ", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("error: could not open config file")
		fmt.Println("-> use default values")
		return config, err
	}

	fmt.Println("parsing config file")
	// Parse JSON in the configuration file.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error: config file not valid")
		return config, err
	}
	fmt.Println("config file successfully loaded :)")

	return config, nil
}
