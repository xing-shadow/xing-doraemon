package assets

import "embed"

//go:embed build/index.html
var index []byte

//go:embed build/*
var dist embed.FS

func GetStaticAssets() embed.FS {
	return dist
}

func GetIndex() []byte {
	return index
}
