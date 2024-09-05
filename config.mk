CC = gcc
LIBS = $(shell pkg-config --cflags glib-2.0)
CFLAGS = -Wall 					\
	 -Wextra				\
	 $(shell pkg-config --libs glib-2.0)

OUTPUT = git-aman

