# kubesh

kubesh allows you to connect to a container running in a  kubernetes pod or run commands against a container.

You don't need to specify the exact name of the pod as there is partial match and you can narrow the search for a pod by using `-n` for the namespace


## Usage

```
kubesh pod-query [flags]
```

The `pod` query is a string with partial match to your full pod name. You can also use the full pod name if you'd like.


### cli flags

<!-- auto generated cli flags begin --->
 flag                        | default   | purpose
-----------------------------|-----------|---------
 `--namespace`, `-n`    | ``   | If present use a specific namespace

Kubesh will use the kubeconfig located in your home directory. 