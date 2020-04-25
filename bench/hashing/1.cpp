#include <iostream>
#include <unordered_map>
#include <ctime>
#include <cstdlib>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
    srand(time(NULL));

    unordered_map<int, int> map;

    int n = 10000000;
    if (argc == 2)
    {
        n = atoi(argv[1]);
    }

    auto start = chrono::steady_clock::now();

    for (int i = 0; i < n; i++)
    {
        map.insert({ rand(), 0 });
    }

    cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

    return 0;
}
