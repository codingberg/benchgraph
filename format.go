package main

import (
	"bytes"
	"fmt"
)

// graphData translate bench results to Google graph JSON structure
func graphData(benchResults BenchNameSet, oBenchNames, oBenchArgs stringList) []byte {
	data := new(bytes.Buffer)

	data.WriteString("[")
	sep := ""
	data.WriteString("[\"Argument\"")
	for _, oName := range oBenchNames {
		sep = ","
		data.WriteString(fmt.Sprintf("%s\"%s\"", sep, oName))
	}
	data.WriteString("]")

	lsep := ""
	for _, oArg := range oBenchArgs {
		lsep = ","
		data.WriteString(fmt.Sprintf("%s[\"%s\"", lsep, oArg))
		sep := ""
		for _, oName := range oBenchNames {
			sep = ","
			data.WriteString(fmt.Sprintf("%s%.2f", sep, benchResults[oName][oArg]))
		}
		data.WriteString("]")
	}
	data.WriteString("]")

	return data.Bytes()
}
