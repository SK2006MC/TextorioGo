# Textorio

Textorio is a Factorio-inspired idle management game that runs in your terminal. It is written in Go and uses the `tview` library for its user interface.

## Status

This project is currently in progress. Some features are not yet implemented, and the game is not yet fully playable.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) (version 1.18 or later)

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/your-username/textorio.git
    cd textorio
    ```

2.  Install the dependencies:
    ```sh
    go mod tidy
    ```

### Running the Game

To run the game, use the following command:

```sh
go run main.go
```

## How to Play

The game is played by entering commands into the input field at the bottom of the screen. The game world is currently very simple, but the following commands are available:

-   `help`: Displays a list of available commands.
-   `inv`: Displays the player's inventory.
-   `craft <item_name>`: Crafts the specified item.
-   `quit`: Exits the game.

## Project Structure

-   `main.go`: The main entry point of the application.
-   `config/`: Contains configuration files for the game.
-   `internal/core/`: Contains the core game logic, such as game state, player, items, and recipes.
-   `internal/ui/`: Contains the user interface logic.
-   `ui/tview/`: The `tview`-based implementation of the UI.
-   `go.mod`, `go.sum`: Go module files.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.
