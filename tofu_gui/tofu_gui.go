package tofu_gui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type Gui struct {
    GuiToEngine chan<- string
    EngineToGui <-chan string
}

func (g *Gui) Start() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Channel Example")

    input := widget.NewEntry()
    input.SetPlaceHolder("Enter something...")
    button := widget.NewButton("Send to Engine", func() {
        g.GuiToEngine <- input.Text
    })

    resultLabel := widget.NewLabel("Waiting for result...")

    go func() {
        for result := range g.EngineToGui {
            fyne.Do(func() {
                resultLabel.SetText(result)
            })
        }
    }()

    content := container.NewVBox(input, button, resultLabel)
    myWindow.SetContent(content)
    myWindow.Resize(fyne.NewSize(400, 300))
    myWindow.ShowAndRun()
}
