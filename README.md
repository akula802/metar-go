# Get METARs with Go

This is a project I'm working on as part of my journey to learn Go. It's pretty simple:

1. Makes a GET request to the FAA's aviationweather.gov REST API
2. Prints the returned content to the console

The compiled executable is called 'metar' and accepts a command-line flag called station. If no station is supplied, it defaults to KBJC.

metar -station=PANC,PAAQ,PATK

I have a similar command mapped as the alias 'localwx' in my ~/.bashrc file.

![screenshot](https://i.imgur.com/hhuSrBH.png)

Feel free to clone this and modify it for your own purposes.

