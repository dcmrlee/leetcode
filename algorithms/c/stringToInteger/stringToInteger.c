#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

#define INT_MIN     (-2147483647 - 1)
#define INT_MAX      2147483647

int atoi(const char *str) {
    if (str==NULL || *str=='\0'){
        return 0;
    }
    
    int ret=0;
    
    for(;isspace(*str); str++);
    
    int neg=1;
    if (*str=='-' || *str=='+') {
        if (*str=='-') {
	    neg=-1;
	}
        str++;
    }
    
    for(; isdigit(*str); str++) {
        int digit = (*str-'0');
        if(neg==-1){
            if( -ret < (INT_MIN + digit)/10 ) {
                return INT_MIN;
            }
        }else{
            if( ret > (INT_MAX - digit) /10 ) {
                return INT_MAX;
            }
        }

        ret = 10*ret + digit ;
    }
    
    return neg*ret;
}


int main()
{
    printf("\"%s\" = %d\n", "123", atoi("123"));
    printf("\"%s\" = %d\n", "   123", atoi("    123"));
    printf("\"%s\" = %d\n", "+123", atoi("+123"));
    printf("\"%s\" = %d\n", " -123", atoi(" -123"));
    printf("\"%s\" = %d\n", "123ABC", atoi("123ABC"));
    printf("\"%s\" = %d\n", " abc123ABC", atoi(" abc123ABC"));
    printf("\"%s\" = %d\n", "2147483647", atoi("2147483647"));
    printf("\"%s\" = %d\n", "-2147483648", atoi("-2147483648"));
    printf("\"%s\" = %d\n", "2147483648", atoi("2147483648"));
    printf("\"%s\" = %d\n", "-2147483649", atoi("-2147483649"));
    printf("\"%s\" = %d\n", "+-2", atoi("+-2"));
    return 0;
}
