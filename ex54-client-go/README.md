## 连接和操作 k8s 集群

ex1-client/

## rest 操作资源

Get 操作，有时有需要排除 NotFound

```
import "k8s.io/apimachinery/pkg/api/errors"

if err != nil && !errors.IsNotFound(err) {
  return err
}
```

相关的是 [client.IgnoreNotFound](https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/client#IgnoreNotFound)。

## 操作 crd

https://pkg.go.dev/github.com/kubernetes-csi/external-snapshotter/client/v6
