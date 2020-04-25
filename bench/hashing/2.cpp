#include <iostream>
#include <fstream>
#include <string>
#include "sha1.hh"
#include <chrono>

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

	auto start = chrono::steady_clock::now();

	sw::sha1::calculate(line);

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
