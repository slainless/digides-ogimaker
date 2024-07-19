export async function getImage(bucket: R2Bucket, key: string): Promise<Uint8Array | null> {
  const object = await bucket.get(key)
  if (object == null) {
    return null
  }

  const buf = await object.arrayBuffer()
  return new Uint8Array(buf)
}