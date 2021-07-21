<div align="center">
    <h1>AnalAsia</h1>
<hr>
<strong>Nothing over family!</strong><br><br>
<img src="https://img.shields.io/github/workflow/status/mathisburger/dom-toretto/Docker?style=for-the-badge">
<img src="https://img.shields.io/github/license/mathisburger/dom-toretto?style=for-the-badge">
<img src="https://img.shields.io/github/go-mod/go-version/mathisburger/dom-toretto?style=for-the-badge">
</div>
<hr>

<div align="center">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" height="100">
</div>

# Project information

This project is just made for fun. Because at the moment (21.07.2021) there are
so many memes out, about dom toretto. I had the idea to create a discord bot,
that responds to every message that contains a version of the word "family".
The response is a random sentence from these memes. 

# Invite dom to your discord (public instance)

If you want to have dom on your server, you can invite him trough this link: 
<a href="https://discord.com/oauth2/authorize?client_id=867443696194289684&scope=bot&permissions=3072">Invitation</a><br>
**NOTE:** Because the public instance is maybe used by many people, it could be a bit slow.
I am going to improve the performance of dom toretto from time to time.

# Setup
There are two different ways, you can setup the project. 
I recommend setting it up with docker, because this is the easier and faster way, if you
already have docker running on your device.

**Docker Setup**

1. First pull the docker image from the github-package-registry.
```shell
docker pull ghcr.io/mathisburger/dom-toretto:latest
```

2. After that you just have to start the container, based on the image you downloaded
a few seconds ago. 
```shell
docker run -d -e botToken=<botToken> ghcr.io/mathisburger/dom-toretto:latest
```
Consider replacing the `botToken` with your own bot token.

**Compile it yourself**

Compiling dom-toretto yourself is a bit more complicated. First your environment
must match following requirements.
- golang 15 or higher
- Linux (recommended, but not needed)
- git

1. clone the code from this repository
```shell
git clone https://github.com/MathisBurger/dom-toretto.git
```

2. go into the directory, you cloned.
```shell
cd dom-toretto
```

3. Compile the executable
```shell
go build -o ./dom-toretto ./cmd/dom.go
```

4. Make the file executable (Linux)
```shell
chmod +x ./dom-toretto
```

5. set the environment variables (not working on Windows. You need to use the GUI)
```shell
export mode=prod
export botToken=<botToken>
```
Consider replacing the `botToken` with your own bot token.

6. start the bot
```shell
sudo ./dom-toretto
```
**NOTE:** if you are using windows, just execute the file with administrator privileges

# Contributing

If you want to contribute to the project, take a look into the <a href="CONTRIBUTING.md">CONTRIBUTING.md</a>.<br>
If you just want to add new synonyms for family, or a new sentence, open a new issue and 
fill out the template for new synonyms.<br>
If you want to request a new feature, just create a new issue and fill out the template for
a feature request.