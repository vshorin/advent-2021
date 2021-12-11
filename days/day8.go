package days

import (
	"bytes"
	"errors"
	"log"
	"regexp"
	"sort"
	"strconv"
)

func Day8() {
	if res, err := day8_1(); err != nil {
		log.Fatalf("crash in day8_1: %v", err)
	} else {
		log.Printf("day8_1 result = %d", res)
	}
	if res, err := day8_2(); err != nil {
		log.Fatalf("crash in day8_2: %v", err)
	} else {
		log.Printf("day8_2 result = %d", res)
	}
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

func day8_2() (int, error) {
	var pipe = []byte(" | ")
	res := 0

	err := ReadLines("./days/inputs/day8_1.txt", func(b []byte) error {
		split := bytes.Split(b, pipe)
		decoded, e := decodeDisplay(split[0])
		if e != nil {
			return e
		}
		numbs := bytes.Split(split[1], []byte(" "))
		var numb string
		for _, n := range numbs {
			sort.Slice(n, func(i, j int) bool {
				return n[i] < n[j]
			})
			// log.Printf("now n = %s and value = %s", n, decoded[string(n)])
			numb += decoded[string(n)]
		}
		n, e := strconv.Atoi(numb)
		// log.Printf("line number is %d", n)
		res += n

		return e
	})
	return res, err
}

func decodeDisplay(coded []byte) (decoded map[string]string, err error) {
	decoded = make(map[string]string)
	// known numbers
	spl := bytes.Split(coded, []byte(" "))
	var unknown [][]byte
	var one []byte
	var four []byte
	for _, s := range spl {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		if len(s) == 2 {
			decoded[string(s)] = "1"
			one = s
		} else if len(s) == 3 {
			decoded[string(s)] = "7"
		} else if len(s) == 4 {
			decoded[string(s)] = "4"
			four = s
		} else if len(s) == 7 {
			decoded[string(s)] = "8"
		} else {
			unknown = append(unknown, s)
		}
	}
	var fByte byte
	var stillUnknown [][]byte
	for _, u := range unknown {
		if len(u) == 5 {
			if byteSliceContains(u, one[0]) && byteSliceContains(u, one[1]) {
				decoded[string(u)] = "3"
				continue
			}
			stillUnknown = append(stillUnknown, u)
			continue
		}
		if len(u) == 6 {
			if byteSliceContains(u, four[0]) && byteSliceContains(u, four[1]) && byteSliceContains(u, four[2]) && byteSliceContains(u, four[3]) {
				decoded[string(u)] = "9"
			} else if byteSliceContains(u, one[0]) && byteSliceContains(u, one[1]) {
				decoded[string(u)] = "0"
			} else {
				decoded[string(u)] = "6"
				if byteSliceContains(u, one[0]) {
					fByte = one[0]
				} else {
					fByte = one[1]
				}
			}
			continue
		}
	}
	for _, u := range stillUnknown {
		if byteSliceContains(u, fByte) {
			decoded[string(u)] = "5"
		} else {
			decoded[string(u)] = "2"
		}
	}
	log.Printf("decoded %v", decoded)
	if len(decoded) != 10 {
		err = errors.New("not all numbers decoded!")
	}

	return
}
