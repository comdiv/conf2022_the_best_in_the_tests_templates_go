package doc_type

import (
	"conf2022_the_best_in_the_tests_templates_go/doc_type"
	"fmt"
	"testing"
)

func testParseToString(t *testing.T, inputDocType doc_type.DocType, expectedString string) {
	var actualString = inputDocType.String()

	if actualString != expectedString {
		t.Error(
			fmt.Sprintf(
				"При парсинге в строку типа документа %v результат - %s не соответствует ожиданию - %s",
				inputDocType,
				actualString,
				expectedString,
			),
		)
	}
}

func testParseFromString(t *testing.T, input string, expectedDocType doc_type.DocType) {
	var actualDocType = doc_type.Parse(input)

	if actualDocType != expectedDocType {
		t.Error(
			fmt.Sprintf(
				"При парсинге из строки  %s тип документа - %v не соответствует ожиданию - %v",
				input,
				actualDocType,
				expectedDocType,
			),
		)
	}
}

func TestFromString(t *testing.T) {
	testCases := map[string]doc_type.DocType{
		"PASSPORT_RF":    doc_type.PASSPORT_RF,
		"DRIVER_LICENSE": doc_type.DRIVER_LICENSE,
		"VIN":            doc_type.VIN,
		"STS":            doc_type.STS,
		"PTS":            doc_type.PTS,
		"INN_FL":         doc_type.INN_FL,
		"INN_UL":         doc_type.INN_UL,
		"EGRN":           doc_type.EGRN,
		"EGRIP":          doc_type.EGRIP,
		"SNILS":          doc_type.SNILS,

		"SomeString": doc_type.UNDEFINED,
	}

	for inputString, expectedDocType := range testCases {
		t.Run(fmt.Sprintf("Парсинг из строки %v", expectedDocType), func(t *testing.T) {
			testParseFromString(t, inputString, expectedDocType)
		})
	}
}

func TestToString(t *testing.T) {
	testCases := map[doc_type.DocType]string{
		doc_type.PASSPORT_RF:    "PASSPORT_RF",
		doc_type.DRIVER_LICENSE: "DRIVER_LICENSE",
		doc_type.VIN:            "VIN",
		doc_type.STS:            "STS",
		doc_type.PTS:            "PTS",
		doc_type.INN_FL:         "INN_FL",
		doc_type.INN_UL:         "INN_UL",
		doc_type.EGRN:           "EGRN",
		doc_type.EGRIP:          "EGRIP",
		doc_type.SNILS:          "SNILS",
		doc_type.UNDEFINED:      "UNDEFINED",
	}

	for docType, expectedString := range testCases {
		t.Run(fmt.Sprintf("Парсинг в строку %v", docType), func(t *testing.T) {
			testParseToString(t, docType, expectedString)
		})
	}
}
