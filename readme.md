# Filter Kubernetes logs for secrets and other sensitive information

This project explores two approaches for filtering logs in Kubernetes applications, each with their own advantages and downsides:

## Running the application with a sidecar container

Here's the approach in a nutshell:

- forward the `stdout` of the main application container to a file
- run a sidecar container and share the file where the main application container writes the logs
- run a filtering process in the sidecar that continuously reads the main application log file and writes the filtered logs to the sidecar `stdout`
- collect the sidecar `stdout` as the application logs

This approach allows us to filter the logs for a single application in the cluster, and assumes the main application can be modified to output its `stdout` to a file.

The main advantage of this approach is that you can enable this only for an application, without running any privileged components in the cluster; the downside is the fact that you need to modify the application to write its logs in a different place.

## Running a privileged DaemonSet

Here's the approach in a nutshell:

- run a privileged Kubernetes DaemonSet that has access to the Kubernetes log file on each worker node
- the DaemonSet will run the same filtering process in place on the log file, so any request to view the logs for a pod will see the modified log stream

This approach has the advantage of not requiring any changes to running applications, together with the fact that it can be enabled for all running applications inside a cluster.
The main downside of this approach is that you need to run a privileged component on every worker node in your cluster - **use at your own risk, and, when possible, always prefer the previous method.**