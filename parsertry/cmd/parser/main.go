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
	theme, err := loader.LoadToolFromTheme(filepath.Join("themes", name, "theme.json"));
	if err != nil {
		panic(fmt.Sprintf("Theme %s not found!", name));
	}

	wb, err := loader.LoadWaybar(filepath.Join("themes", name, "waybar.json"));
	if err != nil {
		panic(fmt.Sprintf("Waybar Preset %s error!", name));
	}

	tools, err := loader.LoadToolMap("assets/map/path.txt");
	if err != nil {
		panic("Map path not found!");
	}

	state, err := loader.LoadState("assets/map/state.json");
	if err != nil {
		panic("State path not found!");
	}

	registerTools := map[string]renderer.Renderer {
		"wlogout": theme.Wlogout,
		"cava"	 : theme.Cava,
		"waybar" : renderer.WaybarRender{
			Waybar: *wb,
			Preset: state.Waybar,
		},
	}

	
	for _, tool := range tools {
		reg, ok := registerTools[tool.Name];
		if !ok {
			fmt.Println("Unknown tool:", tool.Name);
			continue;
		}

		if err := reg.Render(tool.TemplatePath, tool.OutputPath); err != nil {
			panic(err);
		}
	}
}
