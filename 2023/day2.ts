import * as fs from "fs";
import * as rd from "readline";

const cubes = {
  red: 12,
  green: 13,
  blue: 14,
};

function power_set(set: string): number {
  let setCubes = set.split(",");
  let minRed = 0;
  let minGreen = 0;
  let minBlue = 0;

  return minRed * minGreen * minBlue;
}

function check_set(set: string): boolean {
  let setCubes = set.split(",");
  for (let i = 0; i < setCubes.length; i++) {
    let arr = setCubes[i].split(" ");
    let count = arr[1];
    let color = arr[2];
    if (cubes[color] < count) {
      return false;
    }
  }
  return true;
}

export function process_file(filepath: string): number {
  const input = fs.readFileSync(filepath, "utf-8");
  const lines = input.split("\r\n");
  let sum = 0;
  for (let i = 0; i < lines.length; i++) {
    console.log(lines[i]);
    let arr = lines[i].split(":");
    var gameDesc = arr[0];
    let gameId = gameDesc.split(" ");
    var game = arr[1];
    var sets = game.split(";");
    let minRed = 0;
    let minGreen = 0;
    let minBlue = 0;
    for (let j = 0; j < sets.length; j++) {
      let setCubes = sets[j].split(",");
      for (let k = 0; k < setCubes.length; k++) {
        let arr = setCubes[k].split(" ");
        let count = arr[1];
        let color = arr[2];
        switch (color) {
          case "red":
            minRed = Math.max(minRed, Number(count));
            break;
          case "green":
            minGreen = Math.max(minGreen, Number(count));
            break;
          case "blue":
            minBlue = Math.max(minBlue, Number(count));
            break;
        }
      }
    }
    // console.log("Red: " + minRed);
    // console.log("Green: " + minGreen);
    // console.log("Blue: " + minBlue);
    let power = minRed * minGreen * minBlue;
    sum += power;
  }

  return sum;
}

// console.log(process_file("2023/input/day2-1.txt"));
console.log(process_file("2023/input/day2.txt"));
