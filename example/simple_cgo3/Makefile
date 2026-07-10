LIB_NAME = sdk
STATIC_LIB_FILE = lib/static/lib$(LIB_NAME).a
DYNAMIC_LIB_FILE = lib/dynamic/lib$(LIB_NAME).so
CFLAGS = -Iinclude
# 动态库需要 -fPIC
SHARED_CFLAGS = $(CFLAGS) -fPIC
TARGET = main


all: clean $(STATIC_LIB_FILE) $(DYNAMIC_LIB_FILE) $(TARGET) run

$(STATIC_LIB_FILE): csrc/sdk.c include/api.h | lib
	gcc -c $(CFLAGS) csrc/sdk.c -o temp.o
	ar rcs $@ temp.o
	rm temp.o

$(DYNAMIC_LIB_FILE): csrc/sdk.c include/api.h | lib
	gcc -c $(SHARED_CFLAGS) $< -o temp.o
	gcc -shared -o $@ temp.o
	rm -f temp.o

$(TARGET):
	go build -o $(TARGET) main.go

lib:
	mkdir -p lib/dynamic
	mkdir -p lib/static

run:
	./$(TARGET)

clean:
	rm -rf lib $(TARGET) *.o