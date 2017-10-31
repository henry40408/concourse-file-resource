#!/usr/bin/env python3

import json
import sys

from util import calculate_checksum

if len(sys.argv) < 2:
    print("usage: {0} destination".format(sys.argv[0]))
    sys.exit(1)

destination = sys.argv[1]
payload = json.loads(sys.stdin.read())
source = payload.get("source", {})

filename = source.get("filename", "")

if not filename:
    print("source.filename is required")
    sys.exit(1)

content = source.get("content", "")

with open("{0}/{1}".format(destination, filename), "wb") as f:
    f.write(content)
    version = {
        "version": {"sha256sum": calculate_checksum(filename, content)}
    }
    print(json.dumps(version))
