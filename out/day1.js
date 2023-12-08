"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.process_file = void 0;
var fs = require("fs");
var digitWords = {
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
var re_forward = new RegExp("[0-9]|(?:" + Object.keys(digitWords).join("|") + ")");
var re_backward = new RegExp("[0-9]|(?:" +
    Object.keys(digitWords)
        .map(function (word) { return word.split("").reverse().join(""); })
        .join("|") +
    ")");
function process_line(line) {
    var firstDigit = line.match(re_forward)[0];
    var lastDigit = line
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
function process_file(filepath) {
    var input = fs.readFileSync(filepath, "utf-8");
    var lines = input.split("\r\n");
    var sum = 0;
    for (var i = 0; i < lines.length; i++) {
        var num = process_line(lines[i]);
        sum += num;
    }
    return sum;
}
exports.process_file = process_file;
console.log(process_file("2023/input/day1.txt"));
//# sourceMappingURL=day1.js.map