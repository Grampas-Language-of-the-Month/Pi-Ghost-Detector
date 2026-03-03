package main

import (
	sr501_mod "github.com/Low-Level-Devils/HC-SR501-Module_GoLang"
	"github.com/fatih/color"
)

func main() {
	color.Cyan("Starting...")

	pirSensor, err := sr501_mod.NewSensor("GPIO17")

	if err != nil {
		color.Red("Failed to start sensor")

		return
	}

	go pirSensor.Watch()

	color.Blue("Sensor initialized")

	for pirEvent := range pirSensor.Events {
		if pirEvent {
			color.Green("Motion Detected")
		} else {
			color.Green("Sensor reset")
		}
	}
}
