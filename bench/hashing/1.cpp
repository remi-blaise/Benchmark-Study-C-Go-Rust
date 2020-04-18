#include <iostream>
#include <unordered_map>
#include <ctime>
#include <cstdlib>

using namespace std;

int main(int argc, char const *argv[])
{
    srand(time(NULL));

    unordered_map<int, int> map;

    int n = 1000000;
    if (argc == 2)
    {
        n = atoi(argv[1]);
    }

    for (int i = 0; i < n; i++)
    {
        map.insert({ rand(), 0 });
    }

    return 0;
}
