# The sidecar approach

Here's the approach in a nutshell:

- forward the `stdout` of the main application container to a file
- run a sidecar container and share the file where the main application container writes the logs
- run a filtering process in the sidecar that continuously reads the main application log file and writes the filtered logs to the sidecar `stdout`
- collect the sidecar `stdout` as the application logs