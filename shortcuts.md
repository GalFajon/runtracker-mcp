# VS Code keyboard shortcuts for Linux

Start with `Ctrl+P` to find a file, `F12` to follow a definition, `Ctrl+Alt+-` to return, and `Ctrl+1` to focus the first editor group while keeping the terminal visible.

## Open a project and find code

| Shortcut | Action |
| --- | --- |
| `Ctrl+K`, then `Ctrl+O` | Open a project folder. From a terminal in the folder, run `code .`. |
| `Ctrl+P` | Find a file by name. Enter `main.go:42` to jump to line 42. |
| `Ctrl+Tab` | Switch to the last-used editor. Hold Ctrl to choose another recent editor. |
| `Ctrl+Shift+O` | Find a function or symbol in the current file. |
| `Ctrl+T` | Find a symbol across the project. |
| `Ctrl+Shift+F` | Search the project for text, such as an error message or config key. |

## Follow and edit code

| Shortcut | Action |
| --- | --- |
| `F12` | Go to the definition under the cursor. |
| `Ctrl+Alt+-` | Return to the previous location. |
| `Shift+F12` | Show references to the current symbol. |
| `F2` | Rename a symbol and update its references. |
| `Ctrl+.` | Open quick fixes and code actions. |
| `Alt+↑` / `Alt+↓` | Move the current line or selected block. |
| `Ctrl+D` | Select the current word. Repeat to add the next match, then edit them together. |
| `F8` / `Shift+F8` | Jump to the next or previous error or warning. |

## Switch focus and manage files

| Shortcut | Action |
| --- | --- |
| Ctrl + backtick | Toggle the integrated terminal. |
| `Ctrl+1` | Focus the first editor group and keep the terminal visible. |
| `Ctrl+Shift+E` | Focus Explorer. Use arrows to navigate, Enter to open, and F2 to rename. |
| `Ctrl+N`, then `Ctrl+Shift+S` | Create a file and choose its save path. From the integrated terminal, use `code -r main.go` to open a new or existing file in the current window. |
| `Ctrl+Shift+P` | Open the Command Palette and search by action name. On my setup, `F4` also opens it. |
| `Ctrl+K`, then `Ctrl+S` | Open the keyboard shortcut editor. Press the two combinations in sequence. |

## Practice in a Go project

1. Press `Ctrl+P`, type `main.go`, and press Enter.
2. Press `Ctrl+Shift+O` and choose a function.
3. Move the cursor onto a function it calls. Press `F12` to follow it, then `Ctrl+Alt+-` to come back.
4. Open the terminal with Ctrl + backtick and run `go test ./...`.
5. Press `Ctrl+1` to keep editing with the test output still visible.

[Official VS Code Linux keyboard shortcut sheet](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-linux.pdf)
