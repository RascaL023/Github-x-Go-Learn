package main

import (
	"fmt"
	"os"
	"parsertry/internal/loader"
	"parsertry/internal/renderer"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument!");
		os.Exit(4);
	}

	name := os.Args[1];
	theme, err := loader.LoadTheme(filepath.Join("themes", name, "theme.json"));
	if err != nil {
		panic("Theme not found!");
	}

	
	tools := []struct {
		Renderer			renderer.Renderer
		TemplatePath  string
		OutputPath    string
	}{
		{
			Renderer: 		theme.Wlogout,
			TemplatePath: "assets/templates/wlogout.tmpl",
			OutputPath:	  "output/wlogout/source.css",
		},
	}

	for _, tool := range tools {
		if err := tool.Renderer.Render(tool.TemplatePath, tool.OutputPath); err != nil {
			panic(err);
		}
	}
}
