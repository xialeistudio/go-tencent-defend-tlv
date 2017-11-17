package go_tencent_defend_tlv

import (
	"bytes"
	"encoding/binary"
)

func Encode(tag int32, data string) ([]byte, error) {
	buf := new(bytes.Buffer)
	// 写入TAG
	if err := binary.Write(buf, binary.BigEndian, tag); err != nil {
		return nil, err
	}
	dataBuf := []byte(data)
	// 写入length
	if err := binary.Write(buf, binary.BigEndian, int32(len(dataBuf))); err != nil {
		return nil, err
	}
	// 写入数据
	if err := binary.Write(buf, binary.BigEndian, dataBuf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (int32, string, error) {
	buf := bytes.NewBuffer(b)
	var tag, length int32
	// 读取tag
	if err := binary.Read(buf, binary.BigEndian, &tag); err != nil {
		return 0, "", err
	}
	// 读取length
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return 0, "", err
	}
	// 读取数据
	dataBuf := make([]byte, length)
	if err := binary.Read(buf, binary.BigEndian, &dataBuf); err != nil {
		return 0, "", err
	}
	return tag, string(dataBuf), nil
}
