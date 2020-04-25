#include <iostream>
#include <sstream>
#include <string>
#include "cereal/archives/xml.hpp"
#include <chrono>

using namespace std;

struct Data
{
    long int a;
    float b;
    string c;

    // This method lets cereal know which data members to serialize
    template <class Archive>
    void serialize(Archive &archive)
    {
        archive(a, b); // serialize things by passing them to the archive
    }
};

int main(int argc, char const *argv[])
{
    int n = 100000;
    if (argc == 2)
    {
        n = atoi(argv[1]);
    }

    auto start = chrono::steady_clock::now();

    for (int i = 0; i < n; i++)
    {
        // Encode
        stringstream ss;
        {
            cereal::XMLOutputArchive oarchive(ss);

            Data myData = {13273828327, 382283.537749, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet luctus leo sed imperdiet. Morbi ut dolor eu arcu pretium bibendum. Donec eleifend arcu sit amet sodales ultrices. Nam quis diam vel mi hendrerit egestas quis in velit. Aliquam non vulputate magna. Cras et magna bibendum, facilisis magna et, rhoncus."};
            oarchive(myData);
        }
        // Decode
        {
            cereal::XMLInputArchive iarchive(ss);

            Data myData;
            iarchive(myData);
        }
    }

    cout << chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now() - start).count() << endl;

    return 0;
}
