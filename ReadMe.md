# Erpass
A simple but security password generator.Erpass do not use any databases to store your information and password.

Erpass written in GO runs anywhere Go can compile for: Windows, macOS, Linux, ARM, etc and is easy to install.

Erpass is light weight and open source.

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
8. (Optional) Backup the SecretKey and keep it safe.

# How to gererate a new password

The only thing you need to remember is the MasterKey. It should be strong enough.

Input the 'Account' like 'google.com/account1' and your MasterKey.

You can set the 'Style', 'Length' and 'Count' options but you don't have to change them in normal.

'Count' is used to generate different passwords, if you want to change the password but not the 'Account' or the website asks you to do so.

# Command Usage

Erpass: A simple but security password generator.
Usage: erpass [-h][-d true][-host <host>][-port <port>][-log <logfile>][-g account]

	--d    run app as a daemon with -d=true or -d true.
	--host string
		  HTTP listen IP address. (default "127.0.0.1")
	--log string
		  Log file (default "erpass.log")
	--port string
		  HTTP listen port (default "8080")
	-g  	generate a password in terminal.
			Example:  erpass -g github.com
	-h/--help Show this help.

# Security Notice

1. As Erpass do not store anything, if you forget your Master Key or lose SecretKey (generate automatical when Erpass first boot), you may 
never get a chance to retrive your passwords, please BACKUP the SecretKey( store in plain-text in 'erpass.key') and your MasterKey.

2. SecretKey is stored in plain-text, the security of SecretKey is depending on the security of your machine.

3. A security channel (HTTPS) to provide the service is necessary, but Erpass do not provide HTTPS, use a NGINX or deploy Erpass in a trusted network environment.

# Algorithm

There are two parts of keys, one is your master key(you should remember it), and the other is SecretKey(generated at the first time you run the server).

We use PBKDF2, SHA256 and HMAC to generate a middle password and map it 
into final password.
