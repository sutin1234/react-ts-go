/** @type {import('next').NextConfig} */
const nextConfig = {
    output: 'export',
    distDir: 'build',
    env: {
        BASE_URL:"http://localhost",
        PORT:"8000",
        API_FREFIX:"api",
        API_VERSION:"v1"
    }
}

module.exports = nextConfig
