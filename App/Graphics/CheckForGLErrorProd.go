//go:build !dev

package Graphics

func CheckForGLError() bool {
	return false
}
