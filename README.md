# concat

`concat` is a command-line interface (CLI) tool written in Go that concatenates a given string based on a specified separator character. It also provides options to split the input string on a specified character and to preserve or transform the casing of the output string.

## Usage

### Options

- `-s` (string): The string to be concatenated. If not provided, the program will prompt for user input.
- `-sep` (string): The separator character to use for concatenation. Default is an empty string.
- `-split` (string): The character to split the input string on. Default is a space character.
- `-c` (bool): If set to true, the casing of the input string will be preserved. By default, the output string is transformed to lowercase.

If no options are provided, the program will prompt the user to enter the required inputs interactively.

## Examples

1. Concatenate a string with a comma separator and preserve casing:

```sh
concat -s "Hello World" -sep "," -c

Output: `Hello,World`
```
2. Concatenate a string with a hyphen separator, split on the letter 'e', and transform to lowercase:

```sh
concat -s "Hello World" -sep "-" -split "e"

Output: `Hello,World`
```

3. Interactive mode:

```sh
concat
Enter a string to concat: Hello World
Enter a separator: /
Enter a split character: o

Output: `H/ell/W/rld`
```

## Installation

To install `concat`, you'll need to have Go installed on your system. Then, you can clone the repository and build the binary:


```sh
git clone https://github.com/TWolfis/concat.git
cd concat
go build

# install the executable in the directory named by the GOBIN environment variable
go install
```

The compiled binary will be available in the current directory.

## Contributing

Contributions to `concat` are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the project's GitHub repository.
