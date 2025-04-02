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
      $frizzante: "./.frizzante/vite-project/lib",
      $lib: "./lib",
    },
  },
  build: {
    sourcemap: "inline",
    rollupOptions: {
      input: {
        index: "./.frizzante/vite-project/index.html",
      },
    },
  },
});
