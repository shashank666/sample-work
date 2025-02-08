# Envoy Gateway Quickstart Guide

## Prerequisites

- **Kubernetes Cluster**
  - Ensure you have access to a Kubernetes cluster.
  - **Compatibility**: Refer to the [Compatibility Matrix](https://gateway.envoyproxy.io/docs/latest/compatibility/) for supported Kubernetes versions.
  - **LoadBalancer Support**: If your cluster lacks a LoadBalancer, install one like [MetalLB](https://metallb.universe.tf/installation/).

- **Helm**
  - Ensure Helm is installed. [Installation Guide](https://helm.sh/docs/intro/install/).

## Installation

### Install Gateway API CRDs and Envoy Gateway

```bash
helm install eg oci://docker.io/envoyproxy/gateway-helm --version v0.4.0 -n envoy-gateway-system --create-namespace
```

- This installs the Gateway API Custom Resource Definitions (CRDs) and Envoy Gateway in the `envoy-gateway-system` namespace.

### Verify Envoy Gateway Deployment

```bash
kubectl wait --timeout=5m -n envoy-gateway-system deployment/envoy-gateway --for=condition=Available
```

### Deploy GatewayClass, Gateway, HTTPRoute, and Example Application

```bash
kubectl apply -f https://github.com/envoyproxy/gateway/releases/download/v0.4.0/quickstart.yaml -n default
```

- This configures Envoy Gateway to listen for traffic on port 80.

## Testing the Configuration

### With External LoadBalancer Support

1. **Retrieve External IP Address**:

   ```bash
   export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
   ```

2. **Test Connectivity**:

   ```bash
   curl --verbose --header "Host: www.example.com" http://$GATEWAY_HOST/get
   ```

### Without LoadBalancer Support

1. **Identify Envoy Service**:

   ```bash
   export ENVOY_SERVICE=$(kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-namespace=default,gateway.envoyproxy.io/owning-gateway-name=eg -o jsonpath='{.items[0].metadata.name}')
   ```

2. **Port Forward to Envoy Service**:

   ```bash
   kubectl -n envoy-gateway-system port-forward service/${ENVOY_SERVICE} 8888:80 &
   ```

3. **Test Connectivity**:

   ```bash
   curl --verbose --header "Host: www.example.com" http://localhost:8888/get
   ```

## Cleanup

### Remove Deployed Resources

```bash
kubectl delete -f https://github.com/envoyproxy/gateway/releases/download/v0.4.0/quickstart.yaml -n default
```

### Uninstall Envoy Gateway

```bash
helm uninstall eg -n envoy-gateway-system
```

### Delete Namespace

```bash
kubectl delete namespace envoy-gateway-system
```

## Next Steps

You have:

- Installed Envoy Gateway.
- Deployed a backend service and configured a gateway.
- Set up routing using Kubernetes Gateway API resources (`Gateway` and `HTTPRoute`).

### Further Exploration

- **[HTTP Routing](https://gateway.envoyproxy.io/docs/latest/tasks/httproute/)**
- **[Traffic Splitting](https://gateway.envoyproxy.io/docs/latest/tasks/trafficsplitting/)**
- **[Secure Gateways](https://gateway.envoyproxy.io/docs/latest/tasks/tls/)**
- **[Global Rate Limiting](https://gateway.envoyproxy.io/docs/latest/tasks/ratelimiting/)**
- **[gRPC Routing](https://gateway.envoyproxy.io/docs/latest/tasks/grpc/)**
