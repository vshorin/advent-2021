package days

import (
	"bytes"
	"log"
	"strings"
)

func Day12() {
	if res, err := day12_1(); err != nil {
		log.Fatalf("crash in day12_1: %v", err)
	} else {
		log.Printf("day12_2 result = %d", res)
	}
	if res, err := day12_2(); err != nil {
		log.Fatalf("crash in day12_2: %v", err)
	} else {
		log.Printf("day12_2 result = %d", res)
	}
}

func readCaves() (caves map[string][]string, err error) {
	dash := []byte{byte('-')}
	caves = map[string][]string{
		"start": make([]string, 0),
	}

	err = ReadLines("./days/inputs/day12_1.txt", func(b []byte) error {
		spl := bytes.Split(b, dash)
		str0 := string(spl[0])
		str1 := string(spl[1])
		addCave(str0, caves)
		addCave(str1, caves)
		caveRelationship(str0, str1, caves)
		return nil
	})
	if err != nil {
		return
	}
	log.Print("caves:")
	for k, v := range caves {
		log.Printf("%s: relations = %v", k, v)
	}
	return
}

func day12_1() (int, error) {
	caveTravels = make([][]string, 0)

	caves, err := readCaves()
	if err != nil {
		return 0, err
	}
	caveTravel("start", caves, make([]string, 0))
	log.Print("travels:")
	for _, v := range caveTravels {
		log.Printf("\t%v", v)
	}

	return len(caveTravels), err
}

func day12_2() (int, error) {
	caveTravels = make([][]string, 0)

	caves, err := readCaves()
	if err != nil {
		return 0, err
	}
	startCave := travelTracker{
		travel: make([]string, 0),
		dupe:   false,
	}
	caveTravelTracked("start", caves, startCave)
	// log.Print("travels:")
	// for _, v := range caveTravels {
	// 	log.Printf("\t%v", strings.Join(v, ","))
	// }

	return len(caveTravels), err
}

// var knownTravels = make(map[string][][]string)
var caveTravels [][]string

func caveTravel(start string, caves map[string][]string, travel []string) {

	travel = append(travel, start)
	for _, v := range caves[start] {
		if v == "end" {
			caveTravels = append(caveTravels, append(travel, "end"))
			continue
		}
		if byte(v[0]) >= 97 {
			if start != "start" && travel[len(travel)-2] == v {
				continue
			}
			if stringSliceContains(travel, v) {
				continue
			}
		}
		newTravel := make([]string, len(travel))
		copy(newTravel, travel)
		caveTravel(v, caves, newTravel)
	}
}

type travelTracker struct {
	travel []string
	dupe   bool
}

func checkUniqueSmallCave(c string, travel *travelTracker) bool {
	count := 0
	for _, v := range travel.travel {
		if v == c {
			count++
		}
	}
	if count == 0 {
		return true
	}
	if travel.dupe {
		return false
	}
	if count < 2 {
		travel.dupe = true
		return true
	}
	return false
}

func caveTravelTracked(start string, caves map[string][]string, travel travelTracker) {
	// start,A,c,A,b,end
	// start,A,c,A,b,d,b,A,end
	travel.travel = append(travel.travel, start)
	if strings.Join(travel.travel, ",") == "start,A,c,A,b" {
		log.Println("aa")
	}
	for _, v := range caves[start] {
		if v == "end" {
			caveTravels = append(caveTravels, append(travel.travel, "end"))
			continue
		}
		newTravel := make([]string, len(travel.travel))
		copy(newTravel, travel.travel)
		tt := travelTracker{
			dupe:   travel.dupe,
			travel: newTravel,
		}
		if byte(v[0]) >= 97 {
			if start != "start" && !checkUniqueSmallCave(v, &tt) {
				continue
			}
			//if stringSliceContains(travel.travel, v) {
			//	continue
			//}
		}
		caveTravelTracked(v, caves, tt)
	}
}

func addCave(name string, caves map[string][]string) {
	if _, ok := caves[name]; !ok {
		caves[name] = make([]string, 0)
	}
}

func caveRelationship(thisCave string, otherCave string, caves map[string][]string) {
	if otherCave != "start" {
		caves[thisCave] = append(caves[thisCave], otherCave)
	}
	if thisCave != "start" {
		caves[otherCave] = append(caves[otherCave], thisCave)
	}
}
