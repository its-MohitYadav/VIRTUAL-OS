package main

import (
	
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGallery(w fyne.Window){
	
	root_src:="C:\\Users\\Mohit\\Downloads\\wallpapers"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
        log.Fatal(err)
    }
	var picsArr []string;
	for _, file := range files {
        
		if !file.IsDir(){
			extension:=strings.Split(file.Name(),".")[1];
			if extension=="png"||extension=="jpeg"||extension=="jpg"{
				picsArr=append(picsArr,root_src+"\\"+file.Name());
			}
		}
    }
	
	tabs := container.NewAppTabs(
		container.NewTabItem("img", canvas.NewImageFromFile(picsArr[0])),
	)
	for i:=1;i<len(picsArr);i++{
		image:=canvas.NewImageFromFile(picsArr[i])
		tabs.Append(container.NewTabItem("img",image));
	}
	
	tabs.SetTabLocation(container.TabLocationBottom)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,tabs),)
	w.Show()
}
