// TODO:: CATCH THE ERRORS 
// TODO:: ADD STDIN & STDOUT ALIASES 
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var target string
	var output string
	var lineNum string
	var noMsg bool
	var splitArgs string

	flag.StringVar(&target, "t", "", "Target file")
	flag.StringVar(&output, "o", "", "Prints the output to a text file")
	flag.StringVar(&lineNum, "l", "", "Gets that line which is entred")
	flag.StringVar(&splitArgs, "s", "", "It splits the text")
	flag.BoolVar(&noMsg, "noMsg", false, "It returns just output")

	flag.Parse()

	if target == "" {
		fmt.Fprintln(os.Stderr, "ERROR: There is no any target file")
		os.Exit(1)
	}

	text := readFile(target)

	if lineNum != "" {
		text = getLine(lineNum, text)
	}

	if splitArgs != "" {
		text = split(splitArgs, text)
	}


	if !noMsg {
		text += "\nDone"
	}

	if output != "" {
		os.Create(output)

		err := ioutil.WriteFile(output, []byte(text+"\n"), 0444)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: Something wrong during printing output")
			os.Exit(1)
		}
	}

	fmt.Println(text)

}

func readFile(target string) string {
	stream, err := ioutil.ReadFile(target)

	if err != nil {
		log.Fatal(err)
	}

	return string(stream)
}

func getLine(lineNums, text string) string {
	nums := strings.Split(lineNums, ",")
	lines := strings.Split(text, "\n")
	res := []string{}
	for _, i := range nums {
		i, _ := strconv.Atoi(i)
		if len(lines) < i || i <= 0 {
			fmt.Fprintln(os.Stderr, "ERROR: Invalid line", i)
			os.Exit(1)
		}

		res = append(res, lines[i-1])
	}

	return strings.Join(res, "\n") 
}

func split(args, text string) string{
	parsedArgs := strings.Split(args, ",")
	if parsedArgs[0] == "-comma-" {
		parsedArgs[0] = ","
	}
	
	if parsedArgs[0] == "-newLine-" {
		parsedArgs[0] = "\n"
	}

	if parsedArgs[0] == "-space-" {
		parsedArgs[0] = " "
	}

	word := parsedArgs[0]

	indexes := parsedArgs[1:]

	splitted_text := strings.Split(text, word)
	res :=  []string{}
	for _ , i := range indexes {
		i, _ :=  strconv.Atoi(i)
		if len(splitted_text) < i || i <= 0 {
			fmt.Fprintln(os.Stderr, "ERROR: Invalid index", i)
			os.Exit(1)
		}
		res = append(res, splitted_text[i-1])
	}

	return strings.Join(res, "\n")
}

