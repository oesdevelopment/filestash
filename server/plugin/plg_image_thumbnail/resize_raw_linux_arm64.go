package plg_image_thumbnail

import (
	_ "embed"
)

//go:embed dist/raw_linux_arm64.bin
var binaryThumbnailRaw []byte

//go:embed dist/raw_linux_arm64.bin.sha256
var checksumRaw []byte
