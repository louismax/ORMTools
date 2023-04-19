<template>
	<!-- <div class="common-layout" @contextmenu.prevent> -->
	<div class="common-layout">
		<el-container class="container">
			<el-header class="topTitle">
				<Header ref="headerRef" @GetServerList="GetServerListToTree" />
			</el-header>
			<el-container class="MainBox">
				<el-aside class="leftAside">
					<dbTree ref="dbTreeRef" @openServerConfigEdit="openServerConfigEdit"
						@GetTableInfo="GetTableInfoByToHome" />
				</el-aside>
				<DragAdjustWidth />
				<el-main class="rightMain">
					<home ref="homeRef" />
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
	import { useStore } from 'vuex';
	import {
		ref,
		reactive
	} from 'vue'
	import {
		GetUserConfig
	} from '../wailsjs/go/main/App'

	const store = useStore();
	const dbTreeRef = ref()
	const headerRef = ref()
	const homeRef = ref()

GetUserCongifInfo();
	const GetServerListToTree = () => {
		dbTreeRef.value.GetServerList()
	}

	const openServerConfigEdit = (key) => {
		headerRef.value.OpenConfigEdit(key)
	}

	const GetTableInfoByToHome = (key, dbName, tableName, tbComment) => {
		homeRef.value.GetTableInfo(key, dbName, tableName, tbComment)
	}

	function GetUserCongifInfo() {
	  GetUserConfig().then(result => {
		  console.log(result)
	  	if (result.State == true){
			if(result.Data.WindowTheme == "SystemDefault"){
				window.runtime.WindowSetSystemDefaultTheme();
			}else if(result.Data.WindowTheme == "Light"){
				window.runtime.WindowSetLightTheme();
			}else if(result.Data.WindowTheme == "Dark"){
				window.runtime.WindowSetDarkTheme();
			}
			
			document.getElementsByTagName('html')[0].dataset.codeTheme = result.Data.CodeTheme
			
			
			
			store.commit('setUserConfig', result.Data)
		}
	  })
	}
</script>

<style lang="scss">
	@use "sass:meta";


	html[data-code-theme="a11y-dark"] {
		@include meta.load-css("highlight.js/styles/a11y-dark.css");
	}
	html[data-code-theme="a11y-light"] {
		@include meta.load-css("highlight.js/styles/a11y-light.css");
	}
	html[data-code-theme="agate"] {
		@include meta.load-css("highlight.js/styles/agate.css");
	}
	html[data-code-theme="an-old-hope"] {
		@include meta.load-css("highlight.js/styles/an-old-hope.css");
	}
	html[data-code-theme="androidstudio"] {
		@include meta.load-css("highlight.js/styles/androidstudio.css");
	}
	html[data-code-theme="atom-one-dark"] {
		@include meta.load-css("highlight.js/styles/atom-one-dark.css");
	}
	html[data-code-theme="atom-one-light"] {
		@include meta.load-css("highlight.js/styles/atom-one-light.css");
	}
	html[data-code-theme="atelier-cave"] {
		@include meta.load-css("highlight.js/styles/base16/atelier-cave.css");
	}
	html[data-code-theme="paraiso"] {
		@include meta.load-css("highlight.js/styles/base16/paraiso.css");
	}
	html[data-code-theme="github"] {
		@include meta.load-css("highlight.js/styles/github.css");
	}
	html[data-code-theme="github-dark"] {
		@include meta.load-css("highlight.js/styles/github-dark.css");
	}
	html[data-code-theme="github-dark-dimmed"] {
		@include meta.load-css("highlight.js/styles/github-dark-dimmed.css");
	}
	html[data-code-theme="gradient-dark"] {
		@include meta.load-css("highlight.js/styles/gradient-dark.css");
	}
	html[data-code-theme="gradient-light"] {
		@include meta.load-css("highlight.js/styles/gradient-light.css");
	}
	html[data-code-theme="vs"] {
		@include meta.load-css("highlight.js/styles/vs.css");
	}
	html[data-code-theme="vs2015"] {
		@include meta.load-css("highlight.js/styles/vs2015.css");
	}
	
	

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
