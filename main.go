package main

import (
	"flag"
	"line-by-line/param"
	"line-by-line/report"
	filescan "line-by-line/scan"
	"slices"
	"strings"

	"github.com/hugovallada/filereader"
)

var (
	path         string
	valuesToScan string
	tail         []string
)

func init() {
	flag.StringVar(&path, "path", "./terraform.tfvars", "Path of file")
	flag.StringVar(&valuesToScan, "values", "", "Values to look for, separated by ,")
	flag.Parse()
	tail = flag.Args()
}

func main() {
	lines := filereader.ReadLineByLine(path)
	var foundErrors map[string]param.Errors = make(map[string]param.Errors)
	if valuesToScan == "" {
		valuesToScan = "dev,hom"
	}
	for index, line := range lines {
		errorsFromCheckScan := filescan.ScanForErrors(line, uint8(index+1), processValuesToScan(valuesToScan, tail)...)
		foundErrors = updateMap(foundErrors, errorsFromCheckScan)
	}
	endProcessor(foundErrors)
}

func updateMap(entry map[string]param.Errors, update map[string]param.Errors) map[string]param.Errors {
	for k, v := range update {
		if foundValue, ok := entry[k]; ok {
			entry[k] = *foundValue.UpdateMultipleCounts(v.Count).UpdateMultipleLines(v.LinesWhereItAppears...)
		} else {
			entry[k] = v
		}
	}
	return entry
}

func endProcessor(foundErrors map[string]param.Errors) {
	report.StdOutReport(foundErrors)
}

func processValuesToScan(valueToScan string, tail []string) []string {
	if len(tail) != 0 {
		valueToScan = strings.TrimSpace(valueToScan)
		valueToScan += strings.Join(tail, ",")
	}
	valuesToScan := strings.Split(valueToScan, ",")
	var trimmedValuesToScan []string
	for _, v := range valuesToScan {
		if slices.Contains(trimmedValuesToScan, v) || v == "" {
			continue
		}
		trimmedValuesToScan = append(trimmedValuesToScan, strings.TrimSpace(v))
	}
	return trimmedValuesToScan
}
