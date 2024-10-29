package main

import (
    "fmt"
    "strings"
)

func (c *Colony) SimulateAntMovements(path []string) {
    // Each ant starts at the `##start` room.
    antPositions := make([]string, c.Ants)
    for i := range antPositions {
        antPositions[i] = c.StartRoom
    }

    for !allAntsAtEnd(antPositions, c.EndRoom) {
        movements := []string{}
        for ant := range antPositions {
            if antPositions[ant] == c.EndRoom {
                continue // Ant has reached the end, no need to move further
            }

            // Find the next room in the path for this ant
            currentPositionIndex := indexOf(path, antPositions[ant])
            if currentPositionIndex+1 < len(path) {
                nextRoom := path[currentPositionIndex+1]

                // Check if the next room is available
                if !roomOccupied(antPositions, nextRoom, c.StartRoom, c.EndRoom) {
                    antPositions[ant] = nextRoom
                    movements = append(movements, fmt.Sprintf("L%d-%s", ant+1, nextRoom))
                }
            }
        }

        // Print movements in the correct format, without brackets or commas
        if len(movements) > 0 {
            fmt.Println(strings.Join(movements, " "))
        }
    }
}


// Helper function to check if all ants are at the end room
func allAntsAtEnd(positions []string, endRoom string) bool {
    for _, pos := range positions {
        if pos != endRoom {
            return false
        }
    }
    return true
}

// Helper function to find the index of a room in the path
func indexOf(path []string, room string) int {
    for i, r := range path {
        if r == room {
            return i
        }
    }
    return -1
}

// Check if a room is occupied by any ant (except at start and end)
func roomOccupied(positions []string, room, startRoom, endRoom string) bool {
    for _, pos := range positions {
        if pos == room && room != startRoom && room != endRoom {
            return true
        }
    }
    return false
}
