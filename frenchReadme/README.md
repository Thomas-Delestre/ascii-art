# Ascii-Art : 

## Introduction

Ascii-art est un programme qui prend une string pour argument et retourne une représentation graphique de cette string en utilisant l'ASCII dans le terminal.
La string entrée en tant qu'argument peut contenir des chiffres, lettres, espaces, caractères spéciaux et prend en charge les retours à la ligne : ``\n``.
Si une seconde string est entrée, elle permettra de définir la police asii art.
Ce projet à été réalisé dans le cadre de la formation : [Zone01](https://zone01rouennormandie.org/).

Le fichiet test.sh permet de lancer plusieurs commandes dans le terminal pour voir des exemples d'utilisations.

```
Exemple : 

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

## Polices

Avec le progamme se trouvent 4 bibliothèques de caractères qui servent de police d'écriture pour l'ascii-art : 

```
        standard.txt    : standard
        shadow.txt      : shadow
        thinkertoy.txt  : thinkertoy
        varsity.txt     : varsity
```
Pour faire appel à ces polices, il faut la mensionner après l'argument que l'on souhaite transcrire en ascii-art.

```
Exemple : 
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

Si le nom des polices est mal renseigné dans la commande, le programme renvera une erreur.
Si aucune police n'est renseignée, la police par défaut sera `standard`.

****************************************************************************************************************************
## Options

### Help

L'option `--help` permet d'avoir une aide dans le terminal sur toute les options et la façon d'utiliser l'option `--color=<couleur>`.

```
Exemple :
go run . --help
```


### Output

Par défaut, le programme retourne la représentation graphique en ASCII dans le terminal mais l'option `--output=<fileName.txt>` permet d'enregistrer la string dans un fichier au nom renseigné après l'appel de l'option.
Le chemin du fichier sera : `/fonctions/options/output/`
Si l'option `--output=<fileName.txt>` est utilisée en même temps qu'une autre option, le programme renvera une erreur.

```
Usage : 
go run . [OPTION] [STRING] [BANNER]

Exemple :
go run . --output=FileName.txt "hello" shadow
```


### Justify

L'option `--align=<type>` permet de changer l'alignement de l'output dans le terminal.
Seule l'option `--color` peut être utilisée avec l'option `--align=<type>`.
Les options possibles sont : `--align=<left>`, `--align=<right>`, `--align=<center>`, `--align=<justify>`

```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```


### Reverse

L'option `--reverse=<fileName.txt>` permet convertir la représentation graphique en ASCII contenu dans un fichier.
Seule l'option `--color` peut être utilisée avec l'option `--reverse=<fileName.txt>`.
Le fichier à convertir doit se situer dans le dossier `/fonctions/options/reverse/filesToReverse/`

```
Usage : 
go run . [OPTION Reverse] [OPTION Color]

Exemple :
go run . --reverse=fileName.txt --color=red
```


### Color

L'option `--color=<couleur>`permet de changer la couleur de la représentation graphique en ASCII ou celle du résultat de l'option `--reverse` ou de `--align`. L'option `--color` doit être la dernière option entrée.
Les formats possibles avec cette option : 

```
RGB : go run . --color="rgb(x,y,z)" "text"
x,y and z are numbers from 0 to 255.
x = red, y = green, z = blue

HSL : go run . --color="hsl(x%,y%,z%)" "text"
x, y and z are numbers between 0 and 100.
x = red, y = green, z = blue

Hexadecimal : go run . --color="#XXYYZZ" "text"
x,y and z are numbers between 0-9 and A-F.
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


## Co-développeurs

- Delestre Thomas
- Marchais Mickael