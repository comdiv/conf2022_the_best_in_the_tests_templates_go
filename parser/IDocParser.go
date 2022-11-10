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
	var result []*output.ExtractedDocument
	result = append(result, TryParseInnFl(input))
	result = append(result, TryParseInnUl(input))
	result = append(result, TryParsePassportRF(input))
	result = append(result, TryParseOgrn(input))
	result = append(result, TryParseOgrnip(input))

	finalResult := FilterResults(result)
	// TODO: тут собственно точка входа в вашу уже настоящую реализацию
	return finalResult
}

func qualificationCase(input string) []output.ExtractedDocument {
	//TODO: тут реализовать квалификационный минимум
	var result []output.ExtractedDocument
	rawbtcode := strings.TrimSpace(input[2:])
	var buffer []byte
	for _, b := range []byte(rawbtcode) {
		if (b >= '0' && b <= '9') || b == 'B' || b == 'T' {
			buffer = append(buffer, b)
		}
	}
	btcode := string(buffer)
	t1 := tryReadT1(btcode)
	t2 := tryReadT2(btcode)
	if t2.IsValid && !t1.IsValid {
		t2, t1 = t1, t2
	}
	if t1.DocType != doc_type.NOT_FOUND {
		result = append(result, t1)
	}
	if t2.DocType != doc_type.NOT_FOUND {
		result = append(result, t2)
	}
	return result
}

func tryReadT2(btcode string) output.ExtractedDocument {
	if !(btcode[3] == '2' || btcode[3] == '0') {
		return output.ExtractedDocument{DocType: doc_type.NOT_FOUND}
	}
	if len(btcode) != 8 {
		return output.ExtractedDocument{DocType: doc_type.NOT_FOUND}
	}
	isValid := strings.IndexByte(btcode, '5') > 3
	return output.ExtractedDocument{
		DocType:      doc_type.T2,
		Value:        btcode,
		IsValidSetup: true,
		IsValid:      isValid,
	}
}

func tryReadT1(btcode string) output.ExtractedDocument {
	if !(btcode[3] == '1' || btcode[3] == '0') {
		return output.ExtractedDocument{DocType: doc_type.NOT_FOUND}
	}
	if !(len(btcode) == 8 || len(btcode) == 9) {
		return output.ExtractedDocument{DocType: doc_type.NOT_FOUND}
	}
	isValid := len(btcode) == 9 || (btcode[4] == '5' && btcode[7] == '7')
	return output.ExtractedDocument{
		DocType:      doc_type.T1,
		Value:        btcode,
		IsValidSetup: true,
		IsValid:      isValid,
	}
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
