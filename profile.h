#ifndef __PROFILE_H__
#define __PROFILE_H__

typedef struct _profile profile;

struct profile_config {
	char name[25];
	char email[50];
};

struct profile_credential {
	char proto[6];
	char server[25];
	char username[25];
	char secret[50];
};

struct _profile {
	char			  display[25];
	struct profile_config 	  *config;
	struct profile_credential *credential;
};

void add_profile(profile *prof);
profile *get_profile(int pos);
void upt_profile(int pos, profile *prof);
void del_profile(int pos);
int get_prof_len();
void destroy();

#endif /* __PROFILE_H__ */
