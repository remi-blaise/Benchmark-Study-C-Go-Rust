#include <iostream>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	int n = 40;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	int x = 0, y = 1, z = 0;
	for (int i = 0; i < n; i++) {
		z = x + y;
		x = y;
		y = z;
	}

	cout << x << endl;

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
