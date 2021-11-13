%module Swpack
%{
int wsum(int a, int b) {
    return a + b;
}
%}

extern int wsum(int a, int b);