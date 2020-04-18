#include <iostream>
#include <unistd.h>

using namespace std;

int main(int argc, char const *argv[])
{
	int n = 1000; // Number of processes
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	for (int i = 0; i < n; i++)
	{
		fork();
	}

	return 0;
}
