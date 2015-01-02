package matasano

import (
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
	if x := Xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"); x != expected {
		t.Error(x)
	}
}
