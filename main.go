package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/widget"
	"log"
)

func main() {

	a := app.New()
	w := a.NewWindow("Longman Translator Deck creator for Anki")
	w.Resize(fyne.Size{Height: 700,Width: 900})
	w.SetContent(widget.NewButton("Vocabulary words(.csv)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader == nil {
				return
			}
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			fileOpened(reader)
		}, w)
		fd.Resize(fyne.Size{Height: 600,Width: 800})
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		fd.Show()
	}))

	w.ShowAndRun()
	//config, err := app.NewConfig()
	//if err != nil {
	//	panic(err)
	//}
	//application, err := app.New(config)
	//if err != nil {
	//	panic(err)
	//}
	//defer application.Close()
	//application.Run()
}

func fileOpened(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}

	ext := f.URI().Extension()
	if ext == ".png" {
	} else if ext == ".txt" {
	}
	err := f.Close()
	if err != nil {
		fyne.LogError("Failed to close stream", err)
	}
}