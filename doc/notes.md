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
[x] Make a build-all command
[x] Measure bin sizes and output in a csv
[x] Measure execution time and output in a csv
[x] Measure monitoring data with top
[x] Configure all scripts to have a fair duration
[x] Measure critical time
[ ] Make a script that install, build all and execute

Optional:
[ ] Coder 4.cpp
[ ] Compiler sans optimisation hello-world and memory et faire les mesures
[ ] Retapper memory rust pour utiiser des tableaux ?
