#include <iostream>
#include <complex>
#include <ctime>
#include <cstdlib>
#include <chrono>

#define X 99999999.0

using namespace std;

float rand_float()
{
	return static_cast<float>(rand()) / (static_cast<float>(RAND_MAX / X));
}

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	int n = 50000000; // Number of operations
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	auto start = chrono::steady_clock::now();

	for (int i = 0; i < n; i++)
	{
		complex<float> complex1(rand_float(), rand_float());
		complex<float> complex2(rand_float(), rand_float());
		complex<float> mult = complex1 * complex2;
		complex<float> div = complex1 / complex2;
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
