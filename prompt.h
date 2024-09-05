#ifndef __PROMPT_H__
#define __PROMPT_H__

#define OK 	  0
#define NO_INPUT  1
#define TOO_LONG  2
#define TOO_SHORT 3

#include "profile.h"

void get_display_name(profile *prof);
void get_git_name(profile *prof);
void get_git_email(profile *prof);
void get_server_proto(profile *prof);
void get_server_url(profile *prof);
void get_username(profile *prof);

#endif /* __PROMPT_H__ */
