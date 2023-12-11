"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.process_file = void 0;
var fs = require("fs");
function process_file(filepath) {
    var input = fs.readFileSync(filepath, "utf-8");
    var lines = input.split("\r\n");
    var chars;
    for (var i = 0; i < lines.length; i++) {
        console.log(lines[i]);
    }
}
exports.process_file = process_file;
//# sourceMappingURL=day3.js.map