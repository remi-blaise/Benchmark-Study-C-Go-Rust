#include <iostream>
#include <algorithm>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	int S = 25000000;
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

	auto start = chrono::steady_clock::now();

	sort(arr, arr + S);

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
