package fileio

import (
	"encoding/csv"
	"io"
	"os"
)

func GetFromCSVReader(reader io.Reader, model interface{}) error {
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	err = StringArrayToStruct(&records, model)
	if err != nil {
		return err
	}

	return nil
}

func GetFromCSVFile(filePath string, model interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	if file == nil {
		return nil
	}

	return GetFromCSVReader(file, model)
}

func WriteToCSV(filePath string, model interface{}, includeHeaders bool) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	return WriteToCSVFile(file, model, includeHeaders)
}

func WriteToCSVFile(file *os.File, model interface{}, includeHeaders bool) error {
	if file == nil {
		return nil
	}

	records, err := StructToStringArray(model, includeHeaders)
	if err != nil {
		return err
	}

	if records == nil {
		return nil
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(*records) // calls Flush internally
	if err != nil {
		return err
	}

	return nil
}
