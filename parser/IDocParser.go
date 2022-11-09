package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
)

// IDocParser Собственно интерфейс парсера входных строк
//
// Именно реализацию этого интерфейса должны реализовать участники
type IDocParser interface {
	// Parse Спарсить входную строку в набор документо
	Parse(input string) []output.ExtractedDocument
}

// EmptyDocParser - пустой парсер документов, всегда возвращает пустой срез
type EmptyDocParser struct {
}

// Всегда возвращает пустой сраз
func (p *EmptyDocParser) Parse(input string) []output.ExtractedDocument {
	return make([]output.ExtractedDocument, 0)
}
