#include <stdio.h>
#include <string.h>

#include "prompt.h"

static int input_char(char *msg, char *buf, size_t len)
{
	int ch, extra;
	if (msg != NULL) {
		printf("%s", msg);
		fflush(stdout);
	}

	if (fgets(buf, len, stdin) == NULL)
		return NO_INPUT;

	if (strlen(buf) <= 1)
		return TOO_SHORT;

	if (buf[strlen(buf) - 1] != '\n') {
		extra = 0;
		while ((ch = getchar()) != '\n' && (ch != EOF))
			extra = 1;

		return (extra == 1) ? TOO_LONG : OK;
	}

	buf[strlen(buf) - 1] = '\0';
	return OK;
}

void print_err(char *tag, int err)
{
	if (err == NO_INPUT) {
		printf("ERROR: '%s' must not be null\n", tag);
	} else if (err == TOO_LONG) {
		printf("ERROR: '%s' is too long\n", tag);
	} else {
		printf("ERROR: '%s' is too short\n", tag);
	}
}

void get_display_name(profile *prof)
{
	int err, chk = 0;
	do {
		err = input_char("type your display name ~$ ", prof->display, 25);
		if (err != OK) {
			print_err("display name", err);
			continue;
		}

		chk = 1;
	} while (chk == 0);
}

void get_git_name(profile *prof)
{
	int err, chk = 0;
	do {
		err = input_char("type your git 'user.name' ~$ ", prof->config.name, 25);
		if (err != OK) {
			print_err("user.name", err);
			continue;
		}

		chk = 1;
	} while (chk == 0);
}

void get_git_email(profile *prof)
{
	int err, chk = 0;
	do {
		err = input_char("type your git 'user.email' ~$ ", prof->config.email, 50);
		if (err != OK) {
			print_err("user.email", err);
			continue;
		}

		chk = 1;
	} while (chk == 0);
}

int check_proto(char *proto)
{
	if (strcmp(proto, "git") == 0)
		return 1;

	if (strcmp(proto, "http") == 0)
		return 1;

	if (strcmp(proto, "https") == 0)
		return 1;

	return 0;
}

void get_server_proto(profile *prof)
{
	int err, chk = 0;
	do {
		err = input_char("type your git server proto\n[git, http, https ~]$ ",
				prof->credential.proto,
				6);
		if (err != OK) {
			print_err("proto", err);
			continue;
		}

		if (check_proto(prof->credential.proto) != 1) {
			printf("ERROR: invalid proto name\n");
			continue;
		}

		chk = 1;
	} while (chk == 0);
}

void get_server_url(profile *prof)
{
	int err, chk = 0;
	do {
		err = input_char("type your git server url ~$ ", prof->credential.server, 25);
		if (err != OK) {
			print_err("server url", err);
			continue;
		}

		chk = 1;
	} while (chk == 0);
}

void get_username(profile *prof)
{
	int err, chk = 0;

	do {
		err = input_char("type your git username ~$ ", prof->credential.username, 25);
		if (err != OK) {
			print_err("username", err);
			continue;
		}

		chk = 1;
	} while (chk == 0);
}

