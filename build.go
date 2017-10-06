package main

import (
	"bufio"
	"bytes"
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	bt "github.com/mpsejl/behaviortree"
)

type golangBuild struct {
	cli *client.Client
	id  string
	bb  *Blackboard
}

func (t *golangBuild) SetBlackboard(bb *Blackboard) {
	t.bb = bb
}

func (t *golangBuild) Connect() int {
	log.Println("Build - Connect()")
	var err error
	t.cli, err = client.NewEnvClient()
	if err != nil {
		log.Println(err)
		return bt.FAILURE
	}
	return bt.SUCCESS
}

func (t *golangBuild) Create() int {
	log.Println("Build - Create()")
	//var cc *container.Config
	//var hc *container.HostConfig
	//var nc *network.NetworkingConfig

	// Config
	// Volumes
	// Image
	// Cmd
	//

	cc := container.Config{Image: "golang", Cmd: []string{"make", "all"}}
	resp, err := t.cli.ContainerCreate(context.Background(), &cc, nil, nil, "build")
	if err != nil {
		log.Println(err)
		return bt.FAILURE
	}
	t.id = resp.ID
	return bt.SUCCESS
}

func (t *golangBuild) copyTo(file *bytes.Buffer, topath string) bool {
	log.Println("Build - CopyTo()")
	err := t.cli.CopyToContainer(context.Background(), t.id, t.id+":"+topath, bufio.NewReader(file), types.CopyToContainerOptions{})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (t *golangBuild) CopyMakefile() int {
	if t.copyTo(&t.bb.Makefile, "/go/Makefile") {
		return bt.SUCCESS
	}
	return bt.FAILURE
}

func (t *golangBuild) Start() {
	err := t.cli.ContainerStart(context.Background(), t.id, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
}
