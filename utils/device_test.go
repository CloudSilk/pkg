package utils

import (
	"fmt"
	"testing"
)

func TestCPU(t *testing.T) {
	provider := GetCPUInfoProvider()
	if provider == nil {
		t.Log("不支持的操作系统")
		return
	}

	serialNumber, err := provider.GetSerialNumber()
	if err != nil {
		t.Fatal("获取CPU序列号失败:", err)
		return
	}

	fmt.Println("CPU序列号:", serialNumber)

	mbSerial, err := GetMotherboardSerial()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Motherboard Serial: %s\n", mbSerial)

	hdSerial, err := GetHardDriveSerial()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Hard Drive Serial: %s\n", hdSerial)
}
