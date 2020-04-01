#include <iostream>

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

	int **mat = (int **)malloc(M * sizeof(int *));
	for (int i = 0; i < M; i++)
	{
		mat[i] = (int *)malloc(N * sizeof(int));
	}

	// Free the memory
	for (int i = 0; i < M; i++)
	{
		delete[] mat[i];
	}
	delete[] mat;

	return 0;
}
