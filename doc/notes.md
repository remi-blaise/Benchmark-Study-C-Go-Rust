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

TODO:
[ ] Make a build-all command
[x] Measure bin sizes and output in a csv
[ ] Measure execution time and output in a csv

[ ] Make a script that install, build all and execute
