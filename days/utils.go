package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func ReadLines(fn string, cb func([]byte) error) error {
	fmt.Printf("reading file %s...\n", fn)

	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	for {

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're at the EOF, break.
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
		}
		if err == io.EOF {
			fmt.Println("done reading lines")
			break
		}
		err = cb(l)
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	return nil
}

func stringSliceContains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func boolSliceContains(slice []bool, val bool) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func byteSliceContains(slice []byte, val byte) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func maxInt(ints []int) int {
	sort.Ints(ints)
	return ints[len(ints)-1]
}

func minInt(ints []int) int {
	sort.Ints(ints)
	return ints[0]
}

func byteToInt(b byte) int {
	n, e := strconv.Atoi(string(b))
	if e != nil {
		fmt.Printf("can't parse %d as number\n", b)
		panic(e)
	}
	return n
}

func byteArrToInt(b []byte) int {
	n, e := strconv.Atoi(string(b))
	if e != nil {
		fmt.Printf("can't parse %v as number\n", b)
		panic(e)
	}
	return n
}

func stringToInt(s string) int {
	n, e := strconv.Atoi(s)
	if e != nil {
		fmt.Printf("can't parse %v as number\n", s)
		panic(e)
	}
	return n
}
