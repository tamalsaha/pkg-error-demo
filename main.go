package main

import (
	"github.com/pkg/errors"
	"fmt"
)

func f2 () error {
	return errors.WithStack(fmt.Errorf("test err %s", "abc"))
}

func main() {
	panic(f2())
}