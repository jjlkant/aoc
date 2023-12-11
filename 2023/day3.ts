import * as fs from "fs";

function isSymbol(chars: string[][], i: number, j: number): boolean {
    if (isNaN(Number(chars[i][j])) && chars[i][j] != ".") {
        return true;
    }
    return false;
}

function findFullPartNumber(
    chars: string[][],
    x: number,
    y: number,
    seen: boolean[][]
): number {
    let partNumber: string[] = [];
    partNumber.push(chars[x][y]);
    for (let i = 1; i < 3; i++) {
        seen[x][y + i] = true;
        if (isNaN(Number(chars[x][y + i]))) {
            break;
        }
        partNumber.push(chars[x][y + i]);
    }
    for (let i = 1; i < 3; i++) {
        seen[x][y - i] = true;
        if (isNaN(Number(chars[x][y - i]))) {
            break;
        }
        partNumber.unshift(chars[x][y - i]);
    }
    return Number(partNumber.join(""));
}

function findPartNumbers(
    chars: string[][],
    i: number,
    j: number,
    seen: boolean[][],
    gear: boolean
): number[] {
    const steps = [
        [-1, -1],
        [-1, 0],
        [-1, 1],
        [0, -1],
        [0, 1],
        [1, -1],
        [1, 0],
        [1, 1],
    ];
    let partNumbers: number[] = [];
    for (let k = 0; k < steps.length; k++) {
        let x = i + steps[k][0];
        let y = j + steps[k][1];
        if (x < 0 || x >= chars[i].length || y < 0 || y >= chars.length) {
            continue;
        }
        if (seen[x][y] === true) {
            continue;
        }

        seen[x][y] = true;
        if (isNaN(Number(chars[x][y]))) {
            continue;
        }
        partNumbers.push(findFullPartNumber(chars, x, y, seen));
    }
    if (gear && partNumbers.length !== 2) {
        return [];
    }
    return partNumbers;
}

export function process_file(filepath: string, gear: boolean = false): number {
    const input = fs.readFileSync(filepath, "utf-8");
    const chars = input.split("\r\n").map((x) => x.split(""));

    // Part 1
    const seen: boolean[][] = [];
    for (let i = 0; i < chars.length; i++) {
        seen.push(new Array(chars[i].length).fill(false));
    }
    let partNumbers: number[] = [];
    for (let i = 0; i < chars.length; i++) {
        for (let j = 0; j < chars[i].length; j++) {
            if (isSymbol(chars, i, j)) {
                let adjParts: number[] = findPartNumbers(
                    chars,
                    i,
                    j,
                    seen,
                    gear
                );
                if (gear && adjParts.length === 2) {
                    partNumbers.push(adjParts[0] * adjParts[1]);
                } else {
                    partNumbers.push(...adjParts);
                }
            }
        }
    }

    let sum = 0;
    for (let i = 0; i < partNumbers.length; i++) {
        sum += partNumbers[i];
    }
    return sum;
}

console.log(process_file("2023/input/day3-example.txt"));
console.log(process_file("2023/input/day3.txt"));
console.log(process_file("2023/input/day3-example.txt", true));
console.log(process_file("2023/input/day3.txt", true));
