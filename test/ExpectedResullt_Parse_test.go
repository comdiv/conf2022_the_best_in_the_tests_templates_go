package test

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"reflect"
	"testing"
)

const CORRECT_INPUT = "пасспортРФ"
const CORRECT_CONSTRAINTS = "=="
const CORRECT_DOC_TYPE = doc_type.PASSPORT_RF

func testParseDocs(t *testing.T, input string, expectedParsedDocs []output.ExtractedDocument) {
	var parsedResult = output.ParseExpectedResult(input)

	actualEqualsExpected := true

	if len(parsedResult.Docs) != len(expectedParsedDocs) {
		actualEqualsExpected = false
	} else {
		actualEqualsExpected = reflect.DeepEqual(expectedParsedDocs, parsedResult.Docs)
	}

	if !actualEqualsExpected {
		t.Errorf("При парсинге строки %s, результат - %+v\n не совпадает с ожидаемым -  %+v\n",
			input, parsedResult.Docs, expectedParsedDocs)
	}
}

func Test_DeepEqual(t *testing.T) {
	firstSlice := []output.ExtractedDocument{
		{
			DocType: doc_type.PASSPORT_RF,
			Value:   "somePassportRF",
		},
		{
			DocType: doc_type.SNILS,
			Value:   "someSNILS",
		},
	}
	secondSlice := []output.ExtractedDocument{
		{
			DocType: doc_type.PASSPORT_RF,
			Value:   "somePassportRF",
		},
		{
			DocType: doc_type.SNILS,
			Value:   "someSNILS",
		},
	}

	fmt.Println(secondSlice)

	if !reflect.DeepEqual(firstSlice, secondSlice) {
		t.Error("DeepEqual does not work!")
	}
}

func Test_Parse_Constraints(t *testing.T) {
	type ConstraintTestCase struct {
		testName       string
		input          string
		expectedResult output.ExpectedResult
	}

	testCases := []ConstraintTestCase{
		{
			testName:       "'==' - исключительно ожидаемый набор в указанном порядке",
			input:          fmt.Sprintf("%s==%v", CORRECT_INPUT, CORRECT_DOC_TYPE),
			expectedResult: output.ExpectedResult{IsExactly: true, IsOrderRequired: true},
		},
		{
			testName:       "'~=' - ожидаемый набор содержится в итоговой выборке в указанном порядке",
			input:          fmt.Sprintf("%s~=%v", CORRECT_INPUT, CORRECT_DOC_TYPE),
			expectedResult: output.ExpectedResult{IsExactly: false, IsOrderRequired: true},
		},
		{
			testName:       "'=?' - ожидаемый набор содержится в итоговой выборке в указанном порядке",
			input:          fmt.Sprintf("%s=?%v", CORRECT_INPUT, CORRECT_DOC_TYPE),
			expectedResult: output.ExpectedResult{IsExactly: true, IsOrderRequired: false},
		},
		{
			testName:       "'~?' - ожидаемый набор содержится в итоговой выборке в указанном порядке",
			input:          fmt.Sprintf("%s~?%v", CORRECT_INPUT, CORRECT_DOC_TYPE),
			expectedResult: output.ExpectedResult{IsExactly: false, IsOrderRequired: false},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			var parsed = output.ParseExpectedResult(testCase.input)

			if parsed.IsExactly != testCase.expectedResult.IsExactly {
				t.Errorf("При парсинге %s, isExactly ожидалось - %t, получилось - %t", testCase.input, testCase.expectedResult.IsExactly, parsed.IsExactly)
			}

			if parsed.IsOrderRequired != testCase.expectedResult.IsOrderRequired {
				t.Errorf("При парсинге %s, IsOrderRequired ожидалось - %t, получилось - %t", testCase.input, testCase.expectedResult.IsOrderRequired, parsed.IsOrderRequired)
			}
		})
	}
}

func Test_Parse_AllDocTypes(t *testing.T) {
	testCases := map[string]doc_type.DocType{
		"паспорт рф 01239423==PASSPORT_RF": doc_type.PASSPORT_RF,
		"ВУ ==DRIVER_LICENSE":              doc_type.DRIVER_LICENSE,
		"ВИН ==VIN":                        doc_type.VIN,
		"СТС ==STS":                        doc_type.STS,
		"grz==GRZ":                         doc_type.GRZ,
		"инн Данила==INN_FL":               doc_type.INN_FL,
		"инн рога и копыта==INN_UL":        doc_type.INN_UL,
		"ОГРН==OGRN":                       doc_type.OGRN,
		"ОГРНИП==OGRNIP":                   doc_type.OGRNIP,
		"СНИЛС==SNILS":                     doc_type.SNILS,
	}

	for inputString, expectedDocType := range testCases {
		t.Run(fmt.Sprintf("Извлекается документ - %v", expectedDocType), func(t *testing.T) {
			parsed := output.ParseExpectedResult(inputString)

			if len(parsed.Docs) == 0 {
				t.Errorf("Не удалось распарсить строку - %s", inputString)
			} else if parsed.Docs[0].DocType != expectedDocType {
				t.Errorf("При парсинге строки - %s, ожидается документ - %v, а не - %v", inputString, expectedDocType, parsed.Docs[0].DocType)
			}
		})
	}
}

func Test_Parsed(t *testing.T) {
	type TestCase struct {
		testName       string
		input          string
		expectedResult []output.ExtractedDocument
	}

	testCases := []TestCase{
		{
			testName:       "Документы без значений",
			input:          "паспорт рф, инн юл==PASSPORT_RF, INN_UL",
			expectedResult: []output.ExtractedDocument{{DocType: doc_type.PASSPORT_RF, Value: "", IsValidSetup: false}, {DocType: doc_type.INN_UL, Value: "", IsValidSetup: false}},
		},

		{
			testName:       "Некоторые документ со значениями, некоторые без",
			input:          "паспорт рф,  инн юл 0123456789==PASSPORT_RF, INN_UL:0123456789",
			expectedResult: []output.ExtractedDocument{{DocType: doc_type.PASSPORT_RF, Value: "", IsValidSetup: false}, {DocType: doc_type.INN_UL, Value: "0123456789", IsValidSetup: false}},
		},

		{
			testName:       "Все документы со значениями",
			input:          "паспорт рф 9876543210,  инн юл 0123456789==PASSPORT_RF:9876543210, INN_UL:0123456789",
			expectedResult: []output.ExtractedDocument{{DocType: doc_type.PASSPORT_RF, Value: "9876543210"}, {DocType: doc_type.INN_UL, Value: "0123456789"}},
		},

		{
			testName:       "проверка на трим значения",
			input:          "паспорт рф 9876543210==PASSPORT_RF:   9876543210  ",
			expectedResult: []output.ExtractedDocument{{DocType: doc_type.PASSPORT_RF, Value: "9876543210"}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			testParseDocs(t, testCase.input, testCase.expectedResult)
		})
	}
}

func Test_NonNormaliseValue(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Ожидалось исключение при указании номера документа в ненормализованном виде")
		}
	}()

	nonNormaliseValue := "9876 543210"

	if doc_type.PASSPORT_RF.NormaliseValueRegex().MatchString(nonNormaliseValue) {
		t.Fail()
	}

	input := fmt.Sprintf("0123456789==PASSPORT_RF:%s", nonNormaliseValue)

	output.ParseExpectedResult(input)
}

func Test_ValidationIsNotSetup(t *testing.T) {
	input := fmt.Sprintf("0123456789==PASSPORT_RF:0123456789")

	result := output.ParseExpectedResult(input)

	if result.Docs[0].IsValidSetup {
		t.Error("Флаг установки валидации должен быть false")
	}
}

func Test_SetupPositiveValidation(t *testing.T) {
	input := fmt.Sprintf("0123456789==PASSPORT_RF%s:0123456789", output.VALID_DOC_SUFFIX)

	result := output.ParseExpectedResult(input)

	if !result.Docs[0].IsValidSetup || !result.Docs[0].IsValid {
		t.Error("Флаг установки валидации должен быть true, а сама валидация - true")
	}
}

func Test_SetupNegativeValidation(t *testing.T) {
	input := fmt.Sprintf("0123456789==PASSPORT_RF%s:0123456789", output.INVALID_DOC_SUFFIX)

	result := output.ParseExpectedResult(input)

	if !result.Docs[0].IsValidSetup || result.Docs[0].IsValid {
		t.Error("Флаг установки валидации должен быть true, а сама валидация - false")
	}
}
