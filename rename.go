package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var filepath string
	if len(os.Args) > 1 {
		filepath = os.Args[1]
	} else {
		log.Fatal("Please provide the csv file path as an argument")
	}

	csvfile, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		// Read each record from csv
		name, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//oldname := name[0]
		//newname := name[1]
		err = os.Rename(name[0], name[1])
		if err != nil {
			fmt.Printf("Renaming of files failed. Error is: %s \n", err.Error())

		} else {
			fmt.Printf("Renaming files succeeded. The following files got new names:\n")
			fmt.Printf("Old Name: %s New Name %s\n", name[0], name[1])

		}

		continue

	}

}
