package docseq

import (
	"testing"
)

func TestDb(t *testing.T) {
	t.Run("TestDbConnection", func(t *testing.T) {
		_, err := dbConnection("demo")

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestValidateGetNumber(t *testing.T) {
	t.Run("TestValidateGetNumber", func(t *testing.T) {
		
		// negative test cases for client, org, docid, and doctype are required

		_, err := Next("", "", "", "", "", "", "")
		if err != nil {
			if err.Error() != "client is empty" {
				t.Errorf("Expecting 'client is empty', got %v", err)
			}
		}

		_, err = Next("demo", "", "", "", "", "", "")
		if err != nil {
			if err.Error() != "org is empty" {
				t.Errorf("Expecting 'org is empty', got %v", err)
			}
		}

		// seqDocId is not optional
		_, err = Next("demo", "org", "", "", "", "", "")
		if err != nil {
			if err.Error() != "doctype is empty" {
				t.Errorf("Expecting 'doctype is empty', got %v", err)
			}
		}

		// seqDocId is optional
		// DocMustExist = false
		// _, err = Next("demo", "org", "", "", "", "", "")
		// if err != nil {
		// 	if err.Error() != "doctype is empty" {
		// 		t.Errorf("Expecting 'doctype is empty', got %v", err)
		// 	}
		// }

		// prefix and suffix are optional and not tested here
	})
}

// DO NOT DELETE THIS COMMENTED TEST CODE
//
// the tests are commented because it modifies the database value everytime it runs..
// uncomment the tests to run them but be sure you are ready to fix the database values after running the tests

func TestGetNumber(t *testing.T) {
	t.Run("TestGetNumber", func(t *testing.T) {
		uri, err := dbConnection("demo")
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		_, err = getNumber(uri, "jfckl2", "default", "HH")

		if err != nil {
			t.Errorf("Error result from getNumber: %v", err)
		}
	})
}

// func TestNext(t *testing.T) {
// 	t.Run("TestNewSequence", func(t *testing.T) {

// 		doctype := "GRN"
// 		expected := "jfckl2:" + doctype + "/AA/2024/4/30/1000005"
// 		prefix := "jfckl2:" + doctype + "/AA/2024/4/30"

// 		val, err := Next("demo", "jfckl2", "AA", doctype, prefix, "", "/")

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		if len(val) == 0 {
// 			t.Errorf("Expecting a value, got empty string")
// 		}

// 		if val != expected {
// 			t.Errorf("Expecting %v, got %v", expected, val)
// 		}

// 		// without separator
// 		expected = "jfckl2:" + doctype + "AA/2024/4/301000006"
// 		prefix = "jfckl2:" + doctype + "AA/2024/4/30"

// 		val, err = Next("demo", "jfckl2", "AA", doctype, prefix, "", "")

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		if len(val) == 0 {
// 			t.Errorf("Expecting a value, got empty string")
// 		}

// 		if val != expected {
// 			t.Errorf("Expecting %v, got %v", expected, val)
// 		}
// 	})
// }