package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	paths := []string{
		"/home/",
		"/home/user/Documents/../Pictures",
		"/home//foo/",
		"/../",
		"/.../a/../b/c/../d/./",
		"/a/../../b/../c//.//",
	}

	id := 1
	for _, p := range paths {
		fmt.Println(id, ">", simplifyPath(p))
		id++
	}
}

func simplifyPath(path string) string {

	path = filepath.Clean(path)
	isRoot := func(p string) bool{
		return p == "/"
	}

	if isRoot(path){
		return "/"
	}
	hasRoot := path[0] == '/'

	pathStack := []string{}

	var final string

	var tmp string
	for i := 0; i < len(path); i++ {
		if path[i] == '/'{
			if tmp != ""{
				if tmp == ".." && len(pathStack) != 0 {
					pathStack = pathStack[:len(pathStack)-1]
					tmp = ""
					continue
				}

				if strings.Count(tmp, ".") == 1 {
					tmp = ""
					continue
				}
				pathStack = append(pathStack, "/"+tmp)
				tmp = ""
			}
		}else{
			tmp += string(path[i])
			if  i < len(path)-1 && (tmp == ".." && path[i+1] != '.' && len(pathStack) != 0){
				pathStack = pathStack[:len(pathStack)-1]
				tmp = ""
			}
		}
	}

	
	if tmp != "" {
		pathStack = append(pathStack, "/"+tmp)
	}
	
	final = strings.Join(pathStack,"")

	if final == "" && hasRoot {
		return "/"
	}


	final = strings.TrimRight(final,"/")

	final = removeFrontDots(final)
	return final
}


func removeFrontDots(s string) string{
	// var res string
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '.'{
			continue
		}else{
			return s[:i+1]
		}
	}
	return ""
}