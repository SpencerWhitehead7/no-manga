/// <reference types="astro/client" />

interface ImportMetaEnv {
  // only in build files
  readonly VITE_DB_PATH: string;
  // only in astro files
  readonly PUBLIC_DB_PATH: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
