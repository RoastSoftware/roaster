package analyze_test

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze"
)

func TestFlake8(t *testing.T) {
	testJSON := `
{"test/test.py": [{
	"code": "F821",
	"filename": "test/test.py",
	"line_number": 2,
	"column_number": 5,
	"text": "undefined name 'thisfunctiondoesnotexist'",
	"physical_line": "    thisfunctiondoesnotexist(\"Expect an error.\")\n"
}]}`

	var testResult analyze.Flake8Result
	err := json.NewDecoder(strings.NewReader(testJSON)).Decode(&testResult)
	if err != nil {
		log.Println(err)
		return
	}

	testExpected := make(analyze.Flake8Result)
	testMessage := analyze.Flake8Message{
		Code:         "F821",
		Filename:     "test/test.py",
		LineNumber:   2,
		ColumnNumber: 5,
		Text:         "undefined name 'thisfunctiondoesnotexist'",
		PhysicalLine: "    thisfunctiondoesnotexist(\"Expect an error.\")\n",
	}
	testExpected["test/test.py"] = []analyze.Flake8Message{testMessage}

	if !reflect.DeepEqual(testResult, testExpected) {
		t.Errorf("unexpected result: %v", testResult)
	}
}
