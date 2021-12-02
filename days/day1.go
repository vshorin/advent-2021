package days

import (
	"log"
	"strconv"
)

func Day1() {
	if _, err := day1_1(); err != nil {
		log.Fatalf("crash in day1_1: %v", err)
	}
	if _, err := day1_2(); err != nil {
		log.Fatalf("crash in day1_2: %v", err)
	}
}

func day1_1() (int, error) {
	var prevVal int
	count := 0
	err := ReadLines("./days/inputs/day1_1.txt", func(b []byte) error {
		newVal, err := strconv.Atoi(string(b))
		if err != nil {
			return err
		}
		//log.Printf("newVal = %d; prevVal = %d", newVal, prevVal)
		if prevVal > 0 && newVal > prevVal {
			//log.Printf("increase!")
			count++
		}
		prevVal = newVal
		return nil
	})
	log.Printf("day1_1 result = %d", count)
	return count, err
}

func day1_2() (int, error) {
	var slice []int
	count := 0
	err := ReadLines("./days/inputs/day1_1.txt", func(b []byte) error {
		newVal, err := strconv.Atoi(string(b))
		if err != nil {
			return err
		}
		var prevSum int
		log.Printf("slice = %v", slice)
		if len(slice) >= 3 {
			prevSum = slice[0] + slice[1] + slice[2]
			slice = slice[1:]
		}
		slice = append(slice, newVal)
		if len(slice) >= 3 {
			newSum := slice[0] + slice[1] + slice[2]
			if prevSum > 0 && newSum > prevSum {
				log.Printf("increase!")
				count++
			}
		}
		return nil
	})
	log.Printf("day1_2 result = %d", count)
	return count, err
}
