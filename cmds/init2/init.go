package main

import (
	"fmt"

	"github.com/u-root/u-root/pkg/init/services/foo"
)

func main() {
	s := foo.New()

	s.Start()
	s.Stop()
	s.Restart()
	s.Reload()
	s.Status()

	fmt.Printf("\nname: %v\n", s.Unit().Name)

}
