package main

import "fmt"

type WorldInteface interface {
	Callme(t *WorldInteface) string
}

type Functions struct {
	World func() WorldInteface
}

type MyWorld struct {
	WorldName string
}

func NewWorldImpl() *MyWorld {
	return &MyWorld{
		WorldName: "Happy",
	}
}

func (w *MyWorld) Callme(t *MyWorld) string {
	return t.WorldName + " World"
}

func main() {
	functions := &Functions{
		World: func() WorldInteface {
			return NewWorldImpl()
		},
	}

	w := functions.World()
	r := w.Callme(&w)

	fmt.Printf("callme return: %s\n", r)
}
