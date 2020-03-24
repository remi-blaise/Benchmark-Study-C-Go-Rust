#include <iostream>

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
	unsigned int n = 100;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	n = fib(n);

	cout << n;

	return 0;
}
