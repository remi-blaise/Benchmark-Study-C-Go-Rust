#include <iostream>

using namespace std;

// Recursion function
int rec(unsigned long long int n)
{
	if (n <= 0)
	{
		return n;
	}

	return rec(n - 1);
}

int main(int argc, char const *argv[])
{
	unsigned long long int n = 1000;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	n = rec(n);

	cout << n << endl;

	return 0;
}
