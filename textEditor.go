package main

import (

	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int=1
func showEditor(){
	
	w:=myApp.NewWindow("Text Editor")

	w.Resize(fyne.NewSize(600,400));

	content:= container.NewVBox(
		container.NewVBox(
			widget.NewLabel("Text Editor"),
		),
	)
	content.Add(widget.NewButton("Add new File",func() {
		content.Add(widget.NewLabel("New File "+ strconv.Itoa(count)))
		count++;
	}))

	input:=widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text....")
	input.Resize(fyne.NewSize(400,400))

	savebtn:=widget.NewButton("Save file" , func() {
		saveFileDialog:=dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData:=[]byte(input.Text)
				uc.Write(textData)
			},w)
		saveFileDialog.SetFileName("New File "+ strconv.Itoa(count-1)+".txt")
		saveFileDialog.Show()
	})

	openbtn:=widget.NewButton("Open file",func() {
		openFileDialog:=dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData ,_:=ioutil.ReadAll(r)

				output:=fyne.NewStaticResource("New File",ReadData);

				viewData:= widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w:=fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(viewData))
				w.Resize(fyne.NewSize(400,400))
				w.Show()
			},w)

			openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
			
			openFileDialog.Show()
	})

	calciContainer:=container.NewVBox(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				savebtn,
				openbtn,
			),
		),
	)

	w.SetContent(
		container.NewBorder(deskBtn,nil,nil,nil,calciContainer),
	)
	w.Show()
}