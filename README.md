# EKS SSO

A service to add SSO to EKS cluster.

## Install

Have helm tiller installed and then run helm command with desired values

`helm install eks-sso/ --name eks-sso`

## Update Helm Chart

- First update your `Chart.yaml` to desired version, then update your package.

1. `helm package eks-sso`
2. `helm repo index .`