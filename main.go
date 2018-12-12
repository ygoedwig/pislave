package main

import (
    "fmt"
    "time"

    "github.com/yryz/ds18b20"
)

func main() {
    sensors, err := ds18b20.Sensors()
    if err != nil {
        panic(err)
    }

    fmt.Printf("sensor IDs: %v\n", sensors)

    for true {
        timex := time.Now()
    	for _, sensor := range sensors {
            t, err := ds18b20.Temperature(sensor)
            if err == nil {
            	fmt.Printf("At %s sensor: %s temperature: %.2f°C\n",timex, sensor, t)
            }
    	}
        time.Sleep(30 * time.Second)
    }
}
