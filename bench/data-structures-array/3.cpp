#include <iostream>
#include <algorithm>
#include <ctime>
#include <cstdlib>

using namespace std;

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	int S = 5000000;
	if (argc == 2)
	{
		S = atoi(argv[1]);
	}

	int *arr = (int *)malloc(S * sizeof(int));

	// Random array
	for (int i = 0; i < S; i++)
	{
		arr[i] = rand();
	}

	sort(arr, arr + S);

	return 0;
}
