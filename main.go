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
        timeStamp := makeTimestamp()
    	for _, sensor := range sensors {
            t, err := ds18b20.Temperature(sensor)
            if err == nil {
            	fmt.Printf("%d,%s,%.2f\n",timeStamp, sensor, t)
            }
    	}
        time.Sleep(30 * time.Second)
    }
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}
