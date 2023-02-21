import { defineConfig } from 'vite'
import autoprefixer from 'autoprefixer'
import postcssScss from 'postcss-scss'
import fontMagician from 'postcss-font-magician'

export default defineConfig({
    base: '/static/',
    css: {
        postcss: {
            plugins: [
                autoprefixer,
                postcssScss,
                fontMagician({}),
            ],
        }
    },
    build: {
        manifest: false,
        outDir: './dist',
        rollupOptions: {
            input: {
                fileshare: './src/fileshare.ts',
                upload: './src/upload.ts',
                files: './src/files.ts',
                admin: './src/admin.ts',
                fileshare_css: './src/css/fileshare.scss',
            },
            output: {
                entryFileNames: `js/[name].js`,
                chunkFileNames: `js/[name].js`,
                assetFileNames: (assetInfo) => {
                    const ext = assetInfo.name.split('.').pop()
                    switch (ext) {
                        case 'js':
                            return `js/[name].js`
                        case 'css':
                            return `css/[name].css`
                        case 'png':
                        case 'jpg':
                        case 'jpeg':
                        case 'gif':
                        case 'svg':
                            return `images/[name].[ext]`
                        case 'woff':
                        case 'woff2':
                        case 'ttf':
                            return `fonts/[name].[ext]`
                        default:
                            return `assets/[name].[ext]`
                    }
                }
            }
        },
    },
})