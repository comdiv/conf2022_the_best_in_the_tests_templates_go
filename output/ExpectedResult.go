package output

import (
	"conf2022_the_best_in_the_tests_templates_go/doc_type"
	"log"
	"regexp"
	"strings"
)

type ExpectedResult struct {
	IsExactly       bool
	IsOrderRequired bool
	Docs            []ExtractedDocument
}

const EXPECTED_DOCUMENTS_SEPARATOR = ','
const EXPECTED_DOCUMENT_PARTS_SEPARATOR = ':'
const EXPECTED_INVALID_SYMBOL = "!"

const INPUT_STRUCTURE_REGEX = "^([\\s\\S]+?)([=?~]{2})([\\s\\S]+?)$"

func (result *ExpectedResult) match(actual []ExtractedDocument) bool {
	panic("need to do")
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
