# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Usage

### Prerequisites
1. Install openapi-generator-cli: `npm i -g @openapitools/openapi-generator-cli`
1. Install the Java runtime (on Ubuntu derivatives, you may use `sudo apt install default-jre`)

## Execution

### For OpenAPI-Generator

1. To generate the .NET package, run from the root of the repo: `openapi-generator-cli generate -i schemas/api.github.com/api.github.com.2022-11-28.json -t templates/dotnet -g csharp -o generated/dotnet -p packageName=Octokit`. Alternately, you may run the provided script `generate.sh`.
1. To build the .NET package, change directories into generated/dotnet and run `chmod +x build.sh`, followed by `./build.sh`.

### For AutoREST

1. `npm install -g autorest`
1. `time autorest --debug --verbose` from the repo root

#### Literate Configuration

The following comment line makes autorest see this file as a source for configuration: `> see https://aka.ms/autorest`

> see https://aka.ms/autorest

Configuration is below:

```yaml
input-file: schemas/api.github.com/api.github.com.json
namespace: Octokit
output-folder: generated
use: "@autorest/csharp@3.0.0-beta.20210210.4"
clear-output-folder: true
modelerfour:
  # this runs a pre-namer step to clean up names
  # defaults to true if not specified
  prenamer: true
  # relaxes schema duplication checks to allow schemas with the
  # same name and renames duplicate schema names with a suffix
  # of "AutoGenerated" with an additional numeric suffix if more
  # than 2 duplicates of the same name are detected.
  #
  # defaults to false if not specified.
  #
  # NOTE: This parameter is a temporary workaround and will be
  # removed in a future release!
  lenient-model-deduplication: true
```

Notes:

- Config options are the same as CLI versions, except for `--debug` and `--verbose` which are only able to be applied through the API.
