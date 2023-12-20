# Tetromino Square Assembler

This is a program written in Go that assembles tetrominoes from a given text file, in order to create the smallest square possible. Each tetromino is identified and printed with uppercase Latin letters (A for the first one, B for the second, and so on).

## Table of Contents

- [Objective](#objective)
- [Instructions](#instructions)
- [Example File](#example-file)
- [Usage](#usage)
- [Requirements](#requirements)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Objective

The objective of this program is to develop a solution that can assemble tetrominoes from a given text file into a square shape. The program should compile successfully, handle different tetromino shapes, and create the smallest possible square using the provided tetrominoes. It should also follow good coding practices and have test files for unit testing.

## Instructions

1. Ensure you have Go installed on your system.

2. Clone this repository to your local machine.

3. Navigate to the project's root directory.

4. Compile the program by running the following command:

   ````shell
   go build
   ```

5. Run the program with a text file containing the tetrominoes as the argument:

   ````shell
   ./tetromino-square-assembler <path-to-text-file>
   ```

6. The program will assemble the tetrominoes and output the smallest square possible with uppercase Latin letters representing each tetromino. If the tetrominoes or file format are incorrect, the program will display an "ERROR" message.

## Example File

An example of a text file containing tetrominoes:

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

If it isn't possible to form a complete square, the program will leave spaces between the tetrominoes. For example:

```
ABB.
ABB.
A...
A...
```

## Usage

1. Ensure you have a text file with the correct tetromino format.

2. Run the program by following the instructions mentioned in the [Instructions](#instructions) section.

3. The program will process the tetrominoes and output the assembled square with uppercase letters representing each tetromino.

## Requirements

- Go 1.16 or higher

## Testing

1. Navigate to the project's root directory.

2. Run the following command to execute the unit tests:

   ````shell
   go test ./...
   ```

3. The test results will be displayed, indicating the success or failure of each test case.

## Contributing

Contributions to this project are welcome. Feel free to open issues or submit pull requests to suggest improvements, add new features, or fix any existing bugs.

## License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, or distribute this software as per the terms of this license.