package set_one

import (
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	hex_str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64_str := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if HexToBase64(hex_str) != base64_str {
		t.Fail()
	}
}

func TestFixedXOR(t *testing.T) {
	hex_one := "1c0111001f010100061a024b53535009181c"
	hex_two := "686974207468652062756c6c277320657965"
	result := "746865206b696420646f6e277420706c6179"
	if FixedXOR(hex_one, hex_two) != result {
		t.Fail()
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	decrypted := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	expected := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	encrypted := RepeatingKeyXOR(decrypted, key)
	if encrypted != expected {
		t.Fail()
	}
}

func TestHamming(t *testing.T) {
	first := "this is a test"
	second := "wokka wokka!!!"
	distance := 37
	calculated := HammingDistance(first, second)
	if calculated != distance {
		fmt.Println(calculated)
		t.Fail()
	}
}

func TestGetKeySize(t *testing.T) {
	calculated := GetKeySize(B64File("data/6.txt"))
	expected := 29
	if calculated != expected {
		t.Fail()
	}
}
