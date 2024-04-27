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
		t.Error("Expected hellther world got", output3)
	}
}
