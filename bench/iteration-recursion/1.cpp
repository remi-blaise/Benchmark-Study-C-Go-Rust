#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	int N = 1000;
	if (argc == 2)
	{
		N = atoi(argv[1]);
	}

	int n = 0;

	for (int i = 0; i < N; i++)
	{
		n++;
	}

	return 0;
}
