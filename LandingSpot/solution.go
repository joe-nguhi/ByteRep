/*
Space Week Day 4: Landing Spot
Problem Statement: Given a matrix of integers, find the landing spot with the lowest sum of neighbors.
*/
package main

import "fmt"

type Point struct {
	x int
	y int
}

const landingSpot = 0

var directions = []Point{
	{0, -1}, // left
	{0, 1},  // right
	{1, 0},  // down
	{-1, 0}, //  up
}

func main() {
	fmt.Println(findLandingSpot([][]int{
		{1, 0},
		{2, 0},
	}))
}

func findLandingSpot(matrix [][]int) [2]int {

	landingSpots := make(map[Point]int)

	for row := 0; row < len(matrix); row++ {
		for y := 0; y < len(matrix[row]); y++ {
			node := matrix[row][y]
			if node != landingSpot {
				continue
			}

			sn := sumOfNeighbors(Point{x: row, y: y}, matrix)
			landingSpots[Point{x: row, y: y}] = sn
		}
	}

	// Extract keys to a slice
	keys := make([]Point, 0, len(landingSpots))

	for k := range landingSpots {
		keys = append(keys, k)
	}

	var selected Point

	for i := 0; i < len(keys); i++ {
		key := keys[i]

		if i == 0 {
			selected = key
			continue
		}

		value := landingSpots[key]

		if value < landingSpots[selected] {
			selected = key
		}
	}

	return [2]int{selected.x, selected.y}
}

func getNeighbors(node Point, matrix [][]int) []Point {
	ns := make([]Point, 0, len(directions))
	for _, direction := range directions {
		neighbor := Point{
			x: node.x + direction.x,
			y: node.y + direction.y,
		}
		if checkInBounds(neighbor, matrix) {
			ns = append(ns, neighbor)
		}
	}

	return ns
}

func checkInBounds(node Point, matrix [][]int) bool {
	return node.x >= 0 && node.x < len(matrix[0]) && node.y >= 0 && node.y < len(matrix)
}

func getNodeValue(node Point, matrix [][]int) int {
	return matrix[node.x][node.y]
}

func sumOfNeighbors(node Point, matrix [][]int) int {
	sum := 0
	for _, neighbor := range getNeighbors(node, matrix) {
		sum += getNodeValue(neighbor, matrix)
	}

	return sum
}
