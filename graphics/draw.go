package graphics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	FAIL = func() { os.Exit(1) }

	PASS = func() { os.Exit(0) }
)

const (
	CurrentWorkingDirectory = "$PWD/"
	PathToUrlFile           = CurrentWorkingDirectory + "plain/urls.txt"
	PathToBannerFiles       = CurrentWorkingDirectory + "banners/"
)

// checkFileExist checks if the specified banner file exists in the banner/ directory,
// offers to download the banner file if it does not yet existCURRENT_WORKING_DIRECTORY
func checkFileExist(fileName string) string {
	FileInformation, err := os.Stat(fileName)

	// Check if the file exists and has a size of 0 if true download automatically
	if FileInformation != nil && FileInformation.Size() == 0 {
		deleteEmptyBanner(fileName)
		download(fileName)
	}

	// if the file does not exist download
	if err != nil {
		download(fileName)
	}

	return fileName
}

// DirectoryExist() checks if a directory exist
func directoryExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	// IsNotExist() will return true iff the directory/file at the given file path does not exist
	return !os.IsNotExist(err)
}

// deleteEmptyBanner removes the specified empty banner file with no ascii art,
// and it is called before calling download to avoid conflict
func deleteEmptyBanner(filename string) {
	filePath := CurrentWorkingDirectory + filename
	filePath = os.ExpandEnv(filePath)
	rm := exec.Command("rm", filePath)
	_, err := rm.CombinedOutput()
	if err != nil {
		log.Fatalf("error deleting %v", err)
	}
}

// download banner files from "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/"
func download(fileName string) {
	rawLink := "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/"
	file := strings.ReplaceAll(fileName, "banners/", "")
	fullPath := rawLink + file
	downloadCommand := exec.Command("wget", fullPath)
	_, err := downloadCommand.CombinedOutput()

	if err != nil {
		log.Fatalf("\n\tdownload attempt failed! check your internet connection and try again.\n")
	} else {

		move := exec.Command("mv", file, "banners/")
		output, err := move.CombinedOutput()
		if err != nil {
			// check error occurred when moving file to banners/
			fmt.Fprintf(os.Stderr, "error: %T\n", err)
			fmt.Println(output)
		}
	}
}

// fixFileExtension removes additional extensions from a file and returns only the first extension
func fixFileExtension(fileName string) string {
	// the purpose of this part is to remove the repetitive extension if any for example file.txt.txt -> file.txt
	// file extension - a group of words occurring after a (.) period character indicating the format of the file
	position := strings.Index(fileName, ".")
	if position != -1 {
		fileName = fileName[:position] + ".txt"
	}
	return fileName
}

// makeDirectory creates the banner directory if it does not exist
func makeDirectory(dir string) {
	path := CurrentWorkingDirectory + dir
	dir = os.ExpandEnv(path)
	err := os.Mkdir(dir, 0o775)
	if err != nil {
		log.Fatalf("\npath error %s\n\t%v\n", path, err)
	}
}

// ReadBanner This function takes the name of a file in string format,
// reads the ascii art inside the file, and maps each ascii art to its rune value
func ReadBanner(filePath string) map[rune][]string {
	fileName := strings.ReplaceAll(filePath, "banners/", "")
	fileName = fixFileExtension(fileName)
	if !(fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt") {
		fmt.Printf(
			"%s not a valid banner file. Download it from your source and add it to banner/ directory.\n",
			fileName,
		)
		os.Exit(1)
	}
	// split the directory and file apart
	dir, _ := filepath.Split(filePath)

	// if the directory does not exist
	if !directoryExist(dir) {
		// create directory regardless
		makeDirectory(dir)
	} else {
		// check for file existence
		filePath = checkFileExist(filePath)
	}

	file, openingFileError := os.Open(filePath)
	// handle error if the file does not exist error and also other possible errors
	if openingFileError != nil {
		if os.IsNotExist(openingFileError) {
			download(fileName)
			file, openingFileError = os.Open(filePath)
			if openingFileError != nil {
				log.Fatalf("Failed to open banner file: \"%s\"", filePath)
			}
		} else {
			log.Fatalf("Failed to open banner file: \"%s\"", filePath)
		}
	}

	defer file.Close()

	// create a scanner object
	scan := bufio.NewScanner(file)
	// create an ascii art map to store the character and its equivalent ascii art
	asciiMap := make(map[rune][]string)

	// variable i is our loop counter, and it starts from 32 which is space character
	for i := 32; i <= 126; i++ {

		currentRune := rune(i)
		// skip one line
		scan.Scan()
		// array to store the current ascii art
		currentArt := make([]string, 8)

		// read the next 8 lines
		for count := 0; count < 8; count++ {
			// if we reach the end of the file prematurely, stop scanning
			if !scan.Scan() {
				break
			} else {
				// append each line we read to the currentArt array
				line := scan.Text()
				currentArt[count] = line
			}
		}

		asciiMap[currentRune] = currentArt
	}
	// Create an array of arrays to represent four spaces, each eight lines long
	return asciiMap
}
