import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [
    react({
      babel: {
        plugins: [
          [
            'babel-plugin-styled-components',
            {
              displayName: true,
              fileName: true
            }
          ]
        ]
      }
    })
  ],
  server: {
    proxy: {
      "/public": {
        target: "http://localhost:7170",
        changeOrigin: true,
      },
      "/api": {
        target: "http://localhost:5000",
        changeOrigin: true,
      },
    },
    open: true,
    middlewareMode: false,
  },
  build: {
    outDir: 'dist',
  }
})
