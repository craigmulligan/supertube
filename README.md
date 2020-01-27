# Supertube 🕳️

> A networked unix pipe utility.

Supertube uses [patchbay.pub](https://patchbay.pub/index.html) to make it easy to pipe your stdout to remote machines.

## Usage

```
Usage: ./supertube [FLAGS] [pipeName]
  -h	Print Usage.
  -r	Reader mode, the  default is Writer mode.
```

**Note**: It's a good idea to make you pipeName relatively unique, if you don't provide a pipeName we'll generate one for you. For more info read patchbay's note on [security](https://patchbay.pub/index.html#security).

## Install

```
go get github.com/hobochild/supertube
```

Or have a look at the [releases](https://github.com/hobochild/supertube/releases)

## Example

Say you want to send your best friend a meme but hate using slack.

First we'll open a named "bestFriendPipe" pipe and stream the file. This command will block until
your friend has drained the pipe on the other side.

```shell
> cat meme.jpeg | supertube bestFriendPipe
```

Next your friend opens the same named pipe on his machine in reader mode and streams it to their filesystem.

```shell
> supertube -r bestFriendPipe > meme.jpeg
```

Thats it! I think you should be able to do everything you can do with a regular pipe with supertube if not [let me know](https://github.com/hobochild/supertube/issues).
