"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var globals_1 = require("@jest/globals");
var day2_1 = require("../day2");
(0, globals_1.describe)("test input", function () {
    (0, globals_1.test)("Day 2.2 - example", function () {
        (0, globals_1.expect)((0, day2_1.process_file)("./2023/input/day2-1.txt")).toBe(2286);
    });
});
//# sourceMappingURL=day2.test.js.map