package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// GhostMouse represents a ghost mouse character
type GhostMouse struct {
	Name string
	X    int
	Y    int
	Icon string
}

// Maze represents the game board
type Maze struct {
	Width  int
	Height int
	Cheese []Position
}

// Position represents a coordinate in the maze
type Position struct {
	X int
	Y int
}

const (
	ghostIcon  = "👻"
	cheeseIcon = "🧀"
	emptyIcon  = "⬜"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║    👻 GHOST MICE GO - Fun Coding App 👻      ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Welcome to Ghost Mice Go!")
	fmt.Println("Help the ghost mouse find all the cheese! 🧀")
	fmt.Println()

	// Initialize the ghost mouse
	mouse := GhostMouse{
		Name: "Boo",
		X:    0,
		Y:    0,
		Icon: ghostIcon,
	}

	// Initialize the maze
	maze := Maze{
		Width:  8,
		Height: 5,
		Cheese: []Position{
			{X: 3, Y: 2},
			{X: 6, Y: 1},
			{X: 7, Y: 4},
		},
	}

	fmt.Printf("Ghost Mouse %s is ready to hunt for cheese!\n\n", mouse.Name)

	// Game loop
	reader := bufio.NewReader(os.Stdin)
	cheeseCollected := 0

	for {
		// Draw the maze
		drawMaze(maze, mouse)
		
		if cheeseCollected == len(maze.Cheese) {
			fmt.Println("\n🎉 Congratulations! You collected all the cheese! 🎉")
			fmt.Printf("Ghost Mouse %s is very happy! 👻✨\n", mouse.Name)
			break
		}

		// Display instructions
		fmt.Println("\n📝 Commands: up, down, left, right, quit")
		fmt.Printf("Cheese collected: %d/%d\n", cheeseCollected, len(maze.Cheese))
		fmt.Print("Enter command: ")

		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		command := strings.TrimSpace(strings.ToLower(input))

		// Process command
		switch command {
		case "up":
			if mouse.Y > 0 {
				mouse.Y--
				animateMove("up")
			} else {
				fmt.Println("💥 Ouch! Can't go through the wall!")
			}
		case "down":
			if mouse.Y < maze.Height-1 {
				mouse.Y++
				animateMove("down")
			} else {
				fmt.Println("💥 Ouch! Can't go through the wall!")
			}
		case "left":
			if mouse.X > 0 {
				mouse.X--
				animateMove("left")
			} else {
				fmt.Println("💥 Ouch! Can't go through the wall!")
			}
		case "right":
			if mouse.X < maze.Width-1 {
				mouse.X++
				animateMove("right")
			} else {
				fmt.Println("💥 Ouch! Can't go through the wall!")
			}
		case "quit", "exit", "q":
			fmt.Println("\n👋 Thanks for playing Ghost Mice Go!")
			return
		default:
			fmt.Println("❓ Unknown command. Try: up, down, left, right, or quit")
			continue
		}

		// Check if mouse found cheese
		for i, cheese := range maze.Cheese {
			if mouse.X == cheese.X && mouse.Y == cheese.Y {
				fmt.Println("🎊 Yum! You found cheese! 🧀")
				maze.Cheese = append(maze.Cheese[:i], maze.Cheese[i+1:]...)
				cheeseCollected++
				break
			}
		}

		fmt.Println()
	}
}

// drawMaze renders the maze with the ghost mouse and cheese
func drawMaze(maze Maze, mouse GhostMouse) {
	fmt.Println("\n╔" + strings.Repeat("══", maze.Width) + "╗")
	
	for y := 0; y < maze.Height; y++ {
		fmt.Print("║")
		for x := 0; x < maze.Width; x++ {
			if x == mouse.X && y == mouse.Y {
				fmt.Print(ghostIcon + " ")
			} else if hasCheese(maze.Cheese, x, y) {
				fmt.Print(cheeseIcon + " ")
			} else {
				fmt.Print(emptyIcon + " ")
			}
		}
		fmt.Println("║")
	}
	
	fmt.Println("╚" + strings.Repeat("══", maze.Width) + "╝")
}

// hasCheese checks if there's cheese at the given position
func hasCheese(cheese []Position, x, y int) bool {
	for _, c := range cheese {
		if c.X == x && c.Y == y {
			return true
		}
	}
	return false
}

// animateMove adds a simple animation effect
func animateMove(direction string) {
	symbols := map[string]string{
		"up":    "↑",
		"down":  "↓",
		"left":  "←",
		"right": "→",
	}
	
	if symbol, ok := symbols[direction]; ok {
		fmt.Printf("Moving %s %s...\n", direction, symbol)
		time.Sleep(200 * time.Millisecond)
	}
}
