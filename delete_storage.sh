#!/bin/bash

kubectl delete ns storage
kubectl delete sc hpe-iscsi 
kubectl get pv | tail -n+2 | awk '{print $1}' | xargs -I{} kubectl patch pv {} -p '{"metadata":{"finalizers": null}}'
kubectl get pv | awk '{print $1}' | xargs kubectl delete pv
kubectl get volumeattachments | tail -n+2 | awk '{print $1}' | xargs -I{} kubectl patch volumeattachments {} -p '{"metadata":{"finalizers": null}}'
kubectl get volumeattachments | awk '{print $1}' | xargs kubectl delete volumeattachments
