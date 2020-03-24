#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	string s = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin sit amet euismod arcu. Duis ligula ex, feugiat mattis porta quis, maximus vitae ex. Praesent dapibus neque arcu, a consectetur sapien malesuada sit amet. Proin tempor ligula sed sem maximus, id.";
	string delimiter = " ";

	size_t pos = 0;
	string token;
	while ((pos = s.find(delimiter)) != string::npos)
	{
		token = s.substr(0, pos);
		cout << token << endl;
		s.erase(0, pos + delimiter.length());
	}
	cout << s << endl;

	return 0;
}
