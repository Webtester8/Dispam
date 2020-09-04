# Dispam
---
## Introduction
Have you ever thought about spamming a discord server because you hated it so much and  you're just too lazy to do it? Congrats! I have been there and I took the high road to doing it!
---
## Install
#### Normal way
Go to the the Versions File and find the file that is your operating system. Then that file in the folder where it is most convenient for you. You will need to run this in the terminal.
#### Unstable way
Make sure you have Golang installed. Then run:
```sh
cd $PATH/Discord_spammer
go run dispam.go
```
---
## Use
1. Double click on the file to run. It should bring up a terminal.
2. You need the Captcha key for discord. So sign in on the website normally but use a HTTP incepter(EX: burp suite) to capture your sign in request. Find your Captcha key. Save this for later.
3. This terminal will ask for you to sign-in into discord. I **Don't Recommend** using this on public wifi since it does exposes your credentials to tools like WireShark.
4. You then have to enter the Server number. This is the first set of number in the URL to the Server(EX: https://discord.com/channels/ `727263403810357258` /727263403810357261 the highlighted section is the server number)
5. You then will have to give the Room Number. This is the rest of the numbers at the end.(EX: https://discord.com/channels/727263403810357258/ `727263403810357261` the highlighted section is the Room number)
6. Now you get to write a message!
7. Now let it run!
---
## License & copyright
Licensed under the [MIT License](LICENSE)
