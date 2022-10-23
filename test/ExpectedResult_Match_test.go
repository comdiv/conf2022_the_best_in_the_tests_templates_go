package test

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/doc_type"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"testing"
)

func testMatch(t *testing.T, result output.ExpectedResult, actualDocs []output.ExtractedDocument, expectedMatch bool) {
	if result.Match(actualDocs) != expectedMatch {
		if expectedMatch {
			t.Errorf("Набор документов %+v должен удолетворять %+v", actualDocs, result)
		} else {
			t.Errorf("Набор документов %+v не должен удолетворять %+v", actualDocs, result)
		}
	}
}

func Test_Match(t *testing.T) {
	expectedResultCases := []output.ExpectedResult{
		{IsExactly: true, IsOrderRequired: true},
		{IsExactly: true, IsOrderRequired: false},
		{IsExactly: false, IsOrderRequired: true},
		{IsExactly: false, IsOrderRequired: false},
	}

	someDoc := output.ExtractedDocument{DocType: doc_type.SNILS, Value: "012345555"}
	anotherSomeDoc := output.ExtractedDocument{DocType: doc_type.DRIVER_LICENSE, Value: "14587999922"}

	for _, resultCase := range expectedResultCases {
		t.Run(fmt.Sprintf("ExpectedResult:%+v ", resultCase), func(t *testing.T) {
			t.Run("Содержит именно тот набор в том же порядке - соответствует результату", func(t *testing.T) {
				result := output.ExpectedResult{
					IsExactly:       resultCase.IsExactly,
					IsOrderRequired: resultCase.IsOrderRequired,
					Docs:            []output.ExtractedDocument{someDoc, anotherSomeDoc},
				}

				actualDocs := []output.ExtractedDocument{someDoc, anotherSomeDoc}

				testMatch(t, result, actualDocs, true)
			})

			t.Run("Не содержит ожидаемый набор - не соответствует результату", func(t *testing.T) {
				result := output.ExpectedResult{
					IsExactly:       resultCase.IsExactly,
					IsOrderRequired: resultCase.IsOrderRequired,
					Docs:            []output.ExtractedDocument{someDoc, anotherSomeDoc},
				}

				actualDocs := []output.ExtractedDocument{someDoc}

				testMatch(t, result, actualDocs, false)
			})

			var testName string

			if resultCase.IsOrderRequired {
				testName = "Содержит именно тот набор в другом порядке - не соответствует результату"
			} else {
				testName = "Содержит именно тот набор в другом порядке - соответствует результату"
			}

			t.Run(testName, func(t *testing.T) {
				result := output.ExpectedResult{
					IsExactly:       resultCase.IsExactly,
					IsOrderRequired: resultCase.IsOrderRequired,
					Docs:            []output.ExtractedDocument{someDoc, anotherSomeDoc},
				}

				actualDocs := []output.ExtractedDocument{anotherSomeDoc, someDoc}

				testMatch(t, result, actualDocs, !resultCase.IsOrderRequired)
			})

			if resultCase.IsOrderRequired || resultCase.IsExactly {
				testName = "Содержит ожидаемый набор в другом порядке и ДОП. элемент - не соответствует результату"
			} else {
				testName = "Содержит ожидаемый набор в другом порядке и ДОП. элемент - соответствует результату"
			}

			t.Run(testName, func(t *testing.T) {
				result := output.ExpectedResult{
					IsExactly:       resultCase.IsExactly,
					IsOrderRequired: resultCase.IsOrderRequired,
					Docs:            []output.ExtractedDocument{someDoc, anotherSomeDoc},
				}

				actualDocs := []output.ExtractedDocument{anotherSomeDoc, {}, someDoc}

				testMatch(t, result, actualDocs, !(resultCase.IsOrderRequired || resultCase.IsExactly))
			})

			t.Run("Ожидается пустой набор - проверяется пустой набор - соответствует результату", func(t *testing.T) {
				result := output.ExpectedResult{
					IsExactly:       resultCase.IsExactly,
					IsOrderRequired: resultCase.IsOrderRequired,
					Docs:            []output.ExtractedDocument{},
				}

				var actualDocs []output.ExtractedDocument

				testMatch(t, result, actualDocs, true)
			})

			if resultCase.IsExactly {
				testName = "Ожидается пустой набор - проверяется не пустой набор - не соответствует результату"
			} else {
				testName = "Ожидается пустой набор - проверяется не пустой набор - соответствует результату"
			}

			t.Run(testName, func(t *testing.T) {
				result := output.ExpectedResult{
					IsExactly:       resultCase.IsExactly,
					IsOrderRequired: resultCase.IsOrderRequired,
					Docs:            []output.ExtractedDocument{},
				}

				actualDocs := []output.ExtractedDocument{someDoc}

				testMatch(t, result, actualDocs, !resultCase.IsExactly)
			})
		})
	}
}
