import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import eslint from 'vite-plugin-eslint';

// https://vite.dev/config/
export default defineConfig(({ mode }) => ({
  plugins: [
    react(),
    eslint({
      lintOnStart: true,
      failOnError: mode === 'production',
    }),
  ],
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8044',
        changeOrigin: true,
        secure: false,
      },
    },
  },
}));
