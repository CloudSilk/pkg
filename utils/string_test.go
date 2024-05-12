package utils

import "testing"

func TestUpperSnakeCase(t *testing.T) {
	t.Log(UpperSnakeCase("EID"))
	t.Log(UpperSnakeCase("eID"))
	t.Log(UpperSnakeCase("Eid"))
	t.Log(UpperSnakeCase("EId"))
	t.Log(UpperSnakeCase("eid"))
	t.Log(UpperSnakeCase("ID"))
	t.Log(UpperSnakeCase("iD"))
	t.Log(UpperSnakeCase("Id"))
	t.Log(UpperSnakeCase("id"))
	t.Log(UpperSnakeCase("HelloWorld"))
}

func TestLowerSnakeCase(t *testing.T) {
	t.Log(LowerSnakeCase("EID"))
	t.Log(LowerSnakeCase("eID"))
	t.Log(LowerSnakeCase("Eid"))
	t.Log(LowerSnakeCase("EId"))
	t.Log(LowerSnakeCase("eid"))
	t.Log(LowerSnakeCase("ID"))
	t.Log(LowerSnakeCase("iD"))
	t.Log(LowerSnakeCase("Id"))
	t.Log(LowerSnakeCase("id"))
	t.Log(LowerSnakeCase("HelloWorld"))
}

func TestLcFirst(t *testing.T) {
	t.Log(LcFirst("HelloWorld"))
	t.Log(LcFirst("Helloworld"))
	t.Log(LcFirst("helloWorld"))
}

func TestBigCamelName(t *testing.T) {
	t.Log(BigCamelName("EID"))
	t.Log(BigCamelName("eID"))
	t.Log(BigCamelName("Eid"))
	t.Log(BigCamelName("EId"))
	t.Log(BigCamelName("eid"))
	t.Log(BigCamelName("ID"))
	t.Log(BigCamelName("iD"))
	t.Log(BigCamelName("Id"))
	t.Log(BigCamelName("id"))
	t.Log(BigCamelName("HelloWorld"))
	t.Log(BigCamelName("Helloworld"))
	t.Log(BigCamelName("helloWorld"))
	t.Log(BigCamelName("hello_World"))
}
