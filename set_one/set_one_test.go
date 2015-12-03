package set_one

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func GetFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return data
}

func TestFindXORLine(t *testing.T) {
	expected := "Now that the party is jumping\n"
	computed := FindXORLine("data/4.txt")
	if expected != computed {
		fmt.Println(computed)
		t.Fail()
	}
}

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

func TestDecryptECB(t *testing.T) {
	ciphertext := B64File("data/7.txt")
	key := []byte("YELLOW SUBMARINE")
	cleartext := DecryptECB(ciphertext, key)
	expected := GetFile("data/7_sol.txt")
	if string(cleartext) != string(expected) {
		t.Fail()
	}
}

func TestDetectECB(t *testing.T) {
	expected := "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a"
	computed := DetectECB("data/8.txt")
	if expected != computed {
		fmt.Println(computed)
		t.Fail()
	}
}
