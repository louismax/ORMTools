<template>
	<div class="banner_box" v-if="editableTabs.length < 1">
		<el-image style="margin: auto; width: 400px;" :src="banner_img" fit="scale-down" />
		<div class="banner_btn">
			<el-button :text="true">
				<svg t="1681490354196" class="icon" viewBox="0 0 1024 1024" version="1.1"
					xmlns="http://www.w3.org/2000/svg" p-id="2602" width="32" height="32">
					<path
						d="M512 12.64c-282.752 0-512 229.216-512 512 0 226.208 146.72 418.144 350.144 485.824 25.6 4.736 35.008-11.104 35.008-24.64 0-12.192-0.48-52.544-0.704-95.328-142.464 30.976-172.512-60.416-172.512-60.416-23.296-59.168-56.832-74.912-56.832-74.912-46.464-31.776 3.52-31.136 3.52-31.136 51.392 3.616 78.464 52.768 78.464 52.768 45.664 78.272 119.776 55.648 148.992 42.56 4.576-33.088 17.856-55.68 32.512-68.48-113.728-12.928-233.28-56.864-233.28-253.024 0-55.904 20-101.568 52.768-137.44-5.312-12.896-22.848-64.96 4.96-135.488 0 0 43.008-13.76 140.832 52.48 40.832-11.36 84.64-17.024 128.16-17.248 43.488 0.192 87.328 5.888 128.256 17.248 97.728-66.24 140.64-52.48 140.64-52.48 27.872 70.528 10.336 122.592 5.024 135.488 32.832 35.84 52.704 81.536 52.704 137.44 0 196.64-119.776 239.936-233.792 252.64 18.368 15.904 34.72 47.04 34.72 94.816 0 68.512-0.608 123.648-0.608 140.512 0 13.632 9.216 29.6 35.168 24.576 203.328-67.776 349.856-259.616 349.856-485.76 0-282.784-229.248-512-512-512z"
						fill="#515151" p-id="2603"></path>
				</svg>
			</el-button>
			<el-button :text="true">
				<svg t="1681490502766" class="icon" viewBox="0 0 1024 1024" version="1.1"
					xmlns="http://www.w3.org/2000/svg" p-id="2960" width="32" height="32">
					<path
						d="M512 1024C229.222 1024 0 794.778 0 512S229.222 0 512 0s512 229.222 512 512-229.222 512-512 512z m259.149-568.883h-290.74a25.293 25.293 0 0 0-25.292 25.293l-0.026 63.206c0 13.952 11.315 25.293 25.267 25.293h177.024c13.978 0 25.293 11.315 25.293 25.267v12.646a75.853 75.853 0 0 1-75.853 75.853h-240.23a25.293 25.293 0 0 1-25.267-25.293V417.203a75.853 75.853 0 0 1 75.827-75.853h353.946a25.293 25.293 0 0 0 25.267-25.292l0.077-63.207a25.293 25.293 0 0 0-25.268-25.293H417.152a189.62 189.62 0 0 0-189.62 189.645V771.15c0 13.977 11.316 25.293 25.294 25.293h372.94a170.65 170.65 0 0 0 170.65-170.65V480.384a25.293 25.293 0 0 0-25.293-25.267z"
						fill="#C71D23" p-id="2961"></path>
				</svg>
			</el-button>
			<el-button :text="true">
				<svg t="1681490628474" class="icon" viewBox="0 0 1024 1024" version="1.1"
					xmlns="http://www.w3.org/2000/svg" p-id="4001" width="32" height="32">
					<path
						d="M388.8 392c9.6-9.8 12.4-10.2 72.8-10.2 54.4 0 56.2 0.2 64.2 4.2 11.6 5.8 16.6 14 16.6 27.2 0 11.8-4.8 20-15.2 26.8-5.6 3.6-9 3.8-62.2 4.2-32.8 0.2-59-0.4-63-1.6-20.6-5.8-28.2-35.4-13.2-50.6z m122.8 189c-107.8 0-111.6 0.4-120.4 8.2-7 6.2-11.4 18.8-10.2 27.8 1.4 9.4 9.6 20.2 18.4 24 4.4 2 28.2 3.4 112.6 2.4l95.8-1.2 18.4-3c18-10.2 21-34.8 6.2-48.8-10.6-9.4-10-9.4-120.8-9.4z m446.8 260.2c-7 56.8-46 100.8-102.2 115-14.4 3.6-19.4 3.8-345.8 3.6-315.6 0-331.8-0.2-344-3.6-16.8-4.4-31.2-11-44.6-20-11.2-7.6-27.8-23.6-34-32.8-7.6-11.2-16.4-30.6-20-44C64.2 846 64 840.6 64 512.6 64 186.4 64 179.4 67.6 165.2 80.2 115.8 119.4 78 170 66.8c14.6-3.2 664.2-3.8 680-0.6 42.4 8.6 75.8 34.2 95.2 72.8 15.4 30.6 14-3 14.6 361.2 0.4 231.6 0 329-1.4 341z m-170.8-370.4c-2.2-10-8.4-19.2-15.4-23-2.2-1.2-16-2.6-31-3.4-24.8-1.2-27.6-1.6-35.6-6.2-12.4-7.2-15.8-15.2-16-36.6 0-40.8-17-78.8-50.6-113-24-24.4-50.6-41-81.2-50.2-7.2-2.2-23.6-3-78.4-3.6-85.8-1-105 0.8-134.2 12.4-54 21.4-92.6 66.8-106.8 124.8-2.6 10.8-3.2 28.4-3.8 128.6-0.8 125.6 0 144.2 8 169 19.4 61.4 74.2 106.8 129.2 116.8 18.4 3.4 244.4 4.2 267.4 1 40.2-5.4 71.8-21.6 101.4-51.8 21.4-21.8 34.8-45.6 43.6-77 6.4-21.8 5.8-176.8 3.4-187.8z"
						fill="#515151" p-id="4002"></path>
				</svg>
			</el-button>
			<el-button :text="true">
				<svg t="1681490715988" class="icon" viewBox="0 0 1024 1024" version="1.1"
					xmlns="http://www.w3.org/2000/svg" p-id="5027" width="32" height="32">
					<path
						d="M695.296 346.112c11.776 0 23.552 1.024 34.816 2.048-31.232-146.432-187.904-254.976-366.592-254.976-199.68 0-363.52 136.192-363.52 308.736 0 99.84 54.272 181.76 145.408 245.248l-36.352 109.056L236.032 692.736c45.568 9.216 81.92 18.432 127.488 18.432 11.264 0 22.528-0.512 33.792-1.536-7.168-24.064-11.264-49.664-11.264-76.288 0.512-158.208 136.704-287.232 309.248-287.232zM497.664 240.64c31.232 0 56.32 25.088 56.32 56.32s-25.088 56.32-56.32 56.32-56.32-25.088-56.32-56.32 25.088-56.32 56.32-56.32zM243.2 353.792c-31.232 0-56.32-25.088-56.32-56.32s25.088-56.32 56.32-56.32 56.32 25.088 56.32 56.32-25.088 56.32-56.32 56.32zM1024.512 630.784c0-145.408-145.408-263.68-308.736-263.68-173.056 0-309.248 118.272-309.248 263.68s136.192 263.68 309.248 263.68c36.352 0 72.704-9.216 109.056-18.432l99.84 54.784-27.136-90.624c72.704-54.784 126.976-127.488 126.976-209.408z m-403.456-40.96c-22.016 0-39.936-17.92-39.936-39.936s17.92-39.936 39.936-39.936 39.936 17.92 39.936 39.936-17.92 39.936-39.936 39.936z m199.68 2.56c-22.016 0-39.936-17.92-39.936-39.936s17.92-39.936 39.936-39.936 39.936 17.92 39.936 39.936-17.92 39.936-39.936 39.936z"
						fill="#69BB64" p-id="5028"></path>
				</svg>
			</el-button>
			<el-button :text="true">
				<svg t="1681490762477" class="icon" viewBox="0 0 1365 1024" version="1.1"
					xmlns="http://www.w3.org/2000/svg" p-id="6054" width="32" height="32">
					<path
						d="M684.00064 607.733333L1282.005973 9.728A133.077333 133.077333 0 0 0 1232.00064 0h-1098.666667c-16.938667 0-33.205333 3.2-48.138666 8.938667l598.805333 598.794666z"
						p-id="6055" fill="#515151"></path>
					<path
						d="M684.00064 728.394667l-664.533333-664.533334A132.48 132.48 0 0 0 0.00064 133.333333v757.333334C0.00064 964.266667 59.733973 1024 133.333973 1024h1098.666667c73.6 0 133.333333-59.733333 133.333333-133.333333v-757.333334a133.013333 133.013333 0 0 0-18.528-67.733333L684.00064 728.394667z"
						p-id="6056" fill="#515151"></path>
				</svg>
			</el-button>

		</div>
	</div>

	<div v-if="editableTabs.length > 0">
		<el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab">
			<el-tab-pane v-for="item in editableTabs" :key="item.name" :label="item.title" :name="item.name">
				<highlightjs class="code_box" language='golang' :code="item.content" />
			</el-tab-pane>
		</el-tabs>
	</div>
</template>

<script setup>
	import banner_img from '@/assets/images/banner.png'
	import {
		ref
	} from 'vue'
	import hljs from 'highlight.js/lib/core';
	import {
		QueryTableFieldList,
	} from '../../wailsjs/go/main/App'


	let hjTag = ref(0)
	let tabIndex = 0
	const editableTabsValue = ref()
	const editableTabs = ref([])
	defineExpose({
		GetTableInfo
	})

	// if (hjTag.value == 0) {
	// 	import('highlight.js/styles/atom-one-dark.css')
	// } else {
	// 	import('highlight.js/styles/atom-one-light.css')
	// }

	const aaa = () => {
		console.log(document.getElementsByTagName('html')[0].dataset)
		if (document.getElementsByTagName('html')[0].dataset.codeTheme == "light") {
			document.getElementsByTagName('html')[0].dataset.codeTheme = "dark"
		} else {
			document.getElementsByTagName('html')[0].dataset.codeTheme = "light"
		}
	}

	const removeTab = (targetName) => {
		const tabs = editableTabs.value
		let activeName = editableTabsValue.value
		if (activeName === targetName) {
			tabs.forEach((tab, index) => {
				if (tab.name === targetName) {
					const nextTab = tabs[index + 1] || tabs[index - 1]
					if (nextTab) {
						activeName = nextTab.name
					}
				}
			})
		}

		editableTabsValue.value = activeName
		editableTabs.value = tabs.filter((tab) => tab.name !== targetName)
	}

	function GetTableInfo(key, dbName, tableName,tbComment) {
		QueryTableFieldList(key, dbName, tableName,tbComment).then(result => {
			if (result.State == true) {
				const newTabName = `${++tabIndex}`
				editableTabs.value.push({
					title: tableName,
					name: newTabName,
					content: result.Data,
				})
				editableTabsValue.value = newTabName
			} else {
				ElMessage.error(result.Message)
			}
		})
	}
</script>

<style lang="scss">
	.banner_box {
		display: flex;
		flex-direction: column;
		padding-top: 10vh;
	}

	.banner_btn {
		margin-top: 8vh;
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;

		& button {
			width: 38px;
		}
	}
	.code_box{
		text-align: left;
		
		& code{
			&::-webkit-scrollbar {
				width: 8px;
				height: 8px;
			}
			
			&::-webkit-scrollbar-track {
				background: rgb(239, 239, 239);
				border-radius: 2px;
			}
			
			&::-webkit-scrollbar-thumb {
				background: #bfbfbf;
				border-radius: 10px;
			}
			
			&::-webkit-scrollbar-thumb:hover {
				background: #7d7d7d;
			}
		}
		
		
	}
</style>
