package graphics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"path/filepath"
	"os/exec"
	"time"
) 

//makedirectory makes banner directory if it does not exist
func makedirectory(dir string) bool {
	path := "$HOME/ascii-art/"+dir
	dir = os.ExpandEnv(path)
	err := os.Mkdir(dir,0775)
	if err != nil {
		log.Fatalf("\nPath error %s\n\t%v\n",path,err)
		return false
	}
	return true
}

func DirectoryExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	//IsNotExist() will return true if the directory does not exist or the error returned is not about file missing 
	return !os.IsNotExist(err)
}

func download(fileName string) {
	rawLink := "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/"
	file := strings.ReplaceAll(fileName, "banners/" ,"")
	fullPath := rawLink+file
	download_command := exec.Command("wget" , fullPath)
	_, err := download_command.CombinedOutput()

	if err != nil {

		log.Fatalf("\n\tDownload attempt failed! Check your internet connection and try again.\n%T\n", err)
		os.Exit(1)

	}else{

		move := exec.Command("mv" , file, "banners/")
		out , err := move.CombinedOutput()

		if err != nil{
			fmt.Fprintf(os.Stderr, "error: %T\n",err)
			fmt.Println(out)
		}
	}
}

//fix extension removes additional extensions from a file and returns only the first extension
func fixFileExtension(file_name string) string {
	//the purpose of this part is to remove the repetitive extension if any for example file.txt.txt -> file.txt

	//file extension - a group of words occurring after a . period character indicating the format of the file
	position := strings.Index(file_name, ".")
	if position != -1 {
		file_name = file_name[:position]+".txt"
	}
	return file_name
}

func checkFileExist(fileName string) string {

		//check if the requested files exist.
		FileInformation, err := os.Stat(fileName)

		if err != nil {
			//Download the files automatically
			var userChoice string
			fmt.Printf("%s does not exist do you wish to download? yes/no\n",strings.ReplaceAll(fileName,"banners/",""))
			fmt.Scan(&userChoice)

			if (strings.ToLower(userChoice)== "yes" || strings.ToLower(userChoice) == "y"){
				download(fileName)

				sourceLink:= "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/"
				fmt.Printf("Wait..fetching resource from: %s\n",sourceLink)
				duration_to_wait := 2 * time.Second
				time.Sleep(duration_to_wait)
				fmt.Printf("%s download success...\n",strings.ReplaceAll(fileName, "banners/", ""))

			}else if (strings.ToLower(userChoice) == "no" || strings.ToLower(userChoice) == "n") {
				fmt.Printf("Cannot continue with the program without: %s\n",strings.ReplaceAll(fileName, "banners/", ""))
				os.Exit(1)
			}else{
				fmt.Println("WRONG user input.Exiting...")
				os.Exit(1)
			}
		// Check if the file exists and has a size of 0
		if FileInformation != nil && FileInformation.Size() == 0 {
			log.Fatalf("%s is an empty file. Does not contain ASCII art!!", fileName)
			os.Exit(1)
		}
	}
	return fileName
}
	



// This function takes the name of a file in string format, reads the ascii art inside the file, and maps each ascii art to its rune value
func ReadBanner(fileName string) map[rune][]string {

	file_name := strings.ReplaceAll(fileName, "banners/", "")
	file_name = fixFileExtension(file_name)
	if !(file_name == "standard.txt" || file_name == "shadow.txt" || file_name == "thinkertoy.txt"){
		fmt.Printf("%s not a valid banner file. Download it from your source and add it to banner/ directory.\n",file_name)
		os.Exit(1)
	}
	//split the directory and file apart
	dir , _ := filepath.Split(fileName)
	//if the directory does not exist
	if !DirectoryExist(dir) {

		var choice string 
		fmt.Printf("%s does not exist. Do you wish to create it in order for the program to run?  yes(y) or no(n).\n",dir)
		fmt.Scan(&choice)

		if strings.ToLower(choice) == "y" || strings.ToLower(choice) == "yes" {
			//make the directory
			if makedirectory(dir) {
				fmt.Printf("%s \ndirectory created!! rerun the program with the required banner files inside.\n\n",strings.ReplaceAll(dir, "/", ""))
			}
		}else if  strings.ToLower(choice) == "n" || strings.ToLower(choice) == "no" {
			fmt.Println("CAUTION: program cannot run without banner/ directory.")
			os.Exit(0)
		}else{
			fmt.Printf("Wrong Input: You entered -> %s Expected -> yes(y) or no(n).\n",choice)
		}


	}else{
		//check
		fileName = checkFileExist(fileName)
	}
	file, possibleError := os.Open(fileName)
	// handle error if the file does not exist error and also other possible errors
	if possibleError != nil {
		if os.IsNotExist(possibleError) {
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
