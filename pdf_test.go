package pdf

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestReadPdf1 (t *testing.T) {
	f, err := Open("testing/Stagwell_COC.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	t.Logf("%+v", f)

	totalPage := f.NumPage()
	t.Logf("pages: %d", totalPage)
	var buf bytes.Buffer

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := f.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		texts := p.Content().Text
		var lastY = 0.0
		line := ""

		for _, text := range texts {
			if lastY != text.Y {
				if lastY > 0 {
					buf.WriteString(line + "\n")
					line = text.S
				} else {
					line += text.S
				}
			} else {
				line += text.S
			}

			lastY = text.Y
		}
		buf.WriteString(line)
	}
	fmt.Println(buf.String())
}


func TestReadPdf2 (t *testing.T) {
	f, err := Open("testing/voice.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	t.Logf("%+v", f)

	totalPage := f.NumPage()
	t.Logf("pages: %d", totalPage)
	var buf bytes.Buffer

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := f.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		texts := p.Content().Text
		var lastY = 0.0
		line := ""

		for _, text := range texts {
			if lastY != text.Y {
				if lastY > 0 {
					buf.WriteString(line + "\n")
					line = text.S
				} else {
					line += text.S
				}
			} else {
				line += text.S
			}

			lastY = text.Y
		}
		buf.WriteString(line)
	}
	fmt.Println(buf.String())
}


func TestReadPdf3 (t *testing.T) {
	f, err := Open("testing/Schibsted-Mansoor.pdf")
	if err != nil {
		t.Error("Doc should not be nil', got ", err)
	}

	// t.Logf("%+v", f)

	totalPage := f.NumPage()
	t.Logf("pages: %d", totalPage)
	
	plainText, err := f.GetPlainText()
	if err != nil { t.Fatal(err) }

	pbyt, err := io.ReadAll(plainText)
	if err != nil { t.Fatal(err) }

	t.Logf("bytes: %d", len(pbyt))

	fmt.Println(string(pbyt[:100]))

}

