# pssh

```
make install
pssh -host h1  -host h2 -c "ls -a"
```

## ansible group support

```
pssh -g oneid -c "sudo tail -100  /data/services/oneid-0.0.26/admin/start.log | grep rpc | grep 14:3"
```
