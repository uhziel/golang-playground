# controller-runtime 试验

```
$ go get sigs.k8s.io/controller-runtime@v0.16.3
```

## 如何编写一个简单的 controller

- [ex1-basic-controller-using-builder](./ex1-basic-controller-using-builder/)
- [ex2-basic-controller-using-builder-namespaced](./ex2-basic-controller-using-builder-namespaced/)
- [ex3-basic-controller-no-builder](./ex3-basic-controller-no-builder/)

针对 crd 的 controller

- [ex4-crd-controller-using-builder](./ex4-crd-controller-using-builder/)

## 如何构建一个简单的 client cli

- [ex5-crd-client](./ex5-crd-client/)
- [ex6-crd-client-with-cache](./ex6-crd-client-with-cache/)
- [ex7-crd-client-with-cache-fieldindexer](./ex7-crd-client-with-cache-fieldindexer/)
- [ex10-crd-client-watch](./ex10-crd-client-watch/)

cluster 是对 client 的更高封装，是 controller-runtime 把纯 client 相关和 mgr、controller 不直接相关的功能进行封装。

- [ex8-crd-cluster](./ex8-crd-cluster/)
- [ex9-crd-cluster-watch](./ex9-crd-cluster-watch/)
