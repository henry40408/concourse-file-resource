# Concourse File Resource

[![CircleCI](https://circleci.com/gh/henry40408/concourse-file-resource.svg?style=shield)](https://circleci.com/gh/henry40408/concourse-file-resource)
[![codecov](https://codecov.io/gh/henry40408/concourse-file-resource/branch/master/graph/badge.svg)](https://codecov.io/gh/henry40408/concourse-file-resource)
[![Go Report Card](https://goreportcard.com/badge/github.com/henry40408/concourse-file-resource)](https://goreportcard.com/report/github.com/henry40408/concourse-file-resource)
[![Docker Repository on Quay](https://quay.io/repository/henry40408/concourse-file-resource/status "Docker Repository on Quay")](https://quay.io/repository/henry40408/concourse-file-resource)
[![GitHub release](https://img.shields.io/github/release/henry40408/concourse-file-resource.svg)](https://github.com/henry40408/concourse-file-resource)
[![license](https://img.shields.io/github/license/henry40408/concourse-file-resource.svg)](https://github.com/henry40408/concourse-file-resource)
![stability-stable](https://img.shields.io/badge/stability-stable-green.svg)

> Resource to put confidential file via payload

## Source Configuration

- `filename`: string, filename of confidential file in workspace
- `content`: string, content of confidential file

## Behavior

`out` command is not available.

### `check`

Print SHA256 checksum of `filename` and `content` on standard output.

### `in`

Create a file called as `filename` and put it in destination directory.

**CAUTION**: Any files with the same name in destination directory would be overwritten.

## Examples

```yaml
---
resource_types:
  - name: file
    type: docker-image
    source:
      repository: quay.io/henry40408/concourse-file-resource

resources:
  - name: confidential-file
    type: file
    source:
      filename: id_rsa
      content: |
        -----BEGIN RSA PRIVATE KEY-----
        MIIJKAIBAAKCAgEA0aDJt9E+v38csI3+FeyiHPU8kmeF7HeSXb62cjOoKcpiwq+L
        ... something very confidential ...
        goHZOH8rALOXE7nUZeYh2RbPE+JYdSQvFEJmjh0EEAni3d6KJXOvm0NTiTk=
        -----END RSA PRIVATE KEY-----

jobs:
  - name: download-private-ssh-key
    plan:
      - get: confidential-file
      # now you can see the file put in confidential-file workspace as id_rsa
```

## License

MIT
