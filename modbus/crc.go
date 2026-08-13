package modbus

import "github.com/howeyc/crc16"

// CRC16 Calculate Cyclical Redundancy Checking.
func CRC16(bs []byte, t CRCType) uint16 {
	switch t {
	case CRCCCTI16:
		return crc16.Checksum(bs, crc16.CCITTTable)
	case CRCNone:
		return crc16.Checksum(bs, crc16.CCITTFalseTable)
	default: // CRCModbus16
		return crc16Modbus(bs)
	}
}

// crc16Modbus computes the CRC-16/MODBUS checksum (a.k.a. CRC-16/ARC):
// reflected polynomial 0xA001, initial value 0xFFFF, final XOR 0x0000.
//
// It is implemented directly because the third-party IBM/Modbus table yields
// incorrect results: its generic update applies a final XOR of 0xFFFF, which is
// wrong for Modbus (xorout must be 0). Verified against the canonical check
// value CRC-16/MODBUS("123456789") == 0x4B37 and standard request frames.
func crc16Modbus(bs []byte) uint16 {
	crc := uint16(0xFFFF)
	for _, b := range bs {
		crc ^= uint16(b)
		for i := 0; i < 8; i++ {
			if crc&0x0001 != 0 {
				crc = (crc >> 1) ^ 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	return crc
}
