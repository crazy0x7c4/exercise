package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sorter/algorithm/bubblesort"
	"strconv"
)

var (
	inputFile  *string = flag.String("i", "inputFIle", "file contains values for sorting")
	outputFile *string = flag.String("o", "outputFile", "file to receive sorted values")
	algorithm  *string = flag.String("a", "qsort", "algorithm")
)

func readValues(inputFile string) (values []int, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Faild to open the input file!", inputFile)
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		value, err1 := strconv.Atoi(string(line))
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Faild to create the output file.", outputFile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		file.WriteString(strconv.Itoa(value) + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if inputFile != nil {
		fmt.Println("inputFile=", *inputFile, "outputFile=", *outputFile, "algorithm=", *algorithm)
	}

	values, readErr := readValues(*inputFile)
	if readErr != nil {
		fmt.Println(readErr)
	}

	if *algorithm == "bubblesort" {
		bubblesort.BubbleSort(values)
	}

	writeErr := writeValues(values, *outputFile)
	if writeErr != nil {
		fmt.Println(writeErr)
	}
}
