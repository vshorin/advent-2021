package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
