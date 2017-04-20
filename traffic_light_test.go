package main

import (
  "testing"
  "strings"
)

var traffic_light TrafficLight = TrafficLight{"Test Light", "Red", []Step{ Step{10, "Green"}, Step{20, "Yellow"}, Step{30, "Red"} } }

func TestTick(t *testing.T) {
  traffic_light.Tick(15)
  if traffic_light.Status != "Green" {
    t.Error("Expected Green, got ", traffic_light.Status)
  }
  traffic_light.Tick(25)
  if traffic_light.Status != "Yellow" {
    t.Error("Expected Yellow, got ", traffic_light.Status)
  }
  traffic_light.Tick(30)
  if traffic_light.Status != "Red" {
    t.Error("Expected Red, got ", traffic_light.Status)
  }
}

func TestString(t *testing.T) {
  str := traffic_light.String()
  if ! strings.Contains(str, traffic_light.Status) {
    t.Error("Expected to include status:", traffic_light.Status)
  }
}
