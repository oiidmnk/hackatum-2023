const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      // Proxy all requests starting with /api to your backend server
      '/recipes': {
        target: 'http://localhost:8080', // The backend server where the API requests should be forwarded
        changeOrigin: true, // This is needed if the backend server is on a different domain
      },
    },
  },
});
