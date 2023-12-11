import { describe, expect, test } from "@jest/globals";
import { process_file } from "../day4";

describe("Day 4 | Part 1", () => {
    test("Example input", () => {
        expect(process_file("./2023/input/day4-example.txt")).toBe(13);
    });
});

describe("Day 4 | Part 2", () => {
    test("Example input", () => {
        expect(process_file("./2023/input/day4-example.txt", true)).toBe(30);
    });
});
