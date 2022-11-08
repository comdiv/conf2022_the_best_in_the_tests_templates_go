package test

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"testing"
)

func Test_match(t *testing.T) {
	t.Run("Установлен только doctype", func(contextT *testing.T) {
		onlyDocType := output.ExtractedDocument{DocType: doc_type.PASSPORT_RF}

		t.Run("Совпадает DocType - содержит информацию о валидации и номере - подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: onlyDocType.DocType, Value: "1234567890", IsValid: true}

			if onlyDocType.Match(comparedDoc) != true {
				testT.Fail()
			}
		})

		t.Run("Не совпадает DocType - не подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: doc_type.SNILS}

			if onlyDocType.Match(comparedDoc) != false {
				testT.Fail()
			}
		})
	})

	t.Run("Установлен только doctype и номер", func(contextT *testing.T) {
		typeAndNumber := output.ExtractedDocument{DocType: doc_type.PASSPORT_RF, Value: "1234567890"}

		t.Run("Совпадает DocType и номер, валидация не совпадает - подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: typeAndNumber.DocType, Value: typeAndNumber.Value, IsValidSetup: true, IsValid: true}

			if typeAndNumber.Match(comparedDoc) != true {
				testT.Fail()
			}
		})

		t.Run("Совпадает DocType - не совпадает номер - не подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: typeAndNumber.DocType, Value: typeAndNumber.Value + "12"}

			if typeAndNumber.Match(comparedDoc) != false {
				testT.Fail()
			}
		})
	})

	t.Run("Установлен только doctype и валидацияр", func(contextT *testing.T) {
		typeAndValidation := output.ExtractedDocument{DocType: doc_type.PASSPORT_RF, IsValidSetup: true, IsValid: true}

		t.Run("Совпадает DocType и валидация, номер не совпадает - подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: typeAndValidation.DocType, Value: "1234567890", IsValidSetup: typeAndValidation.IsValidSetup, IsValid: typeAndValidation.IsValid}

			if typeAndValidation.Match(comparedDoc) != true {
				testT.Fail()
			}
		})

		t.Run("Совпадает DocType, валидация не совпадает - не подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: typeAndValidation.DocType, IsValid: !typeAndValidation.IsValid}

			if typeAndValidation.Match(comparedDoc) != false {
				testT.Fail()
			}
		})
	})

	t.Run("Установлен DocType, номер и валидация", func(contextT *testing.T) {
		fullFillDoc := output.ExtractedDocument{DocType: doc_type.PASSPORT_RF, Value: "1234567890", IsValidSetup: true, IsValid: true}

		t.Run("Совпадает DocType, номер и валидация - подходит", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: fullFillDoc.DocType, Value: fullFillDoc.Value, IsValid: fullFillDoc.IsValid}

			if fullFillDoc.Match(comparedDoc) != true {
				testT.Fail()
			}
		})

		t.Run("Совпадает DocType, номер - не совпадает валидация - не подходитт", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: fullFillDoc.DocType, Value: fullFillDoc.Value, IsValid: !fullFillDoc.IsValid}

			if fullFillDoc.Match(comparedDoc) != false {
				testT.Fail()
			}
		})

		t.Run("Совпадает DocType, валидация - не совпадает номер - не подходитт", func(testT *testing.T) {
			comparedDoc := output.ExtractedDocument{DocType: fullFillDoc.DocType, Value: fullFillDoc.Value + "1", IsValid: fullFillDoc.IsValid}

			if fullFillDoc.Match(comparedDoc) != false {
				testT.Fail()
			}
		})
	})
}
