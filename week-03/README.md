## Implementations

### 6. **Recursive Sorting**
   - This version introduces recursion, specifically in the sorting function used for the word frequencies. The implementation utilizes MergeSort algorithm.
   - **Characteristics**: Focus on recursion in sorting logic.
   - To run code:
     ```bash
     go run wc06.go ../pride-and-prejudice.txt ../stop_words.txt
     ```
   File: [`wc06.go`](./wc06.go)

### 7. **Function Pipeline with Parameter Passing**
   - In this implementation, each function takes an additional parameter: another function. This additional function is applied at the end of the current function and given the output of the current function as input. The larger problem is solved as a pipeline of these functions, where the next function is passed as a parameter to the current function.
   - **Characteristics**: Functional programming style with function chaining via parameters.
   - To run code:
     ```bash
     go run wc07.go ../pride-and-prejudice.txt
     ```
   File: [`wc07.go`](./wc07.go)
   
### 8. **Monadic Style**
   - This implementation uses a monadic abstraction, where values are wrapped, bind to functions to form a sequence, and then unwrapped to examine the final result. The problem is solved as a series of bind functions that operate on the wrapped values, with unwrapping done at the end.
   - **Characteristics**: Functional programming style, monadic abstraction for wrapping and binding values.
   - To run code:
     ```bash
     go run wc08.go ../pride-and-prejudice.txt
     ```
   File: [`wc08.go`](./wc08.go)
