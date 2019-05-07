/*
 * Author: Shyam Govardhan (7 May 2019)
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

const (
	ListSize            int    = 10
	QuitCode            string = "X"
)

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
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				Swapped = true
			}
		}
	}
}

// https://stackoverflow.com/questions/37290693/how-to-remove-redundant-spaces-whitespace-from-a-string-in-golang
func stripSpaces(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := re_leadclose_whtsp.ReplaceAllString(input, "")
	result = re_inside_whtsp.ReplaceAllString(result, " ")
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println("Enter X to exit...")
	fmt.Printf("Please enter a list of integers (space-separated): ")

	for (scanner.Scan()) {
		var strVal string = scanner.Text()
		if strVal == QuitCode {
			os.Exit(0)
		}
		strVal = stripSpaces(strVal)
		var numlist = strings.Split(strVal, " ")
		var intSlice = make([]int, 0)
		for _, i := range numlist {
			j, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("Invalid input! Please try again...")
				fmt.Printf("Please enter a list of integers (space-separated): ")
				continue
			}
			intSlice = append(intSlice, j)
		}
		fmt.Printf("len(intSlice): %d; cap(intSlice): %d\n", len(intSlice), cap(intSlice))
		if (len(intSlice) > ListSize) {
			fmt.Printf("Please enter a maximum of %d integers\n", ListSize)
			fmt.Printf("Please enter a list of integers (space-separated): ")
			continue
		}
		BubbleSort(intSlice)
		fmt.Println(intSlice)
		fmt.Printf("Please enter a list of integers (space-separated): ")
	}
}
