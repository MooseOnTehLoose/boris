# B.O.R.I.S
Button Operated Resupply Initiation System

![Alt text](images/boris-box.jpg?raw=true "Boris")


Boris is an in-development project to enable external control of supply drop pod call ins for Deep Rock Galactic. It consists of two components, an application to read the current value of nitra in game and send value updates, and a microcontroller that is capable of recieving those updates and controlling the ability of the end user to request supply drops based on the current amount of nitra. 

All files and information necessary to build your own Boris will be added to this repository so you can customize it to match your preferences, class colors and even beard. 

## Companion App

![Alt text](images/GUI-Beta.png?raw=true "Companion App")


The companion application that reads in values and transmits them over USB is designed in golang and is compatible with Windows. On launch, it attempts to locate the teensy by listing all open COM ports and then attempting to connect to each one, send PING and recieve ACK. It will maintain a connection to the first COM port that successfully responds and does not support multiple teensys being connected at once. 

## Teensy-LC
The teensy-lc is an arduino compatible microcontroller that manages the user input while playing Deep Rock Galactic. It can recieve information about the current amount of nitra in your team's inventory, and it can fire off a command to launch a supply drop. There is also a SAFE/ARM switch to ensure that you cannot accidentally call in a supply drop should you want to conserve nitra. 

## 3d Printed Enclosure
All the files required to print your own enclosure can be found here: https://www.prusaprinters.org/prints/58566-boris

## Electronics 

![Alt text](images/schematic.png?raw=true "schematic")

To assemble Boris, you'll need to ordrer some components. I provide the examples I use but feel free to grab similar components elsewhere.

1x Teensy-LC: https://www.pjrc.com/teensy/teensyLC.html \
1x 4 Digit Display, TM1637: https://www.amazon.com/gp/product/B01DKISMXK/ \
1x Single Position Single Throw (SPST) toggle switch, 6mm: https://www.amazon.com/MTS-101-Position-Miniature-Toggle-Switch/dp/B0799LBFNY \
2x 5mm ~3v LED: https://www.amazon.com/gp/product/B072BL2VX1 \
2x 220ohm resistor: https://www.amazon.com/gp/product/B072BL2VX1 \
1x Cherry MX switch equivalent: https://www.amazon.com/Cherry-switches-MX1AG1NN-Mechanical-Keyboard/dp/B07RQTNS58 \
