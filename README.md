## Small business app stack

A sample small business app to track order, and quote. 

The technology stacks are:
- api in golang,
- database in mysql
- UI in angular

## kubernetes
To create images and push to a register, do
```console
make build
```

My registry is gitlab. I deployed the apps via helm into my local kube cluster which is set up on raspberry pi.
```console
  helm secrets install sb ./helm -f ./helm/helm_vars/secrets.yaml -f ./helm/values-arm-test.yaml
```
- To uninstall
```console
helm uninstall sb
```

## docker compose
To run in development:
```
make -f makefile.dc dev_up
```

To run in production:
```console
make -f makefile.dc up ARM=false
```

for arm architecture
```console
make -f makefile.dc up 
```

### Note

Do `make` for help
```console
build                          build images and push to gitlab container registry
```

For docker-compose use `make -f makefile.dc`

```console
build                          build images with docker-compose. Pass ARM=false for none ARM os
destroy                        destroy running containers and volumes (docker compose)
dev_build                      to build images with docker-compose in development. Pass ARM=false for none ARM os
dev_up                         run containers in development. Pass ARM=false for none ARM os
down                           remove running containers (docker compose)
up                             run containers. Pass ARM=false for none ARM os
```

### Reference
- api structure is based on https://github.com/MartinHeinz/go-project-blueprint/tree/rest-api
- Some angular UI features copied and repurposed from [here](https://https://github.com/tomastrajan/angular-ngrx-material-starter)
- `iscsi` [doc](https://www.tecmint.com/setup-iscsi-target-and-initiator-on-debian-9/)
  to set up the iscsi` target/server.


