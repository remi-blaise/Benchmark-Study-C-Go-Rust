#include <iostream>
#include <chrono>

using namespace std;

// Recursion function
int fib(unsigned int n)
{
	if (n <= 1)
	{
		return n;
	}

	return fib(n - 1) + fib(n - 2);
}

int main(int argc, char const *argv[])
{
	unsigned int n = 40;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	n = fib(n);

	cout << n << endl;

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
