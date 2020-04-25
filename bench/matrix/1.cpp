#include <iostream>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	int M = 1000; // Rows
	int N = 1000; // Columns
	if (argc >= 2)
	{
		M = atoi(argv[1]);
		if (argc == 3)
		{
			M = atoi(argv[1]);
		}
	}

	auto start = chrono::steady_clock::now();

	int **mat = (int **)malloc(M * sizeof(int *));
	for (int i = 0; i < M; i++)
	{
		mat[i] = (int *)malloc(N * sizeof(int));
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	// Free the memory
	for (int i = 0; i < M; i++)
	{
		delete[] mat[i];
	}
	delete[] mat;

	return 0;
}
