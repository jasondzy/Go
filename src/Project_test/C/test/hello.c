// hello.c

#include <stdio.h>

void pp(char* s);

char* ss = "string test";
char buf[10] = {'a','r','r','a','y'};

struct info {
    char* name;
    int age;
};

void SayHello(char *s)
{
    pp(s);
}