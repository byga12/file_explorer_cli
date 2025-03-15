package tui

import (
	"fmt"
	"io"
	"io/fs"
	"strconv"
)

const SAVE_CURSOR = "\x1b[s"
const RESTORE_CURSOR = "\x1b[u"
const RED_FONT = "\x1b[31m"
const CYAN_FONT = "\x1b[36m"
const RESET_TEXT = "\x1b[0m"
const INVISIBLE = "\x1b[?25l"
const VISIBLE = "\x1b[?25h"
const DELETE_LINE = "\x1b[K"
const BCKG_WHITE = "\x1b[48;5;231m"
const BLACK_TEXT = "\x1b[38;5;232m"
const BOLD_TEXT = "\x1b[1m"
const GRAY_TEXT = "\x1b[38;5;240m"
const UNDERLINE_TEXT = "\x1b[4m"
const GREEN_TEXT = "\x1b[38;5;42m"

type SearchBar string
type FileViewer string
type KeyBoardHelp string

func (sb SearchBar) Render(out io.Writer, keyword string){
	if len(keyword) == 0{
		fmt.Fprintf(out ," 🔎︎ %s%s%s", GRAY_TEXT , "Search entries...", RESET_TEXT)	
	} else {
		fmt.Fprintf(out ," 🔎︎ %s%s %s", keyword, BCKG_WHITE, RESET_TEXT)	
	}
}

func (fv FileViewer) Render(out io.Writer, de []fs.DirEntry, selectedIndex int){
	if selectedIndex < 0 {
		fmt.Fprintf(out, "   (empty)\n")
		return
	}
	if selectedIndex >= len(de) {
		return
	}
	const TRUNC = 10
	if len(de) <= TRUNC {
		for index, dirEntry := range de{
			fmt.Fprintf(out, "  ")
			if(index==selectedIndex){
				fmt.Fprintf(out, "%s> %s%s \n",CYAN_FONT, dirEntry.Name(), RESET_TEXT)
			} else {
				fmt.Fprintf(out, "  %s \n", dirEntry.Name())
			}
		}
	} else {
		if(selectedIndex >= 0 && selectedIndex < TRUNC){
			for index, dirEntry := range de[0:TRUNC]{
				fmt.Fprintf(out, "  ")
				if(index==selectedIndex){
					fmt.Fprintf(out, "%s> %s%s \n",CYAN_FONT, dirEntry.Name(), RESET_TEXT)
				} else {
					fmt.Fprintf(out, "  %s \n", dirEntry.Name())
				}
			}
			fmt.Fprintf(out, "    ...")
		} else {
			var isLast bool = selectedIndex == len(de)-1
			for index, dirEntry := range de[selectedIndex-TRUNC+1:selectedIndex+1]{
				fmt.Fprintf(out, "  ")
				if(index == TRUNC-1){
					fmt.Fprintf(out, "%s> %s%s \n",CYAN_FONT, dirEntry.Name(), RESET_TEXT)
				} else {
					fmt.Fprintf(out, "  %s \n", dirEntry.Name())
				}
			}
			if !isLast{
				fmt.Fprintf(out, "    ...")
			}
		}
	}
}

func (kbh KeyBoardHelp) Render(out io.Writer){
	fmt.Fprintf(out, "%s%s%s ← %s", BOLD_TEXT, BCKG_WHITE, BLACK_TEXT, RESET_TEXT)
	fmt.Fprintf(out, " Back",)
	fmt.Fprintf(out, "%s", CURSOR_RIGHT(6))
	fmt.Fprintf(out, "%s%s%s Enter %s", BOLD_TEXT, BCKG_WHITE, BLACK_TEXT, RESET_TEXT)
	fmt.Fprintf(out, " Change directory",)
	fmt.Fprintf(out, "%s", CURSOR_RIGHT(6))
	fmt.Fprintf(out, "%s%s%s ESC %s", BOLD_TEXT, BCKG_WHITE, BLACK_TEXT, RESET_TEXT)
	fmt.Fprintf(out, " Exit",)
}

func CURSOR_UP(n int) string{
	return "\x1b[" + strconv.Itoa(n) + "A"
}
func CURSOR_DOWN(n int) string{
	return "\x1b[" + strconv.Itoa(n) + "B"
}
func CURSOR_RIGHT(n int) string{
	return "\x1b[" + strconv.Itoa(n) + "C"
}
func CURSOR_LEFT(n int) string{
	return "\x1b[" + strconv.Itoa(n) + "D"
}