package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var copyClipboard []string
	copyClipboard, _ = readLinesFromFile("./db_dummy.txt")

	myApp := app.New()
	myWindow := myApp.NewWindow("Clipboard")

	label := widget.NewLabel("Clipboard text")
	label2 := widget.NewLabel("Clipboard images")
	separator := widget.NewSeparator()

	input := widget.NewEntry()
	input.SetPlaceHolder("Look for your copy")

	results := container.NewVBox()

	for _, e := range copyClipboard {
		results.Add(widget.NewLabel(e))
	}

	scroll := container.NewVScroll(results)

	content := container.NewBorder(
		container.NewVBox(
			label,
			input,
			separator,
		),
		nil,
		nil,
		nil,
		scroll,
	)

	input.OnChanged = func(value string) {
		results.RemoveAll()

		for _, e := range copyClipboard {
			if strings.Contains(
				strings.ToLower(e),
				strings.ToLower(value),
			) {
				results.Add(widget.NewLabel(e))
			}
		}

		results.Refresh()
	}

	content2 := container.NewVBox()
	content2.Add(label2)
	content2.Add(separator)

	grid := container.New(layout.NewGridLayout(2), content, content2)

	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(700, 800))
	myWindow.ShowAndRun()
}

func readLinesFromFile(path string) ([]string, error) {
	fileHandle, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	var textClipboard []string
	scanner := bufio.NewReader(fileHandle)

	for {
		textLine, err := scanner.ReadString('\n')

		if err == io.EOF {
			if len(textLine) != 0 {
				textClipboard = append(textClipboard, textLine)
			}
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading from file: %w", err)
		}

		textClipboard = append(textClipboard, textLine)
	}

	return textClipboard, nil
}
