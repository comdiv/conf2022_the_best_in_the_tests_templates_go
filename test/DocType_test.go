package test

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"testing"
)

func testParseToString(t *testing.T, inputDocType doc_type.DocType, expectedString string) {
	var actualString = inputDocType.String()

	if actualString != expectedString {
		t.Error(
			fmt.Sprintf(
				"При парсинге в строку типа документа %v результат - %s не соответствует ожиданию - %s",
				inputDocType,
				actualString,
				expectedString,
			),
		)
	}
}

func testParseFromString(t *testing.T, input string, expectedDocType doc_type.DocType) {
	var actualDocType = doc_type.Parse(input)

	if actualDocType != expectedDocType {
		t.Error(
			fmt.Sprintf(
				"При парсинге из строки  %s тип документа - %v не соответствует ожиданию - %v",
				input,
				actualDocType,
				expectedDocType,
			),
		)
	}
}

func TestFromString(t *testing.T) {
	testCases := map[string]doc_type.DocType{
		"PASSPORT_RF":    doc_type.PASSPORT_RF,
		"DRIVER_LICENSE": doc_type.DRIVER_LICENSE,
		"VIN":            doc_type.VIN,
		"STS":            doc_type.STS,
		"GRZ":            doc_type.GRZ,
		"INN_FL":         doc_type.INN_FL,
		"INN_UL":         doc_type.INN_UL,
		"OGRN":           doc_type.OGRN,
		"OGRNIP":         doc_type.OGRNIP,
		"SNILS":          doc_type.SNILS,

		"passport_rf":    doc_type.PASSPORT_RF,
		"driver_license": doc_type.DRIVER_LICENSE,
		"vin":            doc_type.VIN,
		"sts":            doc_type.STS,
		"grz":            doc_type.GRZ,
		"inn_fl":         doc_type.INN_FL,
		"inn_ul":         doc_type.INN_UL,
		"ogrn":           doc_type.OGRN,
		"ogrnip":         doc_type.OGRNIP,
		"snils":          doc_type.SNILS,
	}

	for inputString, expectedDocType := range testCases {
		t.Run(fmt.Sprintf("Парсинг из строки %v", expectedDocType), func(t *testing.T) {
			testParseFromString(t, inputString, expectedDocType)
		})
	}
}

// Проверяет, что если пытаешься спарсить строку, которая не соответствует ни одному из документа
// - выкидывается ошибка
func Test_IncorrectString(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Ожидалось исключение при попытке обработки некорректной строки как тип документа")
		}
	}()

	doc_type.Parse("someIncorrectString")
}

func TestToString(t *testing.T) {
	testCases := map[doc_type.DocType]string{
		doc_type.PASSPORT_RF:    "PASSPORT_RF",
		doc_type.DRIVER_LICENSE: "DRIVER_LICENSE",
		doc_type.VIN:            "VIN",
		doc_type.STS:            "STS",
		doc_type.GRZ:            "GRZ",
		doc_type.INN_FL:         "INN_FL",
		doc_type.INN_UL:         "INN_UL",
		doc_type.OGRN:           "OGRN",
		doc_type.OGRNIP:         "OGRNIP",
		doc_type.SNILS:          "SNILS",
		doc_type.UNDEFINED:      "UNDEFINED",
	}

	for docType, expectedString := range testCases {
		t.Run(fmt.Sprintf("Парсинг в строку %v", docType), func(t *testing.T) {
			testParseToString(t, docType, expectedString)
		})
	}
}

func testingNormaliseValueRegex(
	t *testing.T,
	testName string,
	docType doc_type.DocType,
	value string,
	expectedMatch bool,
) {
	t.Run(testName, func(innerT *testing.T) {
		if docType.NormaliseValueRegex().MatchString(value) != expectedMatch {
			innerT.Fail()
		}
	})
}

func Test_NormaliseValueRegex(t *testing.T) {
	t.Run("Паспорт РФ", func(innerT *testing.T) {
		passportRf := doc_type.PASSPORT_RF

		testingNormaliseValueRegex(
			innerT,
			"десять цифр без пробела - валиден",
			passportRf,
			"0123456789",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"десять цифр, содержит пробел - не валиден",
			passportRf,
			"0123 456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"девять цифр, не содержит пробел - не валиден",
			passportRf,
			"123456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"одиннадцать цифр, не содержит пробел - не валиден",
			passportRf,
			"01234567891",
			false,
		)
	})

	t.Run("Водительское удостоверение", func(innerT *testing.T) {
		dl := doc_type.DRIVER_LICENSE

		testingNormaliseValueRegex(
			innerT,
			"десять цифр без пробела - валиден",
			dl,
			"0123456789",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"десять цифр, содержит пробел - не валиден",
			dl,
			"0123 456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"девять цифр, не содержит пробел - не валиден",
			dl,
			"123456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"одиннадцать цифр, не содержит пробел - не валиден",
			dl,
			"01234567891",
			false,
		)
	})

	t.Run("Идентификационный номер транспортного средства", func(innerT *testing.T) {
		vin := doc_type.VIN

		testingNormaliseValueRegex(
			innerT,
			"семнадцать цифр без пробела - валиден",
			vin,
			"12345678901234567",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"семнадцать заглавных латинских букв без пробела - валиден",
			vin,
			"ABCDEFGHIJKLMNOPQ",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"семнадцать строчных латинских букв без пробела - не валиден",
			vin,
			"abcdefghijklmnopq",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"латинские заглавные буквы + цифры - всего 17 - без пробела - валиден",
			vin,
			"ABCDEFGHIJKLM1234",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"латинские заглавные буквы + цифры - всего 17 содержат пробел - не валиден",
			vin,
			"ABCDEFGHIJKLM 1234",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"шестнадцать цифр без пробела - не валиден",
			vin,
			"1234567890123456",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"восемнадцать цифр без пробела - не валиден",
			vin,
			"123456789012345678",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"шестнадцать заглавных латинских букв без пробела - не валиден",
			vin,
			"ABCDEFGHIJKLMNOP",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"восемнадцать заглавных латинских букв без пробела - не валиден",
			vin,
			"ABCDEFGHIJKLMNOPQR",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"семнадцать заглавных букв кириллица без пробела - не валиден",
			vin,
			"АБВГДЕЁЖЗИЙКЛМНОП",
			false,
		)
	})

	t.Run("Государственный регистрационный номер транспортного средства", func(innerT *testing.T) {
		grz := doc_type.GRZ

		testingNormaliseValueRegex(
			innerT,
			"С227НА69 - валиден",
			grz,
			"С227НА69",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"C227HA69 - написан латинскими буквами - не валиден",
			grz,
			"C227HA69",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"С227НА691 - регион из 3 цифр - валиден",
			grz,
			"С227НА691",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"С227НА6 - не заканчивается на две цифры - не валиден",
			grz,
			"С227НА6",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"С227Н69 - не содержит перед двумя цифрами две буквы - не валиден",
			grz,
			"С227Н69",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"НА69 - содержит только две цифры и две буквы - не валиден",
			grz,
			"НА69",
			false,
		)
	})

	t.Run("Свидетельство о регистрации транспортного средства", func(innerT *testing.T) {
		sts := doc_type.STS

		testingNormaliseValueRegex(
			innerT,
			"1234567890 - десять цифр - валиден",
			sts,
			"1234567890",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"12AA567890 - две цифры две буквы заглавные шесть цифр - валиден",
			sts,
			"12AA567890",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"12 AA 567890 - две цифры две буквы заглавные шесть цифр содержит пробел - не валиден",
			sts,
			"12 AA 567890",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"12aa567890 - две цифры две буквы строчные шесть цифр - не валиден",
			sts,
			"12aa567890",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"12AA56789 - две цифры две буквы строчные 5 цифр - не валиден",
			sts,
			"12AA56789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"123456789 - девять цифр - не валиден",
			sts,
			"123456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"12345678901 - одиннадцать цифр - не валиден",
			sts,
			"12345678901",
			false,
		)
	})

	t.Run("ИНН физ.лица", func(innerT *testing.T) {
		innFl := doc_type.INN_FL

		testingNormaliseValueRegex(
			innerT,
			"двенадцать цифр без пробела - валиден",
			innFl,
			"123456789012",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"двенадцать цифр, содержит пробел - не валиден",
			innFl,
			"1234 56789012",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"одиннадцать цифр, не содержит пробел - не валиден",
			innFl,
			"12345678901",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"тринадцать цифр, не содержит пробел - не валиден",
			innFl,
			"1234567890123",
			false,
		)
	})

	t.Run("ИНН юр.лица", func(innerT *testing.T) {
		innUl := doc_type.INN_UL

		testingNormaliseValueRegex(
			innerT,
			"десять цифр без пробела - валиден",
			innUl,
			"0123456789",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"десять цифр, содержит пробел - не валиден",
			innUl,
			"0123 456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"девять цифр, не содержит пробел - не валиден",
			innUl,
			"123456789",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"одиннадцать цифр, не содержит пробел - не валиден",
			innUl,
			"01234567891",
			false,
		)
	})

	t.Run("ОГРН", func(innerT *testing.T) {
		ogrn := doc_type.OGRN

		testingNormaliseValueRegex(
			innerT,
			"тринадцать цифр без пробела - валиден",
			ogrn,
			"1234567890123",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"тринадцать цифр, содержит пробел - не валиден",
			ogrn,
			"1234 567890123",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"двенадцать цифр, не содержит пробел - не валиден",
			ogrn,
			"123456789012",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"четырнадцать цифр, не содержит пробел - не валиден",
			ogrn,
			"12345678901234",
			false,
		)
	})

	t.Run("ОГРНИП", func(innerT *testing.T) {
		ogrnip := doc_type.OGRNIP

		testingNormaliseValueRegex(
			innerT,
			"пятнадцать цифр без пробела - валиден",
			ogrnip,
			"123456789012345",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"пятнадцать цифр, содержит пробел - не валиден",
			ogrnip,
			"12 3456789012345",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"шестнадцать цифр, не содержит пробел - не валиден",
			ogrnip,
			"123456789012",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"четырнадцать цифр, не содержит пробел - не валиден",
			ogrnip,
			"12345678901234",
			false,
		)
	})

	t.Run("СНИЛС", func(innerT *testing.T) {
		snils := doc_type.SNILS

		testingNormaliseValueRegex(
			innerT,
			"Всё через тире - валиден",
			snils,
			"123-456-789-00",
			true,
		)

		testingNormaliseValueRegex(
			innerT,
			"Последние две цифры через пробел - не валиден",
			snils,
			"123-456-789 00",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"Все через тире - но содержит не одиннадцать цифр - не валиден",
			snils,
			"123-46-789-00",
			false,
		)

		testingNormaliseValueRegex(
			innerT,
			"Без тире - не валиден",
			snils,
			"1234678900",
			false,
		)
	})
}
