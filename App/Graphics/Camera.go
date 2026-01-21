package Graphics

import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	config2 "github.com/malanak2/Game/App/config"

	"github.com/go-gl/mathgl/mgl32"
)

type CameraDirection int

const (
	CameraForward CameraDirection = iota
	CameraBackward
	CameraLeft
	CameraRight
	CameraUp
	CameraDown
)

type CameraT struct {
	Pos     mgl32.Vec3
	Front   mgl32.Vec3
	Up      mgl32.Vec3
	Right   mgl32.Vec3
	WorldUp mgl32.Vec3

	MoveSpeed float32

	Yaw   float32
	Pitch float32
	Zoom  float32

	ProjectionMatrix mgl32.Mat4
	ViewMatrix       mgl32.Mat4

	width  int
	height int

	lastX float32
	lastY float32

	sensitivity float32
}

var Camera CameraT

func InitCamera(pos mgl32.Vec3, up mgl32.Vec3) {
	Camera = CameraT{
		pos,
		mgl32.Vec3{0, 0, -1},
		up,
		mgl32.Vec3{},
		up,
		config2.Cfg.Main.CameraMovespeed,
		-90,
		0,
		45,
		mgl32.Mat4{},
		mgl32.Mat4{},
		1920,
		1080,
		960,
		540,
		0.1,
	}
	Camera.updateCameraVectors()
	Camera.Calculate()
}

func (c *CameraT) UpdateScreen(width int, height int) {
	c.width = width
	c.height = height
	c.Calculate()
}

func (c *CameraT) updateCameraVectors() {
	var front mgl32.Vec3
	front[0] = float32(math.Cos(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch))))
	front[1] = float32(math.Sin(float64(mgl32.DegToRad(c.Pitch))))
	front[2] = float32(math.Sin(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch))))
	c.Front = front.Normalize()
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}

func (c *CameraT) MoveCamera(direction CameraDirection, deltaTime float32) {
	velocity := c.MoveSpeed * deltaTime
	switch direction {
	case CameraForward:
		c.Pos = c.Pos.Add(c.Front.Mul(velocity))
	case CameraBackward:
		c.Pos = c.Pos.Sub(c.Front.Mul(velocity))
	case CameraLeft:
		c.Pos = c.Pos.Sub(c.Right.Mul(velocity))
	case CameraRight:
		c.Pos = c.Pos.Add(c.Right.Mul(velocity))
	case CameraUp:
		c.Pos = c.Pos.Add(c.Up.Mul(velocity))
	case CameraDown:
		c.Pos = c.Pos.Sub(c.Up.Mul(velocity))
	}
}

func (c *CameraT) Calculate() {
	c.ProjectionMatrix = mgl32.Perspective(mgl32.DegToRad(c.Zoom), float32(c.width)/float32(c.height), 0.1, 100)
	c.ViewMatrix = mgl32.LookAtV(c.Pos, c.Pos.Add(c.Front), c.Up)
}

func MouseCallback(window *glfw.Window, xposi float64, yposi float64) {
	xpos := float32(xposi)
	ypos := float32(yposi)
	xoffset := xpos - Camera.lastX
	yoffset := ypos - Camera.lastY
	Camera.lastX = xpos
	Camera.lastY = ypos
	xoffset *= Camera.sensitivity
	yoffset *= Camera.sensitivity
	Camera.Yaw += xoffset
	Camera.Pitch -= yoffset
	if Camera.Pitch > 89.0 {
		Camera.Pitch = 89.0
	}
	if Camera.Pitch < -89.0 {
		Camera.Pitch = -89.0
	}
	direction := mgl32.NewVecN(3)
	direction.Set(0, float32(math.Cos(float64(mgl32.DegToRad(Camera.Yaw)))*math.Cos(float64(mgl32.DegToRad(Camera.Pitch)))))
	direction.Set(1, float32(math.Sin(float64(mgl32.DegToRad(Camera.Pitch)))))
	direction.Set(2, float32(math.Sin(float64(mgl32.DegToRad(Camera.Yaw)))*math.Cos(float64(mgl32.DegToRad(Camera.Pitch)))))
	Camera.Front = direction.Vec3().Normalize()

	Camera.updateCameraVectors()
	Camera.Calculate()
}

func ScrollWheelCallback(window *glfw.Window, xOffset float64, yoffset float64) {
	fov := Camera.Zoom
	fov -= float32(yoffset)
	if fov < 1.0 {
		fov = 1.0
	}
	if fov > 45.0 {
		fov = 45.0
	}
	Camera.Zoom = fov
	Camera.Calculate()
}
