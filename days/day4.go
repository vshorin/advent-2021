package days

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Day4() {
	// if res, err := day4_1(); err != nil {
	// 	log.Fatalf("crash in day4_1: %v", err)
	// } else {
	// 	log.Printf("day4_1 result = %d", res)
	// }
	if res, err := day4_2(); err != nil {
		log.Fatalf("crash in day4_2: %v", err)
	} else {
		log.Printf("day4_2 result = %d", res)
	}
}

func fillBingoData() (numbers []string, boards [][][]string, err error) {
	var board [][]string
	err = ReadLines("./days/inputs/day4_1.txt", func(b []byte) error {
		line := string(b)
		if len(numbers) == 0 {
			numbers = strings.Split(line, ",")
			log.Printf("bingo numbers: %v", numbers)
			return nil
		}
		if len(b) == 0 {
			log.Printf("board = %v", board)
			if len(board) > 0 {
				boards = append(boards, board)
				board = make([][]string, 0)
			}
		} else {
			splt := regexp.MustCompile(`\s+`).Split(line, -1)
			if splt[0] == "" {
				splt = splt[1:]
			}
			board = append(board, splt)
		}

		return nil
	})
	return
}

func day4_1() (int, error) {
	numbers, boards, err := fillBingoData()
	if err != nil {
		return 0, err
	}

	var bingoManager [][][]bool = make([][][]bool, len(boards))
	for _, v := range numbers {
		for b, n := range boards {
			if len(bingoManager[b]) == 0 {
				bingoManager[b] = make([][]bool, len(n))
			}
			for bb, nn := range n {
				if len(bingoManager[b][bb]) == 0 {
					bingoManager[b][bb] = make([]bool, len(nn))
				}
				for bbb, nnn := range nn {
					if nnn == v {
						bingoManager[b][bb][bbb] = true
					}
				}
			}
			if checkBingo(bingoManager[b]) {
				log.Printf("bingo at number %s on board %v", v, n)
				return makeBingoResult(bingoManager[b], n, v)
			}
		}
	}
	return 0, err
}

func checkBingo(checker [][]bool) bool {
	for _, v := range checker {
		foundHorizontal := true
		for _, vv := range v {
			if !vv {
				foundHorizontal = false
				break
			}
		}
		if foundHorizontal {
			return true
		}
	}
	for i := range checker[0] {
		foundVertical := true
		for ii := range checker {
			if !checker[ii][i] {
				foundVertical = false
				break
			}
		}
		if foundVertical {
			return true
		}
	}
	return false
}

func makeBingoResult(checker [][]bool, board [][]string, winning string) (res int, err error) {
	var n, w int
	for i, v := range checker {
		for ii, vv := range v {
			if !vv {
				n, err = strconv.Atoi(board[i][ii])
				if err != nil {
					return
				}
				res += n
			}
		}
	}
	w, err = strconv.Atoi(winning)
	if err != nil {
		return
	}
	return res * w, nil
}

func day4_2() (int, error) {
	numbers, boards, err := fillBingoData()
	if err != nil {
		return 0, err
	}
	slowBingoTracker := make([]bool, len(boards))

	var bingoManager [][][]bool = make([][][]bool, len(boards))
	for _, v := range numbers {
		for b, n := range boards {
			if len(bingoManager[b]) == 0 {
				bingoManager[b] = make([][]bool, len(n))
			}
			for bb, nn := range n {
				if len(bingoManager[b][bb]) == 0 {
					bingoManager[b][bb] = make([]bool, len(nn))
				}
				for bbb, nnn := range nn {
					if nnn == v {
						bingoManager[b][bb][bbb] = true
					}
				}
			}
			if checkBingo(bingoManager[b]) {
				log.Printf("bingo at number %s on board %v", v, n)
				slowBingoTracker[b] = true
				if !boolSliceContains(slowBingoTracker, false) {
					return makeBingoResult(bingoManager[b], n, v)
				}
			}
		}
	}

	return 0, err
}
