# Introduction to Data Structures

## Abstract Data Types x Data Structures

Abstract Data Types (ADT) defines how is the data structure behavior.

And a Data Structure is a ADT's implementation

For example, a List needs to be able to add and remove a item. This is the behavior of a list, then it's a ADT

By otherwise, when we have a Linked List or a Dynamic Array, both are Data Structures.

## Complexity Analysis

- How much time does this algorithm needs to finish?
- How much space does this algorithm needs for its computation?

do you know how answer these questions?

### Big-O Notation

Big-O consider the worst input (case) for your algorithm. It analyze and define the algorithm complexity:

n — The size of the input

Complexities ordered in f rom smallest to largest

- Constant: O(1)
- Logarithmic: O(log(n))
- Linear: O(n)
- Linear Logarithmic: O(n log(n))
- Quadric: O(n^3)
- Cubic: O(n^3)
- Exponential: O(b^n), b > 1
- Factorial: O(n!)
