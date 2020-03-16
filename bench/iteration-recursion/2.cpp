#include <iostream>

using namespace std;

// Recursion function
int rec(int n, int N)
{
	if (N <= 0)
	{
		return n;
	}

	return rec(n + 1, N - 1);
}

int main(int argc, char const *argv[])
{
	int N = 1000;
	if (argc == 2)
	{
		N = atoi(argv[1]);
	}

	int n = 0;

	rec(n, N);

	return 0;
}
