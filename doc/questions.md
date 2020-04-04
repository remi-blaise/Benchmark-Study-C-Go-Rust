2 Memory:
- I think we shouldn't activate optimizations

3 Data structures:
- generate the random value files
- BTS: pregenerate a shuffled consecutive list of numbers between 0 and N

5 Matrix
- We notice no differences ixj vs jxi, is it normal?

6 Concurrency
- (no hidden lock)
- tracer un graphe temps en fonction nb de threads. Interesting because decreases without Goroutines but increases with thread.

Measurements:
Run `top -b > measures` during the execution, then
`cat measures | grep "chrome" | perl -pe 's/^\s+//' | tr -s " " | cut -f 6,7,9,10 -d " "`
to retrieve reserved memory, shared memory, %cpu and %ram
