package puppy

import (
	"github.com/shamimhost/dog"
)

func Bark() string {

	return "Woof!"
}

func Barks() string {

	return "Woof! Woof! Woof!"

}

func BigBark() string {
	return dog.WhenGrowUp(Bark())

}
