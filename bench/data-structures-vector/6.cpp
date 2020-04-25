#include <iostream>
#include <vector>
#include <algorithm>
#include <ctime>
#include <cstdlib>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	int S = 100000000;
	if (argc == 2)
	{
		S = atoi(argv[1]);
	}

	vector<int> vec;

	// Random sorted array
	for (int i = 0; i < S; i++)
	{
		vec.push_back(rand());
	}

	auto start = chrono::steady_clock::now();

	sort(vec.begin(), vec.end());

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
