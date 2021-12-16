package days

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func Day13() {
	if res, err := day13_1(); err != nil {
		log.Fatalf("crash in day13_1: %v", err)
	} else {
		log.Printf("day13_2 result = %d", res)
	}
	if res, err := day13_2(); err != nil {
		log.Fatalf("crash in day13_2: %v", err)
	} else {
		log.Printf("day13_2 result = %d", res)
	}
}

type foldInstruction struct {
	axis  string
	value int
}

func makeFoldData() (err error, sheet map[string]bool, foldInstructions []foldInstruction) {
	sheet = make(map[string]bool)
	comma := []byte{byte(',')}
	eq := []byte{byte('=')}
	foldInstructions = make([]foldInstruction, 0)
	firstMode := true
	foldAlong := []byte("fold along ")
	err = ReadLines("./days/inputs/day13_1.txt", func(b []byte) error {
		if len(b) == 0 {
			firstMode = false
			return nil
		}
		if firstMode {
			spl := bytes.Split(b, comma)
			sheet[fmt.Sprintf("%s,%s", string(spl[0]), string(spl[1]))] = true
		} else {
			spl := bytes.Split(b, foldAlong)
			if len(spl) > 1 {
				fis := bytes.Split(spl[1], eq)
				fi := foldInstruction{
					axis:  string(fis[0]),
					value: byteArrToInt(fis[1]),
				}
				foldInstructions = append(foldInstructions, fi)
			}
		}
		return nil
	})
	return
}

func day13_1() (int, error) {
	err, sheet, foldInstructions := makeFoldData()
	if err != nil {
		return 0, err
	}
	// log.Printf("sheet = %v", sheet)
	// log.Print("=============================")
	// log.Printf("instructions = %v", foldInstructions)

	// lastX := 0
	// lastY := 0
	for _, fi := range foldInstructions {
		sheet = fold(sheet, fi)
		// if fi.axis == "x" {
		// 	lastX = fi.value
		// } else {
		// 	lastY = fi.value
		// }
		break
	}
	// if lastX == 0 {
	// 	lastX = findFirstXY(foldInstructions, "x")
	// }
	// if lastY == 0 {
	// 	lastY = findFirstXY(foldInstructions, "y")
	// }
	log.Printf("sheet = %v", sheet)
	// log.Printf("lastX = %d, lastY = %d", lastX, lastY)

	return /*lastX*lastY - */ len(sheet), err
}

func day13_2() (int, error) {
	err, sheet, foldInstructions := makeFoldData()
	if err != nil {
		return 0, err
	}
	lastX := 0
	lastY := 0
	for _, fi := range foldInstructions {
		sheet = fold(sheet, fi)
		if fi.axis == "x" {
			lastX = fi.value
		} else {
			lastY = fi.value
		}
	}

	log.Printf("sheet = %v", sheet)
	xs := make([][]int, lastY)
	for y := 0; y < lastY; y++ {
		xs[y] = make([]int, lastX)
	}
	for k, _ := range sheet {
		spl := strings.Split(k, ",")
		x := stringToInt(spl[0])
		y := stringToInt(spl[1])
		xs[y][x] = 4
	}
	for _, v := range xs {
		log.Print(v)
	}
	// returns UFRZKAUZ
	return len(sheet), err
}

// func findFirstXY(fis []foldInstruction, axis string) int {
// 	for _, v := range fis {
// 		if v.axis == axis {
// 			return v.value * 2
// 		}
// 	}
// 	return 0
// }

func fold(sheet map[string]bool, fi foldInstruction) map[string]bool {
	log.Printf("now folding %s=%d", fi.axis, fi.value)
	newSheet := make(map[string]bool)
	for k, _ := range sheet {
		spl := strings.Split(k, ",")
		x := stringToInt(spl[0])
		y := stringToInt(spl[1])
		if fi.axis == "x" {
			if x < fi.value {
				newSheet[k] = true
				continue
			} else if x > fi.value {
				// 13 -> 1
				// 13-7=6 7-(13-7)=1
				x = fi.value - (x - fi.value)
			}
		} else {
			if y < fi.value {
				newSheet[k] = true
				continue
			} else if y > fi.value {
				y = fi.value - (y - fi.value)
			}
		}
		newSheet[fmt.Sprintf("%d,%d", x, y)] = true
	}
	return newSheet
}
