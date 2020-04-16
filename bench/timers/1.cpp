#include <iostream>
#include <chrono>
#include <unistd.h>

using namespace std;

int main(int argc, char const *argv[])
{
    auto start = chrono::steady_clock::now();
    auto end = chrono::steady_clock::now();
    long long int counter = 0;
    
    while (chrono::duration_cast<chrono::nanoseconds>(end - start).count() < 1000000000) {
        end = chrono::steady_clock::now();
        counter++;
    }

    cout << counter << endl;

    return 0;
}
