#include <iostream>
#include <complex>
#include <ctime>
#include <cstdlib>

#define X 99999999.0

using namespace std;

float rand_float()
{
	return static_cast<float>(rand()) / (static_cast<float>(RAND_MAX / X));
}

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	int n = 30000000; // Number of operations
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	for (int i = 0; i < n; i++)
	{
		complex<float> complex1(rand_float(), rand_float());
		complex<float> complex2(rand_float(), rand_float());
		complex<float> mult = complex1 * complex2;
		complex<float> div = complex1 / complex2;
	}

	return 0;
}
