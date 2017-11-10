const crypto = require('crypto')

exports.calculateChecksum = function calculateChecksum (filename, content) {
  const hasher = crypto.createHash('sha256')

  if (!filename) {
    throw new Error('filename is required')
  }

  hasher.update(filename)

  if (!content) {
    throw new Error('content is required')
  }

  hasher.update(content)

  return hasher.digest('hex')
}

exports.populatePayloadAsync = function populatePayloadAsync () {
  return new Promise((resolve, reject) => {
    let payload = ''

    process.stdin.on('data', data => {
      payload += data
    })

    process.stdin.on('end', () => resolve(JSON.parse(payload)))
  })
}
