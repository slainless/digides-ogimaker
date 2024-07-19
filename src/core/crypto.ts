import { Buffer } from 'node:buffer'

export const decryptCacheKey = Symbol()

export async function generateKey(): Promise<string> {
  const key = await crypto.subtle.generateKey({
    name: "AES-GCM",
    length: 256
  }, true, ['encrypt', 'decrypt'])

  const raw = await crypto.subtle.exportKey("raw", key as CryptoKey)

  return Buffer.from(raw as ArrayBuffer).toString("base64url")
}

export async function consumeKey(key: string): Promise<CryptoKey> {
  const raw = Buffer.from(key, "base64url")
  return crypto.subtle.importKey("raw", raw, {
    name: "AES-GCM",
    length: 256
  }, false, ["encrypt", "decrypt"])
}