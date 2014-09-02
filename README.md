# gogarmincmd

A simple tool to get activity from Garmin Connect and format the result on the console.

It's a good way to learn Go.

## Usage

The command only need the id of the activity, available in the URL of the activity page.

	gogarmincmd id_of_activity

If the URL is `http://connect.garmin.com/modern/activity/1234` then the id is `1234`. Then command line will be :

	gogarmincmd 1234
