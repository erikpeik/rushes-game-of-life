# Game of Life Rush

### Two-day project where the idea was to implement Conway's Game Of Life, and attempt to make implementation as fast as possible.

## Made by emende & vniemi

### Languages used: Go <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" width='30'> and Python <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/python/python-original.svg" width='30'>

### How to install?
Install Go with: `brew install go`

Compile the program with: `make` or manually (introductions down below)

If compiled successfully there should appear 4 executables:
`life, life_animated, life_split2, life_split4`

Usage: `./life initial_state iterations`

## Manual way to compile:
```
go build -o life .

go build -o life_animated other_versions/animated/life_animated.go

go build -o life_split2 other_versions/split2/life_split2.go

go build -o life_split34 other_versions/split4/life_split4.go
```

For comparison you can run the Python version:
`python3 game_of_life.py initial_state iterations`

