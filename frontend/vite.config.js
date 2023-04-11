import * as path from 'path';
import {
	defineConfig
} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {
	ElementPlusResolver
} from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig({
	resolve: {
		alias: {
			'@': path.resolve(__dirname, 'src'),
			"_C": path.resolve(__dirname, "src/components"),
			"views": path.resolve(__dirname, "src/views")
		}
	},
	css: {
		preprocessorOptions: {
			scss: {
				// 自定义的主题色
				additionalData: `@use "@/styles/element/index.scss" as *;`,
			},
		},
	},

	plugins: [
		vue(),
		AutoImport({
			// Auto import functions from Vue, e.g. ref, reactive, toRef...
			// 自动导入 Vue 相关函数，如：ref, reactive, toRef 等
			imports: ['vue'],

			resolvers: [
				ElementPlusResolver(),
			],
		}),
		Components({
			resolvers: [
				ElementPlusResolver({
					importStyle: "sass",
				}),
			],
		}),
	]
})
