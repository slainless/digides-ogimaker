import { WASI } from "@cloudflare/workers-wasi"
import { Hono } from 'hono'
import drawerWasm from '../build/drawer.wasm'
import { createParameters } from './parameter'
import { memo } from './core/memo'
import { consumeKey, decryptMemo } from './core/crypto'
import { assertPayload } from './schema/generated/payload'
import { TypeGuardError } from 'typia'
import { decode } from './core/payload'
import { getImage } from './core/r2'

const RespImageNotFound = (path: string) => new Response("Image not found at " + path, { status: 400 })

const app = new Hono<Env>()

app.get('/', async (c) => {
  const decryptionKey = await memo(decryptMemo, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET))
  const d = c.req.query('d')
  if (!d) return new Response("Empty Payload", { status: 400 })

  let stdin: Uint8Array
  try {
    const payload = await decode(decryptionKey, d)
    assertPayload(payload)

    const [icon, background] = await Promise.all([
      getImage(c.env.R2_ASSETS, payload.icon),
      getImage(c.env.R2_ASSETS, payload.background),
    ])

    if (icon == null) return RespImageNotFound(payload.icon)
    if (background == null) return RespImageNotFound(payload.background)

    stdin = new Uint8Array(createParameters({
      title: payload.title,
      subtitle: payload.subtitle,
      icon,
      background,
    }))
  } catch (e) {
    if (e instanceof TypeGuardError)
      return new Response(e.message, { status: 400 })

    return new Response(e?.toString() ?? JSON.stringify(e), { status: 500 })
  }

  const stdout = new TransformStream()
  const writer = new WritableStream({
    write(chunk: Uint8Array, controller) {
      if (chunk[0] === 0) {
        console.log(chunk.slice(1))
      } else {
        throw new Response("Drawer exited caused by internal error", { status: 500 })
      }
    },
  })

  const wasi = new WASI({
    args: [],
    stdout: stdout.writable,
    stderr: writer,
    stdin: new Blob([stdin]).stream(),
    returnOnExit: true,
  })


  const instance = new WebAssembly.Instance(drawerWasm, {
    wasi_snapshot_preview1: wasi.wasiImport,
  })

  c.executionCtx.waitUntil(wasi.start(instance))
  return new Response(stdout.readable, {
    headers: {
      'Content-Type': 'image/jpeg',
    },
  })
})

export default app
