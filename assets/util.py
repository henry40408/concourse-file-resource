#!/usr/bin/env python3

import hashlib


def calculate_checksum(filename, content):
    hasher = hashlib.sha256()
    hasher.update(filename)
    hasher.update(content)
    return hasher.hexdigest()
