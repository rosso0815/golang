package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

// grep file
func grep(re, filename string) {
	fmt.Println("re=", re)
	regex, err := regexp.Compile(re)
	if err != nil {
		return // there was a problem with the regular expression.
	}

	fh, err := os.Open(filename)
	f := bufio.NewReader(fh)

	if err != nil {
		return // there was a problem opening the file.
	}
	defer fh.Close()

	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			return
		}
		s := string(buf)
		//		fmt.Fprintln("buffer", s)
		if regex.MatchString(s) {
			fmt.Printf("%s\n", string(buf))
		}
	}
}

func test01() {
	fmt.Println("test01")
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}

func main() {

	// read input
	flag.Parse()
	if flag.NArg() == 2 {
		// grep(flag.Arg(0), flag.Arg(1))
		test01()
	} else {
		fmt.Printf("Wrong number of arguments.\n")
	}

	// end
}
