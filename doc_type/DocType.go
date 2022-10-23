package doc_type

import "fmt"

type DocType int

const (
	UNDEFINED      DocType = 0
	PASSPORT_RF    DocType = 1
	DRIVER_LICENSE DocType = 2
	VIN            DocType = 3
	STS            DocType = 4
	PTS            DocType = 5
	INN_FL         DocType = 6
	INN_UL         DocType = 7
	EGRN           DocType = 8
	EGRIP          DocType = 9
	SNILS          DocType = 10
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
	case PTS:
		return "PTS"
	case INN_FL:
		return "INN_FL"
	case INN_UL:
		return "INN_UL"
	case EGRN:
		return "EGRN"
	case EGRIP:
		return "EGRIP"
	case SNILS:
		return "SNILS"
	default:
		return fmt.Sprintf("%d", int(doc))
	}
}

//Parse преобразует строковое представление в экземпляр перечисления
func Parse(input string) DocType {
	switch input {
	case PASSPORT_RF.String():
		return PASSPORT_RF
	case DRIVER_LICENSE.String():
		return DRIVER_LICENSE
	case VIN.String():
		return VIN
	case STS.String():
		return STS
	case PTS.String():
		return PTS
	case INN_FL.String():
		return INN_FL
	case INN_UL.String():
		return INN_UL
	case EGRN.String():
		return EGRN
	case EGRIP.String():
		return EGRIP
	case SNILS.String():
		return SNILS
	default:
		return UNDEFINED
	}
}
