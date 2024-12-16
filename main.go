package main

import (
	"pokedexgo/internal/pokecache"
)

var cache *pokecache.Cache

func main() {
	cache = pokecache.NewCache(5)
	startREPL()
}
