#!/usr/bin/python3
from typing import List

# https://www.hackerrank.com/challenges/apple-and-orange/problem
def countApplesAndOranges(
    s: int, t: int, a: int, b: int, apples: List[int], oranges: List[int]
) -> List[int]:
    """
    Determine how many apples and oranges will fall on a house roof
    given various distances
    s: the location of the leftmost position of the house
    t: the location of the rightmost position of the house
    a: the location of the apple tree (left of the house)
    b: the location of the orange tree (right of the house)
    apples: a list of location of apples relative to the apple tree location
        +d to the right of the tree, -d to the left of the tree
    oranges: same as apples param but an array of oranges from the orange tree
    """
    apples_on_roof = 0
    oranges_on_roof = 0
    for apple in apples:
        d_apple = a + apple
        if s <= d_apple and d_apple <= t:
            apples_on_roof += 1

    for orange in oranges:
        d_orange = b + orange  # negative values will be given for left
        if s <= d_orange and d_orange <= t:
            oranges_on_roof += 1
    print(f"Apples: {apples_on_roof} Oranges: {oranges_on_roof}")
    return [apples_on_roof, oranges_on_roof]


countApplesAndOranges(7, 11, 5, 15, [-2, 2, 1], [5, -6])
