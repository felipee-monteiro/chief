#include <stdio.h>
#include <stdlib.h>
#include <libspeedy.h>

void main() {
   char* result = Connect("sa", "Epilefac57#$!$24042002", "localhost", "sigma", 1433);
   printf("%s", result);
}
