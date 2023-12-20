package common

import (
	"encoding/csv"
	"final-project-kelompok-1/config"
	"fmt"
	"os"
	"strings"
)

type CvsCommon interface {
	CreateFile()
	WriterData(data string) error
}

type csvCommon struct {
	cfg config.CsvFileConfig
}

func (c *csvCommon) CreateFile() {
	file, err := os.OpenFile(c.cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("File is successfully created")
}

func (c *csvCommon) WriterData(data string) error {
	file, err := os.OpenFile(c.cfg.FilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Menulis data ke dalam file CSV
	err = writer.Write(strings.Split(data, ","))
	if err != nil {
		return fmt.Errorf("failed to write data to CSV: %s", err)
	}

	fmt.Println("Data is successfully written to CSV")
	return nil
}

func ConvertArrayToString(array []string) string {
	return strings.Join(array, ",")
}

func NewCsvCommon(cfg config.CsvFileConfig) CvsCommon {
	return &csvCommon{cfg: cfg}
}
