#include <iostream>
#include <vector>
#include <thread>
#include <mutex>

using namespace std;

int x = 0;
mutex x_mutex;

void subroutine()
{
    x_mutex.lock();
	x++;
    x_mutex.unlock();
}

int main(int argc, char const *argv[])
{
	int n = 10000; // Number of threads
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