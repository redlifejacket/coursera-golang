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
	"strconv"
	"strings"
)

type UserPrompt int

const (
	MaxUserPromptLength int    = 60
	QuitCode            string = "quit"
)

const (
	ReadFloatValue UserPrompt = iota
	ShowStringValue
	ShowFloatValue
	ShowIntValue
	ShowErrorConversion
)

func GetUserPromptPadded(up UserPrompt, str string, strFmt string) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	if !(up == ReadFloatValue) {
		paddedStr = paddedStr + "\n"
	}
	return paddedStr
}

func GetUserPrompt(up UserPrompt) string {
	var strReturnVal string
	switch up {
	case ReadFloatValue:
		strReturnVal = GetUserPromptPadded(up, "Please enter a float value (Type 'quit' to Quit):", "")
	case ShowStringValue:
		strReturnVal = GetUserPromptPadded(up, "String Value:", "%s")
	case ShowFloatValue:
		strReturnVal = GetUserPromptPadded(up, "Float Value:", "%f")
	case ShowIntValue:
		strReturnVal = GetUserPromptPadded(up, "Truncated Integer:", "%d")
	case ShowErrorConversion:
		strReturnVal = "Could not convert [%s] to float!"
	}
	return strReturnVal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf(GetUserPrompt(ReadFloatValue))
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strings.ToLower(strVal) == QuitCode {
			os.Exit(0)
		}
		//fmt.Printf(GetUserPrompt(ShowStringValue), strVal)
		var floatVal, err = strconv.ParseFloat(strVal, 32)
		if err != nil {
			log.Printf(GetUserPrompt(ShowErrorConversion), strVal)
			fmt.Printf(GetUserPrompt(ReadFloatValue))
			continue
		}
		fmt.Printf(GetUserPrompt(ShowFloatValue), floatVal)
		var intVal int32 = int32(floatVal)
		fmt.Printf(GetUserPrompt(ShowIntValue), intVal)
		fmt.Printf(GetUserPrompt(ReadFloatValue))
	}
}
