# Advent of Code 2024 - Go Solutions

Solutions to [Advent of Code 2024](https://adventofcode.com/2024) in Go.

## Quick Reference

| Day | Puzzle Name | Part 1 | Part 2 | Algorithm/Technique |
|-----|-------------|--------|--------|---------------------|
| 1 | Historian Hysteria | Calculate pairwise distance sum | Calculate similarity score | Sorting, hash map counting |
| 2 | Red-Nosed Reports | Count safe reports (monotonic) | Count safe with Problem Dampener | Sequence validation, brute force |
| 3 | Mull It Over | Sum mul() statements | Filter with do()/don't() | Regex parsing, conditional filtering |
| 4 | Ceres Search | Find XMAS in all directions | Find X-MAS cross patterns | Grid search, direction vectors |
| 5 | Print Queue | Validate page order | Fix incorrect orders | Topological sort, custom comparator |
| 6 | Guard Gallivant | Trace guard path | Count loop-creating obstacles | Grid simulation, cycle detection |
| 8 | Resonant Collinearity | Find antinodes between antennas | Find all resonant harmonics | Vector math, collinearity |
| 9 | Disk Fragmenter | Compact disk blocks | Move whole files | Array manipulation, fragmentation |
| 10 | Hoof It | Count reachable peaks (9s from 0s) | Count all distinct trails | DFS, path counting |
| 11 | Plutonian Pebbles | Count stones after 25 blinks | Count after 75 blinks | Map-based counting, memoization |
| 12 | Garden Groups | Calculate fence cost (area × perimeter) | Bulk discount (area × sides) | Flood fill, perimeter calculation |
| 15 | Warehouse Woes | Push boxes, calculate GPS sum | Push wide boxes | Grid simulation, movement rules |
| 16 | Reindeer Maze | Find cheapest path (turns cost 1000) | Count all tiles on best paths | Dijkstra with rotation costs, backtracking |
| 17 | Chronospatial Computer | Execute program, get output | Find quine input | Virtual machine, opcode execution |
| 18 | RAM Run | Find shortest path avoiding bytes | Find first blocking byte | Dijkstra, binary search |
| 19 | Linen Layout | Count possible designs | Count all design combinations | Recursive pattern matching, memoization |

## Algorithm Summary

| Category | Days | Techniques |
|----------|------|------------|
| **Graph/Pathfinding** | 10, 16, 18 | DFS, Dijkstra, shortest path, backtracking |
| **Simulation** | 6, 15, 17 | Grid movement, state machines, cycle detection |
| **Pattern Matching** | 3, 4, 19 | Regex, string search, recursive matching |
| **Optimization** | 9, 11, 12 | Memoization, map-based counting, space optimization |
| **Sorting/Ordering** | 1, 5 | Custom sorting, topological sort |
| **Geometry** | 8 | Vector math, collinear points |
| **Validation** | 2 | Sequence constraints, brute force |

## Days Implemented in Other Languages

| Day | Language | Reason |
|-----|----------|--------|
| 7 | Python | day07-python/ |
| 13 | Python | day13-python/ |
| 14 | Python | day14-python/ |

## Running Solutions

```bash
cd dayXX
go run main.go        # Most days
go test              # Days with test files
```

## Implementation Notes

- Prioritizes clarity and correctness
- Uses Go standard library features (maps, slices)
- Memoization used for performance on recursive problems (Days 11, 19)
- Grid problems use 2D rune slices for efficient manipulation
- Custom data structures where beneficial (priority queues, state tracking)
