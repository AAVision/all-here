package main

import (
	"github.com/alexflint/go-arg"
)

type UserInput struct {
	Port string `arg:"-p"`
}

func main() {
	/**
	This is main Function.
	@var nil.
	@return nil.
	*/
	userInput := UserInput{}
	arg.MustParse(&userInput)

	starter(userInput)

}
