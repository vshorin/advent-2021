package days

import (
	"bytes"
	"log"
	"regexp"
)

func Day8() {
	if res, err := day8_1(); err != nil {
		log.Fatalf("crash in day8_1: %v", err)
	} else {
		log.Printf("day8_1 result = %d", res)
	}
	// if res, err := day8_2(); err != nil {
	// 	log.Fatalf("crash in day8_2: %v", err)
	// } else {
	// 	log.Printf("day8_2 result = %d", res)
	// }
}

func day8_1() (int, error) {
	var pipe = []byte(" | ")
	res := 0

	err := ReadLines("./days/inputs/day8_1.txt", func(b []byte) error {
		split := bytes.Split(b, pipe)
		s := string(split[1])
		r := regexp.MustCompile(`\b(\w{2}|\w{4}|\w{3}|\w{7})\b`).FindAllStringIndex(s, -1)
		res += len(r)
		// log.Printf("at line %s result = %d", s, len(r))

		return nil
	})
	return res, err
}

// func day8_2() (int, error) {
// 	var space = []byte(" ")
// 	var hor int = 0
// 	var ver int = 0
// 	var aim int = 0

// 	err := ReadLines("./days/inputs/day8_1.txt", func(b []byte) error {
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
// 	log.Printf("day8_2 result = %d", res)
// 	return res, err
// }
