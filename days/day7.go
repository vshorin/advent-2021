package days

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	if res, err := day7_1(); err != nil {
		log.Fatalf("crash in day7_1: %v", err)
	} else {
		log.Printf("day7_1 result = %d", res)
	}
	if res, err := day7_2(); err != nil {
		log.Fatalf("crash in day7_2: %v", err)
	} else {
		log.Printf("day7_2 result = %d", res)
	}
}

func countFuels(t int, numbers []int) int {
	var counter int
	for _, v := range numbers {
		if diff := (t - v); diff >= 0 {
			counter += diff
		} else {
			counter += -1 * diff
		}
	}
	return counter
}
func countFuelsProg(t int, numbers []int) int {
	var counter int
	for _, v := range numbers {
		if diff := (t - v); diff > 0 {
			counter += fuelProgCounter(1, diff, 1)
		} else if diff < 0 {
			counter += fuelProgCounter(1, -1*diff, 1)
		}
	}
	return counter
}

func fuelProgCounter(current int, max int, res int) int {
	if current < max {
		res += (current + 1)
		current++
		return fuelProgCounter(current, max, res)
	} else {
		//log.Printf("prog counted %d", res)
		return res
	}

}

func crabMinMax(crabs []int) (int, int) {
	sort.Ints(crabs)
	return crabs[0], crabs[len(crabs)-1]
}

func makeCrabs() ([]int, error) {
	var numbers []int
	err := ReadLines("./days/inputs/day7_1.txt", func(b []byte) error {
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
	return numbers, err
}

func day7_1() (int, error) {
	numbers, err := makeCrabs()
	if err != nil {
		return 0, err
	}
	log.Printf("day7_1 numbers = %v", numbers)

	min, max := crabMinMax(numbers)

	smallestFuel := 0
	smallestI := 0
	for i := min; i < max; i++ {
		if c := countFuels(i, numbers); c < smallestFuel || smallestFuel == 0 {
			smallestFuel = c
			smallestI = i
		}
	}
	log.Printf("smallest fuel = %d at i %d", smallestFuel, smallestI)

	return smallestFuel, err
}

func day7_2() (int, error) {
	numbers, err := makeCrabs()
	if err != nil {
		return 0, err
	}
	log.Printf("day7_1 numbers = %v", numbers)

	min, max := crabMinMax(numbers)

	smallestFuel := 0
	smallestI := 0
	for i := min; i < max; i++ {
		c := countFuelsProg(i, numbers)
		log.Printf("now c = %d", c)
		if c < smallestFuel || smallestFuel == 0 {
			log.Printf("-- setting smallest at %d to %d", i, c)
			smallestFuel = c
			smallestI = i
		} else {
			break
		}
	}
	log.Printf("smallest fuel = %d at i %d", smallestFuel, smallestI)

	return smallestFuel, err
}
