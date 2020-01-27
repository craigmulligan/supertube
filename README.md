# supertube

> A p2p unix pipe utility

Supertube uses [p2p]() networking to make unix pipes across machines as easy as possible.

## Usage

```
Usage of supertube:

  supertube [FLAGS] [rendevous]

  flags:
    -h	Print Usage.
    -r init the reading end of the pipe.
    -w init the writing end of the pipe.
```

## Use cases

### You want pipe a file to a friend on another computer or easily show your friend the output of a command.

```shell
> cat myfile.txt | supertube -w paris
```

Then on then other machine use the connect the read end of the pipe.

```shell
> supertube -r paris > myfile.txt
```
