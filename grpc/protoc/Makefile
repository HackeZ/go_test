# current_dir=`pwd`

OS := $(shell uname)

ifeq ($(OS),$(filter $(OS), Darwin Linux))
	files := $(shell find . -name '*.proto')
	rmcmd =-rm -rf `find . -type d -name '*-remote'`
	cleancmd =rm -f `find . -name '*.go'`;rm -rf `find . -type d -empty`;rm -rf `find . -type d -empty`
else
	files := $(wildcard *.proto */*.proto)
	rmcmd =-rm -rf $(basename $<)/*-remote
	cleancmd =-rm -rf $(dirs)
endif

objects := $(foreach file, $(files), $(basename $(file))/constants.go)
dirs := $(basename $(files))
pkg_prefix := go_test/grpc/protoc/

all: $(objects)

$(objects) : %/constants.go : %.proto
	protoc --go_out=plugins=grpc,import_prefix=$(pkg_prefix):. $<

.PHONY : clean
clean :
	$(cleancmd)

