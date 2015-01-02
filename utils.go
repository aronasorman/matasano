package matasano

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

func Xor(str1, str2 string) string {
	str1base10, _ := hex.DecodeString(str1)
	str2base10, _ := hex.DecodeString(str1)

	var out bytes.Buffer
	for i, _ := range str1base10 {
		b := int(str1base10[i]) ^ int(str2base10[i])
		out.WriteByte(byte(b))
	}

	return out.String()
}

func Hex2Base64(base16str string) (string, error) {
	base10, err := hex.DecodeString(base16str)
	if err != nil {
		return "", err
	}

	var outb bytes.Buffer

	encoder := base64.NewEncoder(base64.StdEncoding, &outb)
	defer encoder.Close()

	_, err = encoder.Write(base10)
	if err != nil {
		return "", err
	}

	return outb.String(), nil
}
