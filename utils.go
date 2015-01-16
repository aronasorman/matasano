package matasano

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"math"
)

func Dict() [][]byte {
	b, err := ioutil.ReadFile("/etc/dictionaries-common/words")
	if err != nil {
		panic(err)
	}

	words := bytes.Split(b, []byte("\n"))
	return words
}

func ScoreText(t []byte, dict [][]byte) (score int) {
	for _, d := range dict {
		if found := bytes.Index(t, d); found != -1 {
			score++
		}
	}

	return
}

func XorBytes(plaintext, key []byte) ([]byte, error) {
	var out bytes.Buffer
	for i, _ := range plaintext {
		keyindex := int(math.Mod(float64(i), float64(len(key))))
		b := int(plaintext[i]) ^ int(key[keyindex])
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

func SetBitCount(b byte) (bits int) {
	// Brian Kernighan's algorithm for counting set bits
	for i := int(b); i != 0; i = i & (i - 1) {
		bits++
	}
	return
}

func SplitBySize(b []byte, blocksize int) (blocks [][]byte) {
	blocks = [][]byte{}

	buf := bytes.NewBuffer(b)
	for {
		block := make([]byte, blocksize)
		n, err := buf.Read(block)
		if err == io.EOF {
			break
		}

		blocks = append(blocks, block)
		if n < len(block) {
			break
		}
	}

	return
}

func HammingDistance(b1, b2 []byte) (distance int, err error) {
	if len(b1) != len(b2) {
		return 0, errors.New("b1 and b2 should be equal length!")
	}

	xored, err := XorBytes(b1, b2)
	if err != nil {
		return 0, err
	}

	for _, b := range xored {
		distance += SetBitCount(b)
	}

	return
}
