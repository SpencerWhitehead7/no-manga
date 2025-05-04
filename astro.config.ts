import cloudflare from "@astrojs/cloudflare";
import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  output: "server",
  adapter: cloudflare({
    // TODO: get it working with cloudflare's optimization/CDN service
    imageService: "passthrough",
  }),
  // this is gross but you need it because you can't disable sessions and
  // it nags you to create a cloudflare KV store if you're using the cloudflare adapter
  // "memory" uses an in-memory store, which is probably as close to a noop as I can get
  session: {
    driver: "memory",
  },
});
