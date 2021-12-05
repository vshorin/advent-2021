package days

import (
	"log"
	"math"
)

var zero = []byte("0")[0]

func Day3() {
	if _, err := day3_1(); err != nil {
		log.Fatalf("crash in day2_1: %v", err)
	}
	if _, err := day3_2(); err != nil {
		log.Fatalf("crash in day2_2: %v", err)
	}
}

func day3_1() (int, error) {
	var total int
	var counters [12]int
	var res [12]int
	err := ReadLines("./days/inputs/day3_1.txt", func(b []byte) error {
		total++
		for i, v := range b {
			if v > zero {
				counters[i]++
			}
		}
		return nil
	})
	log.Printf("day3_1 counters = %v", counters)
	var gamma int
	var epsilon int
	for i := len(counters) - 1; i >= 0; i-- {
		res[i] = int(math.Round(float64(counters[i]) / float64(total)))
		if res[i] > 0 {
			gamma += int(math.Pow(2, float64(len(counters)-1-i)))
		} else {
			epsilon += int(math.Pow(2, float64(len(counters)-1-i)))
		}
	}
	log.Printf("day3_1 res = %v", res) // [0 1 0 1 0 0 1 1 1 0 0 1]
	log.Printf("day3_1 gamma = %d, epsilon = %d", gamma, epsilon)
	mult := gamma * epsilon
	log.Printf("day3_1 result = %d", mult)
	return mult, err
}

func day3_2() (int, error) {
	var startw0 [][]byte
	var startw1 [][]byte
	err := ReadLines("./days/inputs/day3_1.txt", func(b []byte) error {
		if b[0] > zero {
			//log.Printf("appending ONES because %d in %v", b[0], b)
			startw1 = append(startw1, b)
		} else {
			//log.Printf("appending ZEROS because %d in %v", b[0], b)
			startw0 = append(startw0, b)
		}
		return nil
	})
	var oxy []byte
	var co2 []byte
	if len(startw0) > len(startw1) {
		log.Print("day3_2 at start, 0 are more common")
		oxy = makeRatings(startw0, 1, true)
		co2 = makeRatings(startw1, 1, false)
	} else {
		log.Print("day3_2 at start, 1 are more common")
		oxy = makeRatings(startw1, 1, true)
		co2 = makeRatings(startw0, 1, false)
	}

	log.Printf("day3_2 oxy = %v, co2 = %v", oxy, co2)
	return 0, err
}

func makeRatings(slice [][]byte, i int, checkMostCommon bool) (ret []byte) {
	var zeros [][]byte
	var ones [][]byte
	for _, v := range slice {
		log.Printf("check checkMostCommon = %v at i = %d for v = %v", checkMostCommon, i, v)
		if v[i] > zero {
			ones = append(ones, v)
		} else {
			zeros = append(zeros, v)
		}
	}
	var res [][]byte
	if len(ones) >= len(zeros) {
		if checkMostCommon {
			res = ones
		} else {
			res = zeros
		}
	}
	log.Printf("day3_2 at i = %d, len res = %d", i, len(res))

	i++
	if i >= len(slice[0]) {
		return res[0]
	} else {
		return makeRatings(res, i, checkMostCommon)
	}
}
