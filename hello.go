package main

import (
	"bytes"
	"fmt"
	"html/template"

	bt "github.com/mpsejl/behaviortree"
)

// Blackboard defines global data for golang build pipeline
type Blackboard struct {
	Name     string
	Fullname string
	Language string
	Repo     string
	Makefile bytes.Buffer
	Building bool
}

func main() {
	fmt.Println("Hello: Docker test!!!!!\n")

	// Create Makefile
	test1()

}

// test1 orchestrates build of golang project using docker golang official image
func test1() {

	bb := Blackboard{Name: "server", Fullname: "mpsejl/server", Language: "Go", Building: false}
	bb.Repo = "github.com/" + bb.Fullname

	createMakefile := func() int {
		t := template.Must(template.New("Makefile").Parse(MAKEFILE))

		err := t.Execute(&bb.Makefile, bb)

		if err != nil {
			fmt.Println("Error Creating Makefile:", err.Error())
			return bt.FAILURE
		}
		fmt.Println("Makefile:", "\n\n", bb.Makefile.String())
		return bt.SUCCESS
	}

	gb := golangBuild{}
	gb.SetBlackboard(&bb)

	root := bt.NewRootNode("Start", nil, false)
	s1 := bt.NewSequenceNode("Build")
	s1.And(bt.NewActionNode(createMakefile)).
		And(bt.NewActionNode(gb.Connect)).
		And(bt.NewActionNode(gb.Create)).
		And(bt.NewActionNode(gb.CopyMakefile))

	root.SetNode(s1)
	root.Run()

}
