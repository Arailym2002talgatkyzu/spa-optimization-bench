package main

import (
	"io"
)
//according to video it will be useful
type User struct {
	Browsers []string `json:"browsers"`
	Company  string   `json:"-"`
	Email    string   `json:"email"`
	Job      string   `json:"-"`
	Name     string   `json:"name"`
	Phone    string   `json:"-"`
}
// вам надо написать более быструю оптимальную этой функции
//уже который раз читаю, смотрю но никак не могу понять где что
//(что именно вообще ищет Search?), и что нужно делать, no idea

func FastSearch(out io.Writer) {
	SlowSearch(out)
}
