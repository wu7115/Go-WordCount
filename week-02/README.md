## Implementations

### 2. **No Abstractions, No Library Functions, Slices Only**
   - This approach prohibits the use of abstractions (no custom functions) and avoids Go's library functions and any advanced data structure. It only utilizes basic functions and slices for data manipulation.
   - **Characteristics**: Simplistic, avoids abstraction layers, restricted use of features.
   - To run code
     ```bash
      go run wc01.go pride-and-prejudice.txt stop_words.txt
   
   File: [`wc02.go`](./wc02.go)

### 3. **Procedural Abstraction**
   - The solution decomposes the problem into smaller procedures, each handling a specific task. The main logic executes these procedures sequentially, solving the problem step by step.
   - **Characteristics**: Modular, structured through procedural decomposition, more organized than the basic approach.
   - To run code
     ```bash
      go run wc01.go pride-and-prejudice.txt stop_words.txt
   
   File: [`wc03.go`](./wc03.go)

### 4. **Functional Abstraction**
   - This implementation uses functional programming principles, decomposing the problem into mathematical functions. It solves the problem as a pipeline of function applications.
   - **Characteristics**: Focus on immutability, function composition, and avoiding side effects.
   - To run code
     ```bash
      go run wc01.go pride-and-prejudice.txt
   
   File: [`wc04.go`](./wc04.go)

### 5. **Minimal Lines of Code**
   - In this version, the goal is to solve the word count problem in as few lines of code as possible while maintaining functionality.
   - **Characteristics**: Compact, concise, and focused on brevity.
   - To run code
     ```bash
      go run wc01.go pride-and-prejudice.txt
   
   File: [`wc05.go`](./wc05.go)
