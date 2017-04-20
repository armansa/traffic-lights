package main

import (
  "encoding/json"
  "os"
  "fmt"
  "time"
)

func main() {
  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  intersection := Intersection{}
  err := decoder.Decode(&intersection)
  if err != nil {
    fmt.Println("error:", err)
  }
  ticker := time.NewTicker(time.Second)
  go func() {
      println(intersection.String())
      for range ticker.C {
        if intersection.Tick() {
          println(intersection.String())
        }
      }
  }()
  time.Sleep(time.Second * 1801)
  ticker.Stop()
}
