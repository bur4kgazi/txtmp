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
	var split string
	var lineNum string
	var noMsg bool

	flag.StringVar(&target, "t", "", "Target file")
	flag.StringVar(&output, "o", "", "Prints the output to a text file")
	flag.StringVar(&split, "s", "", "It splits the text")
	flag.StringVar(&lineNum, "l", "", "Gets that line which is entred")
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

	if split != "" {
		text = splitText(split, text)
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

func splitText(args, text string) string {
	argsParsed := strings.Split(args, ",")
	if argsParsed[0] == "<comma>" {
		argsParsed[0] = ","
	}

	indexes := argsParsed[1:]

	result := ""
	for _, el := range indexes {
		i, err := strconv.Atoi(el)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: You must entre an numerical value")
		}

		result += strings.Split(text, argsParsed[0])[i-1] + "\n"

	}
	return result
}
