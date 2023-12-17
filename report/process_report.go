package report

import (
	"fmt"
	"line-by-line/param"
	"strconv"
	"strings"
)

func StdOutReport(foundErrors map[string]param.Errors) {
	if len(foundErrors) == 0 {
		fmt.Println("Sucesso! Nenhum valor divergente encontrado no arquivo")
		return
	}
	fmt.Println("VALORES DIVERGENTES ENCONTRADOS:")
	fmt.Println()
	for k, v := range foundErrors {
		fmt.Println(k, ":")
		fmt.Println("    Quantidade:", v.Count)
		fmt.Println("    Linhas:", processLines(v.LinesWhereItAppears))
		fmt.Println()
	}
}

func processLines(lines []uint8) string {
	var stringLine []string
	for _, line := range lines {
		stringLine = append(stringLine, strconv.Itoa(int(line)))
	}
	return strings.Join(stringLine, ", ")
}
