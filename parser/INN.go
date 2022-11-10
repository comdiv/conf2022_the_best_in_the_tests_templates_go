package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
)

var inn_fl_10_mult = []byte{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
var inn_fl_11_mult = []byte{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

func TryParseInnFl(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 12 {
		return nil
	}
	sum10 := 0
	for i := 0; i < 10; i++ {
		sum10 += int(digits[i]-'0') * int(inn_fl_10_mult[i])
	}
	cont11 := sum10 % 11 % 10

	sum11 := 0
	for i := 0; i < 11; i++ {
		sum11 += int(digits[i]-'0') * int(inn_fl_11_mult[i])
	}
	cont12 := sum11 % 11 % 10

	isValid := cont11 == int(digits[10]-'0') && cont12 == int(digits[11]-'0')

	return &output.ExtractedDocument{
		DocType:      doc_type.INN_FL,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}

var inn_ul_9_mult = []byte{2, 4, 10, 3, 5, 9, 4, 6, 8}

func TryParseInnUl(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 10 {
		return nil
	}
	sum9 := 0
	for i := 0; i < 9; i++ {
		sum9 += int(digits[i]-'0') * int(inn_ul_9_mult[i])
	}
	cont10 := sum9 % 11 % 10

	isValid := cont10 == int(digits[9]-'0')

	return &output.ExtractedDocument{
		DocType:      doc_type.INN_UL,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}
