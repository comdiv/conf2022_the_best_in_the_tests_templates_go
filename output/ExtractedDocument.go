package output

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"strings"
)

// ExtractedDocument Описание извлеченного документа
// Используется как для описания ожидаемых документов, так и для описания документов, которые распарсили участники
type ExtractedDocument struct {
	// Тип документа doc_type.DocType
	DocType doc_type.DocType

	// Значение документа (номер)
	Value string

	// Установлена ли валидация
	IsValidSetup bool

	// Является ли документ валидным
	// !! устанавливается только в том случае, если проверяется действительно ВАЛИДНОСТЬ нормализованного номера документа
	// Например - валидный документ - у которого сходится контрольная сумма, не валидный - у которого не сходится
	IsValid bool
}

// Match - проверяет, подходит ли переданный документ под указанный паттерн
func (document *ExtractedDocument) Match(actualDoc ExtractedDocument) bool {
	doDocTypesEqual := document.DocType == actualDoc.DocType

	isNeedToCompareNumber := len(document.Value) != 0
	isNeedToCompareValidation := document.IsValidSetup

	return doDocTypesEqual &&
		(!isNeedToCompareNumber || document.Value == actualDoc.Value) &&
		(!isNeedToCompareValidation || (document.IsValidSetup && document.IsValid == actualDoc.IsValid))
}

func (document *ExtractedDocument) IsNormal() bool {
	return len(document.Value) == 0 || document.DocType.NormaliseValueRegex().MatchString(document.Value)
}

func (d *ExtractedDocument) ToShortString() string {
	var sb strings.Builder

	sb.WriteString(d.DocType.String())
	if d.IsValidSetup {
		if d.IsValid {
			sb.WriteByte('+')
		} else {
			sb.WriteByte('-')
		}
	} else {
		sb.WriteByte('*')
	}
	sb.WriteByte(':')
	if len(strings.TrimSpace(d.Value)) == 0 {
		sb.WriteByte('*')
	} else {
		sb.WriteString(strings.TrimSpace(d.Value))
	}
	sb.WriteByte(',')

	return sb.String()
}
