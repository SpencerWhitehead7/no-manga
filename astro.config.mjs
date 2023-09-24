import { defineConfig } from "astro/config";

import { buildRedirects } from "./src/buildRedirects";

// https://astro.build/config
export default defineConfig({
  redirects: buildRedirects(),
});
