#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	int S = 1000000;
	if (argc == 2)
	{
		S = atoi(argv[1]);
	}

	int *arr = (int *)malloc(S * sizeof(int));

	// Already sorted array
	for (int i = 0; i < S; i++)
	{
		arr[i] = i;
	}

	sort(arr, arr + S);

	return 0;
}
