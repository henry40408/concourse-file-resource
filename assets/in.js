#!/usr/bin/env node

const fs = require('fs')

const lib = require('./lib')

async function main () {
  try {
    const payload = await lib.populatePayloadAsync()
    return putFilesInWorkspace(payload)
  } catch (e) {
    console.error(e.message)
    return process.exit(1)
  }
}

function putFilesInWorkspace (input) {
  if (process.argv.length < 3) {
    console.error(`usage: ${process.argv[0]} destination`)
    return process.exit(1)
  }

  // NOTE node abc.js foobar
  // 0: node
  // 1: abc.js
  // 2: foobar
  const destination = process.argv[2]

  const { source = {} } = input
  const { filename, content } = source

  const digest = lib.calculateChecksum(filename, content)

  const stream = fs.createWriteStream(`${destination}/${filename}`)
  stream.write(content)

  console.log(JSON.stringify({ version: { sha256: digest }, metadata: [] }))
}

main()
