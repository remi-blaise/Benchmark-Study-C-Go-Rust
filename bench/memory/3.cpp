#include <iostream>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	int S = 1000000; // 1 MB
	int N = 1000;
	if (argc >= 2)
	{
		S = atoi(argv[1]);
		if (argc == 3)
		{
			N = atoi(argv[2]);
		}
	}

	S = S / sizeof(int);

	auto start = chrono::steady_clock::now();

	for (int i = 0; i < N; i++)
	{
		int *p = (int *)malloc(S * sizeof(int));
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
