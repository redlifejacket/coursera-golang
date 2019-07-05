/*
 * Author: Shyam Govardhan (16 July 2019)
 * Coursera Functions, methods and Interfaces in Go (UCI)
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

const (
	QuitCode string = "X"
)

var animalMap map[string]*Animal
var animalArray = []string{"cow", "bird", "snake"}
var requestArray = []string{"eat", "move", "speak"}

func usage() {
	fmt.Println("Usage: <animal> <request>")
	fmt.Printf("where animal is one of (%s)\n", strings.Join(animalArray, ","))
	fmt.Printf("where request is one of (%s)\n", strings.Join(requestArray, ","))
	fmt.Println("Enter X to exit...")
}

func init() {
	// https://golang.org/doc/effective_go.html#allocation_new
	cow := new(Animal)
	bird := new(Animal)
	snake := new(Animal)
	animalMap = make(map[string]*Animal)
	animalMap[animalArray[0]] = cow.InitMe("grass", "walk", "moo")
	animalMap[animalArray[1]] = bird.InitMe("worms", "fly", "peep")
	animalMap[animalArray[2]] = snake.InitMe("mice", "slither", "hsss")
}

func test() {
	fmt.Printf(animalArray[0])
	animalMap[animalArray[0]].Eat()
	fmt.Printf(animalArray[1])
	animalMap[animalArray[1]].Move()
	fmt.Printf(animalArray[2])
	animalMap[animalArray[2]].Speak()
}

func processRequest(strAnimal string, strRequest string) {
	animal := animalMap[strAnimal]
	switch strRequest {
	case requestArray[0]:
		animal.Eat()
	case requestArray[1]:
		animal.Move()
	case requestArray[2]:
		animal.Speak()
	default:
		fmt.Printf("%s: Invalid Request.\n", strRequest)
		usage()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	usage()
	fmt.Println(">")
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == QuitCode {
			os.Exit(0)
		}
		strVal = stripSpaces(strVal)
		strVal = strings.ToLower(strVal)
		var list = strings.Split(strVal, " ")
		if len(list) != 2 {
			usage()
			continue
		}
		animal := list[0]
		request := list[1]
		if !(itemIsInArray(animal, animalArray) && itemIsInArray(request, requestArray)) {
			usage()
			continue
		}
		processRequest(animal, request)
		fmt.Println(">")
	}
}

// Animal Object
type Animal struct{ food, locomotion, noise string }

func (a *Animal) Eat() {
	fmt.Printf("%s\n", a.food)
}

func (a *Animal) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Printf("%s\n", a.noise)
}

func (a *Animal) InitMe(f, l, n string) *Animal {
	a.food = f
	a.locomotion = l
	a.noise = n
	return a
}

func (a *Animal) PrintMe() {
	fmt.Printf("food: %s\n", a.food)
	fmt.Printf("locomotion: %s\n", a.locomotion)
	fmt.Printf("noise: %s\n", a.noise)
}

// Utility Methods
func stripSpaces(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := re_leadclose_whtsp.ReplaceAllString(input, "")
	result = re_inside_whtsp.ReplaceAllString(result, " ")
	return result
}

func itemIsInArray(item string, array []string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}
