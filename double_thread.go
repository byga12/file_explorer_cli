package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"

	Fe "github.com/byga12/file_explorer_cli/file_explorer"
	Widgets "github.com/byga12/file_explorer_cli/widgets"
	Kb "github.com/eiannone/keyboard"
)

type Data struct {
	keyword string
	fileExplorer *Fe.FileExplorer
	selectedIndex int
	filteredList []fs.DirEntry		
}

type Key struct {
	key Kb.Key
	char rune
}

func main(){
	defer logFile.Close()
	Logger.Println("test")
	// Setup channels
	//// From main To keyHandler
	keyChannel := make(chan Key)
	dataChannel := make(chan *Data)
	//// From keyHandler To main
	renderChannel := make(chan bool)
	// Setup Data
	fe, err := Fe.NewFileExplorer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	data := &Data{
		keyword: "",
		fileExplorer: &fe,
		selectedIndex: 0,
		filteredList: fe.GetDirectoryEntries(),
	}

	// Setup widgets
	out := os.Stdout
	var searchBar Widgets.SearchBar
	var entriesViewer Widgets.FileViewer
	searchBar.Render(out, "")
	entriesViewer.Render(out, fe.GetDirectoryEntries(), 0)
	fmt.Fprintf(out,Widgets.INVISIBLE)
	
	// Render initial widgets
	render(searchBar, entriesViewer, out, data)
	
	// Init thread
	go initThreadKeyHandler(keyChannel, dataChannel, renderChannel)
	
	// Setup keyboard listener
	if err := Kb.Open(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer Kb.Close()
	
	
	LOOP:
	for {
		char, key, err := Kb.GetKey()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		
		if(key == Kb.KeyEsc){
			break LOOP
			} 
			keyChannel <- Key{key: key, char: char}
			dataChannel <- data
			shouldRender := <- renderChannel
			if shouldRender {
				render(searchBar, entriesViewer, out, data)
				os.Chdir(data.fileExplorer.CurrentPath)
			}
		}
		
	close(keyChannel)
	close(dataChannel)
	close(renderChannel)
	fmt.Fprintf(out,"\n%s %s \x1b[H\x1b[2J\x1b[0m", Widgets.VISIBLE, Widgets.RESET_TEXT)
	fmt.Fprintf(out, "%s", data.fileExplorer.CurrentPath)
	defer out.Close()
}

func initThreadKeyHandler(keyChannel chan Key, dataChannel chan *Data, renderChannel chan bool) {
	for {
		var keyStruct Key = <- keyChannel
		var data *Data = <- dataChannel
		if(data==nil){
			continue
		}
		switch {
		case keyStruct.key == Kb.KeyEsc:
			return
		case keyStruct.key == Kb.KeyBackspace2:
			if(len(data.keyword)>0){
				data.keyword = data.keyword[:len(data.keyword)-1]
				data.filteredList = data.fileExplorer.SearchInPath(data.keyword)
			}
			renderChannel <- true
		case keyStruct.key == Kb.KeySpace:
			data.keyword += " "
			data.filteredList = data.fileExplorer.SearchInPath(data.keyword)
			renderChannel <- true
		case keyStruct.key == Kb.KeyArrowUp:
			if(data.selectedIndex > 0){
				data.selectedIndex--
			}
			renderChannel <- true

		case keyStruct.key == Kb.KeyArrowDown:
			if(data.selectedIndex <= len(data.filteredList)-2){
				data.selectedIndex++
			}
			renderChannel <- true
		case keyStruct.key == Kb.KeyEnter:
			if(data.filteredList[data.selectedIndex].IsDir()){
				data.fileExplorer.ChangeDirectory(data.filteredList[data.selectedIndex].Name())
				data.keyword = ""
				newEntries := data.fileExplorer.GetDirectoryEntries()
				if len(newEntries)==0{
					data.selectedIndex = -1
				} else {
					data.selectedIndex = 0
				}
				data.filteredList = newEntries
				renderChannel <- true
			} else {
				renderChannel <- false
			}
		case keyStruct.key == Kb.KeyArrowLeft:
			renderChannel <- false
		case keyStruct.key == Kb.KeyArrowRight:
			renderChannel <- false
		default:
			data.keyword += string(keyStruct.char)
			data.filteredList = data.fileExplorer.SearchInPath(data.keyword)
			renderChannel <- true
		}
	}
}

func render(searchBar Widgets.SearchBar, entriesViewer Widgets.FileViewer, out io.Writer, data *Data){
	fmt.Fprintf(out,"\x1b[H\x1b[2J")
	
	searchBar.Render(out, data.keyword)
	fmt.Fprintf(out, "\n")
	entriesViewer.Render(out, data.filteredList, data.selectedIndex)
}