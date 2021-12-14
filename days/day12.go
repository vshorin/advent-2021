package days

import (
	"bytes"
	"log"
)

func Day12() {
	if res, err := day12_1(); err != nil {
		log.Fatalf("crash in day12_1: %v", err)
	} else {
		log.Printf("day12_2 result = %d", res)
	}
	// if res, err := day12_2(); err != nil {
	// 	log.Fatalf("crash in day12_2: %v", err)
	// } else {
	// 	log.Printf("day12_2 result = %d", res)
	// }
}

type cave struct {
	name     string
	parents  []*cave
	children []*cave
}

var startCave = cave{
	name:     "start",
	children: []*cave{},
}
var endCave = cave{
	name:    "end",
	parents: []*cave{},
}

func day12_1() (int, error) {
	dash := []byte{byte('-')}

	var caves []*cave
	err := ReadLines("./days/inputs/day12_x.txt", func(b []byte) error {
		spl := bytes.Split(b, dash)
		str0 := string(spl[0])
		str1 := string(spl[1])
		newCave0 := addCave(caves, str0, nil)
		newCave1 := addCave(caves, str1, nil)
		caveRelationship(newCave0, newCave1)
		return nil
	})
	log.Printf("caves: %v", caves)

	return 0, err
}

func addCave(caves []*cave, name string, existing *cave) *cave {
	for _, c := range caves {
		if c.name == name {
			return c
		}
	}
	if existing != nil {
		caves = append(caves, existing)
		return existing
	}
	var newCave *cave
	if name == "start" {
		newCave = &startCave
	} else if name == "end" {
		newCave = &endCave
	} else {
		newCave = &cave{
			name:     name,
			parents:  make([]*cave, 0),
			children: make([]*cave, 0),
		}
	}
	caves = append(caves, newCave)
	return newCave
}

func caveRelationship(thisCave *cave, otherCave *cave) {
	// only start and capital letter cave can have child caves
	if thisCave.name == "start" || byte(thisCave.name[0]) < 97 {
		addCave(thisCave.children, otherCave.name, otherCave)
	}
	if otherCave.name == "start" || byte(otherCave.name[0]) < 97 {
		addCave(otherCave.children, thisCave.name, thisCave)
	}
	// all caves except start can have parent caves
	if thisCave.name != "start" {
		addCave(thisCave.parents, otherCave.name, otherCave)
	}
	if otherCave.name != "start" {
		addCave(otherCave.parents, thisCave.name, thisCave)
	}
}

// func day12_2() (int, error) {
// 	octopuses := make(map[string]*octopus)
// 	nr := 0
// 	err := ReadLines("./days/inputs/day12_1.txt", func(b []byte) error {
// 		for i, v := range b {
// 			octopuses[fmt.Sprintf("%d:%d", nr, i)] = &octopus{charge: byteToInt(v)}
// 		}
// 		nr++
// 		return nil
// 	})
// 	log.Print("octopuses start")
// 	printOctopus(octopuses)
// 	res := 0
// 	for {
// 		blinks := octoBlink(octopuses)
// 		res++
// 		if blinks == 100 {
// 			break
// 		}
// 	}
// 	log.Print("octopuses end")
// 	printOctopus(octopuses)
// 	return res, err
// }
