package graphics

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	CurrentWorkingDirectory = "$PWD/"
	bannersDir              = "banners/"
)

// checkFileExist checks if the specified banner file exists in the banner/ directory and that it is not an empty file;
// offers to download the banner file if it does not yet exist
func checkFileExist(filePath, fileName string) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		// file does not exist, download
		download(fileName)
	} else if fileInfo != nil && fileInfo.Size() == 0 {
		// download the banner file if it exists and is an empty file
		deleteEmptyBanner(filePath)
		download(fileName)
	}
}

// directoryExist checks if the given directory exists
func directoryExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	// IsNotExist() will return true iff the directory/file at the given file path does not exist
	return !os.IsNotExist(err)
}

// deleteEmptyBanner removes the specified empty banner file, possibly
// before calling download to avoid conflict
func deleteEmptyBanner(fileName string) {
	filePath := CurrentWorkingDirectory + fileName
	filePath = os.ExpandEnv(filePath)
	rm := exec.Command("rm", filePath)
	_, err := rm.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to delete empty banner: %q\n%v\n", fileName, err)
	}
}

// download banner files from https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/
func download(fileName string) {
	if !Contains([]string{"standard.txt", "shadow.txt", "thinkertoy.txt"}, fileName) {
		log.Fatalf(
			"%q is not a known banner; download the banner file from your source and add it to the %q directory\n",
			strings.TrimSuffix(fileName, ".txt"), bannersDir,
		)
	}

	rawLink := "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/"
	fullPath := rawLink + fileName
	downloadCommand := exec.Command("wget", fullPath)
	output, err := downloadCommand.CombinedOutput()

	if err != nil {
		log.Fatalf("failed to download banner: %q\n%s\n", fileName, output)
	} else {
		move := exec.Command("mv", fileName, "banners/")
		_, err := move.CombinedOutput()
		if err != nil {
			log.Fatalf("failed to move banner: %q to %q directory\n%v\n", fileName, bannersDir, err)
		}
	}
}

// makeDirectory creates the banner directory if it does not exist
func makeDirectory(dir string) {
	path := CurrentWorkingDirectory + dir
	dir = os.ExpandEnv(path)
	err := os.Mkdir(dir, 0o775)
	if err != nil {
		log.Fatalf("failed to create banners directory at %q\n%v\n", path, err)
	}
}

// ReadBanner This function takes the name of a file in string format,
// reads the ascii art inside the file, and maps each ascii art to its rune value
func ReadBanner(fileName string) map[rune][]string {
	filePath := bannersDir + fileName

	// create banners directory if it does not exist
	if !directoryExist(bannersDir) {
		makeDirectory(bannersDir)
	}
	checkFileExist(filePath, fileName)

	file, openingFileError := os.Open(filePath)
	// handle error if the file does not exist error and also other possible errors
	if openingFileError != nil {
		log.Fatalf("failed to open banner file: %q\n%v\n", filePath, openingFileError)
	}
	defer Close(file)

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
	return asciiMap
}

// Close the given file
func Close(file *os.File) {
	_ = file.Close()
}

// Contains checks whether the string s is an element of the given array
func Contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
