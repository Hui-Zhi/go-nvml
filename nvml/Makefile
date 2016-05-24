VERSION=1.0

all: libgonvml-$(VERSION).so

CC=gcc
#CFLAGS=-I /usr/include/nvidia/gdk/
LDFLAGS=-lnvidia-ml -L /usr/src/gdk/nvml/lib/


gonvml.o: gonvml.c
	$(CC) -c -Wall -Werror -fpic $< $(CFLAGS) $(LDFLAGS)

libgonvml-$(VERSION).so: gonvml.o
	$(CC) --shared -o $@ $<

.PHONY: clean

install:
	cp libgonvml-$(VERSION).so /usr/lib64/
	ln -s /usr/lib64/libgonvml-$(VERSION).so /usr/lib64/libgonvml.so

clean:
	rm *.o *.so /usr/lib64/libgonvml.so /usr/lib64/libgonvml-$(VERSION).so

