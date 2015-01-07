package matasano

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestChallenge1(t *testing.T) {

	teststr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	if stringout, _ := Hex2Base64(teststr); stringout != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error(stringout)
	}
}

func TestChallenge2(t *testing.T) {
	expected := "746865206b696420646f6e277420706c6179"
	if x, _ := Xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"); x != expected {
		t.Error(x)
	}
}

func TestChallenge3(t *testing.T) {
	teststr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	var highest []byte
	var highestscore int

	dict := Dict()

	for _, letter := range "1234567890abcdefghijklmnopqrstuvwxyz" {
		chrkey := hex.EncodeToString([]byte(string(letter)))
		key := strings.Repeat(chrkey, len(teststr)*2)
		xored, err := Xor(teststr, key)
		if err != nil {
			t.Error(err)
		}

		out, err := hex.DecodeString(xored)
		if err != nil {
			t.Error(err)
		}

		if score := ScoreText(out, dict); score > highestscore {
			highest = out
			highestscore = score
		}
	}

	fmt.Printf("Got %s with score %d\n", highest, highestscore)
}
