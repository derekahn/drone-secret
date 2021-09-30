# 🤫 Drone Secret

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/derekahn/drone-secret)](https://goreportcard.com/report/github.com/derekahn/drone-secret)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/derekahn/drone-secret)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/derekahn/drone-secret)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Sourcegraph](https://sourcegraph.com/github.com/gorilla/mux/-/badge.svg)](https://sourcegraph.com/github.com/derekahn/drone-secret?badge)

A plugin to "to interpolate base64 secrets".

## Usage

The following settings changes this plugin's behavior.

- Secrets (required) takes a stringified map (refer to [envconfig](https://github.com/kelseyhightower/envconfig)) and base64 encodes the values while it finds and replaces.
- Directory (optional) is the targeted directory to recursively interpolate
- DenyList (optional) takes a list of files to ignore

Below is an example `.drone.yml` that uses this plugin.

```yaml
kind: pipeline
name: default

steps:
  - name: run "derekahn/drone-secret" plugin
    image: "derekahn/drone-secret"
    pull: if-not-exists
    settings:
      secrets: "${FOO}:alpha,${BAR}:bravo,${BAZ}:charlie"
      directory: "deployments"
      denyList: "deployment.yaml"
```

Below is an **input** example of a file in `deployments/` **before** plugin execution:

```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: test-secret
type: Opaque
data:
  FOO: ${FOO}
  BAR: ${BAR}
  BAZ: ${BAZ}
```

Below is an **output** example of a file in `deployments/` **after** plugin execution:

```yaml
# deployments/secret.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: test-secret
type: Opaque
data:
  FOO: YWxwaGE=     # alpha
  BAR: YnJhdm8=     # bravo
  BAZ: Y2hhcmxpZQ== # charlie
```

> Instantiated the project with [boilr-plugin](https://github.com/drone/boilr-plugin) 👏🏽

## 🚀 Building

Build the plugin binary:

```bash
scripts/build.sh
```

## 🔬 Testing

Execute the plugin from your current working directory:

```bash
docker build -t "derekahn/drone-secret" -f docker/Dockerfile .

docker run --rm \
  -e DRONE_COMMIT_SHA=8f51ad7884c5eb69c11d260a31da7a745e6b78e2 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_BUILD_NUMBER=43 \
  -e DRONE_BUILD_STATUS=success \
  -e PLUGIN_LOG_LEVEL=debug \
  -e PLUGIN_DIRECTORY=test/ignore/ \
  -e PLUGIN_SECRETS='${AWS_REGION}:alpha,${AWS_ACCESS_KEY_ID}:bravo' \
  -w /drone/src \
  -v $(pwd):/drone/src \
  "derekahn/drone-secret"
```

After execution be sure to revert the file `test/ignore/dev_test.yaml` !

## 🔒 Security

This plugin is **highly** secure with just a single dependency ([envconfig](https://github.com/kelseyhightower/envconfig/blob/master/LICENSE)) and a hardened and locked down 🐳 container.

Utilizing [trivy](https://github.com/aquasecurity/trivy) for vulnerability scans.

#### Container (🐳 image) scan for misconfiguration and vulnerabilities:

```bash
$ trivy image derekahn/drone-secret:v1

2021-09-29T17:08:59.613-0700    INFO    Need to update DB
2021-09-29T17:08:59.613-0700    INFO    Downloading DB...
24.18 MiB / 24.18 MiB [----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 4.90 MiB p/s 5s
2021-09-29T17:09:05.657-0700    INFO    Number of language-specific files: 1
2021-09-29T17:09:05.657-0700    INFO    Detecting gobinary vulnerabilities...

bin/plugin (gobinary)
=====================
Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)
```

#### Project scan for misconfiguration and vulnerabilities:

```bash
$ trivy fs --security-checks vuln,config .

2021-09-30T09:20:23.213-0700    INFO    Number of language-specific files: 1
2021-09-30T09:20:23.213-0700    INFO    Detecting gomod vulnerabilities...
2021-09-30T09:20:23.213-0700    INFO    Detected config files: 1

go.sum (gomod)
==============
Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)


docker/Dockerfile (dockerfile)
==============================
Tests: 23 (SUCCESSES: 23, FAILURES: 0, EXCEPTIONS: 0)
Failures: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)
```

## 📦 Licenses

- [x] [MIT](https://github.com/kelseyhightower/envconfig/blob/master/LICENSE) github.com/kelseyhightower/envconfig
