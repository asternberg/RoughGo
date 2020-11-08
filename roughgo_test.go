package main

import "testing"

//TestInitSize: Tests creation of new board from command line args (fixed size or encoded start)
func TestInitSize(t *testing.T) {

	args := []string{"path", "13"} // Intialized with values
	b, blackToMove, err := initBoard(args)

	if err != nil {
		t.Error("Error Initializing:", err)
	}

	if b.Size != 13 {
		t.Error("Error initializing new board with 13 squares")
	}

	if !blackToMove {
		t.Error("Error initializing new board with 13 squares")
	}
	return

}

//TestInitEncodBlack: Tests creation of new board from encoded board
func TestInitEncodBlack(t *testing.T) {

	args := []string{"path", "b", "DWIgDTDhkmAkWpCgUdgBIAAA__8="} // Intialized with encoded board
	b, blackToMove, err := initBoard(args)

	if err != nil {
		t.Error("Error Initializing:", err)
	}

	if b.Size != 13 {
		t.Error("Error initializing new board with 13 squares")
	}

	if !blackToMove {
		t.Error("Error initializing new board with 13 squares")
	}
	return
}

//TestInitEncodeWhite: Tests creation of new board from encoded board
func TestInitEncodeWhite(t *testing.T) {

	args := []string{"path", "w", "DWIgDTDhkmAkWpCgUdgBIAAA__8="} // Intialized with encoded board
	b, blackToMove, err := initBoard(args)

	if err != nil {
		t.Error("Error Initializing:", err)
	}

	if b.Size != 13 {
		t.Error("Error initializing new board with 13 squares")
	}

	if blackToMove {
		t.Error("Error initializing new board with 13 squares")
	}
	return
}
