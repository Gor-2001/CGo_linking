#include "string_utils.h"

int str_length(const char* s) {
    int len = 0;
    while (s[len] != '\0') len++;
    return len;
}
