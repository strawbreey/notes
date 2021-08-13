---
title: "Linux Vscode"
date: 2021-08-10T20:52:56+08:00
draft: false
---

## vs code 远程登录

### 设置设置服务器支持ssh-public-key登录

完成代理设置之后，还需要设置服务器支持ssh-public-key以及AllowTcpForwarding

登录服务器，在root账户下（或使用root权限），vim /etc/ssh/sshd_config 把以下五项配置都打开：

```bash
AllowAgentForwarding yes    # vscode使用TcpForwarding模式，需要把这两项配置打开
AllowTcpForwarding yes      # vscode使用TcpForwarding模式，需要把这两项配置打开
RSAAuthentication yes       # 打开RSA签名认证模式
PubkeyAuthentication yes    # 打开公钥认证模式

# The default is to check both .ssh/authorized_keys and .ssh/authorized_keys2
# but this is overridden so installations will only check .ssh/authorized_keys

AuthorizedKeysFile .ssh/authorized_keys 

# 设置默认的公钥保存文件，设置的值是.ssh/authorized_keys，设置后实际读取的文件为~/.ssh/authorized_keys
```

### 然后重启sshd使配置生效：

```bash
sudo service sshd restart
```

