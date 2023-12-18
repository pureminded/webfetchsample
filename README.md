# Binary Data Transmission Project

## Overview
This project demonstrates the transmission of binary data using the Fetch API in a browser and GoLang backend, based on the library `github.com/valyala/fasthttp`. 

## Functionality
- **Data Generation**: Upon clicking a button in the browser, random data is generated.
- **Data Hashing and Transmission**: The generated data is hashed and sent to the server using the Fetch API with the POST method.
- **Server Processing**: The server receives the data, hashes it, and sends the hash back to the client.
- **Data Verification**: The client compares the received hash with the original hash to confirm the accuracy of the data transmission.

## Client-Server Interaction
1. **Client Side**: On button click, random data is generated, hashed, and transmitted to the server.
2. **Server Side**: The server hashes the received data and responds with the hash.
3. **Client Side**: The client compares the hashes and notifies the user about the correctness of the data transfer.

## Technologies Used
- Fetch API
- GoLang
- fasthttp library

## Goal
The primary goal of this project is to ensure secure and accurate binary data transmission between a client and a server.

## Contributors
- Sergei Kirnitskii

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

*This README is subject to changes and updates.*
