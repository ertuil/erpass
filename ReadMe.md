# Erpass
A simple but security password generator.

# LICENSE 

Copyright @ 2019, ertuil. Released under the MIT license.

# Installation

1. Git clone this repo.
2. Install Golang and Make.
3. (Optional) if you want to generate a new static package, Install go-bindata and run `make static`
4. run `make darwin` or `make windows` or `make linux` to build
5. run `erpass-* --host 0.0.0.0 --port 80` (could be erpass-linux or erpass-windows.exe) 
6. Open page with your browser.
7. Generate a new Secret Key or import one.(Key it safe, never lose it).

# Usage



# Algorithm

There are two parts of keys, one is your master key(you should remember it), and the other is SecretKey(generated at the first time you run the server).

We use PBKDF2, SHA256 and HMAC to generate a middle password and map it 
into final password.
