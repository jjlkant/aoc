import * as fs from "fs";

function processScratchCard(
    scratchcard: string,
    power: boolean = true
): number {
    let game = scratchcard.split(": ")[1].split(" | ");

    const winningNumbers = game[0].split(" ").filter((n) => n);
    const ownNumbers = game[1].split(" ").filter((n) => n);

    const ownWinningNumbers = winningNumbers.filter((value) =>
        ownNumbers.includes(value)
    );

    if (ownWinningNumbers.length < 1) {
        return 0;
    }

    if (power) {
        return Math.pow(2, ownWinningNumbers.length - 1);
    } else {
        return ownWinningNumbers.length;
    }
}

function processFileForPoints(filepath: string): number {
    const input = fs.readFileSync(filepath, "utf-8");
    const lines = input.split("\r\n");
    let ret = 0;
    for (let i = 0; i < lines.length; i++) {
        let value = processScratchCard(lines[i]);
        ret += value;
    }
    return ret;
}

function processFileWithCopies(filepath: string): number {
    const input = fs.readFileSync(filepath, "utf-8");
    const lines = input.split("\r\n");
    let copies: number[] = new Array(lines.length).fill(1);
    for (let i = 0; i < lines.length; i++) {
        let value = processScratchCard(lines[i], false);
        for (let j = 1; j <= value; j++) {
            copies[i + j] += copies[i];
        }
    }
    let sum = 0;
    for (let i = 0; i < copies.length; i++) {
        sum += copies[i];
    }
    return sum;
}

export function process_file(
    filepath: string,
    copies: boolean = false
): number {
    if (copies) {
        return processFileWithCopies(filepath);
    }
    return processFileForPoints(filepath);
}

console.log(process_file("2023/input/day4-example.txt"));
console.log(process_file("2023/input/day4.txt"));
console.log(process_file("2023/input/day4-example.txt", true));
console.log(process_file("2023/input/day4.txt", true));
