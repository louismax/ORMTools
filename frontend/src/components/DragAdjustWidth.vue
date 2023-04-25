<template>
	<div class="resize" title="收缩侧边栏">⋮</div>
</template>
<script setup>
	import {
		ref,
		onMounted,
		getCurrentInstance
	} from 'vue';

	const divDom = ref(null);

	onMounted(() => {
		dragControllerDiv();
	});

	function dragControllerDiv() {
		var resize = document.getElementsByClassName("resize");
		var left = document.getElementsByClassName("leftAside");
		var mid = document.getElementsByClassName("rightMain");
		var box = document.getElementsByClassName("MainBox");
		var treeSpan = document.getElementsByClassName('treeNodetitleSpan');
		let that = this;
		for (let i = 0; i < resize.length; i++) {
			// 鼠标按下事件
			resize[i].onmousedown = function(e) {
				//颜色改变提醒
				resize[i].style.background = "#818181";
				var startX = e.clientX;
				resize[i].left = resize[i].offsetLeft;
				compatibleStyle(mid[0], "none");
				// 鼠标拖动事件
				document.onmousemove = function(e) {
					var endX = e.clientX;
					var moveLen = resize[i].left + (endX -
						startX); // （endx-startx）=移动的距离。resize[i].left+移动的距离=左边区域最后的宽度
					var maxT = box[i].clientWidth - resize[i].offsetWidth; // 容器宽度 - 左边区域的宽度 = 右边区域的宽度
					if (moveLen < 100) moveLen = 100; // 左边区域的最小宽度为32px
					if (moveLen > maxT - 520) moveLen = maxT - 520; //右边区域最小宽度为150px
					resize[i].style.left = moveLen; // 设置左侧区域的宽度

					for (let j = 0; j < left.length; j++) {
						left[j].style.width = moveLen + "px";
						mid[j].style.width =
							box[i].clientWidth - moveLen - 10 + "px";
					}

					for (let j = 0; j < treeSpan.length; j++) {
						if(treeSpan[j].parentNode.parentNode.clientWidth >=330){
							treeSpan[j].className = "treeNodetitleSpan"
						}else{
							if(treeSpan[j].dataset.iscurrent == "true"){
								console.log(treeSpan[j].dataset)
								treeSpan[j].className = "treeNodetitleSpan tree-node-title-span"
							}
						}
					}

				};
				// 鼠标松开事件
				document.onmouseup = function(evt) {
					compatibleStyle(mid[0], "auto");
					//颜色恢复
					resize[i].style.background = "#d6d6d6";
					document.onmousemove = null;
					document.onmouseup = null;
					resize[i].releaseCapture && resize[i]
						.releaseCapture(); //当你不在需要继续获得鼠标消息就要应该调用ReleaseCapture()释放掉
				};
				resize[i].setCapture && resize[i].setCapture(); //该函数在属于当前线程的指定窗口里设置鼠标捕获
				return false;
			};
		}
	}

	function compatibleStyle(node, info) {
		if (node.getElementsByTagName("iframe").length) {
			node.getElementsByTagName(
				"iframe"
			)[0].style.pointerEvents = info;
		}
		if (node.getElementsByTagName("embed").length) {
			node.getElementsByTagName(
				"embed"
			)[0].style.pointerEvents = info;
		}
	}
</script>
<style scoped>
	/*拖拽区div样式*/
	.resize {
		cursor: col-resize;
		left: 5px;
		position: relative;
		top: 45vh;
		background-color: #d6d6d6;
		border-radius: 5px;
		margin-top: -10px;
		width: 10px;
		height: 40px;
		line-height: 40px;
		background-size: cover;
		background-position: center;
		/*z-index: 99999;*/
		font-size: 32px;
		color: white;
	}

	/*拖拽区鼠标悬停样式*/
	.resize:hover {
		color: #444444;
	}
</style>
