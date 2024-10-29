## Implementations

### 9. **Smalltalk-Style Object Message Passing**
   - This version decomposes the problem into "things" (or objects) that make sense for the problem domain. Each object acts as a capsule of data, exposing only a single procedure: the ability to receive and dispatch messages. Messages can also be sent from one capsule to another, creating a communication chain between these objects. This approach mirrors the "Smalltalk" style of object-oriented programming, emphasizing encapsulation and message-passing.
   - **Characteristics**: Object-oriented style, message dispatch between encapsulated objects.
   - To run code:
     ```bash
     go run wc09.go ../pride-and-prejudice.txt ../stop_words.txt
     ```
   File: [`wc09.go`](./wc09.go)

### 10. **JavaScript-Style Map-Based Objects**
   - In this implementation, each "thing" in the program is represented as a map, where keys correspond to properties or values, some of which are procedures/functions. This approach allows for a flexible structure, where objects can be dynamically composed and manipulated in a manner similar to JavaScript objects.
   - **Characteristics**: Map-based structure with values as procedures, inspired by JavaScript objects.
   - To run code:
     ```bash
     go run wc10.go ../pride-and-prejudice.txt ../stop_words.txt
     ```
   File: [`wc10.go`](./wc10.go)

### 11. **Callback-Driven Abstraction with Entity Registration**
   - This version decomposes the larger problem into abstract entities that do not directly execute actions. Instead, entities provide interfaces allowing other entities to register callbacks, enabling interaction without direct control. During computation, registered callbacks are invoked at specific points, creating an event-driven structure. An additional function is included to print all words containing the letter "z."
   - **Characteristics**: Abstraction with callback registration, event-driven approach, additional functionality for filtering words with "z."
   - To run code:
     ```bash
     go run wc11.go ../pride-and-prejudice.txt ../stop_words.txt
     ```
   File: [`wc11.go`](./wc11.go)
