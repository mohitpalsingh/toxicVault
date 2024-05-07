# ToxicVault -> A distributed filesystem that never goes down.

## To Run:
1. clone the repo
2. in the root of the repo, run "make run"


## Functionality:
A server can locally save a file with hashed filename with decrypted file content. As soon as the save is successful, it encryptes the content of the file with a private key and broadcast it too all the peers in the network which will save it with hashed filename and encrypted file content. The server can get the file on a GET request from it's local storage but in case of data deletion, the server will ask from all the peers in the network for the same file and then locally save it again and serve to the user.


## Description:
P2P library: The p2p library has all the methods that are needed for communication to be done between two nodes on the network, this includes messaging, encoding/decoding and transport interface.

Cryptography: each message file is encrypted using a randomly generated 32-bit key. AES-32 is used to encypt the file content while MD5 is used to hash the file name.

Store: Store has all the methods that are required to read and write to the disk. It also provides functionality for writing a stream and copying while encrypting/decrypting from some source to some destination.

Server: Server runs in a go routine which allows saving and fetching of a file, along which it opens a port for accepting remote procedure calls from fellow nodes which can command the same, storing or fetching of a file.