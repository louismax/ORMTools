<template>
	<!-- <div class="common-layout" @contextmenu.prevent> -->
	<div class="common-layout">
		<el-container class="container">
			<el-header class="topTitle">
				<Header ref="headerRef" @GetServerList="GetServerListToTree" />
			</el-header>
			<el-container class="MainBox">
				<el-aside class="leftAside">
					<dbTree ref="dbTreeRef" @openServerConfigEdit="openServerConfigEdit" />
				</el-aside>
				<DragAdjustWidth />
				<el-main class="rightMain">Main
					<div id="result" class="result">{{ data.resultText }}</div>
				</el-main>
			</el-container>
		</el-container>
	</div>
</template>

<script setup>
	import DragAdjustWidth from './components/DragAdjustWidth.vue'
	import Home from './components/Home.vue'
	import Header from './components/Header.vue'
	import dbTree from './components/dbTree.vue'
	import {
		ref,
		reactive
	} from 'vue'
	import {
		GetUserAppDataPath
	} from '../wailsjs/go/main/App'

	const dbTreeRef = ref()
	const headerRef = ref()

	const GetServerListToTree = () => {
		dbTreeRef.value.GetServerList()
	}

	const openServerConfigEdit = (key) => {
		headerRef.value.OpenConfigEdit(key)
	}


	const data = reactive({
		name: "",
		resultText: "",
	})
	greet();

	function greet() {
		GetUserAppDataPath().then(result => {
			data.resultText = result
		})
	}
</script>

<style>
	.container {
		display: flex;
		flex-flow: column nowrap;
		overflow: hidden;
		height: 100vh;
	}

	.topTitle {
		height: 45px;
		line-height: 45px;
		min-height: 45px;
		max-height: 45px;
		padding: 0;
		border-bottom: 1px solid #999999;
	}

	.MainBox {
		height: 95vh;
	}

	.leftAside {
		min-width: 320px;
		margin-bottom: 10px;
		
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

	.rightMain {
		border-left: 1px solid #999999;
	}
</style>
