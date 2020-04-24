#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(int argc, char const *argv[])
{
	int S = 1000000;
	if (argc == 2)
	{
		S = atoi(argv[1]);
	}

	vector<int> vec;

	// Reverse sorted array
	for (int i = 0; i < S; i++)
	{
		vec.push_back(S - i);
	}

	sort(vec.begin(), vec.end());

	return 0;
}
