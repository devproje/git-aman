#include <glib-2.0/glib.h>

#include "profile.h";

static GList *list = NULL;

void add_profile(profile *prof)
{
	list = g_list_append(list, prof);
}

profile *get_profile(int pos)
{
	// TODO: read
}

void upt_profile(int pos, profile *prof)
{
	// TODO: update
}

void del_profile(int pos)
{
	// TODO: delete
}

