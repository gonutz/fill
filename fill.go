package fill

import "github.com/gonutz/matrix"

// Fill fills a 2D image of given size width by height. The start point is given
// in seedX,seedY. This point is always filled. From there out, all neighbors
// are recursively inspected and all for which toFill returns true will be
// filled as well. The given fill function will be applied for the pixels to be
// filled.
func Fill(
	seedX, seedY int,
	width, height int,
	toFill func(x, y int) bool,
	fill func(x, y int),
	neighbors func(x, y int) [][2]int,
) {
	// By default we use the 4-neighborhood.
	if neighbors == nil {
		neighbors = Neighbors4
	}

	if toFill == nil {
		return // We don't know what to fill.
	}
	if fill == nil {
		return // Filling does nothing.
	}
	if seedX < 0 || seedX >= width || seedY < 0 || seedY >= height {
		return // Seed is outside of the image.
	}

	// Done stores for each pixel whether it is already queued for inspection or
	// if it was already filled.
	done := matrix.NewBitBoolMat(width, height)

	// q is the queue that holds all pixels to be filled.
	q := [][2]int{
		[2]int{seedX, seedY},
	}
	done.Set(seedX, seedY, true)

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		fill(x, y)
		for _, n := range neighbors(x, y) {
			nx, ny := n[0], n[1]
			if nx >= 0 && nx < width && ny >= 0 && ny < height && !done.Get(nx, ny) {
				done.Set(nx, ny, true)
				if toFill(nx, ny) {
					q = append(q, n)
				}
			}
		}
	}
}

// Neighbors4 defines the direct horizontal and vertical neghbors of x,y.
func Neighbors4(x, y int) [][2]int {
	return [][2]int{
		[2]int{x - 1, y},
		[2]int{x + 1, y},
		[2]int{x, y - 1},
		[2]int{x, y + 1},
	}
}

// Neighbors8 defines the direct diagonal, horizontal and vertical neghbors of
// x,y.
func Neighbors8(x, y int) [][2]int {
	return [][2]int{
		[2]int{x + 1, y},
		[2]int{x + 1, y + 1},
		[2]int{x + 1, y - 1},
		[2]int{x - 1, y},
		[2]int{x - 1, y + 1},
		[2]int{x - 1, y - 1},
		[2]int{x, y + 1},
		[2]int{x, y - 1},
	}
}
