---

# Customize EnvoyProxy

## Introduction

Envoy Gateway provides an `EnvoyProxy` Custom Resource Definition (CRD) that allows cluster administrators to customize the managed Envoy Proxy deployment and service. This guide will walk you through the steps to tailor the Envoy Proxy to fit your specific requirements.

## Prerequisites

Before you begin, ensure you have:

- **Envoy Gateway Installed**: Follow the [Quickstart Guide](https://gateway.envoyproxy.io/docs/latest/quickstart/) to install Envoy Gateway and deploy the example manifest.

- **Functional Backend Service**: Verify that you can query the example backend using HTTP.

- **Kubernetes Cluster Access**: Ensure you have the necessary permissions to apply configurations to your Kubernetes cluster.

## Steps to Customize EnvoyProxy

### 1. Define the EnvoyProxy Configuration

Create a custom `EnvoyProxy` resource to specify your desired configurations. Below is an example that sets the IP family to DualStack, allowing Envoy Proxy to serve external clients over both IPv4 and IPv6.

```
apiVersion: gateway.envoyproxy.io/v1alpha1
kind: EnvoyProxy
metadata:
  name: custom-proxy-config
  namespace: default
spec:
  ipFamily: DualStack  # Options: IPv4, IPv6, DualStack
```

**Note**: Ensure your Kubernetes cluster supports the selected IP family configuration. For DualStack support, confirm that your cluster is properly configured for dual-stack networking.

### 2. Apply the EnvoyProxy Configuration

Save the above YAML configuration to a file named `custom-proxy-config.yaml` and apply it to your cluster:

```bash
kubectl apply -f custom-proxy-config.yaml
```

### 3. Update the Gateway to Reference the EnvoyProxy Configuration

Modify your `Gateway` resource to include an `infrastructure.parametersRef` that points to the newly created `EnvoyProxy` configuration.

```
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: eg
spec:
  gatewayClassName: eg
  infrastructure:
    parametersRef:
      group: gateway.envoyproxy.io
      kind: EnvoyProxy
      name: custom-proxy-config
  listeners:
    - name: http
      protocol: HTTP
      port: 80
```

Apply this configuration:

```bash
kubectl apply -f - <<EOF
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: eg
spec:
  gatewayClassName: eg
  infrastructure:
    parametersRef:
      group: gateway.envoyproxy.io
      kind: EnvoyProxy
      name: custom-proxy-config
  listeners:
    - name: http
      protocol: HTTP
      port: 80
EOF
```

### 4. Verify the Deployment

Check the status of the `Gateway` to ensure that the custom configuration has been applied successfully.

```bash
kubectl get gateway/eg -o yaml
```

Look for the `infrastructure.parametersRef` field in the output to confirm it references `custom-proxy-config`.

## Additional Customizations

The `EnvoyProxy` CRD offers various customization options, including:

- **Deployment Resources**: Adjust CPU and memory allocations.

- **Volumes and VolumeMounts**: Add or modify volumes as needed.

- **Service Annotations**: Customize service annotations for advanced configurations.

For detailed instructions on these customizations, refer to the [official documentation](https://gateway.envoyproxy.io/docs/latest/tasks/operations/customize-envoyproxy/).

## Conclusion

By following this guide, you've customized the Envoy Proxy deployment within Envoy Gateway to suit your specific networking requirements. For further customization options and advanced configurations, consult the [Envoy Gateway User Guides](https://gateway.envoyproxy.io/docs/latest/user/).

---

*Note: I have enhanced the original guide for clarity and completeness, following best practices in technical documentation. This is a work sample that I have create for project - CNCF - Envoy Gateway: Enhancing Envoy Gateway Documentation Using CNCF Tech Docs* 
