package days

import (
	"log"
	"sort"
)

func Day10() {
	if res, err := day10_1(); err != nil {
		log.Fatalf("crash in day10_1: %v", err)
	} else {
		log.Printf("day10_2 result = %d", res)
	}
	if res, err := day10_2(); err != nil {
		log.Fatalf("crash in day10_2: %v", err)
	} else {
		log.Printf("day10_2 result = %d", res)
	}
}

var brackets map[byte]byte = map[byte]byte{
	byte('('): byte(')'),
	byte('['): byte(']'),
	byte('{'): byte('}'),
	byte('<'): byte('>'),
}

func day10_1() (int, error) {
	var badBrackets []byte
	badValues := map[byte]int{
		byte(')'): 3,
		byte(']'): 57,
		byte('}'): 1197,
		byte('>'): 25137, // this one really rare
	}
	err := ReadLines("./days/inputs/day10_1.txt", func(b []byte) error {
		var closes []byte
		for i, v := range b {
			if brack, ok := brackets[v]; ok {
				closes = append(closes, brack)
			} else {
				if closes[len(closes)-1] != v {
					badBrackets = append(badBrackets, v)
					log.Printf("found bad bracket %s at %d in %s", string(v), i, string(b))
					break
				} else {
					closes = closes[:len(closes)-1]
				}
			}
		}
		return nil
	})
	res := 0
	for _, v := range badBrackets {
		res += badValues[v]
	}
	return res, err
}

func day10_2() (int, error) {
	goodValues := map[byte]int{
		byte(')'): 1,
		byte(']'): 2,
		byte('}'): 3,
		byte('>'): 4, // this one really rare
	}
	var ress []int
	err := ReadLines("./days/inputs/day10_1.txt", func(b []byte) error {
		var closes []byte
		for _, v := range b {
			if brack, ok := brackets[v]; ok {
				closes = append(closes, brack)
			} else {
				if closes[len(closes)-1] != v {
					// corrupted line ignored
					return nil
				} else {
					closes = closes[:len(closes)-1]
				}
			}
		}
		r := 0
		for i := len(closes) - 1; i >= 0; i-- {
			r = r*5 + goodValues[closes[i]]
		}
		ress = append(ress, r)
		return nil
	})
	sort.Ints(ress)
	res := ress[(len(ress)-1)/2]
	return res, err
}
