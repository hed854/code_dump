import serial
import serial.tools.list_ports

ports = [serial.tools.list_ports.main()]
if ports[0]:
    ser = serial.Serial(ports[0])
    print(ser.name)
