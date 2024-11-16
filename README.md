# quotaCompare

`quotaCompare` is a tool designed to compare requested resource quotas against the actual resource usage in Kubernetes namespaces. This can help ensure that your Kubernetes clusters are properly configured and that resources are being utilized as expected.

## Overview

The project is structured as an example Git repository where each directory represents a Kubernetes namespace. Within each namespace directory, there is a `resourceQuota.yaml` or `resourceQuota.json` file that specifies the requested resource quotas for that namespace.

## Features

- **Namespace-based Quota Comparison**: Each directory corresponds to a namespace, allowing for organized and straightforward quota management.
- **Support for YAML and JSON**: Resource quotas can be defined in either [YAML](https://en.wikipedia.org/wiki/YAML) or [JSON](https://www.json.org/json-en.html) format.
- **Kubernetes Integration**: The tool checks the requested quotas against the actual usage in Kubernetes.

## Getting Started

### Prerequisites

- [Go programming language](https://go.dev/) installed on your machine.
- Access to a Kubernetes cluster where you have permissions to view resource quotas.

### Installation

1. Clone the repository:
```bash
   git clone [https://github.com/jmainguy/quotaCompare.git](https://github.com/jmainguy/quotaCompare.git)
   cd quotaCompare
```

2. Build the project:
```bash
go build
```
## Usage

Run the tool by specifying the path to the directory containing your namespace directories:

```bash
./quotaCompare /path/to/namespaces
```

The tool will read each resourceQuota.yaml or resourceQuota.json file and compare the requested quotas with the actual usage in the Kubernetes cluster.

## Example

An example directory structure might look like this:

```bash
exampleQuotaGitRepo/
├── namespace1/
│   └── resourceQuota.yaml
├── namespace2/
│   └── resourceQuota.json
```
