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
import ElementPlus from 'unplugin-element-plus/vite'

export default defineConfig({
	resolve: {
		alias: [{
			find: '@',
			replacement: path.resolve(__dirname, 'src'),
		}, ],
	},
	css: {
		preprocessorOptions: {
			scss: {
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
				ElementPlusResolver({
					importStyle: 'sass',
				})
			],
		}),
		Components({
			resolvers: [
				ElementPlusResolver({
					importStyle: 'sass',
				}),
			],
		}),
		ElementPlus({
			useSource: true,
		}),
	]
})
