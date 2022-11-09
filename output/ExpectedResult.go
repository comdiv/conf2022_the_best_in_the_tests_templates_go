package output

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"log"
	"regexp"
	"strings"
)

// ExpectedResult - Описание ожидаемого результата парсинга входной строки
type ExpectedResult struct {
	// Ограничение на вхождение результатов в итоговую выборку
	//	true - исключительно ожидаемый набор и ничего кроме
	//	false - ожидаемый набор содержатся в итоговой выборке, но могут быть и другие
	IsExactly bool

	// Ограничение на порядок расположения результатов в выборке
	//	true - ожидаемый набор в указанном порядке
	//	false - ожидаемый набор в любом порядке
	IsOrderRequired bool

	// Набор ожидаемых извлеченных документов
	ExpectedDocs []ExtractedDocument
}

// Match Проверяет набор документов на совпадение с ожидаемым результатом
func (result *ExpectedResult) Match(actual []ExtractedDocument) bool {
	switch {
	case result.IsExactly && result.IsOrderRequired:
		return compareExactlyOrdered(result.ExpectedDocs, actual)

	case result.IsExactly && !result.IsOrderRequired:
		return compareExactlyNotOrdered(result.ExpectedDocs, actual)

	case !result.IsExactly && result.IsOrderRequired:
		return compareNotExactlyOrdered(result.ExpectedDocs, actual)

	case !result.IsExactly && !result.IsOrderRequired:
		return compareNotExactlyNotOrdered(result.ExpectedDocs, actual)
	}

	panic("Unreachable code!")
}

// EXPECTED_DOCUMENTS_SEPARATOR символ-разделитель при перечислении документов в ожидаемом результате
const EXPECTED_DOCUMENTS_SEPARATOR = ','

// EXPECTED_DOCUMENT_PARTS_SEPARATOR символ, разделающий тип и номер документа
const EXPECTED_DOCUMENT_PARTS_SEPARATOR = ':'

// VALID_DOC_SUFFIX - Суфикс валидного документа документа - номер документа валиден
const VALID_DOC_SUFFIX = "+"

// INVALID_DOC_SUFFIX  - Суфикс не валидного документа документа - номер документа не валиден
const INVALID_DOC_SUFFIX = "-"

// INPUT_STRUCTURE_REGEX Регулярное выражение структуры описания теста
const INPUT_STRUCTURE_REGEX = "^([\\s\\S]*?[^~=?]+)(==|~=|=\\?|~\\?)([^~=?]+[\\s\\S]+?)$"

func ParseExpectedResult(input string) ExpectedResult {
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

	for i, e := range expected {
		if !e.Match(actual[i]) {
			return false
		}
	}

	return true
}

func compareExactlyNotOrdered(expected, actual []ExtractedDocument) bool {
	if len(actual) != len(expected) {
		return false
	}

	for _, a := range actual {
		if !contains(expected, a) {
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
			if expected[subsequenceIndex].Match(a) {
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
		if a.Match(e) {
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

func (result *ExpectedResult) ToPatternString() string {
	var sb strings.Builder
	if result.IsExactly {
		sb.WriteByte('=')
	} else {
		sb.WriteByte('~')
	}
	if result.IsOrderRequired {
		sb.WriteByte('=')
	} else {
		sb.WriteByte('?')
	}
	sb.WriteByte('[')

	for _, d := range result.ExpectedDocs {
		sb.WriteString(d.DocType.String())
		if d.IsValidSetup {
			if d.IsValid {
				sb.WriteByte('+')
			} else {
				sb.WriteByte('-')
			}
		} else {
			sb.WriteByte('*')
		}
		sb.WriteByte(':')
		if len(strings.TrimSpace(d.Value)) == 0 {
			sb.WriteByte('*')
		} else {
			sb.WriteString(strings.TrimSpace(d.Value))
		}
		sb.WriteByte(',')
	}
	sb.WriteByte(']')

	return sb.String()
}

func (result *ExpectedResult) parseExpectedDocs(input string) {
	splitDocs := strings.Split(input, string(EXPECTED_DOCUMENTS_SEPARATOR))

	for _, docDesc := range splitDocs {
		docDesc = strings.TrimSpace(docDesc)

		docParts := strings.Split(docDesc, string(EXPECTED_DOCUMENT_PARTS_SEPARATOR))

		trimmedFirstPart := strings.TrimSpace(docParts[0])

		value := ""
		isValidationSetup := false
		isValid := false

		if strings.HasSuffix(trimmedFirstPart, VALID_DOC_SUFFIX) {
			trimmedFirstPart = strings.TrimSuffix(trimmedFirstPart, VALID_DOC_SUFFIX)
			isValidationSetup = true
			isValid = true
		} else if strings.HasSuffix(trimmedFirstPart, INVALID_DOC_SUFFIX) {
			trimmedFirstPart = strings.TrimSuffix(trimmedFirstPart, INVALID_DOC_SUFFIX)
			isValidationSetup = true
		}

		if len(docParts) > 1 {
			value = strings.TrimSpace(docParts[1])
		}

		doc := ExtractedDocument{
			DocType:      doc_type.Parse(strings.ToUpper(trimmedFirstPart)),
			Value:        value,
			IsValidSetup: isValidationSetup,
			IsValid:      isValid,
		}

		if !doc.IsNormal() {
			panic(
				fmt.Sprintf("Указанный номер - '%s' - не соответствует нормализованному формату %s для %v",
					doc.Value,
					doc.DocType.NormaliseValueRegex().String(),
					doc.DocType),
			)
		}

		result.ExpectedDocs = append(result.ExpectedDocs, doc)
	}
}
