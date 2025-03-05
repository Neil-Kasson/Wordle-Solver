package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var words []string

func init(){
	//	Populate the word list
	f, _ := os.Open("all_words.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		words = append(words, scanner.Text())
	}
	slices.Sort(words)
}

func main(){
	var again = true
	var guess = ""
	var pattern = ""
	var temp1 = ""
	var temp2 = ""
	var goodLetters = ""

	for again{
		fmt.Print("What word did you enter? ")
		fmt.Scanln(&guess)
		fmt.Printf("Using the following, what colors were each letter?\n green  ->  +\n yellow ->  -\n grey   ->  *\n%s\n", guess)
		fmt.Scanln(&pattern)

		// Remove impossible words
		for i:=0; i<5; i++{
			if(string(pattern[i]) == "+"){
				words = slices.DeleteFunc(words, func(s string) bool{
					return !(string(s[i]) == string(guess[i]))
				})
				goodLetters += string(guess[i])
			} else if(string(pattern[i]) == "-"){
				words = slices.DeleteFunc(words, func(s string) bool{
					return (!strings.Contains(s, string(guess[i]))) || (string(s[i]) == string(guess[i]))
				})
				goodLetters += string(guess[i])
			} else if(string(pattern[i]) == "*"){
				words = slices.DeleteFunc(words, func(s string) bool{
					return strings.Contains(s, string(guess[i])) && !strings.Contains(goodLetters, string(guess[i]))
				})
			}
		}

		// Prompt to show answers
		fmt.Print("\nWould you like to see the ", len(words), " possible answers? (y/n): ")
		fmt.Scanln(&temp1)
		if(temp1 == "y"){
			fmt.Println("\n", words)
		}

		// Prompt to enter another word
		fmt.Print("\nWould you like to enter another word? (y/n): ")
		fmt.Scanln(&temp2)
		if(temp2 == "n"){
			again = false
		}

	}
	// End of for loop

	// https://pkg.go.dev/slices#example-DeleteFunc
}

