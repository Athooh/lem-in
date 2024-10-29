package main

type Room struct {
    Name string
    X, Y int
    Links []string // To store connected rooms
}

type Tunnel struct {
    Room1, Room2 string
}

type Colony struct {
    StartRoom, EndRoom string
    Rooms  map[string]*Room
    Tunnels []Tunnel
    Ants    int
}
