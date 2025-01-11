import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import { fileURLToPath } from 'url'

// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
      $lib: './lib',
    },
  },
  build: {
    rollupOptions: {
      input: {
        frozen: fileURLToPath(new URL('./frizzante/vite-project/index.spa.html', import.meta.url)),
        index: fileURLToPath(new URL('./frizzante/vite-project/index.html', import.meta.url)),
      },
    },
  },
});
