#include <iostream>
#include <string>

using namespace std;

int main(int argc, char const *argv[])
{
	int n = 1000000;
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	string str = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam id ante non enim sodales interdum. In imperdiet varius dui tincidunt pulvinar. Aenean enim lorem, posuere et nibh eu, scelerisque facilisis turpis. Cras efficitur aliquam pellentesque. Donec eleifend velit vel magna pellentesque, ut malesuada mauris tempor. Pellentesque quis augue dapibus, pulvinar eros et, posuere eros. Praesent suscipit lacinia lorem, quis feugiat elit pulvinar sed. Nullam purus est, porta in ante mattis, tempus pellentesque leo. Cras quam nisi, pellentesque et est sed, vestibulum faucibus felis. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.";

	for (int i = 0; i < n; i++)
	{
		hash<string> h;
		h(str);
	}

	return 0;
}