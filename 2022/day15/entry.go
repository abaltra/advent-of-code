package day15

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
  x int
  y int
}

func (p Position) ToString() string {
  return fmt.Sprintf("%d,%d", p.x, p.y)
}

func parseLine(l string) (*Position, *Position) {
 	parts := strings.Split(l, ":")
  sensorRaw := strings.TrimSpace(strings.Replace(parts[0], "Sensor at", "", 1))
  beaconRaw := strings.TrimSpace(strings.Replace(parts[1], "closest beacon is at", "", 1))

  sensorRawCoords := strings.Split(sensorRaw, ",")
  beaconRawCoords := strings.Split(beaconRaw, ",")

  beaconX, _ := strconv.Atoi(strings.TrimSpace(strings.Split(beaconRawCoords[0], "=")[1]))
  beaconY, _ := strconv.Atoi(strings.TrimSpace(strings.Split(beaconRawCoords[1], "=")[1]))

  sensorX, _ := strconv.Atoi(strings.TrimSpace(strings.Split(sensorRawCoords[0], "=")[1]))
  sensorY, _ := strconv.Atoi(strings.TrimSpace(strings.Split(sensorRawCoords[1], "=")[1]))

  return &Position{x: sensorX, y: sensorY}, &Position{x: beaconX, y: beaconY}


}

func Run(lines []string) {
  println("day15")
  for _, l := range lines {
    sensor, beacon := parseLine(l)
    fmt.Printf("Sensor at (%s), closest beacon at (%s)\n", sensor.ToString(), beacon.ToString())
  }
}
