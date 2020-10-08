---
title: "Hugo Install"
date: 2020-09-27T15:42:31+08:00
draft: false
---


If you are on a Windows machine and use Chocolatey for package management, you can install Hugo with the following one-liner

```shell
# install-with-chocolatey.ps1
choco install hugo -confirm

# Or if you need the “extended” Sass/SCSS version:
choco install hugo-extended -confirm
```

If you are on a Windows machine and use Scoop for package management, you can install Hugo with the following one-liner:

```shell
scoop install hugo
# Or install the extended version with:
scoop install hugo-extended
```

### Source

```shell
mkdir $HOME/src
cd $HOME/src
git clone https://github.com/gohugoio/hugo.git
cd hugo
# Remove --tags extended if you do not want/need Sass/SCSS support.
go install --tags extended

```
