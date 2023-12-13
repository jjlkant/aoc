import { describe, expect, test } from "@jest/globals";
import { process_file } from "../day6";

describe("Day 6 | Part 1", () => {
    test("Example input", () => {
        expect(process_file("./2023/input/day6-example.txt")).toBe(288);
    });
});

describe("Day 6 | Part 2", () => {
    test("Example input", () => {
        expect(process_file("./2023/input/day6-example.txt", false)).toBe(71503);
    });
});
