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

	allwords := bytes.Split(b, []byte("\n"))
	filteredwords := make([][]byte, 1000)

	for _, word := range allwords {
		if len(word) >= 4 {
			filteredwords = append(filteredwords, word)
		}
	}

	return filteredwords
}

func ScoreText(t []byte, dict [][]byte) (score int) {
	for _, d := range dict {
		if found := bytes.Index(t, d); found != -1 {
			score++
		}
	}

	return
}

func XorBytes(b1, b2 []byte) ([]byte, error) {
	var out bytes.Buffer
	for i, _ := range b1 {
		b := int(b1[i]) ^ int(b2[i])
		out.WriteByte(byte(b))
	}

	return out.Bytes(), nil
}

func Xor(str1, str2 string) (string, error) {
	str1base10, err := hex.DecodeString(str1)
	if err != nil {
		return "", err
	}
	str2base10, err := hex.DecodeString(str2)
	if err != nil {
		return "", err
	}

	out, err := XorBytes(str1base10, str2base10)

	return hex.EncodeToString(out), nil
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
