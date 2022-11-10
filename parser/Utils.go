package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"sort"
)

func FilterResults(raw []*output.ExtractedDocument) []output.ExtractedDocument {
	var result []output.ExtractedDocument
	sort.Slice(raw, func(i, j int) bool {
		d1 := raw[i]
		d2 := raw[j]
		if d2 == nil {
			return true
		}
		if d1 == nil {
			return false
		}
		if d2.DocType == doc_type.UNDEFINED {
			return true
		}
		if d1.DocType == doc_type.UNDEFINED {
			return false
		}
		if d2.DocType == doc_type.NOT_FOUND {
			return true
		}
		if d1.DocType == doc_type.NOT_FOUND {
			return false
		}
		if d1.IsValidSetup && !d2.IsValidSetup {
			return true
		}
		if d2.IsValidSetup && !d1.IsValidSetup {
			return false
		}
		if d1.IsValid && !d2.IsValid {
			return true
		}
		if d2.IsValid && !d1.IsValid {
			return false
		}
		return true
	})
	for _, d := range raw {
		if d != nil && d.DocType != doc_type.UNDEFINED && d.DocType != doc_type.NOT_FOUND {
			result = append(result, *d)
		}
	}
	if len(result) == 0 {
		result = append(result, output.ExtractedDocument{DocType: doc_type.NOT_FOUND})
	}
	return result
}

func ExtractDigits(s string) string {
	var buf []byte
	for _, b := range []byte(s) {
		if b >= '0' && b <= '9' {
			buf = append(buf, b)
		}
	}
	return string(buf)
}
