package types

import "github.com/go-gl/glfw/v3.3/glfw"

var KeybindManager KeybindManager_T

func InitKeybindManager() {
	KeybindManager = KeybindManager_T{make(map[glfw.Key]*KeyBind)}
}

func (k *KeybindManager_T) ProcessInput(window *glfw.Window) {

}

type ff func()

type KeybindManager_T struct {
	binds map[glfw.Key]*KeyBind
}

func (k *KeybindManager_T) AddOnPressed(key glfw.Key, pressed ff) {
	k.binds[key].AddOnPressed(pressed)
}

func (k *KeybindManager_T) AddOnHeld(key glfw.Key, held ff) {
	k.binds[key].AddOnHeld(held)
}

func (k *KeybindManager_T) AddOnReleased(key glfw.Key, released ff) {
	k.binds[key].AddOnReleased(released)
}

type KeyBind struct {
	lastState glfw.Action

	onPressed  []ff
	onHeld     []ff
	onReleased []ff
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
