# Copyright (c) 2022-present, Ukama Inc.
# All rights reserved.

#
# Makefile for the srvc_router
#
XCC_PATH = ../../nodes/ukamaOS/distro/tools/musl-cross-make/output
VERSION := $(shell git rev-list --abbrev-commit --tags --max-count=1)
# some generic defs.
CC   = gcc
ARCH = X86_64
XGCC = $(XCC_PATH)/bin/x86_64-linux-musl-gcc
CC = $(XGCC)
XLD  = ld
XGXX = g++
HOST = $(shell gcc -dumpmachine)

OPENSSLTARGET = linux-generic64

TARGET_EXEC = srvc_router

BUILD = ./build

VENDOR_LIB=../../nodes/ukamaOS/distro/vendor

# Setting up various compile and link flags
CFLAGS+=-c
CFLAGS+=-g
CFLAGS+=-Wall
CFLAGS+=-O0 -fPIE
CFLAGS+=-I./inc
CFLAGS+=-I$(VENDOR_LIB)/build/include
CFLAGS+=-D_REENTRANT

# Libs
# Libraries
JSON_LIB:=jansson

ULFIUS_LIB_DEP+=-lulfius
ULFIUS_LIB_DEP+=-lorcania 
ULFIUS_LIB_DEP+=-lyder 
ULFIUS_LIB_DEP+=-lmicrohttpd
ULFIUS_LIB_DEP+=-lgnutls
ULFIUS_LIB_DEP+=-lnettle
ULFIUS_LIB_DEP+=-lhogweed 
ULFIUS_LIB_DEP+=-lp11-kit
ULFIUS_LIB_DEP+=-lz


CURL_LIB_DEP=-lcurl 
CURL_LIB_DEP+=-lssl
CURL_LIB_DEP+=-lcrypto

LIBS+=-lpthread
LIBS+=-lrt
LIBS+=-lm

LIBS+=-l$(JSON_LIB)

LIBS+=$(ULFIUS_LIB_DEP)
LIBS+=$(CURL_LIB_DEP)

LIBS+=-luuid

LDFLAGS+=-L$(VENDOR_LIB)/build/lib
LDFLAGS+=-L$(VENDOR_LIB)/build/lib64

#LDFLAGS+=-L/opt/lib/lib

# Soruce files
CFILES   = $(wildcard ./src/*.c)
OBJFILES = $(CFILES:.c=.o)

.PHONY: $(TARGET_EXEC) container

srvc_router: $(TARGET_EXEC)

$(TARGET_EXEC): $(OBJFILES)
	rm -rf $(BUILD); mkdir $(BUILD)
	$(CC) -o $(BUILD)/$(TARGET_EXEC) $(OBJFILES) $(COMM_OBJFILES) $(LDFLAGS) $(LIBS)
	echo "Build Done."

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

all: srvc_router

container: $(TARGET_EXEC)
	@echo Getting required libs
	cp -rf $(VENDOR_LIB)/build/lib $(VENDOR_LIB)/build/lib64 $(BUILD)
	cp -rf $(XCC_PATH)/x86_64-linux-musl/lib/libc.so $(BUILD)/lib
	cp -rf $(XCC_PATH)/x86_64-linux-musl/lib/libatomic.so* $(BUILD)/lib
	@echo Building container
	sudo buildah bud -f . -t ukamaos/service_router:${VERSION}


clean:
	rm -f $(TARGET_EXEC) $(OBJFILES)
	rm -rf $(BUILD)
