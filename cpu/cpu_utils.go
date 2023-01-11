package cpu

func getMostSignificantByte(value uint16) byte {
	return uint8(value >> 8)
}

func getLeastSignificantByte(value uint16) byte {
	return uint8(value & 0xFF)
}

func setMostSignificantByte(value *uint16, newByte byte) {
	*value = (*value & 0x00FF) | (uint16(newByte) << 8)
}

func setLeastSignificantByte(value *uint16, newByte byte) {
	*value = (*value & 0xFF00) | uint16(newByte)
}
