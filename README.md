# Ascii-Art:

## Introduction

Ascii-art is a program that takes a string as an argument and returns a graphical representation of that string using ASCII characters in the terminal. The input string can contain numbers, letters, spaces, special characters, and supports newline characters (`\n`). If a second string is entered, it will define the ascii art font.
This project was carried out as part of the training : [Zone01](https://zone01rouennormandie.org/).

The test.sh file allows you to run several commands in the terminal to see examples of uses.

```
Example:

go run . "Hello\n\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```

****************************************************************************************************************************

## Fonts

With the program are 4 libraries of characters that serve as a writing font for ascii art:

```
        standard.txt    : standard
        shadow.txt      : shadow
        thinkertoy.txt  : thinkertoy
        varsity.txt     : varsity
```
To call these fonts, you must mention them after the argument you want to transcribe into ascii art.

```
Example:
go run . "hello" shadow | cat -e
                                 $
_|                _| _|          $
_|_|_|     _|_|   _| _|   _|_|   $
_|    _| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $

go run . "hello" thinkertoy | cat -e
                 $
o        o o     $
|        | |     $
O--o o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
```

If the name of the fonts is entered incorrectly in the command, the program will return an error. If no font is specified, the default font will be `standard`.

****************************************************************************************************************************

## Options

### Help

The `--help` option provides assistance in the terminal with all the options and how to use the `--color=<color>` option.

```
Example:
go run . --help
```


### Output

By default, the program returns the ASCII graphical representation in the terminal, but the `--output=<fileName.txt>` option allows you to save the string to a file with the name entered after the option call.
The file path will be: `/fonctions/options/output/`
If the `--output=<fileName.txt>` option is used along with another option, the program will return an error.

```
Usage : 
go run . [OPTION] [STRING] [BANNER]

Example:
go run . --output=FileName.txt "hello" shadow
```


### Justify

The `--align=<type>` option allows you to change the alignment of the output in the terminal.
Only the `--color` option can be used with the `--align=<type>` option.
Possible options are: `--align=<left>`, `--align=<right>`, `--align=<center>`, `--align=<justify>`

```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```


### Reverse

The `--reverse=<fileName.txt>` option allows you to convert the ASCII graphical representation contained in a file.
Only the `--color` option can be used with the `--reverse=<fileName.txt>` option.
The file to be converted must be located in the `/fonctions/options/reverse/filesToReverse/` folder.

```
Usage : 
go run . [OPTION Reverse] [OPTION Color]

Example :
go run . --reverse=fileName.txt --color=red
```


### Color

The `--color=<color>` option allows you to change the color of the ASCII graphical representation or that of the result of the `--reverse` or `--align` option. The `--color` option must be the last option entered.
The possible formats with this option:

```
RGB: go run . --color="rgb(x,y,z)" "text"
x, y, and z are numbers from 0 to 255.
x = red, y = green, z = blue

HSL: go run . --color="hsl(x%,y%,z%)" "text"
x, y, and z are numbers between 0 and 100.
x = red, y = green, z = blue

Hexadecimal: go run . --color="#XXYYZZ" "text"
x, y, and z are numbers between 0-9 and A-F.
x = red, y = green, z = blue

ANSI Extended: go run . --color="ansi(38;5;XYZ)" "text"
XYZ is a number between 0 and 255.
XYZ is the ID of the color in the extended ANSI color encoding format.

ANSI Standard: go run . --color="ansi(XY)" "text"
XY is the ID of the color in the standard ANSI color encoding format.
```


```
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> "something"
```


## Co-developers

- [Delestre Thomas](https://github.com/Thomas-Delestre)
- [Marchais Mickael](https://github.com/Jeancrock)