package types

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var KeybindManager KeybindManager_T

func InitKeybindManager() {
	KeybindManager = KeybindManager_T{make(map[glfw.Key]*KeyBind)}
}

type ff func()

type KeybindManager_T struct {
	binds map[glfw.Key]*KeyBind
}

func (k *KeybindManager_T) HandleInput(window *glfw.Window) {
	for a, b := range k.binds {
		b.HandleKeyBind(window.GetKey(a))
	}
}

func (k *KeybindManager_T) AddOnPressed(key glfw.Key, pressed ff) {
	_, ok := k.binds[key]
	if !ok {
		k.binds[key] = NewKeybind()
	}
	k.binds[key].AddOnPressed(pressed)
}

func (k *KeybindManager_T) AddOnHeld(key glfw.Key, held ff) {
	_, ok := k.binds[key]
	if !ok {
		k.binds[key] = NewKeybind()
	}
	k.binds[key].AddOnHeld(held)
}

func (k *KeybindManager_T) AddOnReleased(key glfw.Key, released ff) {
	_, ok := k.binds[key]
	if !ok {
		k.binds[key] = NewKeybind()
	}
	k.binds[key].AddOnReleased(released)
}

type KeyBind struct {
	lastState glfw.Action

	onPressed  []ff
	onHeld     []ff
	onReleased []ff
}

func NewKeybind() *KeyBind {
	return &KeyBind{glfw.Release, make([]ff, 0), make([]ff, 0), make([]ff, 0)}
}

func (k *KeyBind) AddOnPressed(pressed ff) {
	k.onPressed = append(k.onPressed, pressed)
}
func (k *KeyBind) AddOnHeld(held ff) {
	k.onHeld = append(k.onPressed, held)
}
func (k *KeyBind) AddOnReleased(released ff) {
	k.onReleased = append(k.onReleased, released)
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
