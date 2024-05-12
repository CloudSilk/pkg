package excel

import (
	"fmt"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

func WriteToExcel(sheetName string, profiles []ExportProfile, obj interface{}) ([]byte, error) {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	WriteHeaderToExcel(f, sheetName, profiles)
	val := reflect.ValueOf(obj)
	if val.IsNil() {
		buf, err := f.WriteToBuffer()
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return nil, fmt.Errorf("obj not array:%v", val.Kind())
	}

	n := val.Len()
	for i := 0; i < n; i++ {
		el := val.Index(i)
		if el.Kind() == reflect.Ptr {
			el = el.Elem()
		}
		err := writeToExcel(f, sheetName, profiles, i+2, el)
		if err != nil {
			return nil, err
		}
	}
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func WriteHeaderToExcel(f *excelize.File, sheetName string, profiles []ExportProfile) {
	for _, p := range profiles {
		f.SetCellValue(sheetName, indexToAxis(p.Index, 1), p.DisplayName)
	}
}

func writeToExcel(f *excelize.File, sheetName string, profiles []ExportProfile, line int, val reflect.Value) error {
	for _, p := range profiles {
		var fieldVal reflect.Value
		if val.Kind() == reflect.Struct || val.Kind() == reflect.Ptr {
			fieldVal = val.FieldByName(p.FieldName)
		} else {
			keys := val.MapKeys()
			for _, key := range keys {
				if key.String() == p.FieldName {
					fieldVal = val.MapIndex(key)
				}
			}
		}

		if !fieldVal.IsValid() {
			continue
		}
		var data interface{}
		convertFn, ok := exportConvertFuncList[p.Convert]
		if ok {
			var err error
			data, err = convertFn(p, fieldVal)
			if err != nil {
				return err
			}
		} else {
			data = fieldVal.Interface()
		}

		err := f.SetCellValue(sheetName, indexToAxis(p.Index, line), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func indexToAxis(index, line int) string {
	return fmt.Sprintf("%c%d", 'A'+index-1, line)
}

type ExportProfile struct {
	Index       int
	Hide        bool
	FieldName   string
	DisplayName string
	Convert     string
}

var exportConvertFuncList = make(map[string]ConvertFunc)

func init() {
	exportConvertFuncList["TimeStamp"] = TimeStampToDate
}

type ConvertFunc func(ExportProfile, reflect.Value) (interface{}, error)

func TimeStampToDate(p ExportProfile, val reflect.Value) (interface{}, error) {
	if !val.IsValid() || val.IsZero() {
		return "", nil
	}

	var d int64
	if val.Kind() == reflect.Interface {
		ok := false
		d, ok = (val.Interface()).(int64)
		if !ok {
			return "", errors.Errorf("Kind is not Int64.")
		}
	} else if val.Kind() != reflect.Int64 {
		return "", errors.Errorf("Kind is not Int64.")
	} else {
		d = (val.Interface()).(int64)
	}
	return time.Unix(d, 0).Format(time.RFC3339), nil
}

func RegisterConvertFunc(name string, fn ConvertFunc) {
	exportConvertFuncList[name] = fn
}
