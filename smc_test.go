package pdf

import (
	"fmt"
	"testing"
	"os"
)


func TestMoneySMC1 (t *testing.T) {
	f, err := os.Open("testing/Schibsted-Mansoor.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str)
}

func TestMoneySMC2 (t *testing.T) {
	f, err := os.Open("testing/Stagwell_COC.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str)
}

func TestMoneySMC3 (t *testing.T) {
	f, err := os.Open("testing/voice.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str)
}
