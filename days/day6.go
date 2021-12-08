package days

import (
	"log"
	"strconv"
	"strings"
)

func Day6() {
	if res, err := day6_1(); err != nil {
		log.Fatalf("crash in day6_1: %v", err)
	} else {
		log.Printf("day6_1 result = %d", res)
	}
	if res, err := day6_2(); err != nil {
		log.Fatalf("crash in day6_2: %v", err)
	} else {
		log.Printf("day6_2 result = %d", res)
	}
}

func calculateFish(days int) (int, error) {
	fishes := make([]string, 0)

	err := ReadLines("./days/inputs/day6_1.txt", func(b []byte) error {
		fishes = strings.Split(string(b), ",")
		return nil
	})
	log.Printf("fishes amount = %d", len(fishes))

	fishTracker := make(map[int]int)
	for _, v := range fishes {
		var age int
		age, err = strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		fishTracker[age]++
	}
	log.Printf("fish tracker initially = %v", fishTracker)

	// days := 80 // 256
	ageNew := 8
	ageOld := 6
	for i := 0; i < days; i++ {
		fishTrackerNext := make(map[int]int)
		fishTrackerNext[ageNew] = fishTracker[0]
		for j := 0; j < ageNew; j++ {
			fishTrackerNext[j] = fishTracker[j+1]
			if j == ageOld {
				fishTrackerNext[j] += fishTracker[0]
			}
		}
		fishTracker = fishTrackerNext
	}

	log.Printf("fish tracker after %d days = %v", days, fishTracker)
	var res int
	for _, v := range fishTracker {
		res += v
	}

	return res, err
}

func day6_1() (int, error) {
	return calculateFish(80)
}

func day6_2() (int, error) {
	return calculateFish(256)
}
