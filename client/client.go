package client

import "embed"

// Embed a directory
//
//go:embed dist
var EmbedDirStatic embed.FS
