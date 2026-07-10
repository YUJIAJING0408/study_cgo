LIB_NAME = sdk
LIB_FILE = lib/lib$(LIB_NAME).a
TARGET = main
CFLAGS = -Iinclude

all: clean $(LIB_FILE) $(TARGET) run

$(LIB_FILE): csrc/sdk.c include/api.h | lib
	gcc -c $(CFLAGS) csrc/sdk.c -o temp.o
	ar rcs $@ temp.o
	rm temp.o

$(TARGET):
	go build -o $(TARGET) main.go

lib:
	mkdir -p lib

run:
	./$(TARGET)

clean:
	rm -rf lib $(TARGET) *.o