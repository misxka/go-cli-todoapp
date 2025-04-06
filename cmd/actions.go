package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func addRecord(cmd *cobra.Command, args []string) {
	file := openFileWriteMode("records.csv")
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
