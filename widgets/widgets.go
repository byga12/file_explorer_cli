package tui

import (
	"fmt"
	"io"
	"io/fs"
)

const SAVE_CURSOR = "\x1b[s"
const RESTORE_CURSOR = "\x1b[u"
const RED_FONT = "\x1b[31m"
const CYAN_FONT = "\x1b[36m"
const RESET_TEXT = "\x1b[0m"
const INVISIBLE = "\x1b[?25l"
const VISIBLE = "\x1b[?25h"
const DELETE_LINE = "\x1b[K"

type SearchBar string
type FileViewer string


func (sb SearchBar) Render(out io.Writer, keyword string){
	fmt.Fprintf(out ," 🔍 %s", keyword)	
}

func (fv FileViewer) Render(out io.Writer, de []fs.DirEntry, selectedIndex int){
	if selectedIndex < 0 {
		fmt.Fprintf(out, "   (empty)\n")
		return
	}
	if selectedIndex < 0 || selectedIndex >= len(de) {
		return
	}
	fmt.Fprintf(out, "   Seleccionado: %s\n", de[selectedIndex].Name())
	for index, dirEntry := range de{
		fmt.Fprintf(out, "  ")
		if(index==selectedIndex){
			fmt.Fprintf(out, "%s> %s%s \n",CYAN_FONT, dirEntry.Name(), RESET_TEXT)
		} else {
			fmt.Fprintf(out, "  %s \n", dirEntry.Name())
		}
	}
}

