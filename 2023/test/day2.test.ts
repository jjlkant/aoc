import { describe, expect, test } from "@jest/globals";
import { process_file } from "../day2";

describe("test input", () => {
  test("Day 2.2 - example", () => {
    expect(process_file("./2023/input/day2-1.txt")).toBe(2286);
  });
});
