package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day11() {
	if res, err := day11_1(); err != nil {
		log.Fatalf("crash in day11_1: %v", err)
	} else {
		log.Printf("day11_2 result = %d", res)
	}
	if res, err := day11_2(); err != nil {
		log.Fatalf("crash in day11_2: %v", err)
	} else {
		log.Printf("day11_2 result = %d", res)
	}
}

type octopus struct {
	charge int
	glowed bool
}

func printOctopus(o map[string]*octopus) {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Print(o[fmt.Sprintf("%d:%d", x, y)].charge)
		}
		fmt.Print("\n")
	}
}

func day11_1() (int, error) {
	octopuses := make(map[string]*octopus)
	nr := 0
	err := ReadLines("./days/inputs/day11_1.txt", func(b []byte) error {
		for i, v := range b {
			octopuses[fmt.Sprintf("%d:%d", nr, i)] = &octopus{charge: byteToInt(v)}
		}
		nr++
		return nil
	})
	log.Print("octopuses start")
	printOctopus(octopuses)
	res := 0
	for i := 0; i < 100; i++ {
		res += octoBlink(octopuses)
	}
	log.Print("octopuses end")
	printOctopus(octopuses)
	return res, err
}

func day11_2() (int, error) {
	octopuses := make(map[string]*octopus)
	nr := 0
	err := ReadLines("./days/inputs/day11_1.txt", func(b []byte) error {
		for i, v := range b {
			octopuses[fmt.Sprintf("%d:%d", nr, i)] = &octopus{charge: byteToInt(v)}
		}
		nr++
		return nil
	})
	log.Print("octopuses start")
	printOctopus(octopuses)
	res := 0
	for {
		blinks := octoBlink(octopuses)
		res++
		if blinks == 100 {
			break
		}
	}
	log.Print("octopuses end")
	printOctopus(octopuses)
	return res, err
}

func octoBlink(octopuses map[string]*octopus) (blinked int) {
	// increase energy of all octopuses
	octoCharge(octopuses, true)
	// blink those octopuses who are charged
	for _, r := range octopuses {
		if r.charge > 9 {
			blinked++
			r.charge = 0
			r.glowed = false
		}
	}

	return
}

func octoCharge(octopuses map[string]*octopus, init bool) {
	charged := false
	if init {
		for _, v := range octopuses {
			v.charge++
		}
	}
	for k, v := range octopuses {
		spl := strings.Split(k, ":")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])
		if v.charge >= 10 && !v.glowed {
			charged = true
			v.glowed = true
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x-1, y-1)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x-1, y)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x-1, y+1)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x, y-1)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x, y)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x, y+1)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x+1, y-1)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x+1, y)]; ok {
				c.charge++
			}
			if c, ok := octopuses[fmt.Sprintf("%d:%d", x+1, y+1)]; ok {
				c.charge++
			}
		}
	}
	if !charged {
		return
	} else {
		octoCharge(octopuses, false)
	}

}
