# pssh
```
make install
pssh -host h1  -host h2 -c "ls -a"
```

## ansible group support

```
pssh -a ansible_file -g oneid -c "sudo tail -100  /data/services/oneid-0.0.26/admin/start.log | grep rpc | grep 14:3"
```

- `ansible_file` is ansible format file,defaults to `/etc/ansible/hosts`, eg:

```
[oneid]
10.221.112.153
10.221.112.154
10.221.112.155

[clickhouse]
10.221.112.153
10.221.112.154
10.221.112.153
10.221.112.154
```

## More information

```
pssh -h
```