package matasano

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

func Dict() [][]byte {
	b, err := ioutil.ReadFile("/etc/dictionaries-common/words")
	if err != nil {
		panic(err)
	}

	return bytes.Split(b, []byte("\n"))
}

func ScoreText(t []byte, dict [][]byte) int {
	var score int
	for _, d := range dict {
		if found := bytes.Index(t, d); found != -1 {
			score++
		}
	}

	return score
}

func Xor(str1, str2 string) (string, error) {
	str1base10, err := hex.DecodeString(str1)
	if err != nil {
		return "", err
	}
	str2base10, _ := hex.DecodeString(str2)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	for i, _ := range str1base10 {
		b := int(str1base10[i]) ^ int(str2base10[i])
		out.WriteByte(byte(b))
	}

	return hex.EncodeToString(out.Bytes()), nil
}

func ToBase64(base10 []byte) (string, error) {
	var outb bytes.Buffer

	encoder := base64.NewEncoder(base64.StdEncoding, &outb)
	defer encoder.Close()

	_, err := encoder.Write(base10)
	if err != nil {
		return "", err
	}

	return outb.String(), nil
}

func Hex2Base64(base16str string) (string, error) {
	base10, err := hex.DecodeString(base16str)
	if err != nil {
		return "", err
	} else {
		return ToBase64(base10)
	}
}
