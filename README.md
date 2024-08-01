# Inspeqtor

[![GoDoc](https://godoc.org/github.com/mperham/inspeqtor?status.svg)](https://godoc.org/github.com/mperham/inspeqtor)

Inspeqtor monitors your application infrastructure.  It gathers and verifies key metrics
from all the moving parts in your application and alerts you when something
looks wrong.  It understands the application deployment workflow so it
won't bother you during a deploy.

What it does:

* Monitor init.d-, systemd-, upstart-, runit- or launchd-managed services
* Monitor process memory and CPU usage
* Monitor daemon-specific metrics (e.g. redis, memcached, mysql, nginx...)
* Monitor and alert based on host CPU, load, swap and disk usage
* Alert or restart a process if a rule threshold is breached
* Alert if a process disappears or changes PID
* Signal deploy start/stop to silence alerts during deploy

What it doesn't:

* monitor or control arbitrary processes, services must be init-managed
* have *any* runtime dependencies at all, not even libc.

If you've used `monit` before, Inspeqtor will look familiar.  Same
high-level goals but in a modern package.

## Status

Inspeqtor is **feature complete**, reliable and (mostly?) bug-free.  This repo
does not see a lot of code changes because of this, not because it is unmaintained.


## Installation

See the [Inspeqtor wiki](https://github.com/mperham/inspeqtor/wiki) for complete documentation.


## Requirements

Linux 3.0+.  It will run on OS X.  FreeBSD is untested.  It uses about 5-10MB of RAM at runtime.


## License

GPLv3.


## Want to Help?

See the [Development](https://github.com/mperham/inspeqtor/wiki/Development) wiki page for details on how
to get the source code and build Inspeqtor locally.


# Author

Inspeqtor is written by [Mike Perham](http://twitter.com/mperham) of [Contributed Systems](http://contribsys.com).  We build awesome open source-based infrastructure to help you build awesome apps.

We also develop [Sidekiq](http://sidekiq.org) and sell [Sidekiq Pro](https://sidekiq.org/products/pro.html), the best Ruby background job processing system.



Hack around:
Run code in docker:
```
 docker run -it -v $(pwd):/go/src/github.com/mperham/inspeqtor golang:1.10.3 bash
```

Install deps:
```
 go get github.com/jteeuwen/go-bindata/...
 go get github.com/goccmack/gocc/...
 cd src/github.com/goccmack/gocc/
 git checkout 0e2cfc030005b281b2e5a2de04fa7fe1d5063722
 cd /go
 go get github.com/goccmack/gocc/...
 cd src/github.com/mperham/inspeqtor/
 make assets
 GOOS=linux GOARCH=amd64 go build -o inspeqtor cmd/main.go
```
Yes, we need to checkout old version of code to match go version, sigh

