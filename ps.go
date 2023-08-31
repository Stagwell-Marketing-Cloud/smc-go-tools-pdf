// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pdf

import (
	"fmt"
	"io"
)

// A Stack represents a stack of values.
type Stack struct {
	stack []Value
}

func (stk *Stack) Len() int {
	return len(stk.stack)
}

func (stk *Stack) Push(v Value) {
	stk.stack = append(stk.stack, v)
}

func (stk *Stack) Pop() Value {
	n := len(stk.stack)
	if n == 0 {
		return Value{}
	}
	v := stk.stack[n-1]
	stk.stack[n-1] = Value{}
	stk.stack = stk.stack[:n-1]
	return v
}

func newDict() Value {
	return Value{nil, objptr{}, make(dict)}
}

// Interpret interprets the content in a stream as a basic PostScript program,
// pushing values onto a stack and then calling the do function to execute
// operators. The do function may push or pop values from the stack as needed
// to implement op.
//
// Interpret handles the operators "dict", "currentdict", "begin", "end", "def", and "pop" itself.
//
// Interpret is not a full-blown PostScript interpreter. Its job is to handle the
// very limited PostScript found in certain supporting file formats embedded
// in PDF files, such as cmap files that describe the mapping from font code
// points to Unicode code points.
//
// There is no support for executable blocks, among other limitations.
//
func Interpret(strm Value, do func(stk *Stack, op string)) {
	fmt.Println("interpret A")
	rd := strm.Reader()
	fmt.Println("interpret B")
	b := newBuffer(rd, 0)
	fmt.Println("interpret C")
	b.allowEOF = true
	b.allowObjptr = false
	b.allowStream = false
	var stk Stack
	var dicts []dict
	fmt.Println("interpret D")

Reading:
	
	for {
		fmt.Println("interpret E")
		tok := b.readToken()
		fmt.Println("interpret E1")
		if tok == io.EOF {
			break
		}
		if kw, ok := tok.(keyword); ok {
			fmt.Println("interpret F")
			switch kw {
			case "null", "[", "]", "<<", ">>":
				break
			default:
				fmt.Println("interpret H")
				for i := len(dicts) - 1; i >= 0; i-- {
					fmt.Println("interpret H", i)
					if v, ok := dicts[i][name(kw)]; ok {
						stk.Push(Value{nil, objptr{}, v})
						continue Reading
					}
				}
				do(&stk, string(kw))
				continue
			case "dict":
				fmt.Println("interpret G")
				stk.Pop()
				stk.Push(Value{nil, objptr{}, make(dict)})
				continue
			case "currentdict":
				fmt.Println("interpret I")
				if len(dicts) == 0 {
					panic("no current dictionary")
				}
				stk.Push(Value{nil, objptr{}, dicts[len(dicts)-1]})
				continue
			case "begin":
				fmt.Println("interpret J")
				d := stk.Pop()
				if d.Kind() != Dict {
					panic("cannot begin non-dict")
				}
				dicts = append(dicts, d.data.(dict))
				continue
			case "end":
				fmt.Println("interpret K")
				if len(dicts) <= 0 {
					panic("mismatched begin/end")
				}
				dicts = dicts[:len(dicts)-1]
				continue
			case "def":
				fmt.Println("interpret L")
				if len(dicts) <= 0 {
					panic("def without open dict")
				}
				val := stk.Pop()
				key, ok := stk.Pop().data.(name)
				if !ok {
					panic("def of non-name")
				}
				dicts[len(dicts)-1][key] = val.data
				continue
			case "pop":
				fmt.Println("interpret M")
				stk.Pop()
				continue
			}
		}
		fmt.Println("interpret N")
		b.unreadToken(tok)
		fmt.Println("interpret O")
		obj, err := b.readObject()
		fmt.Println("interpret P")
		if err != nil {
			return
		}
		fmt.Println("interpret Q")
		stk.Push(Value{nil, objptr{}, obj})
	}
}

type seqReader struct {
	rd     io.Reader
	offset int64
}

func (r *seqReader) ReadAt(buf []byte, offset int64) (int, error) {
	if offset != r.offset {
		return 0, fmt.Errorf("non-sequential read of stream")
	}
	n, err := io.ReadFull(r.rd, buf)
	r.offset += int64(n)
	return n, err
}
