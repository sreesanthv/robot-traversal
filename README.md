# Robot Navigation on a Rectangular Plane

This project implements a robot movement simulation in Go. The robot moves across a rectangular 2D grid while avoiding obstacles (particles), boundaries, and previously visited positions.

## Problem Description

A team of young engineers has built a robot that travels on a rectangular plane. The robot receives a set of movement instructions and must follow them while observing a few rules.

### Plane Specifications

*   The rectangular plane is defined by its top-right corner coordinate `(M, N)`. The bottom-left is assumed to be `(0, 0)`.
*   The plane contains **particles** (obstacles) at specific coordinates which the robot must avoid.

### Robot Movement

*   The robot's position is defined by `(x, y, D)` where `x` and `y` are coordinates and `D` is one of the four cardinal directions: `N`, `S`, `E`, `W`.
*   It accepts a command string containing characters:
    *   `L`: Turn 90° left
    *   `R`: Turn 90° right
    *   `M`: Move forward one unit in the current direction
*   The robot stops if:
    *   It encounters a **particle**
    *   It attempts to **revisit a previous location**
    *   It tries to **move outside** the defined rectangular plane

## Input Format

1.  First line: Top-right corner of the rectangular plane `(M N)`
2.  Second line: Initial position and direction of the robot `(X Y D)`
3.  Third line: A string of movement commands (`L`, `R`, `M`)
4.  _(Optional)_: Additional lines to define particles (each with `X Y`)

### Example Input

```
4 4
0 0 N
MMMRMMLM
```

## Output Format

Final position and direction of the robot as a single line:

```
X Y D
```

### Example Output

```
2 4 N
```

## Running the Project

1.  Make sure you have Go installed.
2.  Clone this repository.
3.  Navigate into the project directory.
4.  Run the program:

```
go run main.go
```

Provide the input via standard input (stdin).

## Project Structure

```
.
├── main.go         # Entry point of the application
├── robot.go        # Robot movement logic
├── grid.go         # Grid and particle tracking
├── README.md       # Project documentation
```

## Features

*   Handles direction changes and movement logic
*   Validates against boundaries and previously visited positions
*   Extensible for adding particles or multiple robots
