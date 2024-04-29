package graphics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func DisplayArt(asciiMap map[rune][8]string){
	letter := 'a'
	for key, value := range asciiMap{
		if key == letter {
			for _, val := range value {
				fmt.Println(val)
			}
		}
	}
}
// 
func ReadBanner(fileName string) map[rune][8]string {
	link1 := "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	link2 := "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	link3 := "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	file, possibleError := os.Open(fileName)
	// handle error if the file does not exist error and also other possible errors
	if possibleError != nil {
		if os.IsNotExist(possibleError) {
			fmt.Fprintf(os.Stderr, "\"%s\" file does not exist: download the required files from the links below: \n\t%s\n\t%s\n\t%s\n", fileName, link1, link2, link3)
			os.Exit(2)
		}
		log.Fatalf("Failed to open banner file: \"%s\"", fileName)
	}
	defer file.Close()
	// create a scanner object
	scan := bufio.NewScanner(file)
	// create ascii art map to store the character and its equivalent ascii art
	asciiMap := make(map[rune][8]string)
	// i is our loop variable and it starts from 32 which is space character
	for i := 32; i <= 126; i++ {
		currentRune := rune(i)
		// skip one line
		scan.Scan()
		// array to store the current ascii art
		var currentArt [8]string
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
	}
	// Create an array of arrays to represent four spaces, each eight lines long
	return asciiMap
}
