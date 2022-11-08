package doc_type

import (
	"fmt"
	"regexp"
	"strings"
)

// DocType - тип документа
type DocType int

const (
	UNDEFINED DocType = 0

	// PASSPORT_RF - Паспорт РФ
	PASSPORT_RF DocType = 1

	// DRIVER_LICENSE - Водительское удостоверение
	DRIVER_LICENSE DocType = 2

	// VIN - Идентификационный номер транспортного средства
	VIN DocType = 3

	// STS - Свидетельство о регистрации транспортного средства
	STS DocType = 4

	// GRZ - Государственный регистрационный номер транспортного средства
	GRZ DocType = 5

	// INN_FL - ИНН Юр. лица
	INN_FL DocType = 6

	// INN_UL - ИНН Физ. лица
	INN_UL DocType = 7

	// OGRN - ОГРН
	OGRN DocType = 8

	// OGRNIP - ОГРНИП
	OGRNIP DocType = 9

	// SNILS - СНИЛС
	SNILS DocType = 10
)

// String Получает строковое представление экземпляра перечисления
func (doc DocType) String() string {
	switch doc {
	case UNDEFINED:
		return "UNDEFINED"
	case PASSPORT_RF:
		return "PASSPORT_RF"
	case DRIVER_LICENSE:
		return "DRIVER_LICENSE"
	case VIN:
		return "VIN"
	case STS:
		return "STS"
	case GRZ:
		return "GRZ"
	case INN_FL:
		return "INN_FL"
	case INN_UL:
		return "INN_UL"
	case OGRN:
		return "OGRN"
	case OGRNIP:
		return "OGRNIP"
	case SNILS:
		return "SNILS"
	default:
		panic(fmt.Sprintf("попытка получения строкового представления неизвестного типа документа - %v", doc))
	}
}

// Parse преобразует строковое представление в экземпляр перечисления
func Parse(docTypeAsString string) DocType {
	switch strings.ToUpper(docTypeAsString) {
	case PASSPORT_RF.String():
		return PASSPORT_RF
	case DRIVER_LICENSE.String():
		return DRIVER_LICENSE
	case VIN.String():
		return VIN
	case STS.String():
		return STS
	case GRZ.String():
		return GRZ
	case INN_FL.String():
		return INN_FL
	case INN_UL.String():
		return INN_UL
	case OGRN.String():
		return OGRN
	case OGRNIP.String():
		return OGRNIP
	case SNILS.String():
		return SNILS
	case UNDEFINED.String():
		return UNDEFINED
	default:
		panic(fmt.Sprintf("попытка парсинга неизвестного типа документа - %s", docTypeAsString))
	}
}

// NormaliseValueRegex - получает регулярное выражение для проверки нормализованного номера документа
func (doc DocType) NormaliseValueRegex() *regexp.Regexp {
	var pattern string

	switch doc {
	case PASSPORT_RF:
		pattern = "^\\d{10}$"
	case DRIVER_LICENSE:
		pattern = "^\\d{10}$"
	case VIN:
		pattern = "^[A-Z0-9]{17}$"
	case STS:
		pattern = "^\\d{2}[А-ЯA-Z0-9]{2}\\d{6}$"
	case GRZ:
		pattern = "^[АВЕКМНОРСТУХ]\\d{3}[АВЕКМНОРСТУХ]{2}\\d{2,3}$"
	case INN_FL:
		pattern = "^\\d{12}$"
	case INN_UL:
		pattern = "^\\d{10}$"
	case OGRN:
		pattern = "^\\d{13}$"
	case OGRNIP:
		pattern = "^\\d{15}$"
	case SNILS:
		pattern = "^\\d{3}-\\d{3}-\\d{3}-\\d{2}$"
	default:
		panic(fmt.Sprintf("попытка получения регулярное выражение для нормализации неизвестного типа документа - %v", doc))
	}

	result, e := regexp.Compile(pattern)

	if e != nil {
		panic(e)
	}

	return result
}
