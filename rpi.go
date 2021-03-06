package main

import (
	"flag"
	"fmt"
	"github.com/Che4ter/rpi_brain/arduino"
	"github.com/Che4ter/rpi_brain/brain"
	"github.com/Che4ter/rpi_brain/configuration"
	"github.com/Che4ter/rpi_brain/ipc"
	"os"
	"os/signal"
	"time"
)

//Define Arguments
//var argSerialPort = flag.String("serialport", "/dev/null", "ardino serial port")
var argConfigurationFile = flag.String("config", "configuration.json", "configuration file path")

var configFile string

func init() {
	//Parse Arguments
	flag.Parse()
	fmt.Println("***********************")
	fmt.Println("Pren 2 - Team 9")
	fmt.Println("== The Caterpillar ==")
	fmt.Println("**** v2.0 - FINAL **** ")
	fmt.Println("***********************")
	fmt.Println("initializing...")

	configFile = *argConfigurationFile
}

func main() {
	conf, _ := configuration.ParseConfiguration(configFile)

	//Define Channels
	arduinoSendingBridge := make(chan arduino.ArduinoPacket)
	arduinoReceivingBridge := make(chan arduino.ArduinoPacket)
	brainBridge := make(chan int)
	ipcBridge := make(chan string)
	doneBridge := make(chan bool)

	//Start Routines
	go ipc.RunUnixSocketServer(ipcBridge, conf)
	go arduino.RunArduinoServer(arduinoSendingBridge, arduinoReceivingBridge, conf)

	go brain.StartBrain(brainBridge, ipcBridge, conf, doneBridge, arduinoSendingBridge, arduinoReceivingBridge)
	//Wait for Stop
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	for {
		select {
		default:
			time.Sleep(500 * time.Millisecond)
		case <-quit:
			fmt.Println("shutdown caterpillar... ")
			fmt.Println("bye :'( ")
			return
		}
	}
}