"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var globals_1 = require("@jest/globals");
var day3_1 = require("../day3");
(0, globals_1.describe)("test input", function () {
    (0, globals_1.test)("Day 3.1 - example", function () {
        (0, globals_1.expect)((0, day3_1.process_file)("./2023/input/day3-example.txt")).toBe(4361);
    });
});
//# sourceMappingURL=day3.test.js.map