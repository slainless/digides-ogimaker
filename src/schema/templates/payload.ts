import { createAssert } from 'typia'

export interface Payload {
  title: string
  subtitle: string
  icon: string
  background: string
  titleFont?: string
  subtitleFont?: string
}

export const assertPayload = createAssert<Payload>()