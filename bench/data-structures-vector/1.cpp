#include <iostream>
#include <vector>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	vector<int> vec;

	int n = 100000000;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	for (int i = 0; i < n; i++)
	{
		vec.push_back(42);
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
