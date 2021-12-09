package days

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Day7() {
	if res, err := day7_1(); err != nil {
		log.Fatalf("crash in day7_1: %v", err)
	} else {
		log.Printf("day7_1 result = %d", res)
	}
	// if res, err := day7_2(); err != nil {
	// 	log.Fatalf("crash in day7_2: %v", err)
	// } else {
	// 	log.Printf("day7_2 result = %d", res)
	// }
}

func countFuels(t int, numbers []int) int {
	var counter int
	for _, v := range numbers {
		counter += int(math.Abs(float64(t - v)))
	}
	return counter
}

func day7_1() (int, error) {
	var numbers []int
	err := ReadLines("./days/inputs/day7_x.txt", func(b []byte) error {
		splt := strings.Split(string(b), ",")
		for _, v := range splt {
			if nr, err := strconv.Atoi(v); err != nil {
				return err
			} else {
				numbers = append(numbers, nr)
			}

		}
		return nil
	})
	log.Printf("day7_1 numbers = %v", numbers)
	var total int
	for _, v := range numbers {
		total += v
	}
	mean := total / len(numbers)
	log.Printf("mean value = %d", mean)

	prev := 0

	for i := mean; i >= 0; i-- {
		now := countFuels(i, numbers)
		if prev != 0 {
			if prev <= now {
				log.Printf("prev >= now at %d (prev was %d)", i, i+1)
				break
			}
		} else {
			prev = now
		}
	}

	return 0, err
}

// func day7_2() (int, error) {
// 	var startw0 [][]byte
// 	var startw1 [][]byte
// 	err := ReadLines("./days/inputs/day7_1.txt", func(b []byte) error {
// 		bx := append([]byte(nil), b...) // appending with b directly messes things up completely, so do a fresh copy
// 		if b[0] > zero {
// 			//log.Printf("appending ONES because %d in %v", b[0], b)
// 			startw1 = append(startw1, bx)
// 		} else {
// 			//log.Printf("appending ZEROS because %d in %v", b[0], b)
// 			startw0 = append(startw0, bx)
// 		}
// 		return nil
// 	})
// 	var oxy []byte
// 	var co2 []byte
// 	if len(startw0) > len(startw1) {
// 		log.Print("day7_2 at start, 0 are more common")
// 		oxy = makeRatings(startw0, 1, true)
// 		co2 = makeRatings(startw1, 1, false)
// 	} else {
// 		log.Print("day7_2 at start, 1 are more common")
// 		oxy = makeRatings(startw1, 1, true)
// 		co2 = makeRatings(startw0, 1, false)
// 	}

// 	log.Printf("day7_2 oxy = %v, co2 = %v", oxy, co2)
// 	mult := calcExp(12, oxy) * calcExp(12, co2)
// 	log.Printf("day7_2 result = %d", mult)
// 	return mult, err
// }
