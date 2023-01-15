package main

import (
	"fmt"

	"github.com/rknizzle/simplemapreduce/pkg/coordinator"
	"github.com/rknizzle/simplemapreduce/pkg/storage"
)

func main() {
	localFS := storage.LocalFS{}

	c := coordinator.NewCoordinator(localFS)

	_, err := c.StartJob("TODO-TEST-PLUGIN", []string{"TODO-INPUT-1", "TODO-INPUT-2"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", c.Jobs[0])
}
