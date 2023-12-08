import * as fs from "fs";
import * as rd from "readline";

const digitWords = {
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
};

const re_forward = new RegExp(
  "[0-9]|(?:" + Object.keys(digitWords).join("|") + ")"
);
const re_backward = new RegExp(
  "[0-9]|(?:" +
    Object.keys(digitWords)
      .map((word) => word.split("").reverse().join(""))
      .join("|") +
    ")"
);

function process_line(line: string): number {
  let firstDigit = line.match(re_forward)[0];
  let lastDigit = line
    .split("")
    .reverse()
    .join("")
    .match(re_backward)[0]
    .split("")
    .reverse()
    .join("");
  if (firstDigit.length > 1) {
    firstDigit = digitWords[String(firstDigit)];
  }
  if (lastDigit.length > 1) {
    lastDigit = digitWords[String(lastDigit)];
  }
  return Number([firstDigit, lastDigit].join(""));
}

export function process_file(filepath: string): number {
  const input = fs.readFileSync(filepath, "utf-8");
  const lines = input.split("\r\n");

  var sum: number = 0;
  for (let i = 0; i < lines.length; i++) {
    var num = process_line(lines[i]);
    sum += num;
  }
  return sum;
}

console.log(process_file("2023/input/day1.txt"));
