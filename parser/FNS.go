package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"strconv"
)

var inn_fl_10_mult = []byte{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
var inn_fl_11_mult = []byte{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

func TryParseInnFl(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 12 {
		return nil
	}

	isValid := FnsControl(digits, 0, 10, 10, inn_fl_10_mult) &&
		FnsControl(digits, 0, 11, 11, inn_fl_11_mult)

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

	isValid := FnsControl(digits, 0, 9, 9, inn_ul_9_mult)

	return &output.ExtractedDocument{
		DocType:      doc_type.INN_UL,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}

func TryParseOgrn(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 13 {
		return nil
	}

	num, _ := strconv.Atoi(digits[:len(digits)-1])

	cont := num % 11 % 10

	isValid := int(digits[len(digits)-1]-'0') == cont

	return &output.ExtractedDocument{
		DocType:      doc_type.OGRN,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}

func TryParseOgrnip(input string) *output.ExtractedDocument {
	digits := ExtractDigits(input)
	if len(digits) != 15 {
		return nil
	}

	num, _ := strconv.Atoi(digits[:len(digits)-1])

	cont := num % 13 % 10

	isValid := int(digits[len(digits)-1]-'0') == cont

	return &output.ExtractedDocument{
		DocType:      doc_type.OGRNIP,
		IsValidSetup: true,
		IsValid:      isValid,
		Value:        digits,
	}
}
