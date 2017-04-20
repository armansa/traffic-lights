package main

import "fmt"

type Intersection struct {
  CycleLength int
  Time int
  Lights []TrafficLight
}

func (in *Intersection) Tick() (changed bool) {
  changed = false
  in.Time++
  for index := range in.Lights {
    changed = in.Lights[index].Tick(in.Time % in.CycleLength) || changed
  }
  return
}

func (in *Intersection) String() (result string) {
  result = fmt.Sprintf("Time %4d:  ", in.Time)
  for _, light := range in.Lights {
    result += light.String()
  }
  return
}
