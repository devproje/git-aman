include config.mk

OBJS = main.o profile.o input.o prompt.o shell.o
SRCS = $(OBJS:.o=.c)

all: $(OUTPUT)

$(OUTPUT): $(OBJS)
	$(CC) $(LIBS) -o $@ $^ $(CFLAGS)

%.o: %.c
	$(CC) $(LIBS) -c -o $@ $< $(CFLAGS)

clean:
	rm -f $(OBJS)
	rm -f $(OUTPUT)

.PHONY: all, clean

