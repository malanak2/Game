# Game

This is a basic project that renders stuff in a window using opengl and glfw

### TODO:
- Text Rendering
- Kind of modular shape
- Moving stuff around the screen

## Concepts

### Managers
- Managers are global variables, that manage the operational elements of the gameengine
- TextureManager
  - Used for loading textures
  - Manages duplicate loads, etc.
- ShaderManager
  - The same essentially