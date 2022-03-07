- decode base64
```console
echo cGFzc3dvcmQh | base64 --decode
```
- Copy gpg, https://medium.com/@Devopscontinens/encrypting-helm-secrets-7f37a0ccabeb
```console
ssh binyam@192.168.1.10 /usr/local/bin/gpg --export-secret-key bdarge | /usr/local/bin/gpg --import
```

