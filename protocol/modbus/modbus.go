package modbus

import (
	"encoding/binary"
	"fmt"
	"math"
	"sync"

	modbus "github.com/CloudSilk/pkg/modbus"
)

/*
BytesToFloat32 字节转float32

字节序是大端，低字节在低地址，高字节在高地址

等同于c语言

data = buff[i++]; data <<= 8;

data += buff[i++];data <<= 8;

data += buff[i++] ; data <<=  8;

data += buff[i++];

val= *(float*)&data;
*/
func BytesToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(b))
}

/*
Float32ToBytes 浮点数转成字节 按大端字节序，高位在低地址，地位在高地址
*/
func Float32ToBytes(f float32) [4]byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], math.Float32bits(f))
	return [4]byte{b[2], b[3], b[0], b[1]}
}

func NewModbusClient(addr string) (*ModbusClient, error) {
	c := &ModbusClient{
		addr: addr,
	}
	err := c.Connect()
	return c, err
}

type ModbusClient struct {
	addr string
	c    modbus.Client
	lock sync.Mutex
}

// Connect 连接modbus
func (client *ModbusClient) Connect() error {
	client.lock.Lock()
	defer client.lock.Unlock()

	if client.c != nil && client.c.IsConnected() {
		return nil
	}

	if client.c != nil {
		client.c.Close()
		client.c = nil
	}

	c := modbus.NewClient(modbus.NewTCPClientProvider(client.addr, modbus.WithEnableLogger(), modbus.WidthCRC(modbus.CRCModbus16, false)))
	err := c.Connect()
	if err != nil {
		return err
	}
	client.c = c
	return nil
}

// Close 关闭
func (client *ModbusClient) Close() {
	client.lock.Lock()
	defer client.lock.Unlock()
	if client.c != nil {
		client.c.Close()
		client.c = nil
	}
}

// ReadHoldingRegistersBytes 读取寄存器值
func (client *ModbusClient) ReadHoldingRegistersBytes(slaveID byte, address, quantity uint16) ([]byte, error) {
	// 重连
	err := client.Connect()
	if err != nil {
		client.Close()
		return nil, err
	}

	result, err := client.c.ReadHoldingRegistersBytes(slaveID, address, quantity)
	if err != nil {
		fmt.Println("ReadHoldingRegistersBytes error:", err)
		return nil, err
	}
	return result, nil
}
