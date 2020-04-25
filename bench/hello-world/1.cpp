#include <iostream>
#include <chrono>

using namespace std;

int main()
{
    auto start = chrono::steady_clock::now();

	cout << "Hello, world!" << endl;

    cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
