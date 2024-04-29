package special

import "testing"

func TestSlashB(t *testing.T) {
	output := SlashB("hello\\bthere\\bworld\\b")
	expected := "helltherworld"
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}

	output2 := SlashB("hello\\bthere\\bworld")
	expected2 := "helltherworld"
	if output2 != expected2 {
		t.Errorf("Expected %v got %v", expected2, output2)
	}

	output3 := SlashB("hello\\bthere \\bworld")
	expected3 := "hellthereworld"
	if output3 != expected3 {
		t.Errorf("Expected %v got %v", expected3, output3)
	}
}

func TestSlashZero(t *testing.T) {
	output := SlashZero("hey\\0there\\0 our world \\0")
	expected := "heythere our world "
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}

func TestSlashR(t *testing.T) {
	output := SlashR("hello\\rworld")
	expected := "world"
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}

	output2 := SlashR("hello\\rworldew\\rworldewno\\rhey")
	expected2 := "heyldewno"
	if output2 != expected2 {
		t.Errorf("Expected %v got %v", expected2, output2)
	}

	output3 := SlashR("hello\\rworldew\\rhey\\r")
	expected3 := "heyldew"
	if output3 != expected3 {
		t.Errorf("Expected %v got %v", expected3, output3)
	}
}
