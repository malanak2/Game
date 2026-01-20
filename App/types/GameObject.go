package types

import (
	"reflect"
)

type Update func()

type GameObject struct {
	isActive bool

	update []Update
}

func (g *GameObject) AddUpdateMethod(method Update) {
	g.update = append(g.update, method)
}

func (g *GameObject) RemoveUpdateMethod(method Update) {
	for i := range g.update {
		if reflect.ValueOf(g.update[i]).Pointer() == reflect.ValueOf(method).Pointer() {
			g.update = append(g.update[:i], g.update[i+1:]...)
		}
	}

}

func (g *GameObject) Update() {
	if g.isActive {
		for i := range g.update {
			g.update[i]()
		}
	}
}
