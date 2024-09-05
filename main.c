#include <stdio.h>
#include <stdlib.h>
#include <error.h>

#include "profile.h"

#define VERSION "1.0.0-alpha.3"

int main(int argc, char **argv)
{
	printf("git-aman %s\n\n", VERSION);
	printf("checking git system...\n");
	int res = system("git --version");
	if (res != 0) {
		perror("git not installed");
		return res;
	}

	return 0;
}

