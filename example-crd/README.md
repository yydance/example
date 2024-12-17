# k8sservice-crd
练习使用kubebuilder构建crd
## Description
验证基本流程，包含的功能：
- 定义多版本api，编写controller的Reconcile逻辑
- webhook，实现default和ValidateCreate、ValidateUpdate、ValidateDelete
- 添加RBAC权限
- 健康检查，healthz和readyz
- metrics
- TODO

## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/k8sservice-crd:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/k8sservice-crd:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/k8sservice-crd:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/k8sservice-crd/<tag or branch>/dist/install.yaml
```

## NOTICE

### RBAC权限
默认情况下，项目自动生成controller自身的RBAC权限，如果控制器中需要操作k8s其他资源对象，则需要手动添加对应的权限，方式如下：
```internal/controller/getservice_controller.go
// +kubebuilder:rbac:groups=k8sservice.example.cn,resources=getservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8sservice.example.cn,resources=getservices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=k8sservice.example.cn,resources=getservices/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch
```
添加好注释后，生成crd资源
```
make manifests
make install
```

### 自定义资源watch
k8s 资源对象的watch是通过informer实现的，默认go api库里的informer watch的是既有的资源对象，对于自定义的crd资源对象的watch，需要先实现cache的informer，然后在informer上AddEventHandler。
  
需要注意，自定义的crd资源对象，在实现cache ListWatch时，需要使用支持非结构化(unstructured)的函数，否则会报错，这里使用cache.NewSharedInformer，而不是NewListWatchFromClient(只适用于原生资源对象structured)，详见watcher.go里InformerGetService。

