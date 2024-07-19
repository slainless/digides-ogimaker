declare const navigator: import('@cloudflare/workers-types/experimental').Navigator

declare interface Env {
  Bindings: {
    PAYLOAD_ENCRYPTION_SECRET: string
    R2_ASSETS: R2Bucket
    IMAGE_QUALITY: string
    [K: string]: undefined | any
  }
}

declare module "*.wasm" {
  const url: ArrayBuffer
  export default url
}

declare const Go: any