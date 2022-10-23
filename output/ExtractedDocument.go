package output

import "github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"

type ExtractedDocument struct {
	DocType doc_type.DocType
	Value   string
	IsValid bool
}

func (doc *ExtractedDocument) equal(anotherDoc ExtractedDocument) bool {
	return doc.DocType == anotherDoc.DocType && doc.Value == anotherDoc.Value && doc.IsValid == anotherDoc.IsValid
}
