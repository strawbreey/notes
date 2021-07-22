---
title: "Phpdoc"
date: 2021-07-22T15:49:29+08:00
draft: false
---

There are 3 ways to install phpDocumentor:

- Using phive (recommended)
- Using the PHAR (manual install)
- Via Docker
- Via Composer

### Using Phive
```php
$ phive install --force-accept-unsigned phpDocumentor
```
For more information about phive have a look at their website. Now you have phpDocumentor installed, it can be executed like this:

php tools/phpDocumentor

### Using the PHAR

Download the phar file from https://github.com/phpDocumentor/phpDocumentor/releases
You can execute the phar like this: php phpDocumentor.phar

### Via Docker
$ docker pull phpdoc/phpdoc
$ docker run --rm -v $(pwd):/data phpdoc/phpdoc

### Via Composer (not recommended)


### How to use phpDocumentor?

The easiest way to run phpDocumentor is by running the following command:

$ phpdoc run -d <SOURCE_DIRECTORY> -t <TARGET_DIRECTORY>
This command will parse the source code provided using the -d argument and output it to the folder indicated by the -t argument.

phpDocumentor supports a whole range of options to configure the output of your documentation. You can execute the following command, or check our website, for a more detailed listing of available command-line options.

$ phpdoc run -h