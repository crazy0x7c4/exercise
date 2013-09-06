package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

func (user *user) play() {
	fmt.Println("play a game")
}

type human interface {
	play()
}

func main() {
	reflectBaseType()
}

func reflectBaseType() {
	i := "hello world!"
	iType := reflect.TypeOf(&i)
	iValue := reflect.ValueOf(&i)
	fmt.Println("Type:", iType)
	fmt.Println("Value:", iValue)
}
