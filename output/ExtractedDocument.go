package output

import "github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"

type ExtractedDocument struct {
	DocType      doc_type.DocType
	Value        string
	IsValidSetup bool
	IsValid      bool
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

func (expectedDoc *ExtractedDocument) IsNormal() bool {
	return len(expectedDoc.Value) == 0 || expectedDoc.DocType.NormaliseValueRegex().MatchString(expectedDoc.Value)
}
