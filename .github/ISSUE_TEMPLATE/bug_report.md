---
name: Report an issue
about: Create a bug report about an existing issue.
title: ''
labels: 'bug'
assignees: ''

---

**Please do not report security vulnerabilities here**. See the [Responsible Disclosure Program](https://github.com/openfga/go-sdk/blob/main/.github/SECURITY.md).

**Thank you in advance for helping us to improve this library!** Please read through the template below and answer all relevant questions. Your additional work here is greatly appreciated and will help us respond as quickly as possible.

By submitting an issue to this repository, you agree to the terms within the [OpenFGA Code of Conduct](https://github.com/openfga/rfcs/blob/main/CODE-OF-CONDUCT.md).

### Description

> Provide a clear and concise description of the issue, including what you expected to happen.

### Version of SDK

> v0.2.0

### Version of OpenFGA (if known)

> v1.1.0

### OpenFGA Flags/Custom Configuration Applicable

>    environment:
>      - OPENFGA_DATASTORE_ENGINE=postgres
>      - OPENFGA_DATASTORE_URI=postgres://postgres:password@postgres:5432/postgres?sslmode=disable
>      - OPENFGA_TRACE_ENABLED=true
>      - OPENFGA_TRACE_SAMPLE_RATIO=1
>      - OPENFGA_TRACE_OTLP_ENDPOINT=otel-collector:4317
>      - OPENFGA_METRICS_ENABLE_RPC_HISTOGRAMS=true

### Reproduction

> Detail the steps taken to reproduce this error, what was expected, and whether this issue can be reproduced consistently or if it is intermittent.
>
> 1. Initialize OpenFgaClient with openfga_sdk.ClientConfiguration parameter api_host=127.0.0.1, credentials method client_credentials
> 2. Invoke method read_authorization_models
> 3. See exception thrown

### Sample Code the Produces Issues

>
> ```
> <code snippet>
> ```

### Backtrace (if applicable)

> ```
> <backtrace>
> ```


### Expected behavior
> A clear and concise description of what you expected to happen.

### Additional context
> Add any other context about the problem here.