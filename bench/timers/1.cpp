#include <iostream>
#include <chrono>
#include <unistd.h>

using namespace std;

int main(int argc, char const *argv[])
{
    int n = 1000000000; // 1 second
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

    auto start = chrono::steady_clock::now();
    auto end = chrono::steady_clock::now();
    long long int counter = 0;
    
    while (chrono::duration_cast<chrono::nanoseconds>(end - start).count() < n) {
        end = chrono::steady_clock::now();
        counter++;
    }

    cout << counter << endl;

    return 0;
}
