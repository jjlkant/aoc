"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var globals_1 = require("@jest/globals");
var day1_1 = require("../day1");
(0, globals_1.describe)("test input", function () {
    (0, globals_1.test)("Day 1.1 - example", function () {
        (0, globals_1.expect)((0, day1_1.process_file)("./2023/input/day1-1.txt")).toBe(142);
    });
});
//# sourceMappingURL=day1.test.js.map