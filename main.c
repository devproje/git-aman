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

	struct profile_config conf = {
		.name = "Project_IO",
		.email = "me@projecttl.net"
	};

	struct profile_credential cred = {
		.proto = "https",
		.server = "github.com",
		.username = "devproje",
		.secret = "test"
	};

	profile prof = {
		.display = "Default",
		.config = &conf,
		.credential = &cred
	};

	add_profile(&prof);
	profile *data = get_profile(0);
	
	printf("BEFORE: %d\n", get_prof_len());
	printf("%s\n", data->display);

	del_profile(0);
	printf("AFTER: %d\n", get_prof_len());

	destroy();
	return 0;
}

