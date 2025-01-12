import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    svelte({
      compilerOptions: {
        css: "injected",
      },
    }),
  ],
  resolve: {
    alias: {
      $lib: './lib',
    },
  },
  build: {
    rollupOptions: {
      input: {
        index: "./.frizzante/vite-project/index.html",
      },
    },
  },
});
