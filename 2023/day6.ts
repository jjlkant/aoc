import * as fs from "fs";

function waysToBeat(time: number, distance: number): number {
    let minTime = 0;
    let maxTime = time;
    for (let i = 0; i < (time + 1) / 2; i++) {
        // distance = x * (time - x)
        if (minTime * (time - minTime) <= distance) {
            minTime += 1;
        }
        if (maxTime * (time - maxTime) <= distance) {
            maxTime -= 1;
        }
    }
    return maxTime - minTime + 1;
}

export function process_file(filepath: string, spaces: boolean = true): number {
    const input = fs.readFileSync(filepath, "utf-8");
    const lines = input.split("\r\n");
    if (spaces) {
        const times = lines[0]
            .split(":")[1]
            .split(" ")
            .filter((x) => x !== "")
            .map(Number);
        const distances = lines[1]
            .split(":")[1]
            .split(" ")
            .filter((x) => x !== "")
            .map(Number);
        let options = 1;
        for (let i = 0; i < times.length; i++) {
            options *= waysToBeat(times[i], distances[i]);
        }
        return options;
    }
    const time = Number(
        lines[0]
            .split(":")[1]
            .split(" ")
            .filter((x) => x !== "")
            .join("")
    );
    const distance = Number(
        lines[1]
            .split(":")[1]
            .split(" ")
            .filter((x) => x !== "")
            .join("")
    );
    return waysToBeat(time, distance);
}

console.log(process_file("2023/input/day6-example.txt"));
console.log(process_file("2023/input/day6.txt"));
console.log(process_file("2023/input/day6-example.txt", false));
console.log(process_file("2023/input/day6.txt", false));
