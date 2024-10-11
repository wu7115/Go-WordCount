# Word Count Implementations in Go

Welcome to my **Go Word Count Project**! This repository showcases my exploration of various programming styles to solve the same problem: counting word frequencies in a text file. Each solution is implemented in Go with a distinct approach, demonstrating different paradigms in programming.

## Problem Statement

The objective of this project is to count the frequency of words in a given text file, filtering out stop words, and outputting the top 25 most frequent words. Each implementation addresses this challenge using a different programming style.

## Implementations

### 1. **No Constraints (Standard Approach)**
   - This version uses the full power of Go’s standard library and data structures (e.g., maps) without any imposed restrictions.
   - **Characteristics**: Standard Go features, efficient, straightforward solution.
   - To run code
     ```bash
   go run wc01.go pride-and-prejudice.txt
   
   File: [`wc01.go`](./wc01.go)

### 2. **No Abstractions, No Library Functions, Slices Only**
   - This approach prohibits the use of abstractions (no custom functions) and avoids Go's library functions and any advanced data structure. It only utilizes basic functions and slices for data manipulation.
   - **Characteristics**: Simplistic, avoids abstraction layers, restricted use of features.
   
   File: [`wc02.go`](./wc02.go)

### 3. **Procedural Abstraction**
   - The solution decomposes the problem into smaller procedures, each handling a specific task. The main logic executes these procedures sequentially, solving the problem step by step.
   - **Characteristics**: Modular, structured through procedural decomposition, more organized than the basic approach.
   
   File: [`wc03.go`](./wc03.go)

### 4. **Functional Abstraction**
   - This implementation uses functional programming principles, decomposing the problem into mathematical functions. It solves the problem as a pipeline of function applications.
   - **Characteristics**: Focus on immutability, function composition, and avoiding side effects.
   
   File: [`wc04.go`](./wc04.go)

### 5. **Minimal Lines of Code**
   - In this version, the goal is to solve the word count problem in as few lines of code as possible while maintaining functionality.
   - **Characteristics**: Compact, concise, and focused on brevity.
   
   File: [`wc05.go`](./wc05.go)

## How to Run the Code

Each implementation requires two input files: the text file to analyze and a stop words file. The stop words file contains common words to exclude from the word count.

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/word-count-go.git
   cd word-count-go
