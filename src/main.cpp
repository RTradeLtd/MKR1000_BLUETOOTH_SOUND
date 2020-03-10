#include <Arduino.h>
/****************************************
Example Sound Level Sketch for the 
Adafruit Microphone Amplifier
****************************************/
void setup() 
{
   Serial.begin(9600);
}
 
 
void loop() 
{
   unsigned int sample = analogRead(0);
   Serial.println(sample);
}