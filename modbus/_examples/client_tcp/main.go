package main

import (
	"encoding/binary"
	"fmt"
	"time"

	modbus "github.com/nooocode/pkg/modbus"
	modbus2 "github.com/nooocode/pkg/protocol/modbus"
)

const serverAddr = "192.168.1.128:10020"

// const serverAddr = "localhost:502"

func main() {
	//modbus.WithEnableLogger(),
	p := modbus.NewTCPClientProvider(serverAddr,
		modbus.WidthCRC(modbus.CRCModbus16, false),
		modbus.WithTCPTimeout(5*time.Second))
	client := modbus.NewClient(p)
	err := client.Connect()
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	defer client.Close()

	fmt.Println("starting")
	for {
		version, err := GetVersion(client, 1)
		if err != nil {
			panic(err.Error())
		}
		var pressure, altitude, temperature float32
		fmt.Printf("当前时间:%s，当前版本：%s\t",time.Now().Format(time.RFC3339) version)

		results, err := client.ReadInputRegistersBytes(1, 0xcc, 0x02)
		if err != nil {
			panic(err.Error())
		} else {
			pressure = modbus2.BytesToFloat32(results)
			fmt.Printf("当前大气压：%f\t", pressure)
		}

		results, err = client.ReadInputRegistersBytes(1, 0xc8, 0x02)
		if err != nil {
			panic(err.Error())
		} else {
			temperature = modbus2.BytesToFloat32(results)
			fmt.Printf("当前温度：%f\t", temperature)
		}

		results, err = client.ReadInputRegistersBytes(1, 0xc7, 0x01)
		if err != nil {
			panic(err.Error())
		} else {

			data := binary.BigEndian.Uint16(results)
			fmt.Printf("当前大气压楼层：%d\t", data)
		}

		results, err = client.ReadInputRegistersBytes(1, 0x94, 0x02)
		if err != nil {
			panic(err.Error())
		} else {
			altitude = modbus2.BytesToFloat32(results[:4])
			fmt.Printf("当前海拔：%f\r\n", altitude)
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

func GetSystemcConfig(client modbus.Client, slaveID byte) {
	version, err := GetVersion(client, slaveID)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("版本号：%s\r\n", version)

	for i, value := range standardValues {
		standardValue, err := GetStandardValue(client, slaveID, uint16(i+1))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%d楼==>采集卡中的标准值： [%f],默认配置的标准值:[%f]\r\n", i+1, standardValue, value)
			time.Sleep(time.Millisecond * 10)
		}

	}
	standardValueCount, err := GetStandardValueCount(client, slaveID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("标准值数量:%d\r\n", standardValueCount)
		time.Sleep(time.Millisecond * 10)
	}

	results, err := client.ReadInputRegistersBytes(slaveID, 0xc1, 0x01)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		data := binary.BigEndian.Uint16(results)
		fmt.Printf("地下楼层数：%d\r\n", data, results)
		time.Sleep(time.Millisecond * 10)
	}

	results, err = client.ReadWriteWithSubFuncCode(slaveID, 0x42, 0x02, []byte{})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("采集卡IP地址：[% x]\r\n", results)
}

var standardValues = []float32{17.10, 17.30, 17.50, 17.65}

func SetSystemConfig(client modbus.Client, slaveID byte) {
	for i, value := range standardValues {
		err := SetStandardValue(client, slaveID, uint16(i+1), value)
		if err != nil {
			fmt.Printf("设置%d楼的基准值失败:%v\n", i+1, err)
		}
	}
	err := SetStandardValueCount(client, 1, uint32(len(standardValues)))
	if err != nil {
		fmt.Printf("设置基准值数量失败:%vn", err)
	}

	err = SetNegativeFloorNum(client, 1, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	client.ReadWrite(1, 0x41, []byte{0x0c})
}

func SetStandardValue(client modbus.Client, slaveID byte, floor uint16, value float32) error {
	b := modbus2.Float32ToBytes(value)
	return client.WriteMultipleRegistersBytes(slaveID, 0x100+(floor-1)*8, 0x02, b[:])
}

func GetStandardValue(client modbus.Client, slaveID byte, floor uint16) (float32, error) {
	results, err := client.ReadInputRegistersBytes(slaveID, 0x100+(floor-1)*8, 0x08)
	if err != nil {
		return 0, err
	} else {
		return modbus2.BytesToFloat32(results[:4]), nil
	}
}

func GetStandardValueCount(client modbus.Client, slaveID byte) (uint32, error) {
	results, err := client.ReadInputRegistersBytes(slaveID, 0xfc, 0x02)
	if err != nil {
		return 0, err
	} else {
		return binary.BigEndian.Uint32(results), nil
	}
}

func SetStandardValueCount(client modbus.Client, slaveID byte, n uint32) error {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], n)
	b = [4]byte{b[2], b[3], b[0], b[1]}
	return client.WriteMultipleRegistersBytes(slaveID, 0xfc, 0x02, b[:])
}

func SetNegativeFloorNum(client modbus.Client, slaveID byte, n uint16) error {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], n)
	return client.WriteMultipleRegistersBytes(slaveID, 0xc1, 0x01, b[:])
}

func GetVersion(client modbus.Client, slaveID byte) (string, error) {
	results, err := client.ReadWrite(slaveID, 0x2b, []byte{0x01})
	if err != nil {
		return "", err
	}

	return string(results[:len(results)-4]), err
}
