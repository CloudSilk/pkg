package modbus

import "github.com/howeyc/crc16"

// CRC16 Calculate Cyclical Redundancy Checking.
func CRC16(bs []byte, t CRCType) uint16 {
	var tab *crc16.Table = crc16.CCITTFalseTable
	switch t {
	case CRCCCTI16:
		tab = crc16.CCITTTable
	}
	return crc16.Checksum(bs, tab)
}
