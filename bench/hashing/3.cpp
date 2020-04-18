#include <iostream>
#include <fstream>
#include <string>
#include "sha512.hh"

using namespace std;

int main(int argc, char const *argv[])
{
    // Read passwords
    string line;
    ifstream myfile("asset/1000000passwords");
    if (myfile.is_open())
    {
        while (getline(myfile, line))
        {
            sw::sha512::calculate(line);
        }
        myfile.close();
    }
    else
    {
        cout << "Unable to open file";
    }

    return 0;
}
