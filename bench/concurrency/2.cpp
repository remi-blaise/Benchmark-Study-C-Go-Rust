#include <iostream>
#include <vector>
#include <thread>
#include <mutex>
#include <chrono>

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
	int n = 25000; // Number of threads
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	vector<thread> ths;
	for (int i = 0; i < n; i++)
	{
		ths.push_back(thread(&subroutine));
	}

	for (auto &th : ths)
	{
		th.join();
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
