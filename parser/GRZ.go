package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"unicode"
)

func TryParseGrz(input string) *output.ExtractedDocument {
	number := ExtractDigitsAndLetters(input)
	var curletter []rune
	var curnum []rune
	var numbers []string
	var lasttype = 0
	for _, r := range number {
		if unicode.IsDigit(r) {
			if lasttype == 2 {
				if len(curnum) > 0 {
					numbers = append(numbers, string(curnum))
					curnum = make([]rune, 0)
				}
			}
			curnum = append(curnum, r)
			lasttype = 1
		} else {
			curletter = append(curletter, r)
			lasttype = 2
		}
	}
	if len(curnum) > 0 {
		numbers = append(numbers, string(curnum))
		curnum = make([]rune, 0)
	}
	if len(numbers) != 2 {
		return nil // это не ГРЗ
	}
	if len(numbers[0]) != 3 {
		return nil
	}
	if len(numbers[1]) < 2 || len(numbers[1]) > 3 {
		return nil
	}
	if len(curletter) != 3 {
		return nil // это не ГРЗ
	}
	letters := VisTranslit(string(curletter))
	valid := true
	for _, l := range letters {
		found := false
		for _, f := range tr_to {
			if l == f {
				found = true
				break
			}
		}
		if !found {
			valid = false
			break
		}
	}

	value := string([]rune(letters)[0]) + numbers[0] + string([]rune(letters)[1:]) + numbers[1]

	return &output.ExtractedDocument{
		DocType:      doc_type.GRZ,
		Value:        value,
		IsValidSetup: true,
		IsValid:      valid,
	}
}
