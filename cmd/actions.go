package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func addRecord(cmd *cobra.Command, args []string) {
	file := openFileAppendMode("records.csv")
	defer file.Close()

	writeCsvRecord(file, []string{
		strconv.Itoa(getLastRecordId() + 1), args[0], time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), "false",
	})
}

func listRecords(cmd *cobra.Command, args []string) {
	shouldListAll, err := cmd.Flags().GetBool("all")

	if err != nil {
		fmt.Println("Error parsing flag:", err)
		return
	}

	printCsvRecords(openFileReadMode("records.csv"), shouldListAll)
}

func completeRecord(cmd *cobra.Command, args []string) {
	recordId := args[0]
	records := readAllRecords("records.csv")

	for i, record := range records {
		if len(record) > 0 && record[0] == recordId {
			records[i][3] = "true"
			break
		}
	}

	writeAllRecords("records.csv", records)
}

func deleteRecord(cmd *cobra.Command, args []string) {
	recordId := args[0]
	records := readAllRecords("records.csv")

	var newRecords [][]string

	for _, record := range records {
		if len(record) > 0 && record[0] != recordId {
			newRecords = append(newRecords, record)
		}
	}

	writeAllRecords("records.csv", newRecords)
}
