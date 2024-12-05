package utils

import (
	"os"
)

func Opener(filename string) (*os.File,error) {
	fd , err := os.Open(filename)
	if err != nil {
		nil,err
	}
	return fd,nil
}

func Reader(fd *os.File) (string, byte) {
	defer fd.Close()
	scanner := Scanner(fd)
	for scanner.Scan(){
		return scanner.Text(), 0
	}
	return "" , 1
}

func Scanner(fd *os.File) bufio.NewScanner {
	scanner := bufio.NewScanner(fd)
	return scanner
}
