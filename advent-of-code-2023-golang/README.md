# Advent of Code 2023 - Go Solutions

Solutions to [Advent of Code 2023](https://adventofcode.com/2023) in Go.

## Quick Reference

| Day | Puzzle Name | Part 1 | Part 2 | Algorithm/Technique |
|-----|-------------|--------|--------|---------------------|
| 1 | Trebuchet?! | Extract first/last digits | Include spelled-out numbers | Regex, string replacement |
| 2 | Cube Conundrum | Validate games against limits | Find minimum cube set | Parsing, max tracking |
| 3 | Gear Ratios | Sum numbers adjacent to symbols | Find gear ratios (* with 2 numbers) | Grid traversal, neighbor scanning |
| 4 | Scratchcards | Score matching numbers (2^n-1) | Card copying cascades | Set intersection, dynamic counting |
| 5 | Seed Fertilizer | Map seeds through ranges | Treat seeds as ranges | Range mapping, interval arithmetic |
| 6 | Wait For It | Count winning hold times | Single combined race | Quadratic formula / brute force |
| 7 | Camel Cards | Rank poker hands | J as joker (lowest value) | Hand classification, custom sorting |
| 8 | Haunted Wasteland | Navigate AAA → ZZZ | All **A → **Z simultaneously | Ring buffer, LCM cycle detection |
| 9 | Mirage Maintenance | Predict next in sequence | Predict previous in sequence | Difference pyramid, backtracking |
| 10 | Pipe Maze | Find farthest point in loop | Count enclosed tiles | Pipe navigation, area detection |
| 11 | Cosmic Expansion | Galaxy distances (2x expansion) | 1,000,000x expansion | Manhattan distance, weighted expansion |
| 12 | Hot Springs | Count valid arrangements | 5x unfolded patterns | Dynamic programming, memoization |
| 13 | Point of Incidence | Find reflection lines | Find reflection after fixing smudge | Mirror detection, brute force smudges |
| 14 | Parabolic Reflector | Tilt north, calculate load | 1 billion spin cycles | Grid transpose, string manipulation |
| 15 | Lens Library | HASH algorithm sum | Lens box arrangement | Custom hash, ordered lists |
| 16 | Floor Lava | Trace beam from top-left | Max energization from any edge | Beam simulation, direction tracking |
| 17 | Clumsy Crucible | Min path (max 3 straight) | Min path (4-10 straight blocks) | Dijkstra with constraints |
| 18 | Lavaduct Lagoon | Polygon area from RGB | Hex-encoded instructions | Shoelace formula + Pick's theorem |
| 19 | Aplenty | Filter parts through workflows | Count all accepted combinations | Workflow parsing, conditionals |
| 20 | Pulse Propagation | Count pulses after 1000 presses | Find when module gets low pulse | Module simulation, queue processing |
| 21 | Step Counter | Reachable plots in 64 steps | Infinite grid, 26M steps | BFS, modulo wrapping |
| 22 | Sand Slabs | *(not implemented)* | *(not implemented)* | - |
| 25 | Never Tell Me Odds | 2D hailstone intersections | Find trajectory hitting all | Line intersection, constraint solving |

## Algorithm Summary

| Category | Days | Techniques |
|----------|------|------------|
| **Graph/Pathfinding** | 8, 10, 17, 21 | BFS, Dijkstra, cycle detection, LCM |
| **Dynamic Programming** | 12 | Memoization, recursive pattern matching |
| **Geometry** | 18, 25 | Shoelace formula, Pick's theorem, line intersection |
| **Grid Manipulation** | 3, 11, 13, 14, 16 | Transpose, rotation, neighbor scanning |
| **Simulation** | 6, 14, 16, 20 | State tracking, queue processing |
| **String/Parsing** | 1, 2, 4, 5, 7 | Regex, custom parsing, sorting |
| **Math** | 8, 11 | GCD, LCM, weighted distances |

## Running Solutions

```bash
cd dayX
go test              # Most days
go run main.go       # Days 1-7, 22 (where applicable)
```

## Implementation Notes

- Prioritizes correctness over optimization
- Uses Go standard library extensively
- Test files include sample verification + actual solutions
- Some solutions use brute force where input size permits
