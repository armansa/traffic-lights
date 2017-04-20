package main

import (
  "strconv"
  "testing"
  "strings"
)

var intersection Intersection = Intersection{40, 10, []TrafficLight{ TrafficLight{"Test Light", "Red", []Step{ Step{10, "Green"}, Step{20, "Yellow"}, Step{30, "Red"} } } } }

func TestIntersectionTick(t *testing.T) {
  changed := intersection.Tick()
  if ! changed {
    t.Error("Expected changed = true, got false")
  }
}

func TestIntersectionString(t *testing.T) {
  str := intersection.String()
  if ! strings.Contains(str, strconv.Itoa(intersection.Time)) {
    t.Error(str, strconv.Itoa(intersection.Time))
  }
}
