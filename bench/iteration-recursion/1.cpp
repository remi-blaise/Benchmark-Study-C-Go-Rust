#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	unsigned long long int n = 1000;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	while (n > 0)
	{
		n--;
	}

	cout << n << endl;

	return 0;
}
