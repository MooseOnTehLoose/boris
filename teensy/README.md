# Teensy Firmware

This firmware is currently a work in progress. It can be flashed to a teensy-lc via arduino with the teensy plugin enabled: https://www.pjrc.com/teensy/td_download.html

## Features
Currently this firmware enables the teensy to:

- emulate a mouse/keyboard to call in a supply drop in deep rock galactic
- recieve PING /send ACK via usb serial
- control LEDs
- update a seven segment display with the current nitra values. 


Other functionality to be added soon. 

## Flashing Setup

The correct settings in arduino for flashing this code to a teensy-lc are:

- Board: "Teensy LC"
- USB Type: "Serial + Keyboard + Mouse + Joystick"
- CPU Speed: "48 MHz"
- Optimize: "Smallest Code"

There is a bug in the SevenSegmentTM1637 library installed via the arduino library manager. 
To fix it, go to your arduino libraries folder, find SevenSegmentTM1637/src/SevenSegmentFun.h and rename:
void bouchingBall(
to 
void bouncingBall(
then save and close. Now this file will complile properly. 
