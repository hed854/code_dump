#include <Arduino.h>
#include <Wire.h>

static int i2cAddress = 0;
static bool isScanDone = false;
static bool isScanSuccess = false;


bool scan_i2c()
{
  byte error, address;
  int nDevices;

  Serial.println("Scanning...");

  nDevices = 0;
  for (address = 1; address < 127; address++)
  {
    // The i2c_scanner uses the return value of
    // the Write.endTransmisstion to see if
    // a device did acknowledge to the address.
    Wire.beginTransmission(address);
    error = Wire.endTransmission();

    if (error == 0)
    {
      Serial.print("I2C device found at address 0x");
      if (address < 16)
        Serial.print("0");
      Serial.println(address, HEX);

      nDevices++;

    }
    else if (error == 4)
    {
      Serial.print("Unknown error at address 0x");
      if (address < 16)
        Serial.print("0");
      Serial.println(address, HEX);
    }
  }
  if (nDevices == 0)
  {
    Serial.println("No I2C devices found\n");
    return false;
  }
  else
  {
    Serial.println("Scan I2C OK\n");
    i2cAddress = address;
    return true;
  }

  delay(5000); // wait 5 seconds for next scan
}
void requestEvent(){
  // Sends a character when the master uses RequestFrom
  Serial.println("master requests something");
  Wire.write("Z");
}

void receiveEvent(int bytes){
  Serial.println("master sends something");
  while (Wire.available())
  {
   char data = Wire.read();
   Serial.print(data);
  }
}
void setup() {
  // Disable internal pullups
  digitalWrite(SDA, 0);
  digitalWrite(SCL, 0);

  // I2C Protocol slave
  // default SDA pin = A4
  // default SCL pin = A5 
  Wire.begin();
  Wire.setClock(400000);
  //Wire.onRequest(requestEvent);
  //Wire.onReceive(receiveEvent);

  // Debug serial
  Serial.begin(9600);
}

void loop() {
  // Scan once plz
  if(!isScanDone)
  {
    // Don't start right away
    delay(10000);
    isScanSuccess = scan_i2c();
    isScanDone = true;
  }

  if(isScanDone && isScanSuccess)
  {
    //read_usb(0x22);
    Wire.requestFrom(0x22, 1);
    while (Wire.available())
    {
     char data = Wire.read();
     Serial.print(data);
    }
    Wire.requestFrom(0x7C, 1);
    while (Wire.available())
    {
     char data = Wire.read();
     Serial.print(data);
    }
  }
}