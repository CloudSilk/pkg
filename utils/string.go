package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
	"gorm.io/gorm/schema"
)

func GenerateOrderString(str, defaultSort string) (string, error) {
	if str == "" {
		return defaultSort, nil
	}
	list := make(map[string]string)
	err := json.Unmarshal([]byte(str), &list)
	if err != nil {
		return "", err
	}
	result := make([]string, 0, len(list))
	for k, v := range list {
		if k == "" {
			continue
		}
		if v == "descend" {
			result = append(result, fmt.Sprintf("`%s` desc", LowerSnakeCase(k)))
		} else {
			result = append(result, fmt.Sprintf("`%s`", LowerSnakeCase(k)))
		}
	}
	if len(result) == 0 {
		return defaultSort, nil
	}
	return strings.Join(result, ","), nil
}

// UpperSnakeCase 转大写蛇式命名
func UpperSnakeCase(s string) string {
	if s == "EID" || s == "eID" || s == "Eid" || s == "EId" || s == "eid" {
		return "EID"
	}

	if s == "ID" || s == "iD" || s == "Id" || s == "id" {
		return "ID"
	}

	str := stringy.New(s)
	snakeCaseStr := str.SnakeCase("?", "")
	return snakeCaseStr.ToUpper()
}

// ObjectToJsonString 对象转json字符串
func ObjectToJsonString(obj interface{}) string {
	buf, _ := json.Marshal(obj)
	return string(buf)
}

// JsonToStringArray json字符串转字符串数组
func JsonToStringArray(str string) []string {
	var result []string
	json.Unmarshal([]byte(str), &result)
	return result
}

// LowerSnakeCase 转小写蛇式命名
func LowerSnakeCase(s string) string {
	if s == "EID" || s == "eID" || s == "Eid" || s == "EId" || s == "eid" {
		return "eid"
	}

	if s == "ID" || s == "iD" || s == "Id" || s == "id" {
		return "id"
	}
	str := stringy.New(s)
	return str.SnakeCase().ToLower()
}

// ToUpper 转大写
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// LcFirst 转小驼峰式命名 例如HelloWorld -> helloWorld
func LcFirst(s string) string {
	if s == "EID" || s == "eID" || s == "Eid" || s == "EId" || s == "eid" {
		return "eid"
	}

	if s == "ID" || s == "iD" || s == "Id" || s == "id" {
		return "id"
	}

	str := stringy.New(s)
	return str.LcFirst()
}

// RemoveLastCharS 移除最后一个s字符
func RemoveLastCharS(s string) string {
	return strings.TrimRight(s, "s")
}

// BigCamelName 转大驼峰式命名 例如helloWorld -> HelloWorld
func BigCamelName(s string) string {
	if s == "EID" || s == "eID" || s == "Eid" || s == "EId" || s == "eid" {
		return "EID"
	}

	if s == "ID" || s == "iD" || s == "Id" || s == "id" {
		return "ID"
	}
	str := stringy.New(s)
	return str.CamelCase()
}

var NamingStrategy schema.NamingStrategy

// ToTableName 转成数据库表名
func ToTableName(str string) string {
	return NamingStrategy.TableName(str)
}

// GetFileExtensionByLanguage 根据编程语言获取文件后缀名
func GetFileExtensionByLanguage(language string) string {
	switch language {
	case "Golang":
		return ".go"
	case "ProtocolBuffer":
		return ".proto"
	case "JSON":
		return ".json"
	case "YAML":
		return ".yaml"
	case "Typescript":
		return ".ts"
	case "Java":
		return ".java"
	case "Properties":
		return ".properties"
	case "Sql":
		return ".sql"
	case "Vue":
		return ".vue"
	case "React":
		return ".tsx"
	default:
		return ""
	}
}

// ArrayToString 数组转字符串，元素之间用逗号隔开
func ArrayToString(array []string) string {
	if len(array) == 0 {
		return ""
	}
	return strings.Join(array, ",")
}

// StringToArray 字符串转字符串数组，元素之间用逗号隔开
func StringToArray(str string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}
