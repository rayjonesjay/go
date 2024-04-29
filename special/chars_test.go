package special_chars

import "testing"

func TestSlashB(t *testing.T) {
	output := SlashB("hello\\bthere\\bworld\\b")
	if output != "helltherworld" {
		t.Error("Expected helltherworld got", output)
	}

	output2 := SlashB("hello\\bthere\\bworld")
	if output2 != "helltherworld" {
		t.Error("Expected helltherworld got", output2)
	}

	output3 := SlashB("hello\\bthere \\bworld")
	if output3 != "hellthereworld" {
		t.Error("Expected hellthereworld got", output3)
	}
}

func TestSlashZero(t *testing.T) {
	output := SlashZero("hey\\0there\\0 our world \\0")
	if output != "heythere our world " {
		t.Error("Expected heythere our world  got", output)
	}
}

func TestSlashR(t *testing.T) {
	output := SlashR("hello\\rworld")
	if output != "world" {
		t.Error("Expected world got", output)
	}

	output2 := SlashR("hello\\rworldew\\rworldewno\\rhey")
	if output2 != "heyldewno" {
		t.Error("Expected heyldewno got", output2)
	}

	output3 := SlashR("hello\\rworldew\\rhey\\r")
	if output3 != "heyldew" {
		t.Error("Expected heyldew got", output3)
	}
}
