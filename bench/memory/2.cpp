#include <iostream>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	int S = 1000000000; // 1 GB
	int tmp;
	if (argc == 2)
	{
		S = atoi(argv[1]);
	}

	S = S / sizeof(int);

	int *p = (int *)malloc(S * sizeof(int));

	auto start = chrono::steady_clock::now();

	for (int i = 0; i < S; i++)
	{
		tmp = p[i];
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	free(p);

	return 0;
}
