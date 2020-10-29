package readCSV

import (
	"bufio"
	"encoding/csv"
	//	"fmt"
	"io"
	"log"
	"os"
)

func Read(file string) int {

	// Load a TXT file.
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	r.Comment = '#'
	for {
		// read line
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		log.Print(record)
		log.Print("len=", len(record))
		for value := range record {
			//fmt.Printf(" = %v\n", record[value])
			log.Print("value=", record[value])
		}

	}

	return 0
}

func Clear() {
	log.Print("csv clear")
}
