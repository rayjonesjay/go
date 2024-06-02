package graphics

import (
	"ascii/colors"
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

	// Check if the file exists and has a size of 0 if true download automatically
	if FileInformation != nil && FileInformation.Size() == 0 {
		deleteEmptyBanner(fileName)
		download(fileName)
	}

	//if the file does not exist download
	if err != nil {
		downloadALL()
		PASS()
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
		//fmt.Printf("%s%q downloaded successfully.. rerun the program\n%s", colors.GREEN, strings.Replace(fileName, "banners/", "", -1), colors.RESET)
	}
}

// downloads all the banner files at once
func downloadALL() {
	urlFile := "plain/urls.txt"
	fileInfo , err := os.Stat(urlFile)

	if fileInfo.Size() != 287 {
		fmt.Println("url file contains wrong links.")
		PASS()
	}

	if os.IsNotExist(err) {
		bannerFiles := []string{"standard.txt","shadow.txt","thinkertoy.txt"}
		for _, file := range bannerFiles{
			download(file)
		}
		fmt.Println("download success rerun the program")
		PASS()
	}


	// first of all delete the banner files
	bannerFiles := []string{"shadow.txt", "standard.txt", "thinkertoy.txt"}
	path := os.ExpandEnv(PathToBannerFiles)


	for _, file := range bannerFiles {
		_, err := os.Stat(file)

		if err != nil {
			if !os.IsNotExist(err) {
				{
				} //do nothing
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
	_, wgetErr := downloadAll.CombinedOutput()

	if wgetErr != nil {
		// check error occurred when moving file to banners/
		_, file := filepath.Split(urlFile)
		fmt.Fprintf(os.Stderr, "wget error with link in %q check the file %v\n",file, wgetErr)
	}


	for _, file := range bannerFiles {
		moveFileToBannerDir(file)
	}

	fmt.Printf("%sDownload success. Rerun the program.\n%s", colors.GREEN, colors.RESET)
	PASS()
}

// after downloadAll() the files are moved to banners directory
func moveFileToBannerDir(file string) {
	path := os.ExpandEnv(PathToBannerFiles)
	move := exec.Command("mv", file, path)
	_, moveError := move.CombinedOutput()
	if moveError != nil {
		fmt.Fprintf(os.Stderr, "%serror moving %q check download link in \"plain/urls.txt if present\" %s\n", colors.RED,file,colors.RESET)
		FAIL()
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
		log.Fatalf("\npath error %s\n\t%v\n", path, err)
		return false
	}
	return true
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

		//create directory
		if makeDirectory(dir) {
			fmt.Printf("%q directory was missing, created successfully rerun the program.\n", strings.ReplaceAll(dir, "/", ""))
		   // PASS()
		}

	} else {
		// check for file existence
		filePath = checkFileExist(filePath)
	}

	file, openingFileError := os.Open(filePath)
	// handle error if the file does not exist error and also other possible errors
	if openingFileError != nil {
		if os.IsNotExist(openingFileError) {
			fmt.Println("lemons")
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
