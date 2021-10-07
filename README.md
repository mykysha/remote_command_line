# Remote console (3rd theme)

## Derhunov Mykyta, K27

**Built and executed on Go v1.17**

## Needed preparations

**Download and install all resources**

install `Go`

## Local run:

For running server in your IDE, in terminal...

```
make go-server
```

For running client in your IDE, in terminal...

```
make go-client
```

## Server:

Port:

```
:1028
```

## Client:

#### Available commands:

* Show current time:

```
date
```

* Echo text that follows:
```
echo ...
```

* Sleep for given amount of time (input format: 1d 2h 3m 5s):
```
timeout ...
```

* Get list of commands:
```
help
```
```
h
```

* Get info about author:
```
who
```

* Exit client:
```
STOP
```

## Useful makefile commands:

To build for windows-amd64 run... (creates folder in bin/win)

```
make go-build-win
```

To build for mac-amd64 run... (creates folder in bin/mac)

```
make go-build-mac
```

To build for linux-amd64 run... (creates folder in bin/linux)

```
make go-build-linux
```

To format/beautify all code run...

```
make go-formatter
```