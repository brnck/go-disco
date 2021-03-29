# go-disco

## Pre-requisites

When using Rpi LED strip as an output, Raspberry Pi itself must be prepared:

### Preparing RPi

#### Installing

You need to have C library installed on your machine before installing this module

* Clone [this](https://github.com/jgarff/rpi_ws281x) repository
* Install Scons (on raspbian, `apt-get install scons`)
* Type `scons` from inside the source directory.
* Copy compiled lib and rest of the includes:
```markdown
cp ws2811.h /usr/local/include/ws2811.h && \
cp rpihw.h /usr/local/include/rpihw.h && \
cp pwm.h /usr/local/include/pwm.h && \
cp libws2811.a /usr/local/lib/libws2811.a
```
