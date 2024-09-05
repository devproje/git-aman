CC = gcc
LIBS = $(shell pkg-config --cflags glib-2.0)
CFLAGS = -Wall 				     		\
	 -Wextra 			     		\
	 $(shell pkg-config --libs glib-2.0)
OBJS = main.o profile.o
SRCS = $(OBJS:.o=.c)

OUTPUT = git-aman

all: $(OUTPUT)

$(OUTPUT): $(OBJS)
	$(CC) $(LIBS) -o $@ $^ $(CFLAGS) $(LDFLAGS)

%.o: %.c
	$(CC) $(LIBS) -c -o $@ $< $(CFLAGS) $(LDFLAGS)

clean:
	rm -f $(OBJS)
	rm -f $(OUTPUT)

.PHONY: all, clean

