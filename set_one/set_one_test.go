package set_one

import (
	"fmt"
	"testing"
)

func TestChallengeOne(t *testing.T) {
	hex_str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64_str := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if hex_to_base64(hex_str) != base64_str {
		t.Fail()
	}
}

func TestChallengeTwo(t *testing.T) {
	hex_str := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	if bytes_to_hex(FixedXOR(hex_str, key)) != expected {
		t.Fail()
	}
}

func TestChallengeThree(t *testing.T) {
	hex_str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	calculated, _ := BreakSingleByteXOR(hex_to_bytes(hex_str), english_scorer())
	if string(calculated) != expected {
		t.Fail()
	}
}

func TestChallengeFour(t *testing.T) {
	computed := FindXORLine("data/4.txt")
	if string(computed) != "Now that the party is jumping\n" {
		fmt.Println(string(computed))
		t.Fail()
	}
}

func TestChallengeFive(t *testing.T) {
	decrypted := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	expected := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	encrypted := RepeatingKeyXOR(decrypted, key)
	if string(encrypted) != expected {
		t.Fail()
	}
}

func TestChallengeSix(t *testing.T) {
	if HammingDistance("this is a test", "wokka wokka!!!") != 37 {
		t.Fail()
	}

	if get_key_size(b64_file("data/6.txt")) != 29 {
		t.Fail()
	}

	computed := BreakRepeatingKeyXOR(b64_file("data/6.txt"))
	expected := get_file("data/6_sol.txt")
	compare_bytes(computed, expected, t)
}

func TestChallengeSeven(t *testing.T) {
	computed := DecryptECB(b64_file("data/7.txt"), []byte("YELLOW SUBMARINE"))
	expected := get_file("data/7_sol.txt")
	compare_bytes(computed, expected, t)
}

func TestChallengeEight(t *testing.T) {
	expected := "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a"
	computed := DetectECB("data/8.txt")
	if expected != bytes_to_hex(computed) {
		fmt.Println(computed)
		t.Fail()
	}
}
func compare_bytes(computed, expected []byte, t *testing.T) {
	if len(computed) != len(expected) {
		fmt.Printf("Computed has length %d, length %d expected\n", len(computed), len(expected))
		t.Fail()
	}

	for idx, b := range computed {
		if b != expected[idx] {
			fmt.Printf("Computed differs from expected at byte %d\n", idx)
			t.Fail()
		}
	}
}
