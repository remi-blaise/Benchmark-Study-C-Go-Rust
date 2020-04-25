#include <iostream>
#include <vector>
#include <algorithm>
#include <ctime>
#include <cstdlib>

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

	sort(vec.begin(), vec.end());

	return 0;
}
