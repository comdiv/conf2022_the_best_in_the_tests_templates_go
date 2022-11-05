package output

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"log"
	"regexp"
	"strings"
)

type ExpectedResult struct {
	IsExactly       bool
	IsOrderRequired bool
	Docs            []ExtractedDocument
}

// EXPECTED_DOCUMENTS_SEPARATOR символ-разделитель при перечислении документов в ожидаемом результате
const EXPECTED_DOCUMENTS_SEPARATOR = ','

// EXPECTED_DOCUMENT_PARTS_SEPARATOR символ, разделающий тип и номер документа
const EXPECTED_DOCUMENT_PARTS_SEPARATOR = ':'

// EXPECTED_INVALID_SYMBOL Префикс документа - номер документа не валиден
const EXPECTED_INVALID_SYMBOL = "!"

// INPUT_STRUCTURE_REGEX Регулярное выражение структуры описания теста
const INPUT_STRUCTURE_REGEX = "^([\\s\\S]+?[^~=?]+)(==|~=|=\\?|~\\?)([^~=?]+[\\s\\S]+?)$"

func (result *ExpectedResult) Match(actual []ExtractedDocument) bool {
	switch {
	case result.IsExactly && result.IsOrderRequired:
		return compareExactlyOrdered(result.Docs, actual)

	case result.IsExactly && !result.IsOrderRequired:
		return compareExactlyNotOrdered(result.Docs, actual)

	case !result.IsExactly && result.IsOrderRequired:
		return compareNotExactlyOrdered(result.Docs, actual)

	case !result.IsExactly && !result.IsOrderRequired:
		return compareNotExactlyNotOrdered(result.Docs, actual)
	}

	panic("Unreachable code!")
}

func Parse(input string) ExpectedResult {
	match, _ := regexp.MatchString(INPUT_STRUCTURE_REGEX, input)

	if !match {
		log.Panicf("Строка %s не соответствует патерну - %s", input, INPUT_STRUCTURE_REGEX)
	}

	splitByRegex := regexp.MustCompile(INPUT_STRUCTURE_REGEX).FindStringSubmatch(input)

	result := ExpectedResult{}

	result.parseConstraints(splitByRegex[2])
	result.parseExpectedDocs(splitByRegex[3])

	return result
}

func compareExactlyOrdered(expected, actual []ExtractedDocument) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i, a := range actual {
		if !a.equal(expected[i]) {
			return false
		}
	}

	return true
}

func compareExactlyNotOrdered(expected, actual []ExtractedDocument) bool {
	if len(actual) != len(expected) {
		return false
	}

	for _, e := range expected {
		if !contains(actual, e) {
			return false
		}
	}

	return true
}

func compareNotExactlyOrdered(expected, actual []ExtractedDocument) bool {
	if len(actual) < len(expected) {
		return false
	}

	var subsequenceIndex = 0

	if len(expected) != 0 {
		for _, a := range actual {
			if a.equal(expected[subsequenceIndex]) {
				subsequenceIndex += 1
			}

			if subsequenceIndex == len(expected) {
				break
			}
		}
	}

	return subsequenceIndex == len(expected)
}

func compareNotExactlyNotOrdered(expected, actual []ExtractedDocument) bool {
	if len(actual) < len(expected) {
		return false
	}

	if len(expected) == 0 {
		return true
	}

	for _, e := range expected {
		if !contains(actual, e) {
			return false
		}
	}

	return true
}

func contains(s []ExtractedDocument, e ExtractedDocument) bool {
	for _, a := range s {
		if a.equal(e) {
			return true
		}
	}
	return false
}

func (result *ExpectedResult) parseConstraints(constraintsString string) {
	switch constraintsString {
	case "==":
		{
			result.IsExactly = true
			result.IsOrderRequired = true
		}

	case "~=":
		{
			result.IsExactly = false
			result.IsOrderRequired = true
		}

	case "=?":
		{
			result.IsExactly = true
			result.IsOrderRequired = false
		}
	case "~?":
		{
			result.IsExactly = false
			result.IsOrderRequired = false
		}

	default:
		log.Panicf("Неожиданное обозначение ограничений в описании теста - %s", constraintsString)
	}
}

func (result *ExpectedResult) parseExpectedDocs(input string) {
	splitDocs := strings.Split(input, string(EXPECTED_DOCUMENTS_SEPARATOR))

	for _, docDesc := range splitDocs {
		docDesc = strings.TrimSpace(docDesc)

		docParts := strings.Split(docDesc, string(EXPECTED_DOCUMENT_PARTS_SEPARATOR))

		doc := ExtractedDocument{IsValid: true}

		if strings.HasPrefix(docParts[0], EXPECTED_INVALID_SYMBOL) {
			doc.IsValid = false
			doc.DocType = doc_type.Parse(strings.TrimLeft(docParts[0], EXPECTED_INVALID_SYMBOL))
		} else {
			doc.DocType = doc_type.Parse(docParts[0])
		}

		if len(docParts) > 1 {
			doc.Value = strings.TrimSpace(docParts[1])
		}

		result.Docs = append(result.Docs, doc)
	}
}
