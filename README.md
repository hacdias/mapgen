# Map Generator

> Create beautiful and pseudo-random terrain-like maps.

![a sample mat bit](https://user-images.githubusercontent.com/5447088/49112619-27fbf600-f28b-11e8-82e6-78198e65929a.png)

[![Build](https://img.shields.io/travis/com/hacdias/mapgen.svg?style=flat-square)](https://travis-ci.com/hacdias/mapgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/hacdias/mapgen?style=flat-square)](https://goreportcard.com/report/hacdias/mapgen)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/hacdias/mapgen)

## Install

To use it as a library, you need to get it through:

```
go get github.com/hacdias/mapgen
```

Soon, more documentation will be added about the final executable.

## Usage

Documentation about using it as a library can be [found on godoc](http://godoc.org/github.com/hacdias/mapgen).

```
Usage of mapgen:
  -t, --color transition    Color Transition in map generation overrites color palette transition
  -c, --colors filename string   File name with the color palette (default "defaultPalette.txt")
  -f, --filename string     File name to output (default "img.png")
  -h, --height int          Height of image (default 100)
  -l, --lacunarity float    Lacunarity (default 2.5)
  -o, --octaves int         Number of octaves (default 5)
  -p, --persistence float   persistence (default 0.5)
  -x, --scale float         Scale (default 20)
  -s, --seed int            Seed to generate the map (default random)
  -w, --width int           Width of image (default 100)
```

## Contributing

PRs accepted.

## License

MIT © [All Contributors](https://github.com/hacdias/mapgen/graphs/contributors)
