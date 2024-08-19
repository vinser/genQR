# Generate a QR code

Simple REST service to generate a QR codes

Based on ["How to Generate a QR Code with Go"](https://www.twilio.com/en-us/blog/generate-qr-code-with-go) tutorial on the Twilio blog.

## Prerequisites

To install and test this service, you will need:

- Golang installed locally
- OpenSSL to generate self-signed certificates (public and private keys) 
- Curl CLI app
- A smartphone with a QR code scanner

## Start the service

To start the service, run the following commands:
1. Build the service
```bash
go build -o genqr genqr.go
cmod 755 genqr
```
2. Generate self-signed certificates
```bash
openssl req -x509 -noenc -days 365 -newkey rsa:2048 -keyout certs/server.key -out certs/server.crt
```

3. Run the service
```bash
./genqr -port=8080 -cert=certs/server.crt -key=certs/server.key

```

## Generate a QR code

To generate a QR code, send a POST request to http://localhost:8080/ with two POST variables:

- **size**: This sets the width and height of the QR code
- **content**: This is the content that the QR code will embed

The curl example, below, shows how to create a QR code 256x256 px that embeds "https://google.com", and outputs the generated QR code to _data/qrcode.png_.

```bash
curl -X POST \
--form "size=256" \
--form "content=https://google.com" \
--output data/qrcode.png \
https://localhost:8080/
```
Open the _data/qrcode.png_ in your image viewer to see the QR code.
Scan the QR code with the QR code scanner on your smartphone and see the content embedded in the QR code.
