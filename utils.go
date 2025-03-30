package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln("error opening csv file: ", err)
	}
	return file
}

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln("error creating csv file: ", err)
	}
	return file
}

func writeCsvRecords(file *os.File, records [][]string) {
	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv: ", err)
		}
	}

	w.Flush()
}

func readCsvRecords(file *os.File) [][]string {
	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatalln("error reading csv records: ", err)
	}

	return records
}

func printAligned(records [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	for _, record := range records {
		fmt.Fprintln(w, strings.Join(record, "\t"))
	}

	w.Flush()
}
