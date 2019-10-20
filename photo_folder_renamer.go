package main

import (
    "fmt"
	"log"
    "os"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	var separator_in = ""
	var separator_in_mod = ""
	var separator_out = ""
	
	if len(os.Args) > 0 {
		separator_in = os.Args[1]
		separator_in_mod = os.Args[1]
	} else {
		separator_in = "-"
	}
	
	if len(os.Args) > 1 {
		separator_out = os.Args[2]
	} else {
		separator_out = "-"
	}
	
	if separator_in_mod == "." {
		separator_in_mod = "\\."
	}

    dirname := "."

    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

	renameFolders(files, separator_in, separator_in_mod, separator_out)
}

func renameFolders(files []os.FileInfo, separator_in string, separator_in_mod string, separator_out string){
	regex := "^[0-9]{2}" + separator_in_mod + "[0-9]{2}" + separator_in_mod + "[0-9]{4}$"
	fmt.Println("Regex:" + regex)
    for _, file := range files {
		fmt.Println(file.Name())
		
		matched, err := regexp.MatchString(regex, file.Name())
		fmt.Println(matched)
		if err != nil {
			log.Fatal(err)
		}
		
		if matched {
			var numbers = strings.Split(file.Name(), separator_in)
			var newName = numbers[2] + separator_out + numbers[1] + separator_out + numbers[0]
			var newNameTmp = newName
			var i = 1
			var renamingInWork = true
			
			for renamingInWork {
				fmt.Println("Rnaming: " + newName)
				err := os.Rename(file.Name(), newName)
				if err != nil {
					newName = newNameTmp + strconv.Itoa(i)
					i = i + 1
				}else {
					renamingInWork = false
				}
			}
			
			if err != nil {
				log.Fatal(err)
			}
		}
    }
} 
