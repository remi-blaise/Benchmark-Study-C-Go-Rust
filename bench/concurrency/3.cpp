#include <iostream>
#include <unistd.h>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	int n = 10000; // Number of processes
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	pid_t pid = getpid();

	for (int i = 0; i < n; i++)
	{
		if (pid == getpid())
		{
			fork();
		}
		else
		{
			exit(0);
		}
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
