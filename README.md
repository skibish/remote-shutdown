# Remote Shutdown

This application shutdowns the **Linux server** with **shutdown** command.
For additional security, you should genereate some secret code (string) and start your application with it.
You should provide this secret code in the request to `GET /shutdown` in **s** parameter.

You should run this application from `root` user. So **use this application at your own risk**.

```sh
curl -X GET 'remote:9898/shutdown?s=fancyCode'
```

## Usage

```sh
$ remote-shutdown -h
Usage of remote-shutdown:
  -port string
    	Port to listen on (default "9898")
  -sec-code string
    	Security code
```

## Motivation

For some small projects, or home things you don't want always to SSH to your machine to turn it off.
This is a solution to turn off machine while you are laying down in bed.

## Build

To build binary for Linux run following:

```sh
GOOS=linux GOARCH=386 go build .
```
