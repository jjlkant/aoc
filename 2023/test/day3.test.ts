import { describe, expect, test } from "@jest/globals";
import { process_file } from "../day3";

describe("Day 3 | Part 1", () => {
    test("Example input", () => {
        expect(process_file("./2023/input/day3-example.txt")).toBe(4361);
    });
});

describe("Day 3 | Part 2", () => {
    test("Example input", () => {
        expect(process_file("./2023/input/day3-example.txt", true)).toBe(
            467835
        );
    });
});
