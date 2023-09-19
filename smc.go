/** ****************************************************************************************************************** **
	SMC version of what we need for this
	Created by NateDogg - Sept 15, 2023

** ****************************************************************************************************************** **/

package pdf

import (
	
	/*
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	*/

	"github.com/pkg/errors"

	"io"
	"os"
	"os/exec"
	"context"
	"time"
	"fmt"
	"strings"
)

  //-----------------------------------------------------------------------------------------------------------------------//
 //----- CONSTS ----------------------------------------------------------------------------------------------------------//
//-----------------------------------------------------------------------------------------------------------------------//

  //-----------------------------------------------------------------------------------------------------------------------//
 //----- STRUCTS ---------------------------------------------------------------------------------------------------------//
//-----------------------------------------------------------------------------------------------------------------------//


  //-----------------------------------------------------------------------------------------------------------------------//
 //----- FUNCTIONS -------------------------------------------------------------------------------------------------------//
//-----------------------------------------------------------------------------------------------------------------------//

/* 
io.ReaderAt
io.ReadSeeker

ReadAt(p []byte, off int64) (n int, err error)
Read(p []byte) (n int, err error)
Seek(offset int64, whence int) (int64, error)

*/
func PlainText (ctx context.Context, f io.ReaderAt, size int64) (string, error) {

	// create a section reader so everyone is happy
	sr := io.NewSectionReader (f, 0, size)

	// try this library first
	pdfR, err := NewReaderEncrypted(sr, size, func() string { return "" })
	if err == nil {
		plainText, err := pdfR.GetPlainText()
		if err == nil {
			pbyt, _ := io.ReadAll(plainText)

			if len(pbyt) > 0 {
				// fmt.Println ("got it", len(pbyt))
				return strings.TrimSpace (string(pbyt)), nil 
			}
		}
	}
	
	// didn't work, so try another library
	if ctx.Err() != nil { return "", errors.WithStack(ctx.Err()) }

	// create a local file
	flName := fmt.Sprintf("pdf%d.pdf", time.Now().Unix())

	data := make([]byte, size)
	_, err = f.ReadAt (data, 0)
	if err != nil { return "", errors.WithStack (err) }

	// fmt.Println("read", cnt, size)

	err = os.WriteFile (flName, data, 0666)
	if err != nil { return "", errors.WithStack (err) }

	output, err := exec.CommandContext (ctx, "lesspipe", flName).CombinedOutput()

	// remove our file
	os.Remove(flName)

	if err != nil { return "", errors.WithStack (err) }

	// fmt.Println ("worked", len(output))
	// fmt.Println(string(output[:100]))

	return strings.TrimSpace (string(output)), nil // we're good

	/*
	
	// https://cloud.unidoc.io/#/dashboard
	// https://github.com/unidoc/unipdf
	pdfReader, err := model.NewPdfReader(sr)
	if err != nil { return "", errors.WithStack (err) }
	
	isEncrypted, err := pdfReader.IsEncrypted()
	if err == nil && isEncrypted {
		_, err = pdfReader.Decrypt(nil)
		if err != nil { return "", errors.WithStack (err) }
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil { return "", errors.WithStack (err) }
	
	ret := ""
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil { return "", errors.WithStack (err) }
		
		ex, err := extractor.New(page)
		if err != nil { return "", errors.WithStack (err) }
		
		text, err := ex.ExtractText()
		if err != nil { return "", errors.WithStack (err) }
		
		ret = ret + " " + text
	}

	return ret, nil // we're good
	*/
}

// i guess you can't keep firing this, do it once at startup
func InitUniDoc (key string) error {
	// didn't work, so try another library
	// https://cloud.unidoc.io/#/dashboard
	// https://github.com/unidoc/unipdf
	// err := license.SetMeteredKey(key)
	// return errors.WithStack (err)
	return nil 
}
