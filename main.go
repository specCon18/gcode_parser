package main

import (
	"regexp"
	"strings"
	"bufio"
	"os"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func readFileLBL(file_path string) {
    file, err := os.Open(file_path)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err.Error() == "EOF" {
                // If it's EOF, we can just break the loop
                break
            }
            fmt.Println("Error reading file:", err)
            return
        }
        parseGcodeLine(line)  // Output the line
    }
}

func removeLineNumber(input string) string {
	// Define the regex pattern for the prefix 'N' followed by one or more digits and an optional space
	re := regexp.MustCompile(`^N\d+\s?`)
	// Replace the matched prefix with an empty string
	return re.ReplaceAllString(input, "")
}

func parseGcodeLine(line string){
	//TODO: P:0 convert string to struct
	if strings.HasPrefix(line,"N"){
		fmt.Println("--------------------")
		fmt.Println("Line has line number")
		fmt.Println("--------------------")
		//l := removeLineNumber(line)
		//parseGcodeLine(l)
	} else if strings.HasPrefix(line, "G") {
		fmt.Println("is a gcode")
	} else if strings.HasPrefix(line,"M"){
		fmt.Println("is a mcode")
	} else if strings.HasPrefix(line,";"){
		fmt.Println("is a comment")
	}
}

func main() {
	//Create a new app instance
	a := app.New()
	//Create a new Window on the app instance
	w := a.NewWindow("GCode Renderer")

	//Create the menu bar items for the render window
	renderMenuItem := fyne.NewMenuItem("Load", func() { readFileLBL("./test.gcode") })

	//Create the render menu bar and add its items
	renderMenuBar := fyne.NewMainMenu(fyne.NewMenu("File",renderMenuItem))
	w.SetMainMenu(renderMenuBar)
	
	//instantiate gc_render as a Label widget
	gc_render := widget.NewLabel("GCODE RENDER HERE")
	
	//update the window with the gc_render widget
	w.SetContent(gc_render)

	//Resize w1 to 480x640
	w.Resize(fyne.NewSize(640, 480))
	
	//Set the first window visable
	w.Show()
	
	//Create a new Window on the app instance to show the gcode being rendered in w
	w1 := a.NewWindow("GCode Visualizer")

	//instantiate gc_vis as a Label widget
	gc_vis := widget.NewLabel("GCODE TO BE RENDERED HERE")

	//update the window with the gc_vis widget
	w1.SetContent(gc_vis)
	
	//Resize w1 to 480x640
	w1.Resize(fyne.NewSize(480, 640))

	//Set the second window visable
	w1.Show()

	//Run the above defined app instance
	a.Run()
}
