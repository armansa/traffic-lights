package main

import "fmt"

type Step struct {
  Time int
  Status string
}

type TrafficLight struct {
  Name string
  Status string
  Program []Step
}

func (tl *TrafficLight) Tick(time int) (changed bool) {
  step := 0
  for step < len(tl.Program) - 1 && tl.Program[step +1].Time <= time {
    step++
  }
  new_status := tl.Program[step].Status
  changed = new_status != tl.Status
  tl.Status = new_status
  return
}

func (tl *TrafficLight) String() string {
  return fmt.Sprintf("%6s: %6s   ", tl.Name, tl.Status)
}
