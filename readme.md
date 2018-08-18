# Filter Kubernetes logs for secrets and other sensitive information

Here's the approach in a nutshell:

- forward the `stdout` of the main application container to a file
- run a sidecar container and share the file where the main application container writes the logs
- run a filtering process in the sidecar that continuously reads the main application log file and writes the filtered logs to the sidecar `stdout`
- collect the sidecar `stdout` as the application logs

This approach allows us to filter the logs for a single application in the cluster, and assumes the main application can be modified to output its `stdout` to a file.


# Using and deploying the sample

In [the `sample` directory](/sample) there is a simple NodeJS application and the Kubernetes manifests to deploy the main application and the filter.

# Building from source and running locally

Prerequisites:

- [the Go toolchain][go]
- [glide][glide]
- [make][make] (optional)

To build from source:

- glide install
- `make build` or `go build` to build the binary for your OS
- if running and testing locally, you must specify a local `LOGS_FILE`, the location of your Kubernetes config file through `KUBECONFIG` and the desired namespace to filter secrets from through `KUBECONFIG`.

[go]: https://golang.org/doc/install
[glide]: https://github.com/Masterminds/glide
[make]: https://www.gnu.org/software/make/
