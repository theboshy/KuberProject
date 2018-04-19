# KuberProject  <img style="display:inline-block" width="40" heigth="40" src="https://png.icons8.com/ios/50/000000/developer.png">

Servidor **RPC** en *golang* com arquitectura de *microservicio* con integracion a contenedores **docker** y administracion **kubernetes**

### Requerimientos 
* [Minikube](https://github.com/kubernetes/minikube) - (mini) servicio local de kubernetes 
* [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - herramienta de línea de comandos de Kubernetes 
* [ProtoBufCompiler](https://github.com/google/protobuf) - compilador de proto buf
* [GoProtoBufCompiler](https://github.com/golang/protobuf) - compilador de proto buf para **golang**


### Build Project
```sh
$ cd ./[<project_path>]
$ protoc -I ./mcs --go_out=plugins=grpc:./pb ./mcs/*.proto

$ eval $(minikube docker-env)
$ docker build -t local/gcd -f Dockerfile.gcd .

$kubectl apply -f gcd-deployment.yaml
```
