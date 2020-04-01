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

	int n = 1000000; // Number of operations
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	for (int i = 0; i < n; i++)
	{
		complex<float> comp(rand_float(), rand_float());
		float a = abs(comp);
	}

	return 0;
}