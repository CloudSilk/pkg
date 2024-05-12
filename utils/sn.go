package utils

import (
	"errors"
	"fmt"
	"hash/crc32"
	"hash/crc64"

	"github.com/sigurn/crc8"
	"github.com/sony/sonyflake"
)

// GenMac 生成Mac地址
func GenMac(mac int64) string {
	var s string
	for i, v := range fmt.Sprintf("%012X", mac) {
		if i == 0 {
			v = '2'
		}
		if i == 1 {
			v = '2'
		}
		if i != 0 && i%2 == 0 {
			s = s + ":"
		}
		s = s + string(v)
	}
	return s
}

func validate(x string) string {
	v := fmt.Sprintf("%02X", crc8.Checksum([]byte(x), crc8.MakeTable(crc8.CRC8_ITU)))
	return v
}

// CheckSN 校验序列号
func CheckSN(snStr string) error {
	if len(snStr) != 10 {
		return errors.New("SN序列号长度不为10")
	}
	if validate(snStr[0:8]) != snStr[8:10] {
		return errors.New("SN不合法")
	}
	return nil
}

// GenSN 生成10位序列号
func GenSN(sn int64) string {
	x := fmt.Sprintf("%08X", crc32.ChecksumIEEE([]byte(fmt.Sprint(sn))))
	return x + validate(x)
}

// GenSN64 生成18位序列号
func GenSN64(sn int64) string {
	t := crc64.MakeTable(crc64.ECMA)
	ss := crc64.Checksum([]byte{byte(sn)}, t)
	x := fmt.Sprintf("%016X", ss)
	return x + validate(x)
}

var nextID = sonyflake.NewSonyflake(sonyflake.Settings{})

// GenShortID 生成15位短ID
func GenShortID(prefix string) string {
	idStr, _ := nextID.NextID()
	if prefix == "" {
		return fmt.Sprintf("%x", idStr)
	}
	return fmt.Sprintf("%s_%x", prefix, idStr)
}
