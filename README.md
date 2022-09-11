# kubectl-count
A simple `kubectl` plugin for counting instances of a limited types of kubernetes resources, including:
1. configmaps
2. daemonsets
3. deployments
4. nodes
5. persistentvolumeclaims
6. persistentvolumes
7. pods
8. replicasets
9. secrets
10. services
11. statefulsets


## Installation
1. Download the appropriate binary for your machine from the [releases](https://github.com/neelanjan00/kubectl-count/releases) page.
2. Extract the binary from the downloaded tar file. You can use the following command to do so:
```bash
tar -xvf <tar-file-path>
```
3. Make the binary executable:
```bash
chmod +x <binary-file-path>
```
4. Move the binary to `/usr/local/bin`
```bash
sudo mv <binary-file-path> /usr/local/bin
```
5. Check installation by running the following command
```bash
kubectl count
```

## Usage
```bash
kubectl count [command]
```
```
Available Commands:
  configmaps             Count configmaps in a namespace.
  daemonsets             Count daemonsets in a namespace.
  deployments            Count deployments in a namespace.
  nodes                  Count nodes in a cluster.
  persistentvolumeclaims Count persistentvolumeclaims in a namespace.
  persistentvolumes      Count persistentvolumes in a cluster.
  pods                   Count pods in a namespace.
  replicasets            Count replicasets in a namespace.
  secrets                Count secrets in a namespace.
  services               Count services in a namespace.
  statefulsets           Count statefulsets in a namespace.
```

- For example, if you want to count all the pods:
```bash
kubectl count pods
```
- Help for any command or subcommand can be accessed using the `--help` flag.
- The kubectl command aliases such as `po` for pods, `cm` for configmaps, etc. are also supported.
- There are additional flags based on the resources, such as:
```bash
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.

  -h, --help               help for pods

  -n, --namespace string   resource namespace (default "default")

  -l, --selector string    Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.
```
