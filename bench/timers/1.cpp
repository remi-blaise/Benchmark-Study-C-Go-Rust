#include <iostream>
#include <chrono>
#include <unistd.h>

using namespace std;

int main(int argc, char const *argv[])
{
    int n = 30000000; // 1 second
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

    auto start = chrono::steady_clock::now();

    for (int i = 0; i < n; ++i)
    {
        chrono::steady_clock::now();
    }

    cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

    return 0;
}
