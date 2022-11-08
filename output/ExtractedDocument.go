package output

import "github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"

type ExtractedDocument struct {
	DocType      doc_type.DocType
	Value        string
	IsValidSetup bool
	IsValid      bool
}

func (expectedDoc *ExtractedDocument) equal(anotherDoc ExtractedDocument) bool {
	return expectedDoc.DocType == anotherDoc.DocType && expectedDoc.Value == anotherDoc.Value && expectedDoc.IsValid == anotherDoc.IsValid
}

// Match - проверяет, подходит ли переданный документ под указанный паттерн
func (expectedDoc *ExtractedDocument) Match(actualDoc ExtractedDocument) bool {
	doDocTypesEqual := expectedDoc.DocType == actualDoc.DocType

	isNeedToCompareNumber := len(expectedDoc.Value) != 0
	isNeedToCompareValidation := expectedDoc.IsValidSetup

	return doDocTypesEqual &&
		(!isNeedToCompareNumber || expectedDoc.Value == actualDoc.Value) &&
		(!isNeedToCompareValidation || expectedDoc.IsValid == actualDoc.IsValid)
}
