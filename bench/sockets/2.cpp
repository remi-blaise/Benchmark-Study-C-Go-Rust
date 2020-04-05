#include <iostream>		// cout
#include <string.h>		// strlen
#include <string>		// string
#include <sys/socket.h> // socket
#include <arpa/inet.h>	// inet_addr
#include <netdb.h>		// hostent

using namespace std;

int main(int argc, char const *argv[])
{
	int n = 10000; // Number of requests
	if (argc == 2)
	{
		n = atoi(argv[1]);
	}

	struct sockaddr_in server;
	server.sin_addr.s_addr = inet_addr("127.0.0.1");
	server.sin_family = AF_INET;
	server.sin_port = htons(8080);

	// Create socket
	int sock = socket(AF_INET, SOCK_STREAM, 0);
	if (sock < 0)
	{
		cout << "Could not create socket" << endl;
	}

	// Connect to remote server
	if (connect(sock, (struct sockaddr *)&server, sizeof(server)) < 0)
	{
		perror("Connect failed. Error");
		return 1;
	}

	for (int i = 0; i < n; i++)
	{

		// Send message
		string message = "Hello from client";
		if (send(sock, message.c_str(), strlen(message.c_str()), 0) < 0)
		{
			perror("Send failed. Error");
			return 1;
		}

		// Receive reply
		char buffer[17];
		string reply;
		if (recv(sock, buffer, sizeof(buffer), 0) < 0)
		{
			perror("recv failed");
			perror("Error");
			return 1;
		}
		reply = buffer;
		cout << "Reply: " << reply << endl;
	}

	return 0;
}
