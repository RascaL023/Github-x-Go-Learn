package tools

import (
	"parsertry/internal/renderer"
	"parsertry/internal/tools/common"
)

type Tool struct {
	Wlogout common.Wlogout  `json:"wlogout"`;
}

func RegisterTool(t Tool) map[string]renderer.Renderer {
	return map[string]renderer.Renderer{
		"wlogout": t.Wlogout,
	}
}
