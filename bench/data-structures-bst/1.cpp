#include <iostream>
#include <ctime>
#include <cstdlib>
#include "cpp-btree/btree_set.h"
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
    srand(time(NULL));

	btree::btree_set<int> btree;

	int n = 1000000;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	for (int i = 0; i < n; i++)
	{
		btree.insert(rand());
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

    return 0;
}
