2 Memory:
- I think we shouldn't activate optimizations

6 Concurrency
- (no hidden lock)
- tracer un graphe temps en fonction nb de threads. Interesting because decreases without Goroutines but increases with thread.

Measurements:
Run `top -b > measures` during the execution, then
`cat measures | grep "chrome" | perl -pe 's/^\s+//' | tr -s " " | cut -f 6,7,9,10 -d " "`
to retrieve reserved memory, shared memory, %cpu and %ram
-> sleep for 2 secondes before and after

-------------------------

Reste a faire :
[x] Go : deplacer et completer les cas de array -> vector
[x] BST: C++ et Go a remplacer par Google lib
[x] C++ Fork: attention a ne pas instancier 2\*\*n forks
[x] C++ & Go Concurrency: completer
[x] Serialization: Implementer C++
[x] Serialization: Go
[x] Hashing: Implementer C++ & Go
[x] Sockets C et Go : mimic Rust
[x] Timers C et Go : passer les ns en params

[x] BST: Rust a remplacer par Set
[x] Rust Matrix: a refaire a la main
[x] Rust Hashing utiliser algo non secure
[x] String: generate asset/lorem100mb

-------------------------

Interrogations :
- Matrices:
    - je pense qu'on peut virer 1
    - usuellement les matrices ne sont pas des tableaux de tableaux, c'est peut-etre la raison pour laquelle ij ji ne change rien
    - 2.go prod[i][j] += a[i][k] * b[k][j] non-initialized ?
    - do we ignore overflow ?
    - je ne pense pas qu'echanger les variables puisse avoir un impact a cause de la symetrie
>>>>>>> Some Rust
