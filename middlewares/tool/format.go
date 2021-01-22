package tool

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func Int64ToInt(i64 int64) (i int) {
	str := strconv.FormatInt(i64, 10)
	i, _ = strconv.Atoi(str)
	return
}

//字节数(大端)组转成int(无符号的)
func BytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}
