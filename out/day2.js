"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.process_file = void 0;
var fs = require("fs");
var cubes = {
    red: 12,
    green: 13,
    blue: 14,
};
function power_set(set) {
    var setCubes = set.split(",");
    var minRed = 0;
    var minGreen = 0;
    var minBlue = 0;
    return minRed * minGreen * minBlue;
}
function check_set(set) {
    var setCubes = set.split(",");
    for (var i = 0; i < setCubes.length; i++) {
        var arr = setCubes[i].split(" ");
        var count = arr[1];
        var color = arr[2];
        if (cubes[color] < count) {
            return false;
        }
    }
    return true;
}
function process_file(filepath) {
    var input = fs.readFileSync(filepath, "utf-8");
    var lines = input.split("\r\n");
    var sum = 0;
    for (var i = 0; i < lines.length; i++) {
        console.log(lines[i]);
        var arr = lines[i].split(":");
        var gameDesc = arr[0];
        var gameId = gameDesc.split(" ");
        var game = arr[1];
        var sets = game.split(";");
        var minRed = 0;
        var minGreen = 0;
        var minBlue = 0;
        for (var j = 0; j < sets.length; j++) {
            var setCubes = sets[j].split(",");
            for (var k = 0; k < setCubes.length; k++) {
                var arr_1 = setCubes[k].split(" ");
                var count = arr_1[1];
                var color = arr_1[2];
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
        var power = minRed * minGreen * minBlue;
        sum += power;
    }
    return sum;
}
exports.process_file = process_file;
// console.log(process_file("2023/input/day2-1.txt"));
console.log(process_file("2023/input/day2.txt"));
//# sourceMappingURL=day2.js.map