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
	// Declaring some variables to use later
	var again = true
	var guess = ""
	var pattern = ""
	var temp1 = ""
	var temp2 = ""
	var goodLetters = ""

	// Loop to prompt/enter/check each word
	for again{

		// Prompts for word and placecemnt of letter correctness
		fmt.Print("\nWhat word did you enter? ")
		fmt.Scanln(&guess)
		fmt.Printf("\nUsing the following, what colors were each letter?\n green  ->  +\n yellow ->  -\n grey   ->  *\n%s\n", guess)
		fmt.Scanln(&pattern)

		// Remove impossible words
		for i:=0; i<5; i++{
			if(string(pattern[i]) == "+"){											// Green letters
				words = slices.DeleteFunc(words, func(s string) bool{
					return !(string(s[i]) == string(guess[i]))
				})
				goodLetters += string(guess[i])
			} else if(string(pattern[i]) == "-"){								// Yellow letters
				words = slices.DeleteFunc(words, func(s string) bool{
					return (!strings.Contains(s, string(guess[i]))) || (string(s[i]) == string(guess[i]))
				})
				goodLetters += string(guess[i])
			} else if(string(pattern[i]) == "*"){								// Grey letters
				words = slices.DeleteFunc(words, func(s string) bool{
					return strings.Contains(s, string(guess[i])) && !strings.Contains(goodLetters, string(guess[i]))
				})
			}
		}

		// Prompt to show answers
		fmt.Print("\nWould you like to see the ", len(words), " possible answers? (y/n): ")
		fmt.Scanln(&temp1)
		if(temp1 == "y"){
			output := ""
			for i, w := range words{
				if (i%10 == 0){
					output += "\n"
				}
				output += w + "\t"
			}
			fmt.Println(output)
		}

		// Prompt to enter another word
		fmt.Print("\nWould you like to enter another word? (y/n): ")
		fmt.Scanln(&temp2)
		if(temp2 == "n"){
			again = false
		}

		// I know these last 2 prompt sections don't account for incorrect entries, 
		// but it accomplishes the main functionality I entended it to... 
		// it'll show the list if you type 'y', or not if you don't
		// and exit the program if you type 'n' in the 2nd prompt, or repeat if not.

	}
	// End of for loop

	// https://pkg.go.dev/slices#example-DeleteFunc
}

