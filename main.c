#include <stdio.h>
#include <stdlib.h>
#include <error.h>

#include "profile.h"
#include "input.h"

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

	profile prof = input_info();
	printf("Display Name: %s\n", prof.display);
	printf("[CONFIG]\n");
	printf("\tuser.name %s\n", prof.config.name);
	printf("\tuser.email %s\n", prof.config.email);
	printf("[CREDENTIAL]\n");
	printf("\tcredential.proto %s\n", prof.credential.proto);
	printf("\tcredential.server %s\n", prof.credential.server);
	printf("\tcredential.username %s\n", prof.credential.username);
	printf("\tcredential.secret %s\n", prof.credential.secret);

	add_profile(&prof);

	destroy();
	return 0;
}

