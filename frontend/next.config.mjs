/** @type {import('next').NextConfig} */
const nextConfig = {
    webpack: (config, { isServer }) => {
        if (!isServer) {
            // クライアントサイドでのビルド時には`.node`ファイルを無視
            config.externals = config.externals || [];
            config.externals.push('.node');
        } else {
            // サーバーサイドでのビルド時に`node-loader`を適用
            config.module.rules.push({
                test: /\.node$/,
                use: 'node-loader'
            });
        }
        return config;
    }
};

export default nextConfig;
