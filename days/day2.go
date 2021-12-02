package days

import (
	"bytes"
	"log"
	"strconv"
)

func Day2() {
	if _, err := day2_1(); err != nil {
		log.Fatalf("crash in day2_1: %v", err)
	}
	if _, err := day2_2(); err != nil {
		log.Fatalf("crash in day2_2: %v", err)
	}
}

func day2_1() (int, error) {
	var space = []byte(" ")
	var hor int = 0
	var ver int = 0

	err := ReadLines("./days/inputs/day2_1.txt", func(b []byte) error {
		split := bytes.Split(b, space)
		val, err := strconv.Atoi(string(split[1]))
		if err != nil {
			return err
		}
		if len(split[0]) == 2 {
			//log.Printf("up %d", val)
			ver -= val
		}
		if len(split[0]) == 4 {
			//log.Printf("down %d", val)
			ver += val
		}
		if len(split[0]) == 7 {
			//log.Printf("forward %d", val)
			hor += val
		}
		return nil
	})
	res := hor * ver
	log.Printf("day2_1 result = %d", res)
	return res, err
}

func day2_2() (int, error) {
	var space = []byte(" ")
	var hor int = 0
	var ver int = 0
	var aim int = 0

	err := ReadLines("./days/inputs/day2_1.txt", func(b []byte) error {
		split := bytes.Split(b, space)
		val, err := strconv.Atoi(string(split[1]))
		if err != nil {
			return err
		}
		if len(split[0]) == 2 {
			aim -= val
			log.Printf("up %d, aim %d", val, aim)
		}
		if len(split[0]) == 4 {
			aim += val
			log.Printf("down %d, aim %d", val, aim)
		}
		if len(split[0]) == 7 {
			log.Printf("forward %d", val)
			hor += val
			ver += (val * aim)
		}
		return nil
	})
	res := hor * ver
	log.Printf("day2_2 result = %d", res)
	return res, err
}
