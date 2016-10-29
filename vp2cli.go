package main

import (
  "fmt"
  "github.com/derekpitt/weather_station"
)


func main() {
  ws, err := weather_station.New("/dev/ttyUSB0")
  if err != nil {
    panic(err)
  }

  // we take out a lock each time we call GetSample. You can call from as many
  // goroutines as you want!
    sample, _ := ws.GetSample()
    fmt.Printf("%#v", sample)

}
