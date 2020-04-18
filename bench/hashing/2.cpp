#include <iostream>
#include <fstream>
#include <string>
#include "sha1.hh"

using namespace std;

int main(int argc, char const *argv[])
{
	// Read string
	string line;
	ifstream myfile("asset/lorem100mb");
	if (myfile.is_open())
	{
		getline(myfile, line);
		myfile.close();
	}
	else
	{
		cout << "Unable to open file";
	}

	sw::sha1::calculate(line);

	return 0;
}
