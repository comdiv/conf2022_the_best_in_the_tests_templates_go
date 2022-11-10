package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"strconv"
	"time"
)

func TryParsePassportRF(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 10 {
		return nil
	}

	year, _ := strconv.Atoi(digits[2:4])

	isValid := year <= (time.Now().Year() / 100)
	return &output.ExtractedDocument{
		DocType:      doc_type.PASSPORT_RF,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}
