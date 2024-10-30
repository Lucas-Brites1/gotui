package terminal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/eiannone/keyboard"
)

type Option struct {
	Description string
	Action      func()
	Children    *[]Option
}

func CreateOption(description string, action func(), children *[]Option) Option {
	return Option{
		Description: description,
		Action:      action,
		Children:    children,
	}
}

func (o *Option) Run() {
	if o.Action != nil {
		o.Action()
		fmt.Println("\nPress any key to return to the menu...")
		_, _, _ = keyboard.GetKey()
	}
}

type Terminal struct {
	RenderOptions []Option
	selected      int
	OS            string
}

func (t *Terminal) setOption(options ...Option) {
	t.RenderOptions = append(t.RenderOptions, options...)
}

func (t *Terminal) Render(options []Option) {
	t.Clear()
	fmt.Println("Use ↑ and ↓ to navigate, Enter to select, Backspace to go back.")
	for i, option := range options {
		if i == t.selected {
			fmt.Printf("\033[1;7m> %s\033[0m\n", option.Description)
		} else {
			fmt.Printf(" %s\n", option.Description)
		}
	}
}

func (t *Terminal) Clear() {
	switch t.OS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (t *Terminal) Start(options []Option) *Terminal {
	t.OS = runtime.GOOS
	t.setOption(options...)
	t.selected = 0
	return t
}

func (t *Terminal) HandleKeys() {
	if err := keyboard.Open(); err != nil {
		log.Fatal("Error opening keyboard listener: ", err)
	}
	defer keyboard.Close()

	var currentOptions *[]Option = &t.RenderOptions

	for {
		t.Render(*currentOptions)
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Println("Error getting key:", err)
			continue
		}

		switch key {
		case keyboard.KeyArrowUp:
			if t.selected > 0 {
				t.selected--
			}
		case keyboard.KeyArrowDown:
			if t.selected < len(*currentOptions)-1 {
				t.selected++
			}
		case keyboard.KeyEnter:
			selectedOption := (*currentOptions)[t.selected]
			if selectedOption.Children != nil {
				currentOptions = selectedOption.Children
				t.selected = 0
			} else {
				selectedOption.Run()
			}
		case keyboard.KeyBackspace2:
			currentOptions = &t.RenderOptions
			t.selected = 0
		case keyboard.KeyEsc:
			fmt.Println("Exiting the program...")
			return
		}
	}
}
