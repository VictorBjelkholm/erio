# erio

> `go get` directory organization for not only `go get`

Tired of having all repositories in one giant folder? Organize them as they are
organized on Github!

```
$ erio ipfs/notes   
Cloning into '/home/victor/projects/ipfs/notes'...
```

## Installation

`go get -u -v github.com/victorbjelkholm/erio`

## Usage

First, set `ERIO_PATH` to whatever path you usually use for storing repositories.

In my case, this is `/home/victor/projects`, so I have `ERIO_PATH` set to that
value in my `~/.zshrc`.

```
export ERIO_PATH=/home/victor/projects
```

Now you can clone any Github project with either the format `git@github.com/ipfs/notes.git`
or `ipfs/notes` and it'll be placed in `$ERIO_PATH/$organization/$repository`.

In my case, it would be stored in `/home/victor/projects/ipfs/notes`.

# License

MIT 2017 - Victor Bjelkholm
