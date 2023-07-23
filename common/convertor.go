package common

import (
	"encoding/binary"
)

func BytesToUint32LE(sliceOfBytes []byte) uint32 {
	if len(sliceOfBytes) != 0 {
		valueUint32 := binary.LittleEndian.Uint32(sliceOfBytes)
		return valueUint32
	}
	return 0
}
