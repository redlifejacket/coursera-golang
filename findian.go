/*
 * Author: Shyam Govardhan (10 March 2019)
 * Coursera Getting Started with Go (UCI)
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type UserPrompt int

const (
	MaxUserPromptLength int    = 60
	QuitCode            string = "quit"
)

const (
	ReadStringValue UserPrompt = iota
	ShowStringValue
	IanFound
	IanNotFound
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
		strReturnVal = GetUserPromptPadded(up, "Please enter a string value (Type 'quit' to Quit):", "")
	case IanFound:
		strReturnVal = GetUserPromptPadded(up, "Found!", "")
	case IanNotFound:
		strReturnVal = GetUserPromptPadded(up, "Not Found!", "")
	}
	return strReturnVal
}

func main() {
	re := regexp.MustCompile(`^[i|I]+.*[a|A]+.*[n|N]+.*`)
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf(GetUserPrompt(ReadStringValue))
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strings.ToLower(strVal) == QuitCode {
			os.Exit(0)
		}
		strValLower := strings.ToLower(strVal)
		match := re.Match([]byte(strValLower))
		if match == true {
			println(GetUserPrompt(IanFound))
		} else {
			println(GetUserPrompt(IanNotFound))
		}
		fmt.Printf(GetUserPrompt(ReadStringValue))
	}
}
