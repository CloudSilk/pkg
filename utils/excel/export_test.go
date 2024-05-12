package excel

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"
)

type testStruct struct {
	A int
	B string
	C bool
	D interface{}
	E int64
	F map[string]interface{}
}

const sheetName = "test"

var profiles = []ExportProfile{
	{Index: 1, FieldName: "A", DisplayName: "A"},
	{Index: 2, FieldName: "B", DisplayName: "B"},
	{Index: 3, FieldName: "C", DisplayName: "C"},
	{Index: 4, FieldName: "D", DisplayName: "D"},
	{Index: 5, FieldName: "E", DisplayName: "E", Convert: "TimeStamp"},
	{Index: 6, FieldName: "F", DisplayName: "F", Convert: "testConvertMap"},
}

func testConvertMap(p ExportProfile, val reflect.Value) (interface{}, error) {
	if !val.IsValid() || val.IsZero() {
		return "", nil
	}
	if val.Kind() != reflect.Map && val.Kind() != reflect.Interface {
		fmt.Println(val.Kind(), p, val.Interface())
		return nil, errors.Errorf("Kind is not Map.")
	}
	data, err := json.Marshal(val.Interface())
	if err != nil {
		return "", err
	}
	return string(data), err
}

func TestWriteToExcel(t *testing.T) {
	RegisterConvertFunc("testConvertMap", testConvertMap)
	_, err := WriteToExcel(sheetName, profiles, []testStruct{
		{A: 1, B: "B1", C: true, D: 22, E: time.Now().Unix(), F: map[string]interface{}{
			"G": 11,
		}},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = WriteToExcel(sheetName, profiles, []*testStruct{
		{A: 1, B: "B1", C: true, D: 22, E: time.Now().Unix(), F: map[string]interface{}{
			"G": 11,
		}},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = WriteToExcel(sheetName, profiles, []map[string]interface{}{
		{"A": 1, "B": "B1", "C": false, "D": 22, "E": time.Now().Unix(), "F": map[string]interface{}{
			"G": 11,
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTimeStampToDate(t *testing.T) {
	var a int64 = 12
	_, err := TimeStampToDate(ExportProfile{}, reflect.ValueOf(a))
	if err != nil {
		t.Log(err)
	}

	var b int64 = time.Now().Unix()
	_, err = TimeStampToDate(ExportProfile{}, reflect.ValueOf(b))
	if err != nil {
		t.Fatal(err)
	}
}
