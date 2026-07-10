LIB_NAME = sdk
LIB_FILE = lib/lib$(LIB_NAME).a
CFLAGS = -Iinclude
CXX_FLAGS = -Wall -O2 -Iinclude
LDFLAGS = -lstdc++
TARGET = main
CPP_SRCS = csrc/sdk.cpp csrc/wrapper.cpp

all: clean $(LIB_FILE) $(TARGET) run

$(LIB_FILE): $(CPP_SRCS) include/sdk.hpp include/wrapper.h | lib
	g++ $(CFLAGS) -c csrc/sdk.cpp -o sdk.o
	g++ $(CFLAGS) -c csrc/wrapper.cpp -o wrapper.o
	ar rcs $@ sdk.o wrapper.o
	rm sdk.o wrapper.o

$(TARGET):
	go build -o $(TARGET) main.go

lib:
	mkdir -p lib

run:
	./$(TARGET)

clean:
	rm -rf lib $(TARGET) *.o