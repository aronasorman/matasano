package matasano

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
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

	fmt.Printf("Set 1 Challenge 3: Got %s with score %d\n", highest, highestscore)
}

func TestChallenge4(t *testing.T) {
	filecontent, err := ioutil.ReadFile("4.txt")
	if err != nil {
		panic(err)
	}

	filestrings := bytes.Split(filecontent, []byte("\n"))
	dict := Dict()

	var highestscore int
	var highest string
	strb := make([]byte, 120)

	// pregenerate the keys so we don't have to generate them every time
	keyseeds := []byte("1234567890abcdefghijklmnopqrstuvwxyz")
	keys := make([][]byte, len(keyseeds))
	for i, letter := range keyseeds {
		keys[i] = bytes.Repeat([]byte{letter}, len(filestrings[0]))
	}

	for _, str := range filestrings {
		for _, key := range keys {

			n, _ := hex.Decode(strb, str)

			xored, err := XorBytes(strb[:n], key)
			if err != nil {
				t.Error(err)
			}

			if score := ScoreText(xored, dict); score > highestscore {
				highest = string(xored)
				highestscore = score
			}
		}

	}
	fmt.Printf("Set 1 Challenge 4:Got '%s' with score %d\n", highest, highestscore)
}

func TestChallenge5(t *testing.T) {
	key := []byte("ICE")
	plaintext1 := []byte("Burning 'em, if you ain't quick and nimble")
	expected, err := hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20")
	enc, err := XorBytes(plaintext1, key)
	if err != nil {
		panic(err)
	}

	if string(enc) != string(expected) {
		t.Errorf("%s not equal to %s", string(enc), string(expected))
	}
}

func TestSetBitCount(t *testing.T) {
	testnum := byte(9)
	if numbits := SetBitCount(testnum); numbits != 2 {
		t.Errorf("Got %d set bits; should've gotten 2!", numbits)
	}
}

func TestHammingDistance(t *testing.T) {
	text1 := []byte("this is a test")
	text2 := []byte("wokka wokka!!!")

	diff, err := HammingDistance(text1, text2)
	if err != nil {
		panic(err)
	}

	if diff != 37 {
		t.Errorf("Distance should be 37! Got %d", diff)
	}
}

func BenchmarkScoreText(b *testing.B) {
	dict := Dict()
	t := []byte("cOOKINGmcSLIKEAPOUNDOFBACON")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ScoreText(t, dict)
	}
}
