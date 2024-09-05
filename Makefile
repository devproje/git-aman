GC = go
SRCS = *.go **/*.go
OUTPUT = git-aman
PREFIX = /opt/git-aman

all: $(OUTPUT)

$(OUTPUT): $(SRCS)
	go build -o $(OUTPUT) main.go

install:
	mkdir -p $(PREFIX)
	cp -f $(OUTPUT) $(PREFIX)
	ln $(PREFIX)/$(OUTPUT) /usr/bin/$(OUTPUT)

uninstall:
	rm -rf $(PREFIX)
	rm -rf /usr/bin/$(OUTPUT)

clean:
	rm -f $(OUTPUT)

.PHONY: all, clean
