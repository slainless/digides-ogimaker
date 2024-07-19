import { Buffer } from 'node:buffer'
import { Payload } from './schema/generated/payload'

interface Parameters extends Omit<Payload, "icon" | "background"> {
  icon: Uint8Array
  background: Uint8Array
}

export function createParameters(parameter: Parameters): Buffer {
  const title = Buffer.from(parameter.title)
  const subtitle = Buffer.from(parameter.subtitle)
  const titleFont = Buffer.from(parameter.titleFont ?? "")
  const subtitleFont = Buffer.from(parameter.subtitleFont ?? "")

  const data = Buffer.allocUnsafe(
    (4 * 6) +
    title.byteLength +
    subtitle.byteLength +
    titleFont.byteLength +
    subtitleFont.byteLength +
    parameter.icon.byteLength +
    parameter.background.byteLength
  )

  data.writeUint32LE(title.byteLength, 0)
  data.writeUint32LE(subtitle.byteLength, 4)
  data.writeUint32LE(parameter.icon.byteLength, 8)
  data.writeUint32LE(parameter.background.byteLength, 12)
  data.writeUint32LE(titleFont.byteLength, 16)
  data.writeUint32LE(subtitleFont.byteLength, 20)

  var l = 24
  l += title.copy(data, l)
  l += subtitle.copy(data, l)
  l += titleFont.copy(data, l)
  l += subtitleFont.copy(data, l)
  l += Buffer.from(parameter.icon).copy(data, l)
  l += Buffer.from(parameter.background).copy(data, l)

  return data
}