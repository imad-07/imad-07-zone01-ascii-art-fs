package main

import (
	"fmt"
	"os"
	
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf("There is something missing\n")
		return
	}

	for i := 0; i < len(args[0]); i++ {
		if args[0][i] < 32 || args[0][i] > 126 {
			fmt.Printf("error in input\n")
			return
		}
	}
	word := split(args[0])

	fileContent, err := os.ReadFile(banner(args[1]))
	if err != nil {
		fmt.Printf("error in stabdard file")
		return
	}

	lettres := getLettres(fileContent)
	writing(lettres, word)
}

// There are three options for writing text and here we know your choice
func banner(arg string) string {
	banner := ""
	if arg == "shadow" {
		banner = "shadow.txt"
	} else if arg == "thinkertoy" {
		banner = "thinkertoy.txt"
	} else if arg == "standard" {
		banner = "standard.txt"
	} else {
		fmt.Println("There is an error in the banner")
		os.Exit(0)
	}
	return banner
}

// You write the code with some modifications
func writing(lettres [][]string, word []string) {
	bl := false
	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if len(word)-2 == l {
				fmt.Printf("\n")
				continue
			}
			if bl && word[l+1] != "\n" {
				continue
			}
			fmt.Printf("\n")
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				fmt.Printf(lettres[word[l][j]-32][i])
			}
			fmt.Print("\n")
		}
		bl = true
	}
}

func split(str string) []string {
	word := ""
	splitedword := []string{}

	for i := 0; i < len(str); i++ {
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""

			splitedword = append(splitedword, "\n")
			i++
			continue
		}
		word = word + string(str[i])
	}
	splitedword = append(splitedword, word)
	return splitedword
}

func getLettres(fileContent []byte) [][]string {
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	filtering := ""
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 13 {
			filtering = filtering + string(fileContent[i])
		}
	}
	for i := 0; i < len(filtering); i++ {
		if i != len(filtering)-1 && filtering[i] == '\n' && filtering[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if filtering[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, filtering[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}
