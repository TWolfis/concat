package main

/*
concat concatenates transforms a string to lowercase and concatonates it based on a seperator char
Author: T.Wolfis
Version: 1.0.0
*/

import (
	concat "concat/internal"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	str       string
	seperator = ""
	split     = " "
	cassing   = false
)

func main() {
	flag.StringVar(&str, "s", "", "String to concat")
	flag.StringVar(&seperator, "sep", seperator, "Seperator to use")
	flag.StringVar(&split, "split", split, "Split the string on this character")
	flag.BoolVar(&cassing, "c", cassing, "keep casing, transform everything to lower-case by default")
	flag.Parse()

	if str == "" {
		// read in the arguments in an interactive way
		fmt.Printf("Enter a string to concat: ")
		_, err := fmt.Scanln(&str)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Enter a seperator: ")
		_, err = fmt.Scanln(&seperator)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Enter a split character: ")
		_, err = fmt.Scanln(&split)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	out := concat.Concat(str, seperator, split)

	if !cassing {
		out = strings.ToLower(out)
	}

	fmt.Println(out)
}
