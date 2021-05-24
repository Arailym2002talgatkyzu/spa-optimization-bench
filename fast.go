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
//I can make only pprof, вроде regex занимает больше времени

func FastSearch(out io.Writer) {
	SlowSearch(out)
}
