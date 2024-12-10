# Advent of Code CLI Tool (aoc-cli)

A powerful CLI tool designed to streamline your Advent of Code journey. With aoc-cli, you can quickly set up and run Advent of Code puzzles without hassle, allowing you to focus on solving challenges!

  
  

## Features

- Puzzle Management: Automatically fetch and organize - puzzles by year and day.

- Project Directory Setup: Create a structured workspace for each day with input files and starter code.

- Run Puzzles: Execute your solutions directly from the CLI.

- Simple Configuration: Easy setup for Advent of Code cookies.

  

## Installation

  

There are two ways to install `aoc-cli`:

### Option 1: Clone and Build

1. Clone the repository:

		git clone <repository-url> 
		cd <repository-dir> 

3. Build the binary:

		go build -o aoc-cli

### Option 2: Download Prebuilt Binary

1. Download the binary from the [releases](https://github.com/Nabhdeep/aoc-cli/releases/tag/v0.0.1) page.

2. Place the binary in a directory of your choice.

  
  

## Setup an Alias

To make the tool globally accessible:

1. Open `~/.bashrc (or ~/.zshrc if using Zsh)`:

		nano ~/.bashrc

2. Add the following line:

		 alias aoc-cli="$HOME/DIR WHERE THE BUILD IS PLACED/aoc-cli" 

4. Refresh your shell configuration:
	
		. ~/.bashrc

  

## Usage

### Configure Your Cookie

Set up your Advent of Code session cookie:

	aoc-cli config

### Fetch a Puzzle

To fetch a puzzle, use:

	aoc-cli get YYYY --d DD

- Replace YYYY with the year (e.g., 2024).

- Replace DD with the day number (e.g., 01).

This will create:

- A directory named after the year (e.g., 2024/).

- Inside it, a directory for the day (e.g., 2024/day01/) containing:

- `input.txt` for the puzzle input.

- `dayX.go` as a starter file.

  

## Run a Puzzle

To run the solution for a specific puzzle:

	aoc-cli run YYYY --d DD

  

## Development

- Language: Go

- Framework: CobraCLI

Feel free to explore the codebase and contribute!

  

## Future Roadmap

- Add support for multiple languages beyond Go.

- Enhanced error handling and debugging features.

- Submitting day's solution

- 
