#include <iostream>
#include <ctime>
#include <cstdlib>
#include "cpp-btree/btree_set.h"

using namespace std;

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	btree::btree_set<int> btree;

	int S = 1000000;
	int n = 1000000;
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
		btree.insert(rand());
	}

	for (int i = 0; i < n; i++)
	{
		btree.erase(rand());
	}

	return 0;
}
