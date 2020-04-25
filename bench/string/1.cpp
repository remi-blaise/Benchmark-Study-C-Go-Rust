#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>
#include <iterator>
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
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

	istringstream iss(line);
	vector<string> results(istream_iterator<string>{iss}, istream_iterator<string>());

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
