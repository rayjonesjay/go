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

	output4 := SlashB("hello")
	expected4 := "hello"
	if output4 != expected4 {
		t.Errorf("Expected %v got %v", expected4, output4)
	}
}

func TestSlashZero(t *testing.T) {
	output := SlashZero("hey\\0there\\0 our world \\0")
	expected := "heythere our world "
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}

	output2 := SlashZero("hey")
	expected2 := "hey"
	if output2 != expected2 {
		t.Errorf("Expected %v got %v", expected2, output2)
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

	output4 := SlashR("hello")
	expected4 := "hello"
	if output4 != expected4 {
		t.Errorf("Expected %v got %v", expected4, output4)
	}
}

func TestSlashFSlashV(t *testing.T) {
	output := SlashFSlashV("hello\fworld\fthere\vthere")
	expected := "hello\n     world\n          there\n               there"
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}

	output2 := SlashFSlashV("hello\fworld")
	expected2 := "hello\n     world"
	if output2 != expected2 {
		t.Errorf("Expected %v got %v", expected2, output2)
	}

	output3 := SlashFSlashV("hello\fworld\fhey\f")
	expected3 := "hello\n     world\n          hey\n             "
	if output3 != expected3 {
		t.Errorf("Expected %v got %v", expected3, output3)
	}

	output4 := SlashFSlashV("hello")
	expected4 := "hello"
	if output4 != expected4 {
		t.Errorf("Expected %v got %v", expected4, output4)
	}
}
