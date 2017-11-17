package go_tencent_defend_tlv

import (
	"testing"
	"encoding/base64"
)

func TestEncode(t *testing.T) {
	buf, err := Encode(1, "1")
	if err != nil {
		t.Error(err)
		return
	}
	if base64.StdEncoding.EncodeToString(buf) != "AAAAAQAAAAEx" {
		t.Error("编码ASCII数据失败")
		return
	}
	buf, err = Encode(1, "天气不错")
	if err != nil {
		t.Error(err)
		return
	}

	if base64.StdEncoding.EncodeToString(buf) != "AAAAAQAAAAzlpKnmsJTkuI3plJk=" {
		t.Error("编码UTF8数据失败")
		return
	}
}

func TestDecode(t *testing.T) {
	buf, err := base64.StdEncoding.DecodeString("AAAAAQAAAAzlpKnmsJTkuI3plJk=")
	if err != nil {
		t.Error(err)
		return
	}
	tag, data, err := Decode(buf)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tag, data)
}
