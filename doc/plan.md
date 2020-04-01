# Benchmarks

1. Hello World: binary and memory size **(DONE)**
2. Memory management: allocating memory, passing to a function, writing, reading, delete **(DONE)**
3. Iteration vs recursion **(DONE)**
4. Data structure operations: creation, insert, random access, delete, lifecycle on vectors (arrays), trees, graphs, stacks, heaps, queues and Standard sorting algorithms  **(WIP)**
5. Matrix operations: creation, multiplication (stock, ijk, transpose) → FLOPS
6. Concurrency mechanisms: creation, execution, suppression of a thread, process, coroutine, goroutine; mutual exclusion mechanisms
7. String: split function and binary representation (ASCII, UTF-8) **(WIP)**
8. Data serialization
    - Rust & Go ?
    - C++: cereal, message pack, protobuf, …
9. Hashing
10. Sockets: ping pong
11. Timers precision and time read
12. Computation: complex numbers

# Measure Tools

We plan to use the `iostat-tool` command to get measurements as well as linux virtual files for memory size measure.

# System environment

## Rust

Language & compiler version: [1.41.1](https://blog.rust-lang.org/2020/02/27/Rust-1.41.1.html)

## Go

Language & compiler version: 1.14

## C++

Language version: [C++17](https://gcc.gnu.org/projects/cxx-status.html#cxx17)

Compiler: [GCC 8.4](https://gcc.gnu.org/gcc-8/)

# Benchmark plan

This document describes benchmarks to realize.

## Random source

To have a predictable (once-seeded) identical source of random numbers for all test cases, a random suite of number will be pregenerated and stored in a file. Test cases will use this file as the random source.

## Compiler optimization flags

We wish to use the common optimization options that developers use for all languages in order to get fair comparisons of what our tests would imply in real-world scripts.

For Rust, we are using the very same flags as the `release` profile does: https://doc.rust-lang.org/cargo/reference/profiles.html#release

## 1 Hello World

Motivation: See the overhead of the simplest program.

Program:
1. Print "Hello World".

Metrics:
- binary size
- memory size
- compilation time
- execution time

## 2 Memory management

Motivation: Compare memory allocation and desallocation.

Metrics:
- memory size
- execution time

### 2.1 Allocating and writing 1GB

Args:
- `S = 1_000_000_000`: Size in byte

Program:
1. Allocate a 1GB space on the heap.
2. Write 1GB with value 42.
3. Free the space.

### 2.2 Reading 1GB

Args:
- `S = 1_000_000_000`: Size in byte

Program:
1. Allocate 1GB on the heap.
2. Read 1GB just by looping over and consecutively storing each of the values into a one byte variable.

### 2.3 Allocating 1000\*1MB

Args:
- `S = 1_000_000`: Size in byte
- `N = 1000`: Iterations

Program:
1. Allocate an array of 1000 spaces of 1MB on the heap.
2. Do a random access into each space to ensure compiler don't optimize by removing code.

## 3 Iteration vs recursion

Motivation: Compare perfomance between interative and recursive approach

### 3.1 Iteration

Program: Increment a variable using a for-loop

Metrics:
 - Execution time

### 3.2 Recursion

Program: Increment a variable using a recursive function

Metrics:
 - Execution time

## 4 Data structures operations

### 4.0 Array

#### 4.0.1 Sort already sorted

Args:
 - `S = 1_000_000` Size of the array

Program: Sort an array already sorted

#### 4.0.2 Sort reverse sorted

Args:
 - `S = 1_000_000` Size of the array

Program: Sort an array that is reverse sorted

#### 4.0.3 Sort random ordered

Args:
 - `S = 1_000_000` Size of the array

Program:
- Generate a random array
- Sort the array

Metrics:
 - Running time

### 4.1 Vector

#### 4.1.1 Insert in vector

Args:
- `N = 1_000_000`: Number of elements

Program:
1. Create an empty vector
2. Do N times: Insert 42 at the end

#### 4.1.2 Random access in vector

Args:
- `S = 1_000_000`: Size of the vector
- `N = 1_000_000`: Number of reads

Program:
1. Create a vector of size S (with values = 42)
2. Do N times: Select a random location and read (store in a one byte variable)

#### 4.1.3 Deletion in vector

Args:
- `S = 1_000_000`: Size of the vector
- `N = 1_000_000`: Number of deletions

Program:
1. Create a vector of size S (with values = 42)
2. Do N times: Select a random location and delete

#### 4.1.4 Sort already sorted

Args:
 - `S = 1000000` Size of the array

Program: Sort an array already sorted

#### 4.1.5 Sort reverse sorted

Args:
 - `S = 1000000` Size of the array

Program: Sort an array that is reverse sorted

#### 4.1.6 Sort random

Args:
 - `S = 1_000_000` Size of the array

Prerequisite:
- Generate an array filled with random integers in a csv file

Program:
- Read the file (name as first arg)
- Sort the array

Metrics:
 - Running time

### 4.2 Binary Search Tree (BST)

#### 4.2.1 Insert

Args:
- `N = 1_000_000`: Number of elements

Program:
1. Create an empty
2. Do N times: Insert a random number

#### 4.2.2 Search

Args:
- `S = 1_000_000`: Size of the Binary tree
- `N = 1_000_000`: Number of reads

Program:
1. Create of size S with random values
2. Do N times: Select a random value and search for it in the BST

#### 4.2.3 Deletion

Args:
- `S = 1_000_000`: Size of the Binary tree
- `N = 1_000_000`: Number of deletions

Program:
1. Create of size S with random values
2. Do N times: Select a random element and delete

### 4.3 Graph (undirected, unweighted)

#### 4.3.1 Create

Args:
- `V, E = 1_000, 1_000`: Number of vertices and edges
- `N = 1_000_000`: Number of reads

Program:
1. Create a graph with V vertices and E random edges

#### 4.3.2 Random access

Args:
- `S = 1_000_000`: Size
- `N = 1_000_000`: Number of reads

Program:
1. Load previous
2. Do N times: Select a random nodes and read children

#### 4.3.3 Deletion

Args:
- `S = 1_000_000`: Size
- `N = 1_000_000`: Number of deletions

Program:
1. Load previous
2. Do N times: Delete a random vertice with its edges

### 4.4 Trie tree

#### 4.4.1 Insert

Args:
- `N = 1_000_000`: Number of elements

Program:
1. Create an empty
2. Do N times: Insert a random 64-byte string

#### 4.4.2 Random access

Args:
- `S = 1_000_000`: Size
- `N = 1_000_000`: Number of reads

Program:
1. Load previous
2. Do N times: Read a random string (among the stored dictonary)

#### 4.4.3 Deletion

Args:
- `S = 1_000_000`: Size
- `N = 1_000_000`: Number of deletions

Program:
1. Load previous
2. Do N times: Delete a random string (among the stored dictonary)

## 5. Matrix operations

Motivation: Compare time to do matrix operations

### 5.1. Allocating a matrix

Args:
 - `M` Number of rows
 - `N` Number of columns

Program: Allocate a matrix of size M\*N with random values

### 5.2. Multiplying matrix (rows)

Args:
 - `M` Number of rows
 - `N` Number of columns

Program: Multiply a matrix of size M\*N and a matrix of size N\*M with a for loop i, j, k

### 5.3. Multiplying matrix (columns)

Args:
 - `M` Number of rows
 - `N` Number of columns

Program: Multiply a matrix of size M\*N and a matrix of size N\*M with a for loop j, i, k

## 6. Concurrency mechanisms

## 7. Strings
## 8. Data serialization
    - Rust & Go ?
    - C++: cereal, message pack, protobuf, …
## 9. Hashing

## 10. Sockets: ping pong

Motivation: Compare efficiency of socket communication
Program: Two programs communication with each other on the same machine via a socket (ping pong)
Metrics:
 - ?

## 11. Timers precision and time read
## 12. Computation: complex numbers

Motivation: Compare Go's standard complex type with complex numbers implementations in Rust and C++
Program: Perform basic operations on complex numbers:
 - Addition/substraction
 - Multiplication/division
 - Modulus
 - Argument
 - Trigonometric form

Metrics:
 - ?
