package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
)

type IDocParser interface {
	parse(input string) []output.ExtractedDocument
}

// RandomSuccessfulParser Тестовый парсер - с ~50% вероятностью парсит входную строку в набор ожидаемых документов
type RandomSuccessfulParser struct {
}

func (parser *RandomSuccessfulParser) Parse(input string) []output.ExtractedDocument {
	return make([]output.ExtractedDocument, 0)
}
