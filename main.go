package main

// used http://stackoverflow.com/a/15997595/909579

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/builder/parser"
	"os"
)

type Instruction struct {
	Node *parser.Node
}

func (i Instruction) MarshalJSON() ([]byte, error) {
	if i.Node.Next == nil {
		return json.Marshal(&struct {
			Instruction string `json:"instruction"`
		}{
			Instruction: i.Node.Value,
		})
	} else {
		return json.Marshal(&struct {
			Instruction string `json:"instruction"`
			Value       string `json:"value"`
		}{
			Instruction: i.Node.Value,
			Value:       i.Node.Next.Value,
		})
	}
}

func main() {
	var f *os.File
	var err error

	if len(os.Args) < 2 {
		fmt.Println("please supply filename(s)")
		os.Exit(1)
	}

	for _, fn := range os.Args[1:] {
		f, err = os.Open(fn)
		if err != nil {
			panic(err)
		}

		ast, err := parser.Parse(f)
		if err != nil {
			panic(err)
		} else {
			instructions := []Instruction{}
			for _, n := range ast.Children {
				in := Instruction{n}
				instructions = append(instructions, in)
			}
			b, _ := json.Marshal(instructions)
			fmt.Println(string(b))

			fmt.Println(ast.Dump())
		}
	}
}
