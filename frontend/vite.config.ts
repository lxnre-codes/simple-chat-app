import { HmrOptions, ProxyOptions, defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { dirname } from "path";
import { fileURLToPath } from "url";
import checker from "vite-plugin-checker"; // for type checking

const getProxyOptions = (ws = false): ProxyOptions => {
  return {
    target: `http://127.0.0.1:8080`,
    changeOrigin: false,
    secure: true,
    ws,
  };
};

const hmrConfig: HmrOptions = {
  protocol: "ws",
  host: "localhost",
  port: 64999,
  clientPort: 64999,
};

// https://vitejs.dev/config/
export default defineConfig({
  root: dirname(fileURLToPath(import.meta.url)),
  plugins: [react(), checker({ typescript: true })],
  server: {
    host: "localhost",
    port: 3000,
    hmr: hmrConfig,
    proxy: {
      "^/(\\?.*)?$": getProxyOptions(),
      "^/api(/|(\\?.*)?$)": getProxyOptions(true),
    },
  },
});
