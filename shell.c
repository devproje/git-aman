#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "prompt.h"
#include "profile.h"

#define PREFIX 	      "? or 'yes':~$ "
#define DEFAULT_SIZE  50

#define NOT_AVAILABLE 2
#define CONFIRMED 3

void help();
int cli_parser(char *cmd, profile *prof);

int command(char *buf, profile *prof)
{
	int ch, extra;
	printf("%s", PREFIX);
	fflush(stdout);

	if (fgets(buf, DEFAULT_SIZE, stdin) == NULL)
		return NO_INPUT;

	if (strcmp(buf, "\n") == 0)
		return NO_INPUT;

	if (buf[strlen(buf) - 1] != '\n') {
		extra = 0;
		while ((ch = getchar()) != '\n' && (ch != EOF))
			extra = 1;

		return (extra == 1) ? TOO_LONG : OK;
	}

	buf[strlen(buf) - 1] = '\0';

	int res = cli_parser(buf, prof);
	if (res == 1) {
		return CONFIRMED;
	}

	return OK;
}

int cli_parser(char *cmd, profile *prof)
{
	if (strcmp(cmd, "display") == 0) {
		get_display_name(prof);
		return 0;
	}

	if (strcmp(cmd, "user.name") == 0) {
		get_git_name(prof);
		return 0;
	}

	if (strcmp(cmd, "user.email") == 0) {
		get_git_email(prof);
		return 0;
	}

	if (strcmp(cmd, "cred.proto") == 0) {
		get_server_proto(prof);
		return 0;
	}

	if (strcmp(cmd, "cred.server") == 0) {
		get_server_url(prof);
		return 0;
	}

	if (strcmp(cmd, "cred.username") == 0) {
		get_username(prof);
		return 0;
	}

	if (strcmp(cmd, "cred.secret") == 0) {
		// TODO: create secret input prompt
		return 0;
	}

	if (strcmp(cmd, "?") == 0) {
		help();
		return 0;
	}

	if (strcmp(cmd, "quit") == 0) {
		exit(0);
	}

	if (strcmp(cmd, "yes") == 0) {
		
		return 1;
	}

	printf("ERROR: not available command: %s\n", cmd);
	printf("If you want to see all commands, please type '?'.\n");
	
	return 0;
}

void help()
{
	printf("Available Commands:\n");
	printf("Config: user.name, user.email\n");
	printf("Credential: cred.proto, cred.server, cred.username, cred.secret\n");
	printf("misc: ?, quit, yes\n");
}

