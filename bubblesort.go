/*
 * Author: Shyam Govardhan (22 April 2019)
 * Coursera Getting Started with Go (UCI)
 * Write a Bubble Sort program in Go.
 * The program should prompt the user to type in a sequence of up to 10 integers.
 * The program should print the integers out on one line, in sorted order, from
 * least to greatest.
 * Use your favorite search tool to find a description of how the bubble sort
 * algorithm works.
 * As part of this program, you should write a function called BubbleSort() which
 * takes a slice of integers as an argument and returns nothing. The BubbleSort()
 * function should modify the slice so that the elements are in sorted order.
 * A recurring operation in the bubble sort algorithm is the Swap operation which swaps
 * the position of two adjacent elements in the slice. You should write a Swap()
 * function which performs this operation. Your Swap() function should take two
 * arguments, a slice of integers and an index value i which indicates a position
 * in the slice.
 * The Swap() function should return nothing, but it should swap the contents of
 * the slice in position i with the contents in position i+1.
 *
 * Based on an earlier assignment.
 * https://github.com/redlifejacket/coursera-golang/blob/master/slice.go
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type UserPrompt int

const (
	ListSize            int    = 10
	MaxUserPromptLength int    = 40
	QuitCode            string = "X"
)

const (
	ReadStringValue UserPrompt = iota
	ShowStringValue
	RegexFound
	RegexNotFound
)

func GetUserPromptPadded(up UserPrompt, str string, strFmt string) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	if !(up == ReadStringValue) {
		paddedStr = paddedStr + "\n"
	}
	return paddedStr
}

func GetUserPrompt(up UserPrompt) string {
	var strReturnVal string
	switch up {
	case ReadStringValue:
		strReturnVal = GetUserPromptPadded(up, "Please enter an integer ('X' to Exit):", "")
	case RegexFound:
		strReturnVal = GetUserPromptPadded(up, "Valid Input!", "")
	case RegexNotFound:
		strReturnVal = GetUserPromptPadded(up, "Invalid Input!", "")
	}
	return strReturnVal
}

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}

func BubbleSort(numbers []int) {
	Swapped := true
	for Swapped {
		Swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i+1] < numbers[i] {
				Swap(numbers, i)
				Swapped = true
			}
		}
	}
}

func main() {
	var intSlice = make([]int, ListSize)

	re := regexp.MustCompile(`^\d+$`)
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf(GetUserPrompt(ReadStringValue))
	intCount := 0
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == QuitCode || intCount == ListSize {
			os.Exit(0)
		}
		match := re.Match([]byte(strVal))
		intVal, err := strconv.Atoi(strVal)
		if match == false || err != nil {
			println(GetUserPrompt(RegexNotFound))
			fmt.Printf(GetUserPrompt(ReadStringValue))
			continue
		}
		intSlice = append(intSlice, intVal)
		BubbleSort(intSlice)
		var result []int = intSlice[ListSize:len(intSlice)]
		fmt.Println(result)
		fmt.Printf(GetUserPrompt(ReadStringValue))
		intCount += 1
	}
}
