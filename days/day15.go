package days

import (
	"fmt"
	"log"
	"strings"
)

func Day15() {
	// if res, err := day15_1(); err != nil {
	// 	log.Fatalf("crash in day15_1: %v", err)
	// } else {
	// 	log.Printf("day15_2 result = %d", res)
	// }
	if res, err := day15_2(); err != nil {
		log.Fatalf("crash in day15_2: %v", err)
	} else {
		log.Printf("day15_2 result = %d", res)
	}
}

type chiton struct {
	totalWeight int
	path        []string
	visited     bool
}

func day15_1() (int, error) {
	return processChitons(1)
}

func day15_2() (int, error) {
	return processChitons(5) // 3040 this takes 7m8.7469925s
}

func processChitons(scale int) (int, error) {
	startWeights := make(map[string]int)
	tracker := make(map[string]chiton)
	rowSize := 0
	count := 0
	err := ReadLines("./days/inputs/day15_1.txt", func(b []byte) error {
		rowSize = len(b)
		for k, v := range b {
			startWeights[fmt.Sprintf("%d:%d", k, count)] = byteToInt(v)
		}
		count++
		return nil
	})
	weights := extendChitonMap(scale, startWeights, rowSize, count)
	count *= scale
	rowSize *= scale
	log.Printf("weights = %v", weights)
	startChit := chiton{
		totalWeight: 0,
		path:        []string{"0:0"},
	}
	tracker["0:0"] = startChit
	point := "0:0"
	prevPoint := ""
	endPoint := fmt.Sprintf("%d:%d", rowSize-1, count-1)
	for {
		calcWeights(point, prevPoint, weights, tracker)
		prevPoint = point
		point = findLowestUnvisitedChiton(tracker)
		if point == endPoint {
			break
		}
	}
	cheapPath := tracker[endPoint]
	log.Printf("path was %v", cheapPath)
	// for i := 0; i < rowSize; i++ {
	// 	for j := 0; j < count; j++ {
	// 		w := getChitWeight(cheapPath.path, fmt.Sprintf("%d:%d", j, i), weights)
	// 		if w < 0 {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Print(w)
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	res := cheapPath.totalWeight
	return res, err
}

func extendChitonMap(amount int, origWeights map[string]int, row int, col int) (weights map[string]int) {
	weights = make(map[string]int)
	for j := 0; j < amount; j++ {
		for i := 0; i < amount; i++ {
			for k, v := range origWeights {
				spl := strings.Split(k, ":")
				x := stringToInt(spl[0])
				y := stringToInt(spl[1])
				val := v + i + j
				if val > 9 {
					val = val - 9
				}
				weights[fmt.Sprintf("%d:%d", x+row*i, y+col*j)] = val
			}
		}
	}
	return
}

func getChitWeight(path []string, key string, weights map[string]int) int {
	for _, v := range path {
		if v == key {
			return weights[v]
		}
	}
	return -1
}

func findLowestUnvisitedChiton(tracker map[string]chiton) string {
	var lowK string
	var lowV int
	for k, v := range tracker {
		if v.visited {
			continue
		}
		if v.totalWeight < lowV || lowV == 0 {
			lowK = k
			lowV = v.totalWeight
		}
	}
	return lowK
}

func calcWeights(thisPoint, fromPoint string, weights map[string]int, tracker map[string]chiton) {
	//log.Printf("calcWeights now %s from %s", thisPoint, fromPoint)
	spl := strings.Split(thisPoint, ":")
	x := stringToInt(spl[0])
	y := stringToInt(spl[1])
	left := fmt.Sprintf("%d:%d", x-1, y)
	right := fmt.Sprintf("%d:%d", x+1, y)
	top := fmt.Sprintf("%d:%d", x, y-1)
	bottom := fmt.Sprintf("%d:%d", x, y+1)
	nearPoints := make(map[string]int)
	if v, ok := weights[left]; ok && left != fromPoint {
		nearPoints[left] = v
	}
	if v, ok := weights[right]; ok && right != fromPoint {
		nearPoints[right] = v
	}
	if v, ok := weights[top]; ok && top != fromPoint {
		nearPoints[top] = v
	}
	if v, ok := weights[bottom]; ok && bottom != fromPoint {
		nearPoints[bottom] = v
	}
	thisPointChit := tracker[thisPoint]
	thisPointChit.visited = true
	tracker[thisPoint] = thisPointChit
	tracker[fromPoint] = chiton{
		visited: true,
	}
	for k, v := range nearPoints {
		weight := thisPointChit.totalWeight + v
		if prev, ok := tracker[k]; ok && (prev.totalWeight <= weight || prev.visited) {
			continue
		}
		weightSlice := make([]string, len(thisPointChit.path))
		copy(weightSlice, thisPointChit.path)
		weightSlice = append(weightSlice, k)
		chit := chiton{path: weightSlice, totalWeight: weight}
		tracker[k] = chit
	}
}
