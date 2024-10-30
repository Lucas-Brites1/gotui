
# GoTUI

**GoTUI** is a Go library designed for creating interactive terminal menus and CLI options. This library offers a simple way to build structured menus with intuitive keyboard navigation, ideal for terminal-based applications.

## Features

- **Interactive Menus**: Create navigable terminal menus with support for arrow keys.
- **Nested Options**: Organize options into submenus for a hierarchical structure.
- **Cross-Platform**: Works on Windows, macOS, and Linux.
- **Simple API**: Minimal setup and easy-to-use functions for quick menu creation.

## Installation

To install GoTUI, use `go get`:

```bash
go get github.com/Lucas-Brites1/gotui
```

## Usage

Here’s a basic example to help you get started:

```go
package main

import (
	"fmt"
	"github.com/Lucas-Brites1/gotui"
)

func main() {
	t := gotui.Terminal{}

	subOptions := &[]gotui.Option{
		gotui.CreateOption("Option 1", func() { fmt.Println("Action for Option 1") }, nil),
		gotui.CreateOption("Option 2", func() { fmt.Println("Action for Option 2") }, nil),
	}

	options := []gotui.Option{
		gotui.CreateOption("Main Menu", nil, subOptions),
	}

	t.Start(options).HandleKeys()
}
```

## Documentation

- **Option**: Represents each menu item, with fields for `Description`, `Action`, and `Children` (for submenus).
- **Terminal**: Manages the menu state, rendering, and keyboard interactions.
- **Methods**:
  - `CreateOption(description string, action func(), children *[]Option) Option`: Creates a new menu option.
  - `Start(options []Option)`: Initializes the menu.
  - `HandleKeys()`: Handles keyboard input for menu navigation.

## Key Bindings

- **↑ / ↓**: Navigate between menu options
- **Enter**: Select an option
- **Backspace**: Return to the previous menu
- **Esc**: Exit the program

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
