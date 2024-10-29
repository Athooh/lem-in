package main

import (
	"container/list"
	"fmt"
)

func (c *Colony) FindShortestPath() ([]string, error) {
	queue := list.New()
	queue.PushBack([]string{c.StartRoom})

	visited := make(map[string]bool)
	visited[c.StartRoom] = true

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]

		if lastRoom == c.EndRoom {
			return path, nil
		}

		for _, neighbor := range c.Rooms[lastRoom].Links {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return nil, fmt.Errorf("no path from start to end")
}
