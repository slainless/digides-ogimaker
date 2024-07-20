import { Buffer } from 'node:buffer'
import { Payload } from './schema/generated/payload'

interface Parameters extends Omit<Payload, "icon" | "background"> {
  icon: Uint8Array
  background: Uint8Array
}

export function createParameters(parameter: Parameters, quality: number): Buffer {
  const q = isNaN(quality) ? 75 : Math.max(0, Math.min(100, quality))

  const title = Buffer.from(parameter.title)
  const subtitle = Buffer.from(parameter.subtitle)
  const titleFont = Buffer.from(parameter.titleFont ?? "")
  const subtitleFont = Buffer.from(parameter.subtitleFont ?? "")

  const data = Buffer.allocUnsafe(
    1 + // quality
    (4 * 6) + // 4 bytes per parameter size x 6 parameters
    title.byteLength +
    subtitle.byteLength +
    titleFont.byteLength +
    subtitleFont.byteLength +
    parameter.icon.byteLength +
    parameter.background.byteLength
  )

  var l = 0
  l = data.writeUint8(q, l)
  l = data.writeUint32LE(title.byteLength, l)
  l = data.writeUint32LE(subtitle.byteLength, l)
  l = data.writeUint32LE(parameter.icon.byteLength, l)
  l = data.writeUint32LE(parameter.background.byteLength, l)
  l = data.writeUint32LE(titleFont.byteLength, l)
  l = data.writeUint32LE(subtitleFont.byteLength, l)
  l += title.copy(data, l)
  l += subtitle.copy(data, l)
  l += titleFont.copy(data, l)
  l += subtitleFont.copy(data, l)
  l += Buffer.from(parameter.icon).copy(data, l)
  l += Buffer.from(parameter.background).copy(data, l)

  return data
}