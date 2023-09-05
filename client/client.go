package client

import "embed"

// EmbedDirStatic is the embedded directory of the static files.
//
//go:embed dist
var EmbedDirStatic embed.FS
