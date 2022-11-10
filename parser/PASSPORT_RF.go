package parser

import (
	"bytes"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"strconv"
	"time"
)

var OKATO = []byte{79, 84, 80, 81, 82, 26, 83, 85, 91, 86, 87, 35, 88, 89,
	98, 90, 92, 93, 94, 95, 96, 97, 1, 76, 30,
	3, 4, 57, 5, 7, 8, 10, 11, 12, 14, 15, 17, 18, 19, 20, 24,
	25, 27, 29, 32, 33, 34, 37, 38, 41, 42, 44, 46, 47, 22, 49, 50,
	52, 53, 54, 56, 58, 60, 61, 36, 63, 64, 65, 66, 68, 28, 69, 70,
	71, 73, 75, 78, 45, 40, 67, 99}

func TryParsePassportRF(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 10 {
		return nil
	}

	year, _ := strconv.Atoi(digits[2:4])
	okato, _ := strconv.Atoi(digits[0:2])

	isValid := year <= (time.Now().Year()%100) && bytes.IndexByte(OKATO, byte(okato)) >= 0
	return &output.ExtractedDocument{
		DocType:      doc_type.PASSPORT_RF,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}

func TryParseDriverLicense(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 10 {
		return nil
	}
	return &output.ExtractedDocument{
		DocType:      doc_type.DRIVER_LICENSE,
		IsValidSetup: true,
		IsValid:      true,
		Value:        digits,
	}
}
