TARGETS = unstructured

.PHONY: all
all: $(TARGETS)

%: %.go
	go build -o $@ $^

.PHONY: clean
clean:
	rm -rf $(TARGETS)

