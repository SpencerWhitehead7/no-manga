import cloudflare from "@astrojs/cloudflare";
import { defineConfig } from "astro/config";

import { buildRedirects } from "./src/buildRedirects";

// https://astro.build/config
export default defineConfig({
  output: "hybrid",
  adapter: cloudflare({
    runtime: {
      mode: "local",
      type: "pages",
      persistTo: ".wrangler/state/v3",
      bindings: {
        DB: {
          type: "d1",
        },
      },
    },
  }),
  redirects: buildRedirects(),
});
