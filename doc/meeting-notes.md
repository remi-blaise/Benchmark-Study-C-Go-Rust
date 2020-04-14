# Meeting 2 (03/02)

Notes of the meeting:
sockets, Iostat, trie trees used for databases, threads, coroutines, processes, asynchronous, matrix ijk algo, crypto focus on hash fn, serialization just one lib, cereal or msgpack, timers precision and time read
iteration, recursion: bfs dfs or find a other algo
Eg. Implement f: x -> f(x-1)
Can also do with matrices
We need to work on strings
For next time: write spec of the project with system environment, program we want to write, metrics and everything

We can write our programs or look for already made benchmarks
We need to be comprehensive

Talk about memory safety features without forgetting it is for a high perf lab

# Meeting 3 (03/09)

We wrote part of this doc.

For next week:
- [ ] finish this doc
- [x] merge points 3 and 5
- [x] swap 4 and 6
- [x] implement 1 & 2
- [x] implement 7 Iteration vs recursion

# Meeting 4 (03/16)

Next week (03/23):

- [x] Compiler optimization
    - Go: https://github.com/golang/go/wiki/CompilerOptimizations
- [x] Add new test for benchmark iteration: Fibonacci (compare compiler optimization)
- [x] Do as much as Data structures and Sorting as possible (4)
- [x] Implement string benchmark (7)
- [x] Look what we need for Matrix (5)
- [ ] Set up measurement commands
- [ ] Reading about concurrency

# Meeting 5 (03/24)

Next week (03/30):

- [x] Matrix (5)
- [x] Concurrency (6):
    - https://www.geeksforgeeks.org/golang-goroutine-vs-thread/
    - https://medium.com/a-journey-with-go/go-goroutine-os-thread-and-cpu-management-2f5a5eaf518a
    - https://golangbot.com/goroutines/
- [x] Hashing (9)
- [x] Complex numbers (12)

# Meeting 6 (04/01)

What Remi did this week on Rust:
- Compiler optimization: https://doc.rust-lang.org/cargo/reference/profiles.html#release
- Memory (2)
- Fibonacci (3)

Next week (04/06):

- [x] Deeper research on concurrency in Go and Rust (threads, processes)
    - https://yourbasic.org/golang/goroutines-explained/
- [x] Search about using C library in Go
    - https://golang.org/cmd/cgo/
- [ ] Rust: Benchmarks 4, 5, 6, 7, 9 and 12
- [ ] C++/Go: Data serialization (8)
- [x] C++/Go: Sockets (10)
- [ ] C++/Go: Timers (11)
    - https://golang.org/pkg/time/
- [x] Set up measurement commands

# Meeting 7 (04/07)

Next week (04/13):

- [ ] Rust only: Benchmarks 6, 9 and 10
- [ ] Data serialization (8)
- [ ] Timers (11)
- [ ] Using C libraries (13)