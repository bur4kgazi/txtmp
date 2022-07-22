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
	var multiLine string
	var split string
	var lineNum int
	var noMsg bool

	flag.StringVar(&target, "t", "", "Target file")
	flag.StringVar(&output, "o", "", "Prints the output to a text file")
	flag.StringVar(&multiLine, "ml", "", "Gets that lines which are entred")
	flag.StringVar(&split, "s", "", "It splits the text")
	flag.IntVar(&lineNum, "l", 0, "Gets that line which is entred")
	flag.BoolVar(&noMsg, "noMsg", false, "It returns just output")

	flag.Parse()

	if target == "" {
		fmt.Fprintln(os.Stderr, "ERROR: There is no any target file")
		os.Exit(1)
	}

	text := readFile(target)

	if lineNum != 0 {
		text = getLine(lineNum, text)
	}

	if multiLine != "" {
		text = getLines(multiLine, text)
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

func getLine(lineNum int, text string) string {
	lines := strings.Split(text, "\n")
	res := 0
	for res, _ = range lines {
		res = res
	}
	if res < lineNum {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid line")
		os.Exit(1)
	}
	return lines[lineNum-1]
}

func getLines(lineNums, text string) string {
	nums := strings.Split(lineNums, ",")
	lines := strings.Split(text, "\n")
	res := 0
	for res, _ = range lines {
		res = res
	}
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	return strings.Join(lines[num1-1:num2-1], "\n")
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
