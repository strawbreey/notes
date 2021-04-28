---
title: "K8s Tutorial"
date: 2020-12-21T15:05:15+08:00
draft: false
---

```bash
# 1. 查看某个名字空间下所有pod
kubectl get pod -n barad -o wide

# 2. 查看所有名字空间下pod
kubectl get pod --all-namespaces

# 3. 在某个pod内开启一个交互式shell（即“进入容器内”）
kubectl exec -it $POD_NAME -n $NAMESPACE -- /bin/bash

# 4. 在pod内执行命令
kubectl exec -it $POD_NAME -n $NAMESPACE -- jps -l

# 5. 查看所有node
kubectl get node

# 6. 查看某个node或pod详细信息
kubectl describe node $NODE_NAME
kubectl describe pod $POD_NAME -n $NAMESPACE

# 7. 查看某个名字空间下service
kubectl get svc -n $NAMESPACE -o wide

# 8. 查看某个service详细信息，后端endpoint等
kubectl describe svc $SVC_NAME -n $NAMESPACE

# 9. 删除某个pod
kubectl delete po $POD_NAME -n $NAMESPACE

# 10. 替换某个service
kubectl delete svc $SVC_NAME -n $NAMESPACE && kubectl create -f xxx.svc.yaml

# 11. 查看某个k8s资源的yaml或json定义
kubectl get po $POD_NAME -n $NAMESPACE -o yaml
kubectl get po $POD_NAME -n $NAMESPACE -o json

# 12.  修改pod的内存大小
kubectl edit deployment  $deployment -n $NAMESPACE
```