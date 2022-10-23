package conf2022_the_best_in_the_tests_templates_go

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/input"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/parser"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

var filePath = filepath.Join(basepath, "local.csv")

func Test_Validate(t *testing.T) {
	validationResult := input.Validate(filePath)

	if !validationResult.IsValid {
		t.Errorf("Файл - %s не валиден. %+v", filePath, validationResult.GetErrors())
	}
}

func Test_Local(t *testing.T) {

	testDescriptions := input.ParseFromFile(filePath)
	p := parser.RandomSuccessfulParser{}

	for _, testDesc := range testDescriptions {
		if !testDesc.IsDisabled {
			t.Run(fmt.Sprintf("%s№%d input: %s", testDesc.Author, testDesc.Number, testDesc.StringToProcessed), func(t *testing.T) {
				expectedResult := output.Parse(testDesc.StringToProcessed)
				actualDocs := p.Parse(testDesc.StringToProcessed)

				if !expectedResult.Match(actualDocs) {
					t.Error(testDesc.CommentOnFailure)
				}
			})
		}
	}
}
