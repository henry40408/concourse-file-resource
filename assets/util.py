#!/usr/bin/env python3

import hashlib


def calculate_checksum(filename, content):
    hasher = hashlib.sha256()
    hasher.update(filename.encode())
    hasher.update(content.encode())
    return hasher.hexdigest()
