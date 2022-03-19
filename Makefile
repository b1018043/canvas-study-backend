FILE=main

.PHONY: clean

all: $(FILE)

$(FILE): $(FILE).go
	go build -o $(FILE) $(FILE).go

clean:
	rm -rf $(FILE)