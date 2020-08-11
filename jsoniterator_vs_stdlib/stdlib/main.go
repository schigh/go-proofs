package main

import (
	"encoding/json"
)

type Foo struct {
	Name string `json:"name"`
}

type Bar struct {
	Name string `json:"name"`
	Foos []Foo  `json:"foos"`
}

type Baz struct {
	Name string `json:"name"`
	Bar  *Bar   `json:"bar"`
}

func (b *Bar) MarshalJSON() ([]byte, error) {
	type alias Bar
	a := (*alias)(b)
	if a.Foos == nil {
		a.Foos = make([]Foo, 0)
	}
	return json.Marshal(a)
}

func (b *Baz) MarshalJSON() ([]byte, error) {
	type alias Baz
	a := (*alias)(b)
	if a.Bar == nil {
		a.Bar = &Bar{}
	}
	return json.Marshal(a)
}

func main() {
	run()
}

func run() {
	var err error
	b := &Bar{Name: "bar"}
	_, err = json.Marshal(b)
	if err != nil {
		panic(err)
	}

	b2 := &Baz{Name: "baz"}
	_, err = json.Marshal(b2)
	if err != nil {
		panic(err)
	}
}
