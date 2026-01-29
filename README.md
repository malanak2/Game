# Game

This is a basic project that renders stuff in a window using opengl and glfw

## Troubleshooting
If windows says that the app cannot be launched on your device, try compiling with these compiler flags
```
-ldflags="-s -w"
```

### TODO:
- Make some kind of protocol for communication between graphical engine and the glue program and implement it properly
- This would most probably pass config from some kind of global config

## Concepts

### Managers
- Managers are global variables, that manage the operational elements of the gameengine
- TextureManager
  - Used for loading textures
  - Manages duplicate loads, etc.
- ShaderManager
  - The same essentially


## Credits
So many thanks to learnopengl.com, a LOT of the opengl code is taken from there, most of it changed up, but this project would not exist if it was not for this site
https://github.com/Samson-Mano/opengl_textrendering