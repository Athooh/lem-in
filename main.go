package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	inputFile := os.Args[1]

	colony, err := ParseInput(inputFile)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	// Display in the expected format
	fmt.Println(colony.Ants) // Number of ants
	fmt.Println("##start")
	fmt.Printf("%s %d %d\n", colony.StartRoom, colony.Rooms[colony.StartRoom].X, colony.Rooms[colony.StartRoom].Y)

	// Print rooms excluding start and end
	for name, room := range colony.Rooms {
		if name != colony.StartRoom && name != colony.EndRoom {
			fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
		}
	}

	fmt.Println("##end")
	fmt.Printf("%s %d %d\n", colony.EndRoom, colony.Rooms[colony.EndRoom].X, colony.Rooms[colony.EndRoom].Y)

	// Print tunnels
	for _, tunnel := range colony.Tunnels {
		fmt.Printf("%s-%s\n", tunnel.Room1, tunnel.Room2)
	}

	fmt.Println() // Line break before ant movements

	// Find shortest path and simulate ant movements
	path, err := colony.FindShortestPath()
	if err != nil {
		fmt.Printf("Error finding path: %v\n", err)
		return
	}

	colony.SimulateAntMovements(path)
}
