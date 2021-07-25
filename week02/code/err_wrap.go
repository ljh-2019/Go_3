package code

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := rap2()
	fmt.Printf("%+v", err)
}

func rap1() error {
	return errors.New("test err")
}

func rap2() error {
	return rap1()
}
