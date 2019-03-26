/*
 * Author: Shyam Govardhan (24 March 2019)
 * Coursera Getting Started with Go (UCI)
 * Write a program which reads information from a file and represents it in a slice of structs.
 * Assume that there is a text file which contains a series of names. Each line of the text
 * file has a first name and a last name, in that order, separated by a single space on the line.
 * Your program will define a name struct which has two fields, fname for the first name, and
 * lname for the last name. Each field will be a string of size 20 (characters).
 * Your program should prompt the user for the name of the text file. Your program will
 * successively read each line of the text file and create a struct which contains the first and
 * last names found in the file. Each struct created will be added to a slice, and after all
 * lines have been read from the file, your program will have a slice containing one struct for
 * each line in the file. After reading all lines from the file, your program should iterate
 * through your slice of structs and print the first and last names found in each struct.
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type UserPrompt int
type Person struct {
	fname string
	lname string
}

const (
	MaxUserPromptLength = 45
	QuitCode            = "X"
	MapKeyName          = "name"
	MapKeyAddress       = "address"
	FieldSize           = 20
	BufferSize          = 42
	ListSize            = 10
)

const (
	PromptReadFilePath UserPrompt = iota
	PromptInvalidFilePath
	PromptSampleFilePath
	PromptExit
)

// This method accepts a UserPrompt as input and pads it with the trailing spaces.
func GetUserPromptPadded(up UserPrompt, str string, strFmt string, newline bool) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	if newline {
		paddedStr += "\n"
	}
	return paddedStr
}

// This method accepts the UserPrompt as input and returns a padded string.
func GetUserPrompt(up UserPrompt) string {
	switch up {
	case PromptReadFilePath:
		return GetUserPromptPadded(up, "Please enter absolute file path:", "", false)
	case PromptInvalidFilePath:
		return GetUserPromptPadded(up, "Invalid path! Please enter a valid file path.", "", true)
	case PromptSampleFilePath:
		msg := "Eg: /dev/mooc/coursera/golang/week4/data.txt"
		return GetUserPromptPadded(up, msg, "", true)
	case PromptExit:
		return GetUserPromptPadded(up, "Enter 'X' to Exit", "", true)
	}
	return ""
}

// Reads and validates a string against a regular expression.
func readValidInput(str string, regex *regexp.Regexp) (string, bool) {
	if str == QuitCode {
		os.Exit(0)
	}
	match := regex.Match([]byte(str))
	if !match {
		return "", false
	}
	return str, true
}

// Prompts the user indefinitely until a valid input is entered or the user chooses to Exit.
func readUntilValid(readPrompt UserPrompt, invalidPrompt UserPrompt, regex *regexp.Regexp) string {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf(GetUserPrompt(readPrompt))
	for scanner.Scan() {
		value, isValid := readValidInput(scanner.Text(), regex)
		if value != "" && isValid {
			return value
		}
		fmt.Printf(GetUserPrompt(invalidPrompt))
		if invalidPrompt == PromptInvalidFilePath {
			fmt.Printf(GetUserPrompt(PromptSampleFilePath))
		}
		fmt.Printf(GetUserPrompt(readPrompt))
	}
	return ""
}

// Displays error and aborts execution.
func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// Checks if a file is accessible for reading.
func getFileHandle(path string) (bool, *os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, nil, err
	}

	if fi, err := file.Stat(); err != nil || fi.IsDir() {
		return false, nil, err
	}
	return true, file, nil
}

// Based on https://kgrz.io/reading-files-in-go-an-overview.html
// This function is not used in the program.
// I created it to understand the difference between reading an entire file
// into memory and reading in chunks.
func getFileAsString(file *os.File) (string, int, error) {
	fileinfo, err := file.Stat()
	if err != nil {
		return "", 0, err
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	byteCount, err := file.Read(buffer)
	if err != nil {
		return "", 0, err
	}

	return string(buffer), byteCount, nil
}

// Based on https://kgrz.io/reading-files-in-go-an-overview.html
// This function is not used in the program.
// I created it to understand the difference between reading an entire file
// into memory and reading in chunks.
func getFileInChunkAsString(file *os.File) (string, int, error) {
	bufferSize := BufferSize
	buffer := make([]byte, bufferSize)
	str := ""
	for {
		byteCount, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return "", 0, err
		} else if err != nil && err == io.EOF {
			return str, byteCount, nil
		}
		content := string(buffer[:byteCount])
		str += content
	}
}

// Based on https://www.socketloop.com/tutorials/golang-bufio-newreader-readline-to-read-file-line-by-line
// This function accepts a file pointer as input.
// Reads one line at a time.
// Splits each line using a space delimiter.
// Populates a Person struct by setting the fname and lname fields.
// Appends the Person struct to a slice.
// Returns the slice with person structs.
func getPersonSliceFromFile(file *os.File) []Person {
	reader := bufio.NewReader(file)
	var personSlice = make([]Person, 0, ListSize)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fields := strings.Split(string(line), " ")
		person := Person{fname: fields[0], lname: fields[1]}
		personSlice = append(personSlice, person)
	}
	return personSlice
}

// Main Program
// Prompts the user to enter a valid file path.
// Verifies that the file path is accessible for reading.
// Opens the file and passes the file pointer to getPersonSliceFromFile().
// Iterates through the person slice and displays the information on screen.
func main() {
	fmt.Printf(GetUserPrompt(PromptExit))
	RegexForFilePath := regexp.MustCompile(`^(.+)/([^/]+)$`)

	filePath := readUntilValid(PromptReadFilePath, PromptInvalidFilePath, RegexForFilePath)
	exist, file, _ := getFileHandle(filePath)
	if !exist {
		fmt.Printf("%s: File DOES NOT exist.\n", filePath)
		os.Exit(0)
	}
	defer file.Close()

	persons := getPersonSliceFromFile(file)
	for _, person := range persons {
		fmt.Printf("fname: [%-20s]; lname: [%-20s]\n", person.fname, person.lname)
	}
}
