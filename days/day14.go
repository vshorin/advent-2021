package days

import (
	"bytes"
	"log"
	"sort"
)

func Day14() {
	if res, err := day14_1(); err != nil {
		log.Fatalf("crash in day14_1: %v", err)
	} else {
		log.Printf("day14_2 result = %d", res)
	}
	if res, err := day14_2(); err != nil {
		log.Fatalf("crash in day14_2: %v", err)
	} else {
		log.Printf("day14_2 result = %d", res)
	}
}

func calcPolymers(steps int) (int, error) {
	var line string
	var rules = make(map[string]string)
	var polymerScore = make(map[string]int)
	readRules := false
	arrow := []byte(" -> ")
	err := ReadLines("./days/inputs/day14_1.txt", func(b []byte) error {
		if len(b) == 0 {
			return nil
		}
		if readRules {
			spl := bytes.Split(b, arrow)
			rules[string(spl[0])] = string(spl[1])
		} else {
			line = string(b)
			readRules = true
		}
		return nil
	})
	// line = line[0:2]
	log.Printf("line = %s", line)
	log.Printf("rules = %v", rules)
	polyRules := calcPolymerRules(rules)
	countInitialPolymers(line, polymerScore)
	log.Printf("poly rules = %v", polyRules)
	pairTracker := make(map[string]int)
	for i := 2; i <= len(line); i++ {
		pairTracker[line[i-2:i]]++
	}
	log.Printf("pair tracker = %v", pairTracker)

	for i := 0; i < steps; i++ {
		pairTracker = polymerStep(pairTracker, polymerScore, polyRules, rules)
		log.Printf("%d pair tracker = %v", i, pairTracker)
	}
	log.Printf("polymer score = %v", polymerScore)

	arr := make([]int, 0)
	for _, v := range polymerScore {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	log.Printf("arr = %v", arr)
	return arr[len(arr)-1] - arr[0], err

	// this is brute force approach that doesn't work for large amount of steps. Need to track pairs instead
	// // line = "NB"
	// // countInitialPolymers(line)
	// for i := 0; i < steps; i++ {
	// 	line = updatePolymer(line, rules)
	// 	if steps < 6 {
	// 		log.Printf("line = %s", line)
	// 	}
	// 	log.Printf("%d line len = %d\t\tscore = %v", i, len(line), polymerScore)
	// }
	// log.Printf("polymer score = %v", polymerScore)
	// arr := make([]int, 0)
	// for _, v := range polymerScore {
	// 	arr = append(arr, v)
	// }
	// sort.Ints(arr)
	// log.Printf("arr = %v", arr)
	// return arr[len(arr)-1] - arr[0], err
}

func polymerStep(pairTracker, polymerScore map[string]int, polyRules map[string]map[string]int, rules map[string]string) (tracker map[string]int) {
	tracker = make(map[string]int)
	for k, v := range pairTracker {
		// if _, ok := tracker[k]; !ok {
		// 	tracker[k] = v
		// }
		if v == 0 {
			continue
		}
		rule, ok := polyRules[k]
		if !ok {
			// unknown pair - keep the same value
			tracker[k] = v
			continue
		}
		//log.Printf("rule for %s is %v", k, rule)
		for ruleK, ruleV := range rule {
			//log.Printf("%s: updating %s from %d to %d (%d)", k, ruleK, v, v*ruleV, tracker[ruleK]+v*ruleV)
			tracker[ruleK] = tracker[ruleK] + v*ruleV
		}
		polymerScore[rules[k]] += v
	}
	return
}

/*
BBBB		BB = 3
BNBNBNB		BB = 0	BN = 3	NB = 3


2021/12/18 08:01:27 0 pair tracker = map[BC:1 CB:0 CH:1 CN:1 HB:1 NB:1 NC:1 NN:0]
NCNBCHB


2021/12/18 08:13:56 1 pair tracker = map[BB:2 BC:2 BH:1 CB:2 CC:1 CH:0 CN:1 HB:0 HC:1 NB:2 NC:0]
NBCCNBBBCBHCB

*/

func day14_1() (int, error) {
	return calcPolymers(10)
}
func day14_2() (int, error) {
	return calcPolymers(40)
}

func countInitialPolymers(line string, polymerScore map[string]int) {
	for _, v := range line {
		polymerScore[string(v)]++
	}
}

func calcPolymerRules(rules map[string]string) (calcs map[string]map[string]int) {
	calcs = make(map[string]map[string]int)
	for k, v := range rules {
		if _, ok := calcs[k]; !ok {
			calcs[k] = make(map[string]int)
		}
		calcs[k][k] = 0
		calcs[k][k[0:1]+v]++
		calcs[k][v+k[1:2]]++
	}
	return
}

/*
NB
B = B*2 + (i + 1)%2
N = N*2 + i%2


reading file ./days/inputs/day14_x.txt...
done reading lines
2021/12/17 21:53:05 line = NNCB
2021/12/17 21:53:05 rules = map[BB:N BC:B BH:H BN:B CB:H CC:N CH:B CN:C HB:C HC:B HH:N HN:C NB:B NC:B NH:C NN:C]
2021/12/17 21:53:05 line = NCNBCHB
2021/12/17 21:53:05 0 line len = 7              score = map[B:2 C:2 H:1 N:2]
2021/12/17 21:53:05 line = NBCCNBBBCBHCB
2021/12/17 21:53:05 1 line len = 13             score = map[B:6 C:4 H:1 N:2]
2021/12/17 21:53:05 line = NBBBCNCCNBBNBNBBCHBHHBCHB
2021/12/17 21:53:05 2 line len = 25             score = map[B:11 C:5 H:4 N:5]
2021/12/17 21:53:05 line = NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
2021/12/17 21:53:05 3 line len = 49             score = map[B:23 C:10 H:5 N:11]
2021/12/17 21:53:05 line = NBBNBBNBBBNBBNBBCNCCNBBBCCNBCNCCNBBNBBNBBNBBNBBNBNBBNBBNBBNBBNBBCHBHHBCHBHHNHCNCHBCHBNBBCHBHHBCHB
2021/12/17 21:53:05 4 line len = 97             score = map[B:46 C:15 H:13 N:23]
2021/12/17 21:53:05 polymer score = map[B:46 C:15 H:13 N:23]
2021/12/17 21:53:05 arr = [13 15 23 46]
2021/12/17 21:53:05 day14_2 result = 33




2021/12/17 21:40:14 1 line len = 3              score = map[B:0			C:1 		N:2]
2021/12/17 21:40:14 2 line len = 5              score = map[B:1 		C:2 		N:2]
2021/12/17 21:40:14 3 line len = 9              score = map[B:3 		C:3 		N:3]
2021/12/17 21:40:14 4 line len = 17             score = map[B:6 		C:5 		N:6]
2021/12/17 21:40:14 5 line len = 33             score = map[B:15 		C:8 		N:10]
2021/12/17 21:40:14 6 line len = 65             score = map[B:31 		C:13 		N:21]
2021/12/17 21:40:14 7 line len = 129            score = map[B:68 		C:21 		N:40]
2021/12/17 21:40:14 8 line len = 257            score = map[B:141 		C:34 		N:82]
2021/12/17 21:40:14 9 line len = 513            score = map[B:295 		C:55 		N:163]
2021/12/17 21:40:14 10 line len = 1025          score = map[B:606 		C:89 		N:330]
2021/12/17 21:40:14 11 line len = 2049          score = map[B:1243 		C:144 		N:662]
2021/12/17 21:40:14 12 line len = 4097          score = map[B:2531 		C:233 		N:1333]
2021/12/17 21:40:14 13 line len = 8193          score = map[B:5140 		C:377 		N:2676]
2021/12/17 21:40:14 14 line len = 16385         score = map[B:10401 	C:610 		N:5374]
2021/12/17 21:40:14 15 line len = 32769         score = map[B:21003 	C:987 		N:10779]
2021/12/17 21:40:14 16 line len = 65537         score = map[B:42326 	C:1597 		N:21614]
2021/12/17 21:40:14 17 line len = 131073        score = map[B:85175 	C:2584 		N:43314]
2021/12/17 21:40:14 18 line len = 262145        score = map[B:171191 	C:4181 		N:86773]
2021/12/17 21:40:14 19 line len = 524289        score = map[B:343748 	C:6765 		N:173776]
2021/12/17 21:40:14 20 line len = 1048577       score = map[B:689701 	C:10946 	N:347930]
2021/12/17 21:40:15 21 line len = 2097153       score = map[B:1382975 	C:17711 	N:696467]
2021/12/17 21:40:15 22 line len = 4194305       score = map[B:2771726 	C:28657 	N:1393922]
2021/12/17 21:40:17 23 line len = 8388609       score = map[B:5552803 	C:46368 	N:2789438]

B = B * 2 + i

*/

// func updatePolymer(line string, rules map[string]string) string {
// 	// this is wrong approach because string grows almost exponentially
// 	indexes := make(map[int]string)
// 	for k, v := range rules {
// 		if k == line {
// 			indexes[1] = v
// 		}
// 		finds := findIndicesInString(line, k)
// 		// log.Printf("finds for %s:%s are %v", k, v, finds)
// 		for _, f := range finds {
// 			i := f + 1
// 			indexes[i] = v
// 		}
// 	}
// 	var sb strings.Builder
// 	for k, v := range line {
// 		if index, ok := indexes[k]; ok {
// 			sb.WriteString(index)
// 			polymerScore[index]++
// 		}
// 		sb.WriteRune(v)
// 	}
// 	return sb.String()
// }

// func day14_2() (int, error) {
// 	octopuses := make(map[string]*octopus)
// 	nr := 0
// 	err := ReadLines("./days/inputs/day14_1.txt", func(b []byte) error {
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
