# maws2json

## Usage examples

### Over serial line (stdin pipe)

Lets assume that Vaisala weather station is connected via RS232 to USB serial dongle in /dev/ttyUSB0:

```sh
cat /dev/ttyUSB0 | maws2json
```

### Over TCP/IP stream (stdin pipe)

Lets assume thatVaisal weather station can be reached from IP address 10.10.10.10 in a port 5000:

```sh
netcat 10.10.10.10 5000 | maws2json
```

### Convert data lines from a file

Lets assume that you have a file called "maws.dat" that contains MAWS datalines:

```sh
maws2json --file maws.dat
```

### Read data and submit it to GCP Pub/Sub

Lets assume that you are reading again from serial port in real time:

```sh
cat /dev/ttyUSB0 | maws2json |xargs -I % gcloud pubsub topics publish projects/antti-peltonens-lab-solita/topics/maws --message='%'
```
