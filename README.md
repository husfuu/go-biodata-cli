# Go-Biodata-CLI

A CLI app that can operate simple create, read, update and delete data with JSON data.

## Running this project locally

1. Clone this project locally
2. Run `go build`
3. Enjoy!

## Usage

The following will explain how to run the program

### Get all biodatas

```
go run biodata-cli get -all
```

result:

```
 Id        Name                   Address                                            Reason
---- ----------------- ------------------------------ ---------------------------------------------------------------------
 1    Skyler Kaufman    780-5712 Feugiat Rd.           In ornare sagittis felis. Donec tempor, est ac mattis semper,
 2    Axel Lambert      P.O. Box 145, 4244 Morbi Ave   Sed id risus quis diam luctus lobortis. Class aptent taciti
 3    Janna Avery       140-1888 Eu Av.                pharetra. Nam ac nulla. In tincidunt congue turpis. In condimentum.
 4    Maxwell Beasley   297-4652 Lobortis Street       Nulla semper tellus id nunc interdum feugiat. Sed nec metus
 5    Chastity Hyde     Ap #393-6549 Aliquet Ave       dictum placerat, augue. Sed molestie. Sed id risus quis diam
```

### Get biodata by Id

```
go run biodata-cli get -id="3"
```

result:

```
 Id      Name           Address                                     Reason
---- ------------- ----------------- ---------------------------------------------------------------------
 3    Janna Avery   140-1888 Eu Av.   pharetra. Nam ac nulla. In tincidunt congue turpis. In condimentum.
```

### Add Biodata

```
go run biodata-cli add -name="husfuu" -address="isekai world" -reason="wannabe anime main character"
```

result:

```
 Id        Name                   Address                                            Reason
---- ----------------- ------------------------------ ---------------------------------------------------------------------
 1    Skyler Kaufman    780-5712 Feugiat Rd.           In ornare sagittis felis. Donec tempor, est ac mattis semper,
 2    Axel Lambert      P.O. Box 145, 4244 Morbi Ave   Sed id risus quis diam luctus lobortis. Class aptent taciti
 3    Janna Avery       140-1888 Eu Av.                pharetra. Nam ac nulla. In tincidunt congue turpis. In condimentum.
 4    Maxwell Beasley   297-4652 Lobortis Street       Nulla semper tellus id nunc interdum feugiat. Sed nec metus
 5    Chastity Hyde     Ap #393-6549 Aliquet Ave       dictum placerat, augue. Sed molestie. Sed id risus quis diam
 6    husfuu            isekai world                   wannabe anime main character

```

### Update Biodata

```
go run biodata-cli edit -id="5"
```

### Delete Biodata

```
go run biodata-cli delete -id="5"
```

## Dev Dependencies

- [flag](https://pkg.go.dev/flag#hdr-Usage)
- [simpletables](https://github.com/alexeyco/simpletable)
