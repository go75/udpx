package udpx

import (
	"github.com/go75/udpx/common/e"
	"github.com/go75/udpx/utils"
)

type Obj struct {
	ID uint32
	Payload []byte
}

func BytesToObj(data []byte) (Obj, error) {
	obj := Obj{}
	if len(data) < 4 {
		return obj, e.DataLenErr
	}
	obj.ID = utils.BytesToUint32(data[:4])
	obj.Payload = data[4:]
	return obj, nil;
}

func ObjtoBytes(obj Obj) []byte {
	data := make([]byte, len(obj.Payload) + 4)
	copy(data[:4], utils.Uint32ToBytes(obj.ID))
	copy(data[4:], obj.Payload)
	return data
}