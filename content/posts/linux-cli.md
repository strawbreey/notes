---
title: "Linux Cli"
date: 2020-09-16T11:42:08+08:00
draft: false
---

#### create a new file

```shell
touch readme.txt
# or 
> 'readme.txt'
```

####  create a new directory
```shell
mkdir documents
mkdir -p docments
```

####  Copy `readme.txt` to the `documents/` directory

```shell
cp readme.txt documents
```

####  duplicate a file
```
cp readme.txt readme.bak.txt
cp readme{,.bak}.txt
# Note: learn how the {} works with touch foo{1,2,3}.txt and see what happens.
```

####  copy a direcatory
```shell
cp -a myMusic myMedia/
# or
cp -a myMusic myMedia/myMusic
```

#### duplicate a directory
```shell
cp -a myMusic/ myMedia/
# or
cp -a myMusic myMedia/
```

#### move a file
```
mv readme.txt documents/
```

#### rename a file
```shell
mv readme.txt README.txt
```

#### move a directory
```shell
mv myMedia myMusic/
# or
mv myMedia/ myMusic/myMedia
```

#### rename d directory
```shell
mv myMedia/ Mysic
```

#### merge directories
```shell
rsync -oa /images/ images/2/ # note: may over-write files with the same name, so be careful! name, so be careful!
```

#### show file/directory size
```shell
du -sh node_modules/
```

#### show file/directory info
```shell
stat -x readme.md # on macOs
stat readme.md  # on linux
```

#### open a file with thie default program
```shell
xdg-pen file
open file
```

#### zip  a directory
```shell
zip -r archive_name.zip folder_to_compress
```

#### unzip a directory
```shell
unzip archive_name.zip
```

#### peek file in a zip file
```shell
zipinfo archive_name.zip
# or
unzip -l archive_name.zip
```

#### remove a file
```
rm readme.txt
```

#### remove a directory
```shell
rm -r documents
```

#### list directory content
```shell
ls documents    # Simple
ls -la shell    # -l: show in list format. -a: show all files, including hidden. -la combines those options.
ls -alrth shell # -r: reverse output. -t: sort by time (modified). -h: output human-readable sizes.
```

#### tree view a directory and its subdirectories
```shell
tree                                                        # on Linux
find . -print | sed -e 's;[^/]*/;|____;g;s;____|; |;g'      # on MacOS
# Note: install homebrew (https://brew.sh) to be able to use (some) Linux utilities such as tree.
# brew install tree
```

#### find a stale file
```shell
find my_folder -mtime +5 # Find all files modified more than 5 days ago

```

#### show a calendar
```shell
cal # show a calendar
cal 11 2018 # show 2018/11 calendar
```

#### find a future date

```shell
# What is todays date?
date +%m/%d/%Y 

# What about a week from now?
date -d "+7 days"                                           # on Linux
date -j -v+7d                                               # on MacOS

```

#### use a calculator
```shell
bc # user calculator
```

#### force quit a program
```shell
killall program_name
killall nginx
```

#### check server response
```shell
curl -i umair.surge.sh
```

#### view content of a file
```
cat apps/settings.py
# if the file is too big to fit on one page, you can send it to a 'pager' (less) which shows you one page at a time.
cat apps/settings.py | less
```

#### search for a text
```shell
rep -i "Query" file.txt
```

#### view an image
```shell
imgcat image.png
```

#### show disk size
```shell
df -h
```

#### check performance of your computer
```shell
 top
```

#### Hotkeys
```
Ctrl + A  Go to the beginning of the line you are currently typing on
Ctrl + E  Go to the end of the line you are currently typing on
Ctrl + L  Clears the Screen, similar to the clear command
Ctrl + U  Clears the line before the cursor position. If you are at the end of the line, clears the entire line.
Ctrl + H  Same as backspace
Ctrl + R  Let’s you search through previously used commands
Ctrl + C  Kill whatever you are running
Ctrl + D  Exit the current shell
Ctrl + Z  Puts whatever you are running into a suspended background process. fg restores it.
Ctrl + W  Delete the word before the cursor
Ctrl + K  Clear the line after the cursor
Ctrl + T  Swap the last two characters before the cursor
Esc + T   Swap the last two words before the cursor
Alt + F   Move cursor forward one word on the current line
Alt + B   Move cursor backward one word on the current line
Tab       Auto-complete files and directory names
```

- [You-Dont-Need-GUI](https://github.com/you-dont-need/You-Dont-Need-GUI)
- [tldr](https://github.com/tldr-pages/tldr)
