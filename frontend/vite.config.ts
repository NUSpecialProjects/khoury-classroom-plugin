import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import eslint from "vite-plugin-eslint2";
import { resolve } from "node:path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), eslint({ cache: false })],
  resolve: {
    alias: { "@": resolve(__dirname, "./src") },
  },
  server: {
    port: 3000,
  },
});
