#include <iostream>		// cout
#include <string.h>		// strlen
#include <string>		// string
#include <sys/socket.h> // socket
#include <arpa/inet.h>	// inet_addr
#include <netdb.h>		// hostent
#include <chrono>

using namespace std;

int main(int argc, char const *argv[])
{
	struct sockaddr_in server;
	server.sin_addr.s_addr = inet_addr("127.0.0.1");
	server.sin_family = AF_INET;
	server.sin_port = htons(8080);

	// Create socket
	int server_fd = socket(AF_INET, SOCK_STREAM, 0);
	if (server_fd < 0)
	{
		cout << "Could not create socket" << endl;
		return 1;
	}

	// Attach socket to the port
	if (::bind(server_fd, (struct sockaddr *)&server, sizeof(server)) < 0)
	{
		perror("bind failed");
		return 1;
	}

	// Listen mode
	if (listen(server_fd, 3) < 0)
	{
		perror("listen");
		return 1;
	}

	// Accept connections
	int servlen = sizeof(server);
	int new_socket = accept(server_fd, (struct sockaddr *)&server, (socklen_t *)&servlen);
	if (new_socket < 0)
	{
		perror("accept");
		return 1;
	}

	auto start = chrono::steady_clock::now();

	while (true)
	{
		// Read
		char buffer[17];
		string reply;
		if (recv(new_socket, buffer, sizeof(buffer), 0) < 0)
		{
			perror("recv failed");
			return 1;
		}
		reply = buffer;
		// cout << "Message received: " << reply << endl;

		// Send message
		string message = "Hello from server";
		if (send(new_socket, message.c_str(), strlen(message.c_str()), 0) < 0)
		{
			perror("Send failed. Error");
			return 1;
		}
	}

	cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

	return 0;
}
