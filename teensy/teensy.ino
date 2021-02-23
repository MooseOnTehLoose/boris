//For use with https://github.com/mooseontehloose/boris
//Designed for Teensy-LC

//pin setup
const byte PIN_CLK = 0;
const byte PIN_DIO = 1;
int firePin = 2;
int disp = 30;
int nitra = 0;

const byte numChars = 32;
char receivedChars[numChars];  
boolean newData = false;


#include "SevenSegmentTM1637.h"
SevenSegmentTM1637    display(PIN_CLK, PIN_DIO);

void setup() {
  pinMode(firePin, INPUT_PULLUP);
  Keyboard.begin();
  Serial.begin(9600);   
  display.begin();           
  display.setBacklight(disp); 
  display.print("drg");
}

void loop() {
  printNum(nitra);
  //Listen for the fire button
  while (digitalRead(firePin) == HIGH) {
    recvWithEndMarker();
    delay(50);
  }
}

void recvWithEndMarker() {
    static byte ndx = 0;
    char endMarker = '\n';
    char rc;
   
    while (Serial.available() > 0 && newData == false) {
        rc = Serial.read();

        if (rc != endMarker) {
            receivedChars[ndx] = rc;
            ndx++;
            if (ndx >= numChars) {
                ndx = numChars - 1;
            }
        }
        else {
            receivedChars[ndx] = '\0'; // terminate the string
            ndx = 0;
            newData = true;
            String data = String(receivedChars);
            if (data == "PING"){
              ack();
            } else {
              update(data);
            }
        }
    }
}


void ack() {
    if (newData == true) {
       Serial.print("ACK\n");
       newData = false;
    }
}

void update(String data) {
    if (newData == true) {
       nitra = data.toInt();
       printNum(nitra);
       newData = false;
    }
}

void printNum(int number) {
  bool negative = false;

  display.clear(); display.setBacklight(disp); 
  if (number < 0) {
    number = -number;
    if (number >= 1000) {
      display.print("----");
      return;
    }
    negative = true;
  }
  if (number > 9999) {
    display.print("----");
  }
  else {
    if (number < 10) {
      display.setCursor(0, 3);
    }
    else if ( number < 100) {
      display.setCursor(0, 2);
    }
    else  if (number < 1000) {
      display.setCursor(0, 1);
    } else {
      display.setCursor(0, 0);
    }
    display.print(number);
  }
  if (negative) {
    display.setCursor(0, 0);
    display.print("-");
  }
}
