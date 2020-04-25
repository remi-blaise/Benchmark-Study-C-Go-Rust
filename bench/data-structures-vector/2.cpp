#include <iostream>
#include <vector>
#include <ctime>
#include <cstdlib>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	vector<int> vec;

	int S = 100000000;
	int n = 10000000;
	if (argc >= 2)
	{
		S = atoi(argv[1]);

		if (argc == 3)
		{
			n = atoi(argv[2]);
		}
	}

	for (int i = 0; i < S; i++)
	{
		vec.push_back(42);
	}

	auto start = chrono::steady_clock::now();

	int sum = 0;

	for (int i = 0; i < n; i++)
	{
		sum += vec.at(rand() % S);
		S -= 1;
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
