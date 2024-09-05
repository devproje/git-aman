#include <glib-2.0/glib.h>

#include "profile.h";

static GList *list = NULL;

void add_profile(profile *prof)
{
	list = g_list_append(list, prof);
}

profile *get_profile(int pos)
{
	return g_list_nth_data(list, pos);
}

void upt_profile(int pos, profile *prof)
{
	// TODO: update
}

void del_profile(int pos)
{
	profile *data = get_profile(pos);
	list = g_list_remove(list, data);
}

int get_prof_len()
{
	return g_list_length(list);
}

void destroy()
{
	g_list_free(list);
}

