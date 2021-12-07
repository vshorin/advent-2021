package days

import (
	"log"
	"regexp"
	"strings"
)

func Day4() {
	if _, err := day4_1(); err != nil {
		log.Fatalf("crash in day4_1: %v", err)
	}
	// if _, err := day4_2(); err != nil {
	// 	log.Fatalf("crash in day4_2: %v", err)
	// }
}

func day4_1() (int, error) {
	var numbers []string
	var board []string

	err := ReadLines("./days/inputs/day4_1.txt", func(b []byte) error {
		line := string(b)
		if len(numbers) == 0 {
			numbers = strings.Split(line, ",")
			log.Printf("bingo numbers: %v", numbers)
			return nil
		}
		if len(b) == 0 {
			log.Print("next board...")
		} else {
			splt, _ := regexp.Split(line)
			board = append(board, regexp.Split(line))
		}

		return nil
	})
	return 0, err
}

func day4_2() (int, error) {

	err := ReadLines("./days/inputs/day4_1.txt", func(b []byte) error {
		return nil
	})
	return 0, err
}
