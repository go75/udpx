package utils

// 大端字节序
func BytesToUint32(data []byte) uint32 {
	return uint32(data[0]) + uint32(data[1]) << 8 + uint32(data[2])<< 16 + uint32(data[3]) << 24
}

// 大端字节序
func Uint32ToBytes(i uint32) []byte {
	return []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)}
}