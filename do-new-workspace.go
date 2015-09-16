package main

import (
	"fmt"
	"log"

	"github.com/proxypoke/i3ipc"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getWorkspaces() []i3ipc.Workspace {
	var i3 *i3ipc.IPCSocket
	var workspaces []i3ipc.Workspace
	var err error

	i3, err = i3ipc.GetIPCSocket()
	must(err)
	defer i3.Close()

	workspaces, err = i3.GetWorkspaces()
	must(err)

	return workspaces

}

func message(action string) {
	var i3 *i3ipc.IPCSocket
	var err error

	i3, err = i3ipc.GetIPCSocket()
	must(err)
	defer i3.Close()

	i3.Command(action)

}

func getFreeWorkspace() int32 {
	var workspaces = getWorkspaces()
	var numbers = make(map[int32]*i3ipc.Workspace)

	for _, w := range workspaces {
		numbers[w.Num] = &w
	}

	var i int32 = 1
	for {
		if numbers[i] == nil {
			return i
		}
		i++
	}

}

func main() {
	fmt.Println("echo blubb")
	var num int32 = getFreeWorkspace()
	var action = fmt.Sprintf("workspace %d", num)

	fmt.Println(action)

	message(action)

}
