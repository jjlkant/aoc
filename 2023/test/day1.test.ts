import { describe, expect, test } from "@jest/globals";
import { process_file } from "../day1";

describe("test input", () => {
  test("Day 1.1 - example", () => {
    expect(process_file("./2023/input/day1-1.txt")).toBe(142);
  });
});
