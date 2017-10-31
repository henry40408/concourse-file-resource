#!/usr/bin/env python3

import hashlib
import json
import sys

from util import calculate_checksum

payload = json.loads(sys.stdin.read().strip())
source = payload.get("source", {})

filename = source.get("filename", "")
content = source.get("content", "")

version = [
    {"sha256sum": calculate_checksum(filename, content)}
]

sys.stdout.write(json.dumps(version))
