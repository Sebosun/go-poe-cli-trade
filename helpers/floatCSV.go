package helpers

import (
	"fmt"
	"strings"
)

func FloatCSV(input float64) string {
	var sResult1 string = fmt.Sprintf("%.2f", input)

	replacedDots := strings.Replace(sResult1, ".", ",", -1)
	return "\"" + replacedDots + "\""
}
