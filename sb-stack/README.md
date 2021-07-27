
- If `nfs` is used do the following, install the following in each node. 
  - install `nfs-common` `sudo apt-get install nfs-common` in each node
  - add the local docker registry in docker daemon.json in each node
    ```json
      {
        "insecure-registries": [
          "nfs.my.home:5000"
        ]
      }
    ```
  Note: Unfortunately, I run to permission error accessing `/var/lib/mysql`. 
  I choose to use `iscsi` instead. I use [this](https://www.tecmint.com/setup-iscsi-target-and-initiator-on-debian-9/) 
  doc to set up the iscsi` target/server. 
  
- Install a chart
```console
helm install sb-stack ./sb-stack --dry-run
```
or
```console
 helm install sb-stack ./sb-stack -f ./sb-stack/values-arm.yaml --dry-run
```
- uninstall 
```console
helm uninstall sb-stack
```
### Build images

- Built mariadb as
```console
docker buildx b -f Dockerfile-mariadb --platform linux/arm64 . -t sm-db-arm64
docker image tag sm-db-arm64 nfs.my.home:5000/sm-db-arm64
```

- api

```console
docker buildx b -f Dockerfile --platform linux/arm64 . -t sm-api-arm64 --target prod --load
```

#### Misc
- decode base64
```console
echo cGFzc3dvcmQh | base64 --decode
```