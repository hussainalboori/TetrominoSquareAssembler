
# Tetromino Assembler

# Description
The Tetromino Assembler project aims to develop a program that assembles tetrominoes in order to create the smallest square possible. Tetrominoes are geometric shapes composed of four squares, similar to the shapes in the game Tetris. The program takes a text file as input, which contains a list of tetrominoes represented by characters. The program organizes these tetrominoes to form a square, minimizing the empty space between them.


## Features

- Assembles tetrominoes to create the smallest square possible
- Identifies each tetromino in the solution with uppercase Latin letters (A, B, C, etc.)
- Handles errors for incorrect tetromino or file format
- Written in Go programming language
- Includes unit tests for ensuring functionality


## Usage
1. Clone the project repository.

```bash
git clone https://github.com/hussainalboori/TetrominoSquareAssembler
```
2. 2. Install Go if not already installed on your system.
3. Run The Project 
``` bash 
go run . <path of Tetromino>
```
## Example

For example, consider the following input text file:

```
#...
#...
#...
#...

....
....
..##
..##
```

Running the Tetromino Assembler program with this file as input will produce the following output:

```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

In this output, each tetromino is represented by a letter, and the resulting square is formed by assembling the tetrominoes.



## Dependencies

The Tetromino Assembler project uses only the standard Go packages, so there are no external dependencies.

## Contributing

Contributions to the Tetromino Assembler project are welcome. If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on the project repository.


## License

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/).


