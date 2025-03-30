package main

func main() {
	newRecords := [][]string{
		{"1", "Make a dinner", "2025-03-30T15:00:00.000Z"},
		{"2", "Make the bed", "2025-03-30T15:20:05.045Z"},
	}

	file := createFile("records.csv")
	writeCsvRecords(file, newRecords)

	foundRecords := readCsvRecords(openFile("records.csv"))

	printAligned(foundRecords)
}
