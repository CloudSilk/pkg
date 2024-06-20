package system

import "testing"

func TestGetSystemUsage(t *testing.T) {
	systemInfo,err:=GetSystemUsage()
	if err!=nil{
		t.Fatal(err)
	}
	systemInfo.Print()
}
