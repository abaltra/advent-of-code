package day15

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Sensor struct {
	position             *Position
	nearestBeacon        *Position
	maxInfluenceDistance int
}

func DistanceBetween(p1 Position, p2 Position) int {
	distX := int(math.Max(float64(p1.x), float64(p2.x)) - math.Min(float64(p1.x), float64(p2.x)))
	distY := int(math.Max(float64(p1.y), float64(p2.y)) - math.Min(float64(p1.y), float64(p2.y)))
	return distX + distY
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

type Segment struct {
	start int
	end   int
}

func processInfluenceArea(influencedSegments map[int][]Segment, sensor *Sensor) {
	fmt.Printf("Processing influence area for sensor (%s)\n", sensor.position.ToString())
	for i := 0; i <= sensor.maxInfluenceDistance; i++ {
		segment := &Segment{
			start: sensor.position.x - (sensor.maxInfluenceDistance - i),
			end:   sensor.position.x + (sensor.maxInfluenceDistance - i),
		}
		influencedSegments[sensor.position.y-i] = append(influencedSegments[sensor.position.y-i], *segment)
		influencedSegments[sensor.position.y+i] = append(influencedSegments[sensor.position.y+i], *segment)
	}
}

func Run(lines []string) {
	println("day15")
	sensors := []*Sensor{}
	rightmostElement := 0
	leftmostElement := math.MaxInt
	influencedSegments := map[int][]Segment{}
	for _, l := range lines {
		sensorPosition, beaconPosition := parseLine(l)
		sensor := &Sensor{position: sensorPosition, nearestBeacon: beaconPosition, maxInfluenceDistance: DistanceBetween(*sensorPosition, *beaconPosition)}
		fmt.Printf("Sensor at (%s), closest beacon at (%s) with a distance of %d\n", sensor.position.ToString(), sensor.nearestBeacon.ToString(), sensor.maxInfluenceDistance)
		sensors = append(sensors, sensor)

		rightmostElement = int(math.Max(float64(sensor.position.x+sensor.maxInfluenceDistance), float64(rightmostElement)))
		leftmostElement = int(math.Min(float64(sensor.position.x-sensor.maxInfluenceDistance), float64(leftmostElement)))

		processInfluenceArea(influencedSegments, sensor)
	}
	println()


	coveredPositions := 0
	fmt.Printf("Leftmost index: %d\n", leftmostElement)
	skippableEntries := map[int]bool{}
	knownSensorPositions := map[string]bool{}
	knownBeaconPositions := map[string]bool{}

	const Y_VALUE int = 2000000
	const MAX_COORD_VALUE int = 4000000

	for i := range influencedSegments {
		sort.Slice(influencedSegments[i], func(_i, _j int) bool {
			return influencedSegments[i][_i].start < influencedSegments[i][_j].start
		})
	}
	for _, sensor := range sensors {
		if sensor.position.y == Y_VALUE {
			skippableEntries[sensor.position.x] = true
		}
		if sensor.nearestBeacon.y == Y_VALUE {
			skippableEntries[sensor.nearestBeacon.x] = true
		}
		knownSensorPositions[sensor.position.ToString()] = true
		knownBeaconPositions[sensor.nearestBeacon.ToString()] = true
	}

	for i := leftmostElement; i <= rightmostElement; i++ {
		pos := Position{
			x: i,
			y: Y_VALUE,
		}

		for _, sensor := range sensors {
			sd := DistanceBetween(*sensor.position, *sensor.nearestBeacon)
			sp := DistanceBetween(*sensor.position, pos)
			if sd >= sp && !skippableEntries[pos.x] {
				coveredPositions++
				break
			}
		}
	}

	fmt.Printf("There are %d covered positions on line y=10\n", coveredPositions)

	println("part2")

	for y := 0; y <= MAX_COORD_VALUE; y++ {
		if y % 1000 == 0 {
			fmt.Printf("Checking Y=%d\n", y)
		}
		for x := 0; x <= MAX_COORD_VALUE; x++ {
			isContained := false
			for _, segment := range influencedSegments[y] {
				if x >= segment.start && x <= segment.end {
					isContained = true
					x = segment.end
					break
				}
			}
			if !isContained {
				fmt.Printf("Found available space at (%d,%d). Tunning frequency: %d\n", x, y, 4000000 * x + y)
				return
			}
		}
	}
}
