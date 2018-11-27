<p><img with="100%" src="docs/socialloot.jpg"/></p>

# Socialloot
Socialloot is a platform to share images, texts and links and rate and comment them. It is strongly inspired by reddit and created by me during a lecture in university. View a demo version [here](https://socialloot.dotcookie.me).

## Features
- 🖼 Upload static images and gifs
- 💬 Comment posts and other comments
- 👍🏼 Up and downvote posts and comments
- 🔍 Search for posts, topics and users 

## Installation
The best way to install socialloot is with docker. Just clone the repository, build the container and run it.

```bash
$ git clone https://github.com/KeKsBoTer/socialloot

$ cd socialloot
  # build docker image
$ docker build -t socialloot .

  # start container and expose port 8080
$ docker run -d -p 8080:8080 socialloot
```
Open http://localhost:8080 in browser