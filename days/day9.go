package days

import (
	"fmt"
	"log"
)

func Day9() {
	if res, err := day9_1(); err != nil {
		log.Fatalf("crash in day9_1: %v", err)
	} else {
		log.Printf("day9_1 result = %d", res)
	}
	// if res, err := day9_2(); err != nil {
	// 	log.Fatalf("crash in day9_2: %v", err)
	// } else {
	// 	log.Printf("day9_2 result = %d", res)
	// }
}

func day9_1() (int, error) {
	var heights [][]int
	err := ReadLines("./days/inputs/day9_1.txt", func(b []byte) error {
		ints := make([]int, len(b))
		for i, v := range b {
			ints[i] = byteToInt(v)
		}
		heights = append(heights, ints)
		return nil
	})
	//log.Printf("heights are %v", heights)
	lpMap := make(map[string]int, 0)
	for i, h := range heights {
		for j, r := range h {
			var s []int
			if i > 0 {
				s = append(s, heights[i-1][j])
			}
			if j < len(h)-1 {
				s = append(s, heights[i][j+1])
			}
			if i < len(heights)-1 {
				s = append(s, heights[i+1][j])
			}
			if j > 0 {
				s = append(s, heights[i][j-1])
			}
			s = append(s, r)
			//log.Printf("at %d:%d slice is %v", i, j, s)
			if minInt(s) == r && s[len(s)-1] != r {
				lpMap[fmt.Sprintf("%d:%d", j, i)] = r
			}
		}
	}
	log.Printf("lowest points: %v", lpMap)
	res := 0
	for _, v := range lpMap {
		res += (v + 1)
	}
	return res, err
}

// func day9_2() (int, error) {
// 	var pipe = []byte(" | ")
// 	res := 0

// 	err := ReadLines("./days/inputs/day9_1.txt", func(b []byte) error {
// 		split := bytes.Split(b, pipe)
// 		decoded, e := decodeDisplay(split[0])
// 		if e != nil {
// 			return e
// 		}
// 		numbs := bytes.Split(split[1], []byte(" "))
// 		var numb string
// 		for _, n := range numbs {
// 			sort.Slice(n, func(i, j int) bool {
// 				return n[i] < n[j]
// 			})
// 			// log.Printf("now n = %s and value = %s", n, decoded[string(n)])
// 			numb += decoded[string(n)]
// 		}
// 		n, e := strconv.Atoi(numb)
// 		// log.Printf("line number is %d", n)
// 		res += n

// 		return e
// 	})
// 	return res, err
// }
