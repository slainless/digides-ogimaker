import { Buffer } from 'node:buffer'
import { Payload, assertPayload } from '../schema/generated/payload'

export async function decode(key: CryptoKey, data: string): Promise<Payload> {
  const arr = Buffer.from(data, "base64url")
  const iv = arr.subarray(0, 12)
  const encrypted = arr.subarray(12)

  const decrypted = await crypto.subtle.decrypt({
    name: "AES-GCM",
    iv
  }, key, encrypted)

  const parsed = JSON.parse(Buffer.from(decrypted).toString("utf8"))
  return assertPayload(parsed)
}

export async function encode(key: CryptoKey, payload: Payload): Promise<string> {
  const data = Buffer.from(JSON.stringify(payload))

  const iv = Buffer.allocUnsafe(12)
  crypto.getRandomValues(iv)

  const encrypted = await crypto.subtle.encrypt({
    name: "AES-GCM",
    iv
  }, key, data)

  const result = Buffer.allocUnsafe(12 + encrypted.byteLength)
  result.set(iv, 0)
  result.set(new Uint8Array(encrypted), 12)

  return result.toString("base64url")
}