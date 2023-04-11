<template>
	<!-- <div class="common-layout" @contextmenu.prevent> -->
	<div class="common-layout">
		<el-container class="container">
			<el-header class="topTitle">
				<Header />
			</el-header>
			<el-container class="MainBox">
				<el-aside class="leftAside">
					<dbTree />
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
		reactive
	} from 'vue'
	import {
		GetUserAppDataPath
	} from '../wailsjs/go/main/App'

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
	}

	.rightMain {
		border-left: 1px solid #999999;
	}
</style>
