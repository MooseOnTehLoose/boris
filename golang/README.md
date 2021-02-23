# Building the Companion App

The companion app utilizes the walk library for golang. In order to build boris.exe successfully, you need to include a manifest file. The generic manifest file included in this repo should work but please file an issue if it does not on your windows machine. MacOS X and Linux are not supported as Deep Rock Galactic does not support running on these platforms natively. 

To start building, place main.go boris.manifest in the same directory, open up a terminal and type:

rsrc -manifest test.manifest -o rsrc.syso

Then compile with:

go build -ldflags="-H windowsgui"

## Work In Progress

Currently the companion app can manually update the nitra values displayed on the teeny. It can automatically connect to a teensy on startup and if you forget to plug in the teensy or unplug it, you can reconnect to it. 

Attempting to reconnect while already connected will throw an error and crash the program and teensy. Restart the teensy before reconnecting. 

The next milestone for the companion app will be to dynamically update the teensy as the team's nitra value changes. 
