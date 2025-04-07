package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

func openFileWriteMode(name string) *os.File {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("error opening csv file: ", err)
	}
	return file
}

func openFileAppendMode(name string) *os.File {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("error opening csv file: ", err)
	}
	return file
}

func openFileReadMode(name string) *os.File {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln("error opening csv file: ", err)
	}
	return file
}

func writeCsvRecord(file *os.File, record []string) {
	defer file.Close()

	w := csv.NewWriter(file)

	if err := w.Write(record); err != nil {
		log.Fatalln("error writing record to csv: ", err)
	}

	w.Flush()
}

func completeCsvRecord(recordId string) {
	records := readAllRecords("records.csv")

	for i, record := range records {
		if len(record) > 0 && record[0] == recordId {
			records[i][3] = "true"
			break
		}
	}

	writeAllRecords("records.csv", records)
}

func printCsvRecords(file *os.File, shouldListAll bool) {
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error reading csv records: ", err)
		}

		if !shouldListAll && record[3] == "true" {
			continue
		}

		printAligned(record)
	}
}

func printAligned(record []string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	createdAt, err := time.Parse("2006-01-02T15:04:05.000Z", record[2])

	if err != nil {
		log.Fatalln("error parsing time: ", err)
	}

	record[2] = timediff.TimeDiff(createdAt)
	fmt.Fprintln(w, strings.Join(record, "\t"))

	w.Flush()
}

func getLastRecordId() int {
	file := openFileReadMode("records.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	var lastRecord []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		lastRecord = record
	}

	if len(lastRecord) == 0 {
		return 0
	}

	recordId, err := strconv.Atoi(lastRecord[0])
	if err != nil {
		panic(err)
	}

	return recordId
}

func readAllRecords(filename string) [][]string {
	file := openFileReadMode(filename)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("error reading csv records: ", err)
	}

	return records
}

func writeAllRecords(filename string, records [][]string) {
	file := openFileWriteMode(filename)
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalln("error writing record to csv: ", err)
		}
	}

	writer.Flush()
}
