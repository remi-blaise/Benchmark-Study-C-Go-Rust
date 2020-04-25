#include <iostream>
#include <fstream>
#include <string>
#include "sha512.hh"
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
    // Read passwords
    string line;
    ifstream myfile("asset/1000000passwords");
    if (myfile.is_open())
    {
        auto start = chrono::steady_clock::now();

        while (getline(myfile, line))
        {
            sw::sha512::calculate(line);
        }
        cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

        myfile.close();
    }
    else
    {
        cout << "Unable to open file";
    }

    return 0;
}
