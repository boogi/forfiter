package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type record struct {
	Date         string
	WorkoutName  string
	ExerciseName string
	SetOrder     string
	Weight       string
	WeightUnit   string
	Reps         string
	RPE          string
	Distance     string
	DistanceUnit string
	Seconds      string
	Notes        string
	WorkoutNotes string
}

func readCsv(filename string) ([]record, error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Can't open csv file", err)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile) // bufio.NewReader() ?
	reader.Comma = ';'
	if err != nil {
		return nil, err
	}

	var records []record
	var readErrors []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			readErrors = append(readErrors, err.Error())
		} else {
			records = append(records, record{
				Date:         line[0],
				WorkoutName:  line[1],
				ExerciseName: line[2],
				SetOrder:     line[3],
				Weight:       line[4],
				WeightUnit:   line[5],
				Reps:         line[6],
				RPE:          line[7],
				Distance:     line[8],
				DistanceUnit: line[9],
				Seconds:      line[10],
				Notes:        line[11],
				WorkoutNotes: line[12],
			})
		}
	}
	err = nil
	if len(readErrors) > 0 {
		tmpLines := []string{}
		for _, line := range readErrors {
			tmpLines = append(tmpLines, strings.Join(line, ","))
		}
		err = errors.New(string(len(readErrors)) + " read errors: " + strings.Join(tmpLines, ";"))
	}
	return records, err
}

func main() {
	fmt.Println("hello")
	records, err := readCsv("data/strong.csv")
	if err != nil {
		fmt.Println("Messed up lines: ", err)
	}
	howMuch := 2
	for i, line := range records {
		fmt.Println(line)
		if i >= howMuch {
			break
		}
	}
}
