#!/bin/bash

# This script is called by 'go generate' to set up a new day's directory.

DAY=$AOC_DAY 

# If AOC_DAY was not set or invalid, set the default day to 1.
if [ -z "$DAY" ] || [ "$DAY" -lt 1 ]; then
    echo "Day argument not provided or invalid. Defaulting to Day 1."
    DAY=1
fi

# Format the day number with leading zero (e.g., 1 -> 01)
DAY_DIR=$(printf "day%02d" $DAY)

echo "--- Advent of Code Day $DAY Setup ---"

# Check if the directory already exists
if [ -d "$DAY_DIR" ]; then
    echo "Error: Directory '$DAY_DIR' already exists. Aborting."
    exit 1
fi

# 1. Create the new day's directory
mkdir "$DAY_DIR"

# 2. Copy the main application file template
cp template/main.go.tmpl "$DAY_DIR/main.go"

# 3. Create the empty input file (for your main puzzle input)
touch "$DAY_DIR/input.txt"

# 4. Copy the test file template
cp template/main_test.go.tmpl "$DAY_DIR/main_test.go"

# 5. Create the empty test input file (for the AoC example input)
touch "$DAY_DIR/test_input.txt"

# 6. Replace the placeholder 'XX' in the new main.go
#    (Note: The sed command works differently on macOS/BSD, but this is fine for Ubuntu)
sed -i "s/Day XX/Day $DAY/" "$DAY_DIR/main.go"

echo "Successfully created:"
echo "- $DAY_DIR/main.go"
echo "- $DAY_DIR/main_test.go"
echo "- $DAY_DIR/input.txt (Main Input)"
echo "- $DAY_DIR/test_input.txt (Example Input)"
echo "------------------------------------"
echo "Next Steps:"
echo "1. Paste example into $DAY_DIR/test_input.txt and answer into main_test.go."
echo "2. Run 'go test ./$DAY_DIR' to verify your solution."