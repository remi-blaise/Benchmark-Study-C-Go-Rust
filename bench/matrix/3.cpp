#include <iostream>
#include <ctime>
#include <cstdlib>

using namespace std;

int main(int argc, char const *argv[])
{
	srand(time(NULL));

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

	int **a = (int **)malloc(M * sizeof(int *)); // Matrix A: M*N
	for (int i = 0; i < M; i++)
	{
		a[i] = (int *)malloc(N * sizeof(int));
		for (int j = 0; j < N; j++)
		{
			a[i][j] = rand();
		}
	}

	int **b = (int **)malloc(N * sizeof(int *)); // Matrix A: N*M
	for (int i = 0; i < N; i++)
	{
		b[i] = (int *)malloc(M * sizeof(int));
		for (int j = 0; j < M; j++)
		{
			b[i][j] = rand();
		}
	}

	// Multiply the matrix
	int **prod = (int **)malloc(M * sizeof(int *)); // Product: M*M
	for (int i = 0; i < M; ++i)
	{
		prod[i] = (int *)malloc(M * sizeof(int));
		for (int j = 0; j < M; ++j)
		{
			prod[i][j] = 0;
		}
	}
	for (int j = 0; j < M; j++)
	{
		for (int i = 0; i < M; i++)
		{
			for (int k = 0; k < N; k++)
			{
				prod[i][j] += a[i][k] * b[k][j];
			}
		}
	}

	// Free the memory
	for (int i = 0; i < M; i++)
	{
		delete[] a[i];
	}
	delete[] a;

	for (int i = 0; i < M; i++)
	{
		delete[] b[i];
	}
	delete[] b;

	return 0;
}
