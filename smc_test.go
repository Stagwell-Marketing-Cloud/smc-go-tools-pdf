package pdf

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
	"context"
	"os"
)


func TestMoneySMC1 (t *testing.T) {
	f, err := os.Open("testing/Schibsted-Mansoor.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	assert.Equal (t, 221289, len(str))

	fmt.Println(len(str), str[:100])
}

func TestMoneySMC2 (t *testing.T) {
	f, err := os.Open("testing/Stagwell_COC.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	assert.Equal (t, 65604, len(str))

	fmt.Println(len(str), str[:100])
}

func TestMoneySMC3 (t *testing.T) {
	f, err := os.Open("testing/voice.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str)
}

func TestMoneySMC4 (t *testing.T) {
	f, err := os.Open("testing/10840.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str[:100])
}

func TestMoneySMC5 (t *testing.T) {
	f, err := os.Open("testing/crash.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str[:100])
}

func TestMoneySMC5a (t *testing.T) {
	f, err := os.Open("testing/twilio.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str), str[:100])
}

// this file gets stuck in a loop
func TestMoneySMC6a (t *testing.T) {
	f, err := os.Open("testing/ACXIOM_CX_Predictions_2024.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	sz, err := f.Stat()
	if err != nil { t.Fatal(err) }
	
	str, err := PlainText (context.Background(), f, sz.Size())
	if err != nil { t.Fatal(err) }

	fmt.Println(len(str)) //, str[:100])
}
