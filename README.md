# digides-ogimaker

<p align="center">
  <img src="./assets/example.jpeg" width="80%"/>
</p>

`digides-ogimaker` A.K.A Digital Desa OG Image Maker, A cloudflare worker to generate OpenGraph banner, configured specifically for Digital Desa Profile Website.

Using Hono for edge framework, while Go is used to draw the image.

Live at https://og-image-generator.digital-desa.workers.dev/.

[Click here to view render example.](https://og-image-generator.digital-desa.workers.dev/?d=JfDxtkTsSNXx8lDF4BqMVcHUla4WE59lF6SeUCmRr61FvhaaPrmDfkz41Zt0C1lUKGX0X-ywirXMyPLTW78arco6dyhnnIl4OWgi6g1Evf8rFSh6iMrn1OPB7U7OKAb7-2lLFa5RfmmmAQ2U9wrtIn5EHVFGi8jnrUewocOO5vK2pdSxXaFJnSlM47ULN3fHxuqnYGdi0KT7RjyZ7et1ZAOzA-GD9PZTLhZBlc-H48FOIx9zmLo-E8UBacdm3hHDKXUuociO8e8VJzEaZ1bLhhb8ttipzIzn7wgrQ3PUUjHCar8eCd06jg)

## Difference with previous repository

This project is built on top of WASIP1 instead of javascript WASM. The flow of data is greatly simplified, reading data from `stdin` and write to `stdout`. Since there is no necessity for memory reading yet, so this project WASI code is pretty simple.

Rendering bug on previous repo is also fixed by default by changing the OS to WASIP1.

This project can also be built with `tinygo` easily with no code change.

## Performance

Managed to hit TTFB at average 0.5ms and streamed the response at average ~1s. The biggest bottleneck in this case is the JPEG decoder.

## Future plan

This repo is production ready, with WASM loaded correctly and service running.

Considering that [cloudflare WebGPU API is not ready yet](https://developers.cloudflare.com/durable-objects/api/webgpu/) 
(and even then, it's only supported on Durable Objects), the only choice is to render using CPU. Not that I know how to program using 
WebGPU (I wish I have time to learn it), but it's the best option for graphic drawing in Cloudflare Worker.

I'm thinking of incorporating libvips either via C binding (if at all possible) or by running wasm runtime on Go to load another wasm, don't know if this is even possible.