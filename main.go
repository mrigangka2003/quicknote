package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool	
}


var (
	vaultDir string
)

func Init() {
	homeDir,err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v", err)
		os.Exit(1)
	}

	vaultDir = fmt.Sprintf("%s/.quicknote", homeDir)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd


	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil
		case "enter":
			return m,nil
		}
	}

	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
		return m, cmd
	}

	return m, cmd
}

//view method 
func (m model) View() string {

	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("243")).
		Background(lipgloss.Color("234")).Padding(1, 2)

	welcome := style.Render("Welcome to QuickNote !")
	instructions := "CTRL+N:new file  CTRL+L:list  ESC:back/save CTRL+S:save CTRL+C/q:quit"


	view := ""
	if m.createFileInputVisible {
		view += "Create a new file:\n\n"
		view += m.newFileInput.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s\n", welcome, view, instructions)
}


func initializeMode() model {

	err := os.MkdirAll(vaultDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating vault directory: %v\n", err)
		os.Exit(1)
	}

	//initialize the text input model
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("222"))
	ti.PromptStyle= lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	return model{
		newFileInput: ti,
		createFileInputVisible: false,
	}
}
func main() {
	fmt.Println("welcome to quicknote")
	p:= tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}