# gogarmincmd

A simple tool to get activity from Garmin Connect and format the result on the console.

It's a good way to learn Go.

## Usage

The command only need the id of the activity, available in the URL of the activity page.

	gogarmincmd id_of_activity

If the URL is `http://connect.garmin.com/modern/activity/1234` then the id is `1234`. Then command line will be :

	gogarmincmd 1234

To have percentage of cardiac frequency you can specify the maximum cardiac frequency with `-fcmax` option :

	gogarmincmd -fcmax=200 1234

Just type `gogarmincmd` or `gogarmincmd -h` to have some help.


## Build

You need Go installed : <http://golang.org/doc/install>

Somewhere on your system :

	mkdir gogarmincmd
	cd gogarmincmd
	export GOPATH=`pwd`
	export PATH=$GOPATH/bin:$PATH
	go get github.com/samonzeweb/gogarmincmd
	
The executable file is now in the `bin` subdirectory.


## Language

gogarmincmd generate data with french text only, but it take only a minute to change `diaplay.go` according to your own language.
