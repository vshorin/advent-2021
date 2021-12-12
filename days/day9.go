package days

import (
	"fmt"
	"log"
	"sort"
)

func Day9() {
	if res, err := day9_1(); err != nil {
		log.Fatalf("crash in day9_1: %v", err)
	} else {
		log.Printf("day9_1 result = %d", res)
	}
	if res, err := day9_2(); err != nil {
		log.Fatalf("crash in day9_2: %v", err)
	} else {
		log.Printf("day9_2 result = %d", res)
	}
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

type region map[string]bool

func day9_2() (int, error) {
	var regions []region
	iter := 0
	nine := byte('9')
	err := ReadLines("./days/inputs/day9_1.txt", func(b []byte) error {
		//ints := make([]int, len(b))
		for i, v := range b {
			//ints[i] = byteToInt(v)
			if v != nine {
				foundIn := -1
				for j, r := range regions {
					// look for existing regions left and above. If found, add to them
					if r[fmt.Sprintf("%d:%d", i-1, iter)] || r[fmt.Sprintf("%d:%d", i, iter-1)] {
						r[fmt.Sprintf("%d:%d", i, iter)] = true
						// if this was already found, need to join regions
						if foundIn >= 0 {
							for k := range r {
								regions[foundIn][k] = true
							}
							foundIn = j
							// delete region once joined
							regions = append(regions[:foundIn], regions[foundIn+1:]...)
							break
						}
						foundIn = j
					}
				}
				if foundIn == -1 {
					// make new region if not found
					reg := make(region)
					reg[fmt.Sprintf("%d:%d", i, iter)] = true
					regions = append(regions, reg)
				}
			}
		}
		iter++
		return nil
	})
	sort.Slice(regions, func(i, j int) bool {
		return len(regions[i]) > len(regions[j])
	})
	log.Printf("regions = %v", regions)
	res := len(regions[0]) * len(regions[1]) * len(regions[2])
	return res, err
}
