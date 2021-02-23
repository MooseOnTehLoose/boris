# B.O.R.I.S
Button Operated Resupply Initiation System

Boris is an in-development project to enable external control of supply drop pod call ins for Deep Rock Galactic. It consists of two components, an application to read the current value of nitra in game and send value updates, and a microcontroller that is capable of recieving those updates and controlling the ability of the end user to request supply drops based on the current amount of nitra. 

## Companion App
The companion application that reads in values and transmits them over USB is designed in golang and is compatible with Windows. On launch, it attempts to locate the teensy by listing all open COM ports and then attempting to connect to each one, send PING and recieve ACK. It will maintain a connection to the first COM port that successfully responds and does not support multiple teensys being connected at once. 

If you lose connection to the teensy, or it was not plugged in when the companion app was launched, you can click "Reconnect" to attempt to scan through all open COM ports and connect to a teensy. Currently, attempting to reconnect while the companion app is already connected to a teensy will crash the program. 

## Teensy-LC
The teensy-lc is an arduino compatible microcontroller that manages the user input while playing Deep Rock Galactic. It can recieve information about the current amount of nitra in your team's inventory, and it can fire off a command to launch a supply drop. There is also a SAFE/ARM switch to ensure that you cannot accidentally call in a supply drop should you want to conserve nitra. 

## But Why Tho?

I play Gunner. Gunner's bubble shield is mapped by default to keyboard key 4. Supply Pod is mapped to keyboard key 5. A majority of bubble shields are used reactively during high stress situations and there is a larger than average possibility of mistakes being made while attempting to use one. Since both supply pods and bubble shields are used in the exact same way with potentially no delay, it is easy to accidentally call in a supply pod instead of a life saving shield. B.O.R.I.S is designed to help correct this by separating the supply pod from the keyboard entirely. 
