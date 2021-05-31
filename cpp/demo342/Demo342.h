//
// Created by Administrator on 2021/5/31.
//

#ifndef CPP_DEMO342_H
#define CPP_DEMO342_H


class Demo342 {
public:
    static bool isPowerOfFour(int n) {
        return n > 0 && (n & (n - 1)) == 0 && n % 3 == 1;
    }
};


#endif //CPP_DEMO342_H
