/*
 * Author: Shyam Govardhan (24 March 2019)
 * Coursera Getting Started with Go (UCI)
 * Write a program which prompts the user to first enter a name, and then enter an address.
 * Your program should create a map and add the name and address to the map using the keys
 * “name” and “address”, respectively. Your program should use Marshal() to create a JSONs
 * object from the map, and then your program should print the JSON object.
 */
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type UserPrompt int

const (
	MaxUserPromptLength = 40
	QuitCode            = "X"
	MapKeyName          = "name"
	MapKeyAddress       = "address"
)

const (
	PromptReadName UserPrompt = iota
	PromptReadAddress
	PromptInvalidName
	PromptInvalidAddress
	PromptSampleAddress
	PromptExit
)

func GetUserPromptPadded(up UserPrompt, str string, strFmt string, newline bool) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	if newline {
		paddedStr += "\n"
	}
	return paddedStr
}

func GetUserPrompt(up UserPrompt) string {
	switch up {
	case PromptReadName:
		return GetUserPromptPadded(up, "Please enter name:", "", false)
	case PromptReadAddress:
		return GetUserPromptPadded(up, "Please enter address:", "", false)
	case PromptInvalidName:
		return GetUserPromptPadded(up, "Invalid Name! Please enter a valid name.", "", true)
	case PromptInvalidAddress:
		msg := "Invalid Address! Please enter address."
		return GetUserPromptPadded(up, msg, "", true)
	case PromptSampleAddress:
		msg := "Eg: 1600 Amphitheatre Pkwy MTV CA 94043"
		return GetUserPromptPadded(up, msg, "", true)
	case PromptExit:
		return GetUserPromptPadded(up, "Enter 'X' to Exit", "", true)
	}
	return ""
}

func readValidInput(up UserPrompt, str string, regex *regexp.Regexp) (string, bool) {
	if str == QuitCode {
		os.Exit(0)
	}
	match := regex.Match([]byte(str))
	if !match {
		return "", false
	}
	return str, true
}

func readUntilValid(readPrompt UserPrompt, invalidPrompt UserPrompt, regex *regexp.Regexp) string {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf(GetUserPrompt(readPrompt))
	for scanner.Scan() {
		value, isValid := readValidInput(readPrompt, scanner.Text(), regex)
		if value != "" && isValid {
			return value
		}
		fmt.Printf(GetUserPrompt(invalidPrompt))
		if invalidPrompt == PromptInvalidAddress {
			fmt.Printf(GetUserPrompt(PromptSampleAddress))
		}
		fmt.Printf(GetUserPrompt(readPrompt))
	}
	return ""
}

func main() {
	fmt.Printf(GetUserPrompt(PromptExit))
	// RegexForName from: https://www.regextester.com/93648
	RegexForName := regexp.MustCompile(`^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$`)
	RegexForAddress := regexp.MustCompile(`^\d{1,5}\s+\w.*$`)

	mapNameAddress := make(map[string]string)
	mapNameAddress[MapKeyName] = readUntilValid(PromptReadName, PromptInvalidName, RegexForName)
	mapNameAddress[MapKeyAddress] = readUntilValid(PromptReadAddress, PromptInvalidAddress, RegexForAddress)
	json, err := json.Marshal(mapNameAddress)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	fmt.Printf("Name:    %s\n", mapNameAddress[MapKeyName])
	fmt.Printf("Address: %s\n", mapNameAddress[MapKeyAddress])
	fmt.Printf("JSON:    %s\n", json)
}
