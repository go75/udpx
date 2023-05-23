package mod

import (
	"net"
	"github.com/go75/udpx/utils"
	"github.com/go75/udpx/common/e"
)

type Request struct {
	Addr *net.UDPAddr
	Obj
}

func Pack(req *Request) []byte {
	//创建一个存放byte字节数据的缓冲
	buf := make([]byte, 4 + len(req.Payload))

	//1.将id写到buf中
	copy(buf[:4], utils.Uint32ToBytes(req.ID))

	//2.将message的负载写到buf中
	copy(buf[4:], req.Payload)

	return buf
}

func Unpack(data []byte) (*Request, error) {
	if len(data) < 4 {
		return nil, e.DataLenErr
	}
	req := Request{}
	//1.读id
	req.ID = utils.BytesToUint32(data[:4])

	//2.读负载
	req.Payload = data[4:]
	return &req, nil
}