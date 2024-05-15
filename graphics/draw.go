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

	// FAIL To be called if a program was run in the wrong way
	FAIL = func() { os.Exit(1) }

	// PASS To be called if a program run successfully
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

	// Check if the file exists and has a size of 0 if its true ask the user to download a new one
	if FileInformation != nil && FileInformation.Size() == 0 {
		fmt.Printf("%s is an empty file. Do you wish to download the file? yes(y)/no(n): ", fileName)

		// read the input from standard input
		reader := bufio.NewReader(os.Stdin)

		// read the string until it encounters a newline character (when you press enter a newline char is appened to string)
		choice, _ := reader.ReadString('\n')

		// remove any trailing and leading white spaces
		choice = removeExtraSpace(choice)

		if strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {

			deleteEmptyBanner(fileName)
			download(fileName)
		} else if strings.ToLower(choice) == "no" || strings.ToLower(choice) == "n" {

			fmt.Printf("Cannot continue with the program without: %s\n", strings.ReplaceAll(fileName, "banners/", ""))
			FAIL()

		} else {

			fmt.Println("WRONG user input! Exiting...")
			FAIL()
		}

	}

	if err != nil {
		// Download the files automatically
		fmt.Printf("%s does not exist do you wish to download or get all files ? yes(y) no(n) all(a): ", strings.ReplaceAll(fileName, "banners/", ""))

		// make a reader
		reader := bufio.NewReader(os.Stdin)

		userChoice, err := reader.ReadString('\n')
		userChoice = removeExtraSpace(userChoice)

		if err != nil {
			log.Fatalf("error reading choice: check delimeter %v\n", err)
		}

		if strings.ToLower(userChoice) == "yes" || strings.ToLower(userChoice) == "y" {
			download(fileName)
		} else if strings.ToLower(userChoice) == "no" || strings.ToLower(userChoice) == "n" {

			fmt.Printf("Cannot continue with the program without: %s\n", strings.ReplaceAll(fileName, "banners/", ""))
			FAIL()

		} else if strings.ToLower(userChoice) == "all" || strings.ToLower(userChoice) == "a" {

			downloadALL()
			PASS()
		} else {

			fmt.Printf("Wrong Input: You entered -> %s Expected -> yes(y) or no(n).\n", userChoice)
			FAIL()
		}
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
		log.Fatalf("\n\tDownload attempt failed! Check your internet connection and try again.\n")
	} else {

		move := exec.Command("mv", file, "banners/")
		output, err := move.CombinedOutput()
		if err != nil {
			// check error occurred when moving file to banners/
			fmt.Fprintf(os.Stderr, "error: %T\n", err)
			fmt.Println(output)
		}
		fmt.Println()
		fmt.Printf("%s downloaded successfully.. rerun the program\n", strings.Replace(fileName, "banners/", "", -1))
		PASS()
	}
}

// downloads all the banner files at once
func downloadALL() {
	urlFile := "plain/urls.txt"
	_, err := os.Stat(urlFile)

	if os.IsNotExist(err) {
		fmt.Println("urls.txt does not exist")
		FAIL()
	}
	// first of all delete the banner files
	bannerFiles := []string{"shadow.txt", "standard.txt", "thinkertoy.txt"}
	path := os.ExpandEnv(PathToBannerFiles)
	for _, file := range bannerFiles {
		_, err := os.Stat(file)
		if err != nil {
			if !os.IsNotExist(err) {
				{
				} // same as do nothing
			}
		} else {

			rm := exec.Command("rm", path+file)
			out, err := rm.CombinedOutput()
			if err != nil {
				log.Fatalf("Error Deleting %s %v\n", out, err)
			}
		}
	}

	downloadAll := exec.Command("wget", "-i", os.ExpandEnv(PathToUrlFile))
	out, wgetErr := downloadAll.CombinedOutput()
	if wgetErr != nil {
		// check error occurred when moving file to banners/
		fmt.Fprintf(os.Stderr, "error: %T\n", wgetErr)
		fmt.Println(string(out))
	}
	for _, file := range bannerFiles {
		moveFileToBannerDir(file)
	}
	fmt.Printf("All banner Files downloaded Successfully... rerun the program\n")
	PASS()
}

// after downloadAll() the files are moved to banners directory
func moveFileToBannerDir(file string) {
	path := os.ExpandEnv(PathToBannerFiles)
	move := exec.Command("mv", file, path)
	fmt.Println(path + file)
	out, err := move.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ErrorA: %v Output: %s ", err, out)
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
func makeDirectory(dir string) bool {
	path := CurrentWorkingDirectory + dir
	dir = os.ExpandEnv(path)
	err := os.Mkdir(dir, 0o775)
	if err != nil {
		log.Fatalf("\nPath error %s\n\t%v\n", path, err)
		return false
	}
	return true
}

// remove extra leading and trailing spaces
func removeExtraSpace(str string) string {
	return strings.TrimSpace(str)
}

// ReadBanner This function takes the name of a file in string format,
// reads the ascii art inside the file, and maps each ascii art to its rune value
func ReadBanner(filePath string) map[rune][]string {
	fileName := strings.ReplaceAll(filePath, "banners/", "")
	fileName = fixFileExtension(fileName)
	if !(fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt") {
		fmt.Printf("%s not a valid banner file. Download it from your source and add it to banner/ directory.\n", fileName)
		os.Exit(1)
	}
	// split the directory and file apart
	dir, _ := filepath.Split(filePath)
	// if the directory does not exist
	if !directoryExist(dir) {

		fmt.Printf("%s does not exist. Do you wish to create it in order for the program to run?  yes(y) or no(n): ", dir)
		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		choice = removeExtraSpace(choice)

		if err != nil {
			log.Fatalf("error reading choice: check delimeter %v\n", err)
		}

		if strings.ToLower(choice) == "y" || strings.ToLower(choice) == "yes" {
			// make the directory
			if makeDirectory(dir) {
				fmt.Printf("%s \ndirectory created!! rerun the program with the required banner files inside.\n\n", strings.ReplaceAll(dir, "/", ""))
				PASS()
			}
		} else if strings.ToLower(choice) == "n" || strings.ToLower(choice) == "no" {
			fmt.Println("CAUTION: program cannot run without banner/ directory.")
			os.Exit(1)
		} else {
			fmt.Printf(" Wrong Input: You entered -> %s Expected -> yes(y) or no(n).\n", choice)
		}

	} else {
		// check for file existence
		filePath = checkFileExist(filePath)
	}

	file, openingFileError := os.Open(filePath)
	// handle error if the file does not exist error and also other possible errors
	if openingFileError != nil {
		if os.IsNotExist(openingFileError) {
			// file not found, exit unsuccessfully with code 1
			os.Exit(1)
		}
		log.Fatalf("Failed to open banner file: \"%s\"", filePath)
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
