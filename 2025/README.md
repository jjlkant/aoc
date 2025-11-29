# 🌟 Advent of Code 2025 - Solutions

This repository contains my solutions for the Advent of Code 2025 puzzles, implemented in **Go**.

## ⚙️ Project Structure

The project is structured to maintain consistency, separate the independent daily puzzles, and maximize code reusability.

```bash
/2025
|--- /day01          # Solution for Day 1
|    |--- main.go
|    |--- main_test.go
|    |--- input.txt       # Main Puzzle Input
|    |--- test_input.txt  # Example Puzzle Input
|
|--- /utils          # Shared Go package for common utilities (input reading, parsing)
|    |--- input.go
|
|--- /template       # Source files for 'go generate'
|
|--- go.mod          # Go module definition
|--- generate.go     # Holds the 'go:generate' directive
|--- README.md
```

-----

## 🚀 Getting Started

### Prerequisites

1. **Go:** Ensure Go (version 1.25.4 or newer) is installed and available in your shell's `$PATH`.
2. **Ubuntu/Zsh:** This setup relies on Bash scripting (`create_day.sh`) for automation.

### Initialization

Navigate to the `2025` directory and initialize the Go module:

```bash
go mod tidy
```

-----

## 💻 Daily Workflow (Automation)

The core of this setup is the automated creation of new solution files using **`go generate`**. This ensures every new day starts with the correct file structure, imports, and test stubs.

### 1\. Generate the New Day

To generate the directory and files for a specific day, use the `AOC_DAY` environment variable:

| Goal | Command |
| :--- | :--- |
| **Generate Day 1** | `go generate` |
| **Generate Day 5** | `AOC_DAY=5 go generate` |

This command executes the `./create_day.sh` script, which:

* Creates the directory (e.g., `day05`).
* Copies `main.go`, `main_test.go`, `input.txt`, and `test_input.txt`.

### 2\. Input and Test Setup

Once the day is generated (e.g., `day05`):

1. **Paste Main Input:** Paste the massive puzzle input into **`day05/input.txt`**.
2. **Paste Example Input:** Paste the small example input from the AoC problem description into **`day05/test_input.txt`**.
3. **Set Expected Answer:** Open **`day05/main_test.go`** and replace the `const expected = 0` placeholder with the known correct answer for the example input.

### 3\. Solve and Test

The solution logic is separated into callable functions (`solvePart1`, `solvePart2`) within `main.go`.

**Develop your solution using Go's built-in testing utility:**

```bash
# Run tests for the current day only
go test ./day05 

# Run all tests across the entire 2025 module
go test ./... 
```

**Once all tests pass, run the final solution:**

```bash
go run ./day05/main.go
```

-----

## 📦 Shared Utilities (`/utils`)

The `/utils` package holds common, reusable logic, primarily input handling.

* **`utils.ReadInputLines()`:** This function is called by every daily `main.go`. It automatically finds the local `input.txt` (or `test_input.txt` when testing) and returns the content as a slice of strings (`[]string`), handling trimming and splitting for you. This allows your daily solution files to focus strictly on the puzzle logic.
