## Small business app stack

A sample small business app to track order, and quote. 

The technology stacks are:
- api in golang,
- database in mysql
- UI in angular

## docker compose
To run in development:
```
make dev_up 
```

To run in production:
```console
make up ARM=false
```

for arm architecture
```console
make up 
```

## kubernetes
To make images, do
```console
make k_build
make push
```
My local registry is located at nfs.my.home:5000. I deployed the apps via helm into my local kube cluster which is set up on raspberry pi.
```console
  helm secrets install sb ./helm -f ./helm/helm_vars/secrets-all.yaml -f ./helm/values-arm-test.yaml
```
- To uninstall
```console
helm uninstall sb
```

### Note

Do `make` for help
```console
api_doc                        create api doc
build                          build images for local k8s
dc_build                       build containers. Pass ARM=false for none ARM os
destroy                        destroy running containers and volumes (docker compose)
dev_dc_build                   to build containers in development. Pass ARM=false for none ARM os
dev_up                         run containers in development. Pass ARM=false for none ARM os
down                           remove running containers (docker compose)
k_api_arm                      build api image for for local k8s
k_db_arm                       build db image for for local k8s
k_ui_arm                       build ui image for for local k8s
push                           push images to local registry
up                             run containers. Pass ARM=false for none ARM os
```

### Reference
- api structure is based on https://github.com/MartinHeinz/go-project-blueprint/tree/rest-api
- Some angular UI features copied and repurposed from [here](https://https://github.com/tomastrajan/angular-ngrx-material-starter)
- `iscsi` [doc](https://www.tecmint.com/setup-iscsi-target-and-initiator-on-debian-9/)
  to set up the iscsi` target/server.


