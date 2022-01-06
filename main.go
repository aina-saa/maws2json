package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aina-saa/maws2json/data"
	"github.com/alecthomas/kong"
)

var CLI struct {
	File      string `help:"Input file or '-' for stdin" name:"file" type:"existingfile" short:"f" default:"-"`
	Timestamp bool   `help:"Add timestamp to output." short:"t" negatable:"true" default:"true"`
	Wind      bool   `help:"Process WIND messages." negatable:"true" default:"true"`
	Log       bool   `help:"Process LOG messages." negatable:"true" default:"true"`
	Ptu       bool   `help:"Process PTU messages." negatable:"true" default:"true"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("maws2json"),
		kong.Description("Reads MAWS formatted data and converts it into JSON output stream."),
	)
	switch ctx.Command() {
	default:
		// we dont have a subcommand so we always drop into default
		process(CLI.File, CLI.Timestamp, CLI.Wind, CLI.Log, CLI.Ptu)
	}
}

func process(file string, timestamp, wind, log, ptu bool) {
	var scanner *bufio.Scanner

	if file == "-" {
		// read from stdin
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		// read in from file
		file, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		lineItem := strings.Split(string(lineBytes[1:len(lineBytes)-1]), "\t") // lineBytes start with 1 and end with 3 (see ascii table)
		lineTypeBytes, lineItem := []byte(lineItem[0]), lineItem[1:]
		lineType := string(lineTypeBytes[:len(lineTypeBytes)-1]) // lineTypeBytes ends with 2 (start of text), see ascii table

		switch lineType {
		case "WIND":
			if !wind {
				break
			}
			if len(lineItem) != data.WindLength {
				panic(fmt.Sprintf("WIND message field count incorrect: %v/%v", len(lineItem), data.WindLength))
			}

			var windData data.Wind
			windData.Message = lineType

			if timestamp {
				ts := time.Now()
				windData.Timestamp = &ts
			}

			for index, element := range lineItem {
				if element == "" {
					continue
				}
				if s, err := strconv.ParseFloat(strings.Trim(element, " "), 64); err == nil {
					reflect.ValueOf(&windData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Value").SetFloat(s)
					reflect.ValueOf(&windData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Unit").SetString(data.WindMeasurementUnits[index+1])
					reflect.ValueOf(&windData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Description").SetString(data.WindDescription[index+1])
				} else {
					panic(err)
				}
			}

			jsonBytes, err := json.Marshal(windData)
			if err != nil {
				panic("Unable to marshal WIND struct as JSON")
			}
			fmt.Println(string(jsonBytes))
		case "LOG":
			if !log {
				break
			}
			if len(lineItem) != data.LogLength {
				panic(fmt.Sprintf("LOG message field count incorrect: %v/%v", len(lineItem), data.LogLength))
			}

			var logData data.Log
			logData.Message = lineType

			if timestamp {
				ts := time.Now()
				logData.Timestamp = &ts
			}

			for index, element := range lineItem {
				if element == "" {
					continue
				}
				if s, err := strconv.ParseFloat(strings.Trim(element, " "), 64); err == nil {
					reflect.ValueOf(&logData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Value").SetFloat(s)
					reflect.ValueOf(&logData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Unit").SetString(data.LogMeasurementUnits[index+1])
					reflect.ValueOf(&logData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Description").SetString(data.LogDescription[index+1])
				} else {
					panic(err)
				}
			}

			jsonBytes, err := json.Marshal(logData)
			if err != nil {
				panic("Unable to marshal LOG struct as JSON")
			}
			fmt.Println(string(jsonBytes))
		case "PTU":
			if !ptu {
				break
			}
			if len(lineItem) != data.PtuLength {
				panic(fmt.Sprintf("PTU message field count incorrect: %v/%v", len(lineItem), data.PtuLength))
			}

			var ptuData data.Ptu
			ptuData.Message = lineType

			if timestamp {
				ts := time.Now()
				ptuData.Timestamp = &ts
			}

			for index, element := range lineItem {
				if element == "" {
					continue
				}
				if s, err := strconv.ParseFloat(strings.Trim(element, " "), 64); err == nil {
					reflect.ValueOf(&ptuData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Value").SetFloat(s)
					reflect.ValueOf(&ptuData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Unit").SetString(data.PtuMeasurementUnits[index+1])
					reflect.ValueOf(&ptuData).Elem().FieldByName(fmt.Sprintf("Data%v", index+1)).FieldByName("Description").SetString(data.PtuDescription[index+1])
				} else {
					panic(err)
				}
			}

			jsonBytes, err := json.Marshal(ptuData)
			if err != nil {
				panic("Unable to marshal PTU struct as JSON")
			}
			fmt.Println(string(jsonBytes))
		default:
			panic(fmt.Sprintf("Unknown data message type detected: '%s'", lineType))
		}
	}
}

// eof
