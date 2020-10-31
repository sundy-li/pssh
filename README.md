# pssh


## Install

- Binary: Download from [Release](https://github.com/sundy-li/pssh/releases)
- By source code: `make install`

## Simple usage

```
pssh -host h1  -host h2 -c "ls -a"
```

## Ansible group support

```
pssh -a ansible_file -g oneid -c "hostname"
```

- `ansible_file` is ansible format file,defaults to `/etc/ansible/hosts`, eg:

```
[oneid]
10.221.112.153
10.221.112.154
10.221.112.155

[clickhouse]
10.221.112.153
10.221.112.157
10.221.112.159
10.221.112.154
```

## pssh with rsync
```
    pssh rsync -g clickhouse  -s dist -d /data1/
```

## More information

```
pssh -h
```