package output

import "github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"

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
func (expectedDoc *ExtractedDocument) Match(actualDoc ExtractedDocument) bool {
	doDocTypesEqual := expectedDoc.DocType == actualDoc.DocType

	isNeedToCompareNumber := len(expectedDoc.Value) != 0
	isNeedToCompareValidation := expectedDoc.IsValidSetup

	return doDocTypesEqual &&
		(!isNeedToCompareNumber || expectedDoc.Value == actualDoc.Value) &&
		(!isNeedToCompareValidation || (expectedDoc.IsValidSetup && expectedDoc.IsValid == actualDoc.IsValid))
}

func (expectedDoc *ExtractedDocument) IsNormal() bool {
	return len(expectedDoc.Value) == 0 || expectedDoc.DocType.NormaliseValueRegex().MatchString(expectedDoc.Value)
}
