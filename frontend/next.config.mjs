/** @type {import('next').NextConfig} */
const nextConfig = {
  webpack: (config, { isServer }) => {
    // バイナリファイルの扱いにnode-loaderを追加
    config.module.rules.push({
      test: /\.node$/,
      use: 'node-loader',
    });

    // サーバーサイドでのみcanvasを利用する場合はここで設定する
    if (isServer) {
      config.externals = ['canvas', ...(config.externals || [])];
    }

    return config;
  },
};

export default nextConfig;
