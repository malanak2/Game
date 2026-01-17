package types

import (
	"errors"
	"log/slog"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var KeybindManager KeybindManagerT

func InitKeybindManager() {
	KeybindManager = KeybindManagerT{make(map[glfw.Key]*KeyBind)}
}

type Func func()

type KeybindManagerT struct {
	binds map[glfw.Key]*KeyBind
}

func (k *KeybindManagerT) HandleInput(window *glfw.Window) {
	for a, b := range k.binds {
		b.HandleKeyBind(window.GetKey(a))
	}
}

func (k *KeybindManagerT) AddOnPressed(key glfw.Key, pressed Func) int {
	_, ok := k.binds[key]
	if !ok {
		k.binds[key] = NewKeybind()
	}
	return k.binds[key].AddOnPressed(pressed)
}

func (k *KeybindManagerT) AddOnHeld(key glfw.Key, held Func) int {
	_, ok := k.binds[key]
	if !ok {
		k.binds[key] = NewKeybind()
	}
	return k.binds[key].AddOnHeld(held)
}

func (k *KeybindManagerT) AddOnReleased(key glfw.Key, released Func) int {
	_, ok := k.binds[key]
	if !ok {
		k.binds[key] = NewKeybind()
	}
	return k.binds[key].AddOnReleased(released)
}

func (k *KeybindManagerT) RemoveOnPressed(key glfw.Key, funcIndex int) error {
	_, ok := k.binds[key]
	if !ok {
		slog.Error("RemoveOnPressed", "err", "KeyBind not initialized")
		return errors.New("key not initialized")
	}
	return k.binds[key].RemoveOnPressed(funcIndex)
}

func (k *KeybindManagerT) RemoveOnHeld(key glfw.Key, funcIndex int) error {
	_, ok := k.binds[key]
	if !ok {
		slog.Error("RemoveOnHeld", "err", "KeyBind not initialized")
		return errors.New("key not initialized")
	}
	return k.binds[key].RemoveOnHeld(funcIndex)
}

func (k *KeybindManagerT) RemoveOnReleased(key glfw.Key, funcIndex int) error {
	_, ok := k.binds[key]
	if !ok {
		slog.Error("RemoveOnReleased", "err", "KeyBind not initialized")
		return errors.New("key not initialized")
	}
	return k.binds[key].RemoveOnReleased(funcIndex)
}

type KeyBind struct {
	lastState glfw.Action

	onPressed  []Func
	onHeld     []Func
	onReleased []Func
}

func NewKeybind() *KeyBind {
	return &KeyBind{glfw.Release, make([]Func, 0), make([]Func, 0), make([]Func, 0)}
}

func (k *KeyBind) AddOnPressed(pressed Func) int {
	k.onPressed = append(k.onPressed, pressed)
	return len(k.onPressed) - 1
}

func (k *KeyBind) AddOnHeld(held Func) int {
	k.onHeld = append(k.onPressed, held)
	return len(k.onHeld) - 1
}

func (k *KeyBind) AddOnReleased(released Func) int {
	k.onReleased = append(k.onReleased, released)
	return len(k.onReleased) - 1
}

func (k *KeyBind) RemoveOnPressed(index int) error {
	slog.Info("RemoveOnPressed", "msg", "Removing function from OnPressed", "index", index)
	if index < 0 || index >= len(k.onPressed) {
		slog.Error("RemoveOnPressed", "err", "Out of bounds error", "index", index)
		return errors.New("index out of bounds")
	}
	if len(k.onPressed) == 1 {
		k.onPressed = make([]Func, 0)
		return nil
	}

	if len(k.onPressed) == index+1 {
		k.onPressed = k.onPressed[:index]
		return nil
	}
	k.onPressed = append(k.onPressed[:index], k.onPressed[index+1:]...)
	return nil
}

func (k *KeyBind) RemoveOnHeld(index int) error {
	k.onHeld = append(k.onHeld[:index], k.onHeld[index+1:]...)
	slog.Info("RemoveOnHeld", "msg", "Removing function from OnHeld", "index", index)
	if index < 0 || index >= len(k.onHeld) {
		slog.Error("RemoveOnHeld", "err", "Out of bounds error", "index", index)
		return errors.New("index out of bounds")
	}
	if len(k.onHeld) == 1 {
		k.onHeld = make([]Func, 0)
		return nil
	}

	if len(k.onHeld) == index+1 {
		k.onHeld = k.onHeld[:index]
		return nil
	}
	k.onHeld = append(k.onHeld[:index], k.onHeld[index+1:]...)
	return nil
}

func (k *KeyBind) RemoveOnReleased(index int) error {
	slog.Info("RemoveOnReleased", "msg", "Removing function from OnReleased", "index", index)
	if index < 0 || index >= len(k.onReleased) {
		slog.Error("RemoveOnReleased", "err", "Out of bounds error", "index", index)
		return errors.New("index out of bounds")
	}
	if len(k.onReleased) == 1 {
		k.onReleased = make([]Func, 0)
		return nil
	}

	if len(k.onReleased) == index+1 {
		k.onReleased = k.onReleased[:index]
		return nil
	}
	k.onReleased = append(k.onReleased[:index], k.onReleased[index+1:]...)
	return nil
}

func (k *KeyBind) HandleKeyBind(key glfw.Action) {
	if k.lastState == glfw.Press { // Was
		k.lastState = key
		for _, v := range k.onHeld {
			v()
		}
	} else if key == glfw.Release {
		k.lastState = key
		for _, v := range k.onReleased {
			v()
		}
	} else {
		k.lastState = key
		for _, v := range k.onPressed {
			v()
		}
	}
}
