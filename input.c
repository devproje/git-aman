#include <stdio.h>
#include <string.h>
#include <termios.h>

#include "input.h"
#include "shell.h"
#include "prompt.h"
#include "profile.h"

void profile_checker(profile *prof)
{
	printf("Profile: %s\n", prof->display);
	
	printf("Config:\n");
	printf("\tuser.name: %s\n", prof->config.name);
	printf("\tuser.email: %s\n", prof->config.email);
	
	printf("Credential:\n");
	printf("\tcredential.proto: %s\n", prof->credential.proto);
	printf("\tcredential.server: %s\n", prof->credential.server);
	printf("\tcredential.username: %s\n", prof->credential.username);
	printf("\tcredential.secret: %s\n", prof->credential.secret);
}

profile input_info()
{
	int confirm = 0;
	profile prof;
	get_display_name(&prof);
	get_git_name(&prof);
	get_git_email(&prof);
	get_server_proto(&prof);
	get_server_url(&prof);
	get_username(&prof);

	do {
		char cmd[50];
		profile_checker(&prof);
		int res = command(cmd, &prof);
		if (res == 3) {
			confirm = 1;
		}
	} while (confirm == 0);

	return prof;
}

profile *edit_info()
{
	// TODO: create edit info function
}

profile *select_prof()
{
	// TODO: create select profile function
}

