package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func ParseInput(filename string) (*Colony, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("could not open file: %w", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    colony := &Colony{Rooms: make(map[string]*Room)}

    // Step 1: Parse number of ants
    if scanner.Scan() {
        ants, err := strconv.Atoi(scanner.Text())
        if err != nil || ants <= 0 {
            return nil, fmt.Errorf("invalid number of ants")
        }
        colony.Ants = ants
    }

    // Step 2: Parse rooms and tunnels
    for scanner.Scan() {
        line := scanner.Text()
        if line == "##start" || line == "##end" {
            if scanner.Scan() {
                room, err := parseRoom(scanner.Text())
                if err != nil {
                    return nil, fmt.Errorf("error parsing room: %w", err)
                }
                colony.Rooms[room.Name] = room
                if line == "##start" {
                    colony.StartRoom = room.Name
                } else {
                    colony.EndRoom = room.Name
                }
            }
        } else if strings.Contains(line, "-") {
            tunnel := parseTunnel(line)
            colony.Tunnels = append(colony.Tunnels, tunnel)

            // Add links to the adjacency list for each room
            colony.Rooms[tunnel.Room1].Links = append(colony.Rooms[tunnel.Room1].Links, tunnel.Room2)
            colony.Rooms[tunnel.Room2].Links = append(colony.Rooms[tunnel.Room2].Links, tunnel.Room1)
        } else if !strings.HasPrefix(line, "#") {
            room, err := parseRoom(line)
            if err != nil {
                return nil, fmt.Errorf("error parsing room: %w", err)
            }
            colony.Rooms[room.Name] = room
        }
    }

    return colony, nil
}

func parseRoom(line string) (*Room, error) {
    parts := strings.Fields(line)
    if len(parts) != 3 {
        return nil, fmt.Errorf("invalid room format")
    }
    x, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil, err
    }
    y, err := strconv.Atoi(parts[2])
    if err != nil {
        return nil, err
    }
    return &Room{Name: parts[0], X: x, Y: y}, nil
}

func parseTunnel(line string) Tunnel {
    parts := strings.Split(line, "-")
    return Tunnel{Room1: parts[0], Room2: parts[1]}
}
