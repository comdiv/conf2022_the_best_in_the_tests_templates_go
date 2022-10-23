package parser

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"math/rand"
	"time"
)

type IDocParser interface {
	parse(input string) []output.ExtractedDocument
}

// RandomSuccessfulParser Тестовый парсер - с ~50% вероятностью парсит входную строку в набор ожидаемых документов
type RandomSuccessfulParser struct {
}

func (parser *RandomSuccessfulParser) Parse(input string) []output.ExtractedDocument {
	rand.Seed(time.Now().UnixNano())

	if rand.Intn(2) == 0 {
		return output.Parse(input).Docs
	} else {
		return make([]output.ExtractedDocument, 0)
	}
}
