# Speedreader

A small terminal based speedreading tool, inspired by [pasky's speedread project](https://github.com/pasky/speedread).

The original project was written in Perl. I'm (trying) to implement it in Golang using [Bubbletea](https://github.com/charmbracelet/bubbletea).

## Installation
**Install script:**
```
curl "https://raw.githubusercontent.com/jonh-a/speedreader/main/install.sh" | bash
```

**Building from source:**
```
go build
```

## Usage

### Input
Speedreader supports **piped input** and **file paths**.

**Piped input:**
```
echo "Hi, I don't have anything interesting to say" | speedreader
```

**File paths:**
```
./speedreader -f ~/Desktop/test.txt
```

### Pausing
By default, the app will start iterating through the text. 

You can pass the `-p` flag to start the reading paused.

```
echo "Hi, I don't have anything interesting to say" | speedreader -p
```

You can pause/unpause the reader with the space key while the program is running.

### Words Per Minute
Provide custom words per minute through a command line argument or through a `~/.speedreaderrc` file. A command line argument will be considered more important than a config file.

**As an argument:**
```
echo "Hi, I don't have anything interesting to say" | speedreader -w 250
```

**In a ~/.speedreaderrc file:**
```
wpm: 250
```

You can also adjust the words per minute with the use of the left and right arrow keys while the program is running.

## Development

### Triggering a new release
Push a new release by tagging your commit and incrementing the version number as necessary.
```
$ git commit -m "commit message goes here"
$ git tag v1.0.0
$ git push --atomic origin main v1.0.0
```