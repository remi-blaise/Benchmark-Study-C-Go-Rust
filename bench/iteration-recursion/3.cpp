#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	unsigned int n = 100;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	int x = 0, y = 1, z = 0;
	for (int i = 0; i < n; i++) {
		z = x + y;
		x = y;
		y = z;
	}

	cout << x << endl;

	return 0;
}
