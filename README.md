[![CI](https://github.com/takumin/ykmgr/actions/workflows/ci.yml/badge.svg)](https://github.com/takumin/ykmgr/actions/workflows/ci.yml)

# ykmgr

Yubikey Manager

# Development

## PCSC Lite

ykmgr requires PCSC Lite.

For MacOS:

No additional package required.

For Windows:

No additional package required.
However, it may not work properly.

For LinuxBrew:

```
brew install pcsc-lite
```

For Alpine:

```
sudo apk add pcsc-lite-dev
```

For Debian/Ubuntu:

```
sudo apt-get install libpcsclite-dev
```

For Fedora:

```
sudo yum install pcsc-lite-devel
```

For CentOS:

```
sudo yum install 'dnf-command(config-manager)'
sudo yum config-manager --set-enabled PowerTools
sudo yum install pcsc-lite-devel
```

For FreeBSD:

```
sudo pkg install pcsc-lite
```
