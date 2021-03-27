# hivemq-demo

This is demo for display mqtt cluster

## Prerequisites

- [terraform](https://www.terraform.io/downloads.html)
- [docker](https://www.docker.com/products/docker-desktop) or [podman](https://podman.io/getting-started/installation)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start#installation)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [helm](https://helm.sh/docs/intro/install/)
- [istioctl](https://istio.io/latest/docs/setup/getting-started/#download) or [brew install istioctl](https://formulae.brew.sh/formula/istioctl)
- [golang](https://golang.org/doc/install)

## Usage

initialize terraform module

```bash
$ terraform init
```

set up enviroment on kind cluster, include istio and hivemq.

```bash
$ terraform apply -auto-approve
```

run examles 

- for publish

```bash
$ cd example/publish && go run main.go
```

- for subscribe

```bash
$ cd examples/subscribe && go run main.go
```

for destroy

```bash
$ terraform destroy -auto-approve
```

![hivemq01](https://github.com/GrassShrimp/hivemq-demo/blob/master/hivemq01.png)