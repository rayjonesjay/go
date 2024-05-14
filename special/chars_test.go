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
	output := SlashF("hello\\fworld\\fthere\\fthere")
	expected := "hello\n     world\n          there\n               there"
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}

	output2 := SlashF("hello\\fworld")
	expected2 := "hello\n     world"
	if output2 != expected2 {
		t.Errorf("Expected %v got %v", expected2, output2)
	}

	output3 := SlashF("hello\\fworld\\fhey\\f")
	expected3 := "hello\n     world\n          hey\n             "
	if output3 != expected3 {
		t.Errorf("Expected %v got %v", expected3, output3)
	}
	// fmt.Println(output3)

	output4 := SlashV("hello")
	expected4 := "hello"
	if output4 != expected4 {
		t.Errorf("Expected %v got %v", expected4, output4)
	}

	output5 := SlashV("\\vhello\\v\\v")
	expected5 := "\nhello\n\n     "
	if output5 != expected5 {
		t.Errorf("Expected %v got %v", expected5, output5)
	}

	output6 := SlashV("hello\\v\\v\\vthere")
	expected6 := "hello\n\n\n     there"
	if output6 != expected6 {
		t.Errorf("Expected %v got %v", expected6, output6)
	}
	// fmt.Println(output6)

	output7 := SlashV("\n")
	expected7 := "\n"
	if output7 != expected7 {
		t.Errorf("Expected %v got %v", expected7, output7)
	}

	output8 := SlashV("")
	expected8 := ""
	if output8 != expected8 {
		t.Errorf("Expected %v got %v", expected8, output8)
	}

	output9 := SlashF("g\\fhello")
	expected9 := "g\n hello"
	if output9 != expected9 {
		t.Errorf("Expected %v got %v", expected9, output9)
	}

	output10 := SlashF("hello\\fg")
	expected10 := "hello\n     g"
	if output10 != expected10 {
		t.Errorf("Expected %v got %v", expected10, output10)
	}
}
