# GoGOL - Conway's Game of Life, written in Golang

This is a very simple implementation of [Conway's Game of
Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) in Go. Nothing
fancy, just to get my feet wet with Go and a little bit of SDL.

No fancy algorithms to be found here, sorry!

## Requirements

This program requires [go-sdl2](https://github.com/veandco/go-sdl2), which in
turn requires [SDL2](http://libsdl.org).

## How to build and run

Simply `go get` this package and then from the package's directory:

```
go build
./gogol
```

## License

This program is distributed under the terms of the [GNU General Public License
Version 3](LICENSE).

This program links against the [go-sdl2](https://github.com/veandco/go-sdl2)
library which is distributed under a [BSD 3-clause
license](https://github.com/veandco/go-sdl2/blob/master/LICENSE).

This program links against the [libsdl2](http://libsdl.org) library, which is
distributed under the terms of the [zlib](http://libsdl.org/license.php)
license.

This program packages the
[WhiteRabbit](http://www.dafont.com/white-rabbit.font) font by Matthew Welch.
WhiteRabbit is distributed under the terms of the license found at
[whitrabt_license.txt](whitrabt_license.txt).

