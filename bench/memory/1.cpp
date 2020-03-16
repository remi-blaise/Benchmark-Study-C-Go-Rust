#include <iostream>
#include <ctime>

using namespace std;

int main(int argc, char const *argv[])
{
	srand(time(0));

	int S = 1000000000 / sizeof(int); // 1 GB

	int *p = (int *)malloc(S * sizeof(int));
	for (int i = 0; i < S; i++)
	{
		p[i] = 42;
	}

	cout << "Allocated!\n";

	int pos = rand() % S;
	cout << "Item at position " << pos << ": " << p[pos] << "\n";

	free(p);

	cout << "Deallocated!\n";

	return 0;
}
