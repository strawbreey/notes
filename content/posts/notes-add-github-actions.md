---
title: "Notes Add Github Actions"
date: 2021-07-29T14:42:47+08:00
draft: false
---

### 博客系统添加github actions

```yml
name: CI
on: push
jobs:
  deploy:
    runs-on: ubuntu-18.04
    steps:
      - name: Git checkout
        uses: actions/checkout@v2

      - name: Update theme
        # (Optional)If you have the theme added as submodule, you can pull it and use the most updated version
        run: git submodule update --init --recursive

      - name: Setup hugo
        uses: peaceiris/actions-hugo@v2
        with:
          hugo-version: "0.64.0"

      - name: Build
        # remove --minify tag if you do not need it
        # docs: https://gohugo.io/hugo-pipes/minification/
        run: hugo --minify

      - name: Deploy
        uses: peaceiris/actions-gh-pages@v3
        with:
          personal_token: ${{ secrets.TOKEN }}
          external_repository: <username>/<username>.github.io
          publish_dir: ./public
          #   keep_files: true
          user_name: <username>
          user_email: <username@email.com>
          publish_branch: master
        #   cname: example.com
```
