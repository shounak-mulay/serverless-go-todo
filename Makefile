.PHONY: build clean deploy remove gomodgen

export GO111MODULE:=on
export FUNCTIONS_ROOT:=functions

build: gomodgen
	scripts/build-functions.sh

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose 

remove:
	sls remove --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
