#include <iostream>
#include "demo342//Demo342.h"
using namespace std;
void demo342()
{
    int number = 2563;
    bool isFour;

    Demo342 demo342 ;
    isFour = Demo342::isPowerOfFour(number);


    cout << "�Ƿ���4���� " << endl;
    cout << isFour << endl;
}

int main() {

    demo342();
    return 0;
}



