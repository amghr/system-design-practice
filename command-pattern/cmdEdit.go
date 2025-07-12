package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Command interface contains Do and Undo methods
type Command interface {
	Do()
	Undo()
}

// UndoRedoManager tracks undo and redo using stacks
type UndoRedoManager struct {
	undoStack []Command
	redoStack []Command
}

// ExecutesCommand executes the the command
func (m *UndoRedoManager) ExecuteCommand(cmd Command) {
	cmd.Do()
	m.undoStack = append(m.undoStack, cmd)
	m.redoStack = nil
}

// Undo pops the last command and undoes it, pushes it onto redo stack
func (m *UndoRedoManager) Undo() {
	if len(m.undoStack) == 0 {
		fmt.Println("Nothing to undo")
		return
	}
	last := m.undoStack[len(m.undoStack)-1]
	m.undoStack = m.undoStack[:len(m.undoStack)-1]
	last.Undo()
	m.redoStack = append(m.redoStack, last)
}

// Redo re-executes the last undone command and pushed it onto undo stack
func (m *UndoRedoManager) Redo() {
	if len(m.redoStack) == 0 {
		fmt.Println("Nothing to redo")
		return
	}
	last := m.redoStack[len(m.redoStack)-1]
	m.redoStack = m.redoStack[:len(m.redoStack)-1]
	last.Do()
	m.undoStack = append(m.undoStack, last)
}

// Editor struct hold Text, FontSize and Color
type Editor struct {
	Text     string
	FontSize int
	Color    string
}

// Display prints the text and its color
func (e *Editor) Display() {
	colorCode := map[string]string{
		"red":   "\033[31m",
		"blue":  "\033[34m",
		"reset": "\033[0m",
	}

	fmt.Println("---- Output ----")
	fmt.Printf("Font Size: %d\n", e.FontSize)
	fmt.Printf("%s%s%s\n", colorCode[e.Color], e.Text, colorCode["reset"])
	fmt.Println("----------------")
}

// TypeCommand changes the text and stores previous text
type TypeCommand struct {
	editor   *Editor
	newText  string
	prevText string
}

// Do sets the editor to new test and prints it
func (c *TypeCommand) Do() {
	c.prevText = c.editor.Text
	c.editor.Text = c.newText
	c.editor.Display()
}

// Undo resets to previous text and displays it
func (c *TypeCommand) Undo() {
	c.editor.Text = c.prevText
	c.editor.Display()
}

// FontSizeCommand changes the fontsize and saves old fontsize
type FontSizeCommand struct {
	editor       *Editor
	newSize      int
	previousSize int
}

// Do sets the text to new fontsize
func (c *FontSizeCommand) Do() {
	c.previousSize = c.editor.FontSize
	c.editor.FontSize = c.newSize
	c.editor.Display()
}

// Undo restores the text to previous fontsize
func (c *FontSizeCommand) Undo() {
	c.editor.FontSize = c.previousSize
	c.editor.Display()
}

// ColorCommand changes the color and stores previous color
type ColorCommand struct {
	editor        *Editor
	newColor      string
	previousColor string
}

// Do sets the color to new color
func (c *ColorCommand) Do() {
	c.previousColor = c.editor.Color
	c.editor.Color = c.newColor
	c.editor.Display()
}

// Undo restores the previous color
func (c *ColorCommand) Undo() {
	c.editor.Color = c.previousColor
	c.editor.Display()
}

// main reads the command from stdin and handles them
func main() {
	editor := &Editor{Text: "", FontSize: 12, Color: "reset"}
	manager := &UndoRedoManager{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Terminal Text Editor (type 'help' for commands)")
	editor.Display()

	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		switch {
		case line == "undo":
			manager.Undo()
		case line == "redo":
			manager.Redo()
		case strings.HasPrefix(line, "type "):
			text := strings.TrimPrefix(line, "type ")
			cmd := &TypeCommand{editor: editor, newText: text}
			manager.ExecuteCommand(cmd)
		case strings.HasPrefix(line, "fontsize "):
			val := strings.TrimPrefix(line, "fontsize ")
			if size, err := strconv.Atoi(val); err == nil {
				cmd := &FontSizeCommand{editor: editor, newSize: size}
				manager.ExecuteCommand(cmd)
			} else {
				fmt.Println("Invalid font size")
			}
		case strings.HasPrefix(line, "color "):
			val := strings.TrimPrefix(line, "color ")
			if val == "red" || val == "blue" || val == "reset" {
				cmd := &ColorCommand{editor: editor, newColor: val}
				manager.ExecuteCommand(cmd)
			} else {
				fmt.Println("Invalid color")
			}
		case line == "help":
			fmt.Println("Commands:")
			fmt.Println("  type <text>       - Set text")
			fmt.Println("  fontsize <number> - Change font size")
			fmt.Println("  color <name>      - Change text color (red, blue, reset)")
			fmt.Println("  undo              - Undo last command")
			fmt.Println("  redo              - Redo last undone command")
		case line == "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
