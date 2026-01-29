package main

import (
	"fmt"
	"os"
	"strings"
)

func parseKeyword(keyword string) string {
	res := "";
	switch strings.ToUpper(keyword) {
	case "HOME"		: res = os.Getenv("HOME");
	case "WAYBAR" : res = "spacer";
	default:
		res = os.Getenv(keyword);
		if res == "" {
			res = "$" + keyword;
		}
	}
	
	return res;
}

func ExpandPath(path string) string {
	for i := 0; i < len(path); i++ {
		if path[i] == '$' {
			start  := i + 1;
			end 	 := start;
			parsed := "";

			for end < len(path) {
				ch := path[end];
				if ch == '/' || ch == '|' || ch == '\\' || ch == '.' {
					parsed = parseKeyword(path[start:end]);
					path = path[:start - 1] + parsed + path[end:];

					i = end - 1;
					break;
				}

				end++;
			}
		}
	}

	return path;
}

func main(){
	// path := "waybar|assets/templates/waybar/$WAYBAR.tmpl";
	path := "$MYEN/waybar";
	ex := ExpandPath(path);
	fmt.Println(ex);
}
