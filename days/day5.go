package days

import (
	"bytes"
	"log"
	"math"
	"strconv"
)

func Day5() {
	if res, err := day5_1(); err != nil {
		log.Fatalf("crash in day5_1: %v", err)
	} else {
		log.Printf("day5_1 result = %d", res)
	}
	if res, err := day5_2(); err != nil {
		log.Fatalf("crash in day5_2: %v", err)
	} else {
		log.Printf("day5_2 result = %d", res)
	}
}

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func makeLines5() (lines []line, err error) {
	del := []byte(" -> ")
	comma := []byte(",")
	err = ReadLines("./days/inputs/day5_1.txt", func(b []byte) error {
		points := bytes.Split(b, del)
		point1 := bytes.Split(points[0], comma)
		point2 := bytes.Split(points[1], comma)
		line := line{}
		if line.x1, err = strconv.Atoi(string(point1[0])); err != nil {
			return err
		}
		if line.y1, err = strconv.Atoi(string(point1[1])); err != nil {
			return err
		}
		if line.x2, err = strconv.Atoi(string(point2[0])); err != nil {
			return err
		}
		if line.y2, err = strconv.Atoi(string(point2[1])); err != nil {
			return err
		}
		lines = append(lines, line)

		return nil
	})
	return
}

func fillIntersections5(lines []line, checkDiagonal bool) (res int, err error) {
	board := make(map[string]bool)
	isections := make(map[string]bool)
	for _, l := range lines {
		if l.x1 == l.x2 {
			amount := int(math.Abs(float64(l.y1 - l.y2)))
			for i := 0; i <= amount; i++ {
				var mult = 1
				if l.y1 > l.y2 {
					mult = -1
				}
				if key := strconv.Itoa(l.x1) + "," + strconv.Itoa(l.y1+mult*i); board[key] {
					isections[key] = true
				} else {
					board[key] = true
				}
			}
		} else if l.y1 == l.y2 {
			amount := int(math.Abs(float64(l.x1 - l.x2)))
			for i := 0; i <= amount; i++ {
				var mult = 1
				if l.x1 > l.x2 {
					mult = -1
				}
				if key := strconv.Itoa(l.x1+mult*i) + "," + strconv.Itoa(l.y1); board[key] {
					isections[key] = true
				} else {
					board[key] = true
				}
			}
		} else if checkDiagonal {
			amount := int(math.Abs(float64(l.x1 - l.x2)))
			for i := 0; i <= amount; i++ {
				var multX = 1
				if l.x1 > l.x2 {
					multX = -1
				}
				var multY = 1
				if l.y1 > l.y2 {
					multY = -1
				}
				if key := strconv.Itoa(l.x1+multX*i) + "," + strconv.Itoa(l.y1+multY*i); board[key] {
					isections[key] = true
				} else {
					board[key] = true
				}
			}
		}
	}

	return len(isections), err
}

func day5_1() (int, error) {
	lines, err := makeLines5()
	if err != nil {
		return 0, err
	}
	return fillIntersections5(lines, false)
}

func day5_2() (int, error) {
	lines, err := makeLines5()
	if err != nil {
		return 0, err
	}
	return fillIntersections5(lines, true)
}

// func day5_2() (int, error) {
// 	var space = []byte(" ")
// 	var hor int = 0
// 	var ver int = 0
// 	var aim int = 0

// 	err := ReadLines("./days/inputs/day5_1.txt", func(b []byte) error {
// 		split := bytes.Split(b, space)
// 		val, err := strconv.Atoi(string(split[1]))
// 		if err != nil {
// 			return err
// 		}
// 		if len(split[0]) == 2 {
// 			aim -= val
// 			log.Printf("up %d, aim %d", val, aim)
// 		}
// 		if len(split[0]) == 4 {
// 			aim += val
// 			log.Printf("down %d, aim %d", val, aim)
// 		}
// 		if len(split[0]) == 7 {
// 			log.Printf("forward %d", val)
// 			hor += val
// 			ver += (val * aim)
// 		}
// 		return nil
// 	})
// 	res := hor * ver
// 	log.Printf("day5_2 result = %d", res)
// 	return res, err
// }
