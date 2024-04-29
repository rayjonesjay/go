package graphics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
) 

// This function takes the name of a file in string format, reads the ascii art inside the file, and maps each ascii art to its rune value
func ReadBanner(fileName string) map[rune][]string {
	link1 := "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	link2 := "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	link3 := "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	file, possibleError := os.Open(fileName)
	// handle error if the file does not exist error and also other possible errors
	if possibleError != nil {
		if os.IsNotExist(possibleError) {
			fmt.Fprintf(os.Stderr, "\"%s\" file does not exist: download the required files from the links below: \n\t%s\n\t%s\n\t%s\n", fileName, link1, link2, link3)
			os.Exit(1) //file not found exit code
		}
		log.Fatalf("Failed to open banner file: \"%s\"", fileName)
	}
	defer file.Close()
	// create a scanner object
	scan := bufio.NewScanner(file)
	// create ascii art map to store the character and its equivalent ascii art
	asciiMap := make(map[rune][]string)
	// i is our loop variable and it starts from 32 which is space character
	for i := 32; i <= 126; i++ {
		currentRune := rune(i)
		// skip one line
		scan.Scan()
		// array to store the current ascii art
		var currentArt []string = make([]string, 8)
		// read the next 8 lines
		for count := 0; count < 8; count++ {
			// if we reach the end of the file prematurely stop scanning
			if !scan.Scan() {
				break
			} else {
				// append each line we read to the currentArt array
				line := scan.Text()
				// Replace tabs with four spaces in the line
				line = strings.ReplaceAll(line, "\n", "    ")
				currentArt[count] = line
			}
		}
		asciiMap[currentRune] = currentArt


		// //handle \t
		// asciiMap[rune(9)] = []string {
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// 	asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0],
		// }
		
		var rows int = 8
		asciiMap[9] = make([]string, rows)
		
		for i := 0 ; i < rows; i++{
			//in each row there is a slice of 4*space
			asciiMap[rune(9)][i] = asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]+asciiMap[32][0]
		}
	}
	// Create an array of arrays to represent four spaces, each eight lines long
	return asciiMap
}
