#!/usr/bin/env node

const lib = require('./lib')

async function main () {
  try {
    const payload = await lib.populatePayloadAsync()
    const { source = {}, version } = payload
    const { filename, content } = source

    const digest = lib.calculateChecksum(filename, content)

    const versions = []
    if (version) {
      versions.push(version)
    }
    versions.push({ sha256sum: digest })

    console.log(JSON.stringify(versions))
  } catch (e) {
    console.error(e.message)
    return process.exit(1)
  }
}

main()
