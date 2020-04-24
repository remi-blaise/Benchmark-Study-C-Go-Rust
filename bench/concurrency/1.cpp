#include <iostream>
#include <vector>
#include <thread>

using namespace std;

int x = 0;

void subroutine()
{
	x++;
}

int main(int argc, char const *argv[])
{
	int n = 25000; // Number of threads
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	vector<thread> ths;
	for (int i = 0; i < n; i++)
	{
		ths.push_back(thread(&subroutine));
	}

	for (auto &th : ths)
	{
		th.join();
	}

	return 0;
}
