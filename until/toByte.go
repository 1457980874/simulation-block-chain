package until

import (
	"bytes"
	"encoding/binary"
)

func IntToByte(num int64)([]byte,error){
	buff:=new(bytes.Buffer)
	err:=binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}

func StringToByte(str string)[]byte{
	return []byte(str)
}
