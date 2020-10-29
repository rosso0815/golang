package main

import (
	"bufio"
	"fmt"
	//	"io/ioutil"
	"os"
	"testing"
)

var fileName = "test.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestWriteStdout(t *testing.T) {

	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")

	w.Flush() // Don't forget to flush!
}

func TestWriteFile(t *testing.T) {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	//d1 := []byte("hello\ngo\n")
	//err := ioutil.WriteFile("text01.txt", d1, 0644)
	//check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create(fileName)
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	//d2 := []byte{115, 111, 109, 101, 10}
	//n2, err := f.Write(d2)
	//check(err)
	//fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("select * from table01;\nselect gugus from hans;\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	//w := bufio.NewWriter(f)
	//n4, err := w.WriteString("buffered\n")
	//fmt.Printf("wrote %d bytes\n", n4)
}

func TestReadFile(t *testing.T) {

	fmt.Println("@@@ TestReadFile")

	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println("line =", scanner.Text())

	}

}
