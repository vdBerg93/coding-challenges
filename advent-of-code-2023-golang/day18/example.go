package main

import (
	"fmt"
)

func calculateSurfaceArea(path [][]int) float64 {
	// Validate path
	if len(path) < 3 {
		fmt.Println("Invalid path. A polygon must have at least 3 vertices.")
		return 0
	}

	// Find bounding box
	minX, minY, maxX, maxY := path[0][0], path[0][1], path[0][0], path[0][1]
	for _, point := range path {
		if point[0] < minX {
			minX = point[0]
		}
		if point[1] < minY {
			minY = point[1]
		}
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
	}

	// Initialize surface area
	area := 0.0

	// Scanline algorithm
	for y := minY; y <= maxY; y++ {
		inside := false
		for i := 0; i < len(path); i++ {
			j := (i + 1) % len(path)
			if (path[i][1] > y) != (path[j][1] > y) &&
				(path[i][0] <= path[j][0]) &&
				(path[j][0] <= path[i][0] || float64(path[i][0]+(y-path[i][1])*(path[j][0]-path[i][0])/(path[j][1]-path[i][1])) <= float64(path[j][0])) {
				inside = !inside
			}
		}

		if inside {
			area += float64(maxX - minX + 1)
		}
	}

	// Account for the edge width
	area *= 1.0

	return area
}

func main() {
	// Example usage
	path := [][]int{{0, 0}, {4, 0}, {4, 4}, {2, 6}, {0, 4}}
	surfaceArea := calculateSurfaceArea(path)
	fmt.Printf("Surface Area: %.2f\n", surfaceArea)
}

/*
Certainly! This section of the code implements the even-odd rule for determining whether a point is inside a polygon or not. Let me break it down:

(path[i][1] > y) != (path[j][1] > y): This part checks whether the y coordinate of the current scanline is between the y coordinates of two consecutive vertices (path[i] and path[j]). The expression (path[i][1] > y) != (path[j][1] > y) evaluates to true if one vertex is above the scanline and the other is below it. This condition is necessary for the edge intersecting the scanline.

(path[i][0] <= path[j][0]): This condition checks whether the x-coordinate of the current vertex (path[i]) is less than or equal to the x-coordinate of the next vertex (path[j]). It ensures that the algorithm processes edges in the correct order from left to right.

(path[j][0] <= path[i][0] || float64(path[i][0]+(y-path[i][1])*(path[j][0]-path[i][0])/(path[j][1]-path[i][1])) <= float64(path[j][0])): This part checks if the scanline intersects with the edge defined by the current and next vertices. It uses linear interpolation to find the x-coordinate of the intersection point on the edge given the y coordinate of the scanline. The condition path[j][0] <= path[i][0] || ... <= path[j][0] ensures that the intersection point is between the two vertices.

If all the above conditions are true, inside = !inside toggles the inside variable. The even-odd rule dictates that if a scanline intersects an odd number of edges, the point is inside the polygon. Toggling inside accounts for the odd number of intersections.

In summary, this section checks whether the current scanline intersects with the edge defined by the current and next vertices, and based on that, it updates the inside variable according to the even-odd rule for determining whether a point is inside the polygon.
*/
