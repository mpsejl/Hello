package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	bt "github.com/mpsejl/behaviortree"
)

// Blackboard defines global data for golang build pipeline
type Blackboard struct {
	Name     string
	Fullname string
	Language string
	Repo     string
	Building bool
}

func main() {
	fmt.Println("Hello: Docker test!!!!!\n")

	// Create Makefile
	test1()

	// Connect to Docker and list running containers

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

}

// test1 orchestrates build of golang project using docker golang official image
func test1() {

	bb := Blackboard{Name: "server", Fullname: "mpsejl/server", Language: "Go", Building: false}
	bb.Repo = "github.com/" + bb.Fullname

	root := bt.NewRootNode("Start", nil, false)

	// Makefile
	t := template.Must(template.New("Makefile").Parse(MAKEFILE))

	var tempbuffer bytes.Buffer
	err := t.Execute(&tempbuffer, bb)

	fmt.Println(tempbuffer.String())

}
