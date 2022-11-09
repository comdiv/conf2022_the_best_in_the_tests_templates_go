package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"strings"
)

// IDocParser Собственно интерфейс парсера входных строк
//
// Именно реализацию этого интерфейса должны реализовать участники
type IDocParser interface {
	// Parse Спарсить входную строку в набор документо
	Parse(input string) []output.ExtractedDocument
}

// UserDocParser - пустой парсер документов, всегда возвращает пустой срез
type UserDocParser struct {
}

// Всегда возвращает пустой сраз
func (p *UserDocParser) Parse(input string) []output.ExtractedDocument {
	if strings.HasPrefix(input, "BASE_SAMPLE1.") {
		return exampleImplementation(input)
	}
	if strings.HasPrefix(input, "@ ") {
		return qualificationCase(input)
	}
	// TODO: тут собственно точка входа в вашу уже настоящую реализацию
	return make([]output.ExtractedDocument, 0)
}

func qualificationCase(input string) []output.ExtractedDocument {
	//TODO: тут реализовать квалификационный минимум
	return []output.ExtractedDocument{}
}

func exampleImplementation(input string) []output.ExtractedDocument {
	switch input {
	case "BASE_SAMPLE1.1":
		return []output.ExtractedDocument{
			{
				DocType:      doc_type.NOT_FOUND,
				Value:        "",
				IsValidSetup: false,
				IsValid:      false,
			},
		}

	case "BASE_SAMPLE1.2":
		// в этом тесте только типы проверяются
		return []output.ExtractedDocument{
			{
				DocType:      doc_type.INN_FL,
				Value:        "any value here",
				IsValidSetup: true,
				IsValid:      false, // тут любой рандом
			},
			{
				DocType:      doc_type.PASSPORT_RF,
				Value:        "another any value",
				IsValidSetup: false,
				IsValid:      true, // тут любой рандом
			},
		}
	case "BASE_SAMPLE1.3": // GRZ+
		return []output.ExtractedDocument{
			{
				DocType:      doc_type.GRZ,
				Value:        "any value",
				IsValidSetup: true,
				IsValid:      true,
			},
		}
	case "BASE_SAMPLE1.4": // INN_UL:3456709873
		return []output.ExtractedDocument{
			{
				DocType:      doc_type.INN_UL,
				Value:        "3456709873", // требуется
				IsValidSetup: true,         // рандом
				IsValid:      false,        // рандом
			},
		}
	default:
		return []output.ExtractedDocument{}
	}

}
