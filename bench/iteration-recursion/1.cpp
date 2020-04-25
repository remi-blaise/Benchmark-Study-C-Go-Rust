#include <iostream>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	unsigned long long int n = 1000;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	while (n > 0)
	{
		n--;
	}

	cout << n << endl;

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
