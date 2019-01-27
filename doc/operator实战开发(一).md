### 要求

- [dep](https://golang.github.io/dep/docs/installation.html) version v0.5.0+.
- [git](https://git-scm.com/downloads)
- [go](https://golang.org/dl/) version v1.10+.
- [docker](https://docs.docker.com/install/) version 17.03+.
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) version v1.11.0+.
- Access to a kubernetes v.1.11.0+ cluster.
- 翻墙: export http_proxy=socks5://127.0.0.1:1080;export https_proxy=socks5://127.0.0.1:1080

### 安装operator-sdk

#### 方式一:源码编译

```bash
$ mkdir -p $GOPATH/src/github.com/operator-framework
$ cd $GOPATH/src/github.com/operator-framework
$ git clone https://github.com/operator-framework/operator-sdk
$ cd operator-sdk
$ git checkout -b v0.4.0
$ make dep
$ make install

# 会生成$GOPATH/bin/operator-sdk二进制文件
# 检查安装
operator-sdk --version
```

#### 方式二:官方下载

```bash
$ wget -O $GOPATH/bin/operator-sdk https://github.com/operator-framework/operator- sdk/releases/download/v0.4.0/operator-sdk-v0.4.0-x86_64-linux-gnu
$ sudo chmod 755 $GOPATH/bin/operator-sdk

# 检查安装
operator-sdk --version
```

### 创建项目

```bash
# 1.初始化项目
cd $GOPATH/src/github.com/fenggolang/ && operator-sdk new app-operator
# 2.添加api资源模型
operator-sdk add api --kind=App --api-version=app.example.com/v1
# 3.添加controller控制器
operator-sdk add controller --kind=App --api-version=app.example.com/v1
```

### 修改api资源模型

```bash
# 主要是修改Spec和Status这２个结构体部分
pkg/apis/app/v1/app_types.go

# api写或者修改完了之后要运行一下生成zz_generated.deepcopy.go
# Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
# 重要说明：修改此文件后，运行"operator-sdk generate k8s"重新生成代码
wangfeng@wangfeng-PC:~/go/src/github.com/fenggolang/app-operator$ operator-sdk generate k8s
INFO[0000] Running code-generation for Custom Resource group versions: [app:[v1], ] 
INFO[0001] Code-generation complete.
```

### 修改controller控制器

```bash
# 此控制器主要做：判断CR中的定义，是否是创建，是否需要更新，是否是删除(如果是删除则删除对应的CRD即可)
...
```

### 启动operator

```bash
oc new-project operator
oc apply -f deploy/role.yaml
oc apply -f deploy/role_binding.yaml
oc apply -f deploy/crds/app_v1_app_crd.yaml
operator-sdk up local
```

### 创建CR

```bash
# 编写CR文件
cat deploy/crds/app_v1_app_cr.yaml
apiVersion: app.example.com/v1
kind: App
metadata:
  name: example-app
spec:
  # Add fields here
  replicas: 2
  image: nginx
  ports:
    - containerPort: 80
  envs:
    - name: DEMO
      value: app
    - name: GOPATH
      value: gopath
  resources:
    limits:
      cpu: 100m
      memory: 100M
    requests:
      cpu: 100m
      memory: 100M
 # 创建CR应用 
 oc apply -f deploy/crds/app_v1_app_cr.yaml 
```

### 注意事项

```bash
api资源模型定义和controller注意事项

一个完整的流程是：
1. 初始化项目
2. 定义API资源模型
3. 写Controller的业务逻辑(可能需要封装k8s各种实例的资源)，逻辑控制大致分为三部分: 拿到自定义api模型之后先判断这个api资源是否存在于k8s集群中，如果存在的话就需要判断它是否需要更新，如果需要更新那就更新，如果不需要更新那就不做任何处理;如果不存在就需要创建它与之关联的资源模型
```

