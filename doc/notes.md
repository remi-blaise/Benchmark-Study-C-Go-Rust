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
[ ] Go : deplacer et completer les cas de array -> vector
[ ] BST: C++ et Go a remplacer par Google lib
[x] BST: Rust a remplacer par Set
[ ] Rust Matrix: a refaire a la main
[ ] C++ Fork: attention a ne pas instancier 2\*\*n forks
[ ] C++ & Go Concurrency: completer
[ ] String: generate asset/lorem100mb
[ ] Serialization: Implementer C++
[ ] Serialization: Go
[ ] Hashing: Implementer C++ & Go
[ ] Rust Hashing changement fichier et utiliser algo non secure
[ ] Sockets C et Go : mimic Rust
[ ] Timers C et Go : passer les ns en params
