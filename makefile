PLUGINS_DIR=./plugins
PLUGINS_SRC=$(wildcard $(PLUGINS_DIR)/*/*.go)
PLUGINS_SO=$(patsubst %.go,%.so,$(PLUGINS_SRC))

all: main $(PLUGINS_SO)

main:
	go build -o main ./main.go

$(PLUGINS_SO): $(PLUGINS_SRC)
	go build -buildmode=plugin -o $@ $<

clean:
	rm -rf main $(PLUGINS_SO)