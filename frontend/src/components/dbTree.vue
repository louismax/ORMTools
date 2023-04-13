<template>
	<el-tree :data="dataSource" node-key="key" empty-text="暂无数据库配置" :highlight-current="true"
		:expand-on-click-node="false" :check-on-click-node="true" @node-click="handleNodeClick"
		@current-change="currentNodeChange" default-expand-all icon-class="none" ref="treeForm">
		<template #default="{ node, data }">
			<span class="custom-tree-node">
				<span class="tree-node-title">
					<svg v-if="data.obj_type == 'connect'" style="width: 16px; height: 16px;padding-right: 4px;"
						t="1680792160481" class="icon" viewBox="0 0 1024 1024" version="1.1"
						xmlns="http://www.w3.org/2000/svg" p-id="3925" width="200" height="200">
						<path
							d="M734.608696 624.788406c-26.713043 0-46.005797-17.808696-46.005797-46.005797s17.808696-46.005797 46.005797-46.005797c26.713043 0 46.005797 17.808696 46.005797 46.005797s-19.292754 46.005797-46.005797 46.005797z"
							p-id="3926" :fill="data.conState?'#9254bf':'#707070'"></path>
						<path
							d="M838.492754 0h-652.985508c-25.228986 0-44.521739 19.292754-44.521739 44.521739v934.956522c0 25.228986 19.292754 44.521739 44.521739 44.521739h652.985508c25.228986 0 44.521739-19.292754 44.521739-44.521739v-934.956522c0-25.228986-19.292754-44.521739-44.521739-44.521739z m-289.391305 608.463768c-20.776812 0-37.101449-16.324638-37.101449-37.101449s16.324638-37.101449 37.101449-37.101449 37.101449 16.324638 37.10145 37.101449-16.324638 37.101449-37.10145 37.101449z m185.507247 44.521739c-40.069565 0-74.202899-32.649275-74.202899-74.202898s32.649275-74.202899 74.202899-74.202899c40.069565 0 74.202899 32.649275 74.202898 74.202899s-34.133333 74.202899-74.202898 74.202898z m74.202898-207.768116h-593.623188v-148.405797h593.623188v148.405797z m0-222.608695h-593.623188v-148.405797h593.623188v148.405797z"
							p-id="3927" :fill="data.conState?'#9254bf':'#707070'"></path>
					</svg>
					<svg v-if="data.obj_type == 'db'" style="width: 16px; height: 16px;padding-right: 4px;"
						t="1680796288505" class="icon" viewBox="0 0 1024 1024" version="1.1"
						xmlns="http://www.w3.org/2000/svg" p-id="5867" width="200" height="200">
						<path
							d="M512 384c-229.8 0-416-57.3-416-128v256c0 70.7 186.2 128 416 128s416-57.3 416-128V256c0 70.7-186.2 128-416 128z"
							fill="#707070" p-id="5868"></path>
						<path
							d="M512 704c-229.8 0-416-57.3-416-128v256c0 70.7 186.2 128 416 128s416-57.3 416-128V576c0 70.7-186.2 128-416 128zM512 320c229.8 0 416-57.3 416-128S741.8 64 512 64 96 121.3 96 192s186.2 128 416 128z"
							fill="#707070" p-id="5869"></path>
					</svg>
					<svg v-if="data.obj_type == 'table'" style="width: 16px; height: 16px;padding-right: 4px;"
						t="1680796372308" class="icon" viewBox="0 0 1024 1024" version="1.1"
						xmlns="http://www.w3.org/2000/svg" p-id="7260" width="200" height="200">
						<path
							d="M329.142857 786.285714v-109.714285q0-8-5.142857-13.142858t-13.142857-5.142857H128q-8 0-13.142857 5.142857t-5.142857 13.142858v109.714285q0 8 5.142857 13.142857t13.142857 5.142858h182.857143q8 0 13.142857-5.142858t5.142857-13.142857z m0-219.428571V457.142857q0-8-5.142857-13.142857t-13.142857-5.142857H128q-8 0-13.142857 5.142857t-5.142857 13.142857v109.714286q0 8 5.142857 13.142857t13.142857 5.142857h182.857143q8 0 13.142857-5.142857t5.142857-13.142857z m292.571429 219.428571v-109.714285q0-8-5.142857-13.142858t-13.142858-5.142857H420.571429q-8 0-13.142858 5.142857t-5.142857 13.142858v109.714285q0 8 5.142857 13.142857t13.142858 5.142858h182.857142q8 0 13.142858-5.142858t5.142857-13.142857zM329.142857 347.428571V237.714286q0-8-5.142857-13.142857t-13.142857-5.142858H128q-8 0-13.142857 5.142858t-5.142857 13.142857v109.714285q0 8 5.142857 13.142858t13.142857 5.142857h182.857143q8 0 13.142857-5.142857t5.142857-13.142858z m292.571429 219.428572V457.142857q0-8-5.142857-13.142857t-13.142858-5.142857H420.571429q-8 0-13.142858 5.142857t-5.142857 13.142857v109.714286q0 8 5.142857 13.142857t13.142858 5.142857h182.857142q8 0 13.142858-5.142857t5.142857-13.142857z m292.571428 219.428571v-109.714285q0-8-5.142857-13.142858t-13.142857-5.142857h-182.857143q-8 0-13.142857 5.142857t-5.142857 13.142858v109.714285q0 8 5.142857 13.142857t13.142857 5.142858h182.857143q8 0 13.142857-5.142858t5.142857-13.142857z m-292.571428-438.857143V237.714286q0-8-5.142857-13.142857t-13.142858-5.142858H420.571429q-8 0-13.142858 5.142858t-5.142857 13.142857v109.714285q0 8 5.142857 13.142858t13.142858 5.142857h182.857142q8 0 13.142858-5.142857t5.142857-13.142858z m292.571428 219.428572V457.142857q0-8-5.142857-13.142857t-13.142857-5.142857h-182.857143q-8 0-13.142857 5.142857t-5.142857 13.142857v109.714286q0 8 5.142857 13.142857t13.142857 5.142857h182.857143q8 0 13.142857-5.142857t5.142857-13.142857z m0-219.428572V237.714286q0-8-5.142857-13.142857t-13.142857-5.142858h-182.857143q-8 0-13.142857 5.142858t-5.142857 13.142857v109.714285q0 8 5.142857 13.142858t13.142857 5.142857h182.857143q8 0 13.142857-5.142857t5.142857-13.142858z m73.142857-182.857142v621.714285q0 37.714286-26.857142 64.571429t-64.571429 26.857143H128q-37.714286 0-64.571429-26.857143t-26.857142-64.571429V164.571429q0-37.714286 26.857142-64.571429t64.571429-26.857143h768q37.714286 0 64.571429 26.857143t26.857142 64.571429z"
							p-id="7261" fill="#707070">
						</path>
					</svg>
					<span>
						{{node.label}}
					</span>
				</span>
				<span v-if="node.isCurrent && data.obj_type == 'connect'">
					<el-tooltip v-if="!data.conState" content="打开连接" placement="bottom" effect="light">
						<el-button :text="true" style="width: 32px;height: 27px;margin: 0;" @click.stop="openDB(data)">
							<svg style="width: 16px; height: 16px;" t="1680797183435" class="icon"
								viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1267"
								width="200" height="200">
								<path
									d="M594.601737 716.593828a10.677965 10.677965 0 0 0-14.993033 0l-154.903626 154.903625c-71.747147 71.67401-192.78846 79.353368-271.995554 0-79.353368-79.353368-71.67401-200.248408 0-271.995554L307.613149 444.67141a10.677965 10.677965 0 0 0 0-15.06617l-53.02414-53.097277a10.677965 10.677965 0 0 0-15.066169 0L84.619214 531.411589a288.231912 288.231912 0 0 0 0 407.956764 288.305049 288.305049 0 0 0 407.956764 0l154.903625-154.903626a10.677965 10.677965 0 0 0 0-15.06617l-52.804729-52.804729z m344.839753-631.974614a288.158775 288.158775 0 0 0-407.956764 0l-155.049899 154.903626a10.677965 10.677965 0 0 0 0 15.066169l52.951003 52.951004c4.095658 4.095658 10.897375 4.095658 14.993033 0l154.976762-154.976763c71.67401-71.67401 192.78846-79.280232 271.922418 0 79.353368 79.353368 71.747147 200.321545 0 272.068691L716.374417 579.462431a10.677965 10.677965 0 0 0 0 15.066169l53.097277 53.02414c4.095658 4.168794 10.970512 4.168794 15.06617 0l154.903626-154.903626a288.670733 288.670733 0 0 0 0-408.103037zM642.798852 325.824199a10.677965 10.677965 0 0 0-15.06617 0L325.751062 627.659546a10.677965 10.677965 0 0 0 0 15.066169l52.80473 52.80473c4.168794 4.095658 10.970512 4.095658 15.06617 0l301.835346-301.908483a10.677965 10.677965 0 0 0 0-15.06617l-52.658456-52.731593z"
									fill="#707070" p-id="1268"></path>
							</svg>
						</el-button>
					</el-tooltip>
					<el-tooltip v-if="data.conState" content="断开连接" placement="bottom" effect="light">
						<el-button :text="true" style="width: 32px;height: 27px;margin: 0;" @click.stop="closeDB(data)">
							<svg style="width: 16px; height: 16px;" t="1680798337609" class="icon"
								viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1529"
								width="200" height="200">
								<path
									d="M939.52 84.626286a288.182857 288.182857 0 0 0-407.990857 0L402.285714 213.796571l68.022857 68.022858 129.170286-129.243429c71.753143-71.68 192.804571-79.286857 272.018286 0 79.286857 79.36 71.68 200.338286 0 272.091429l-129.170286 129.170285 68.096 68.096 129.170286-129.170285a288.841143 288.841143 0 0 0-0.073143-408.137143z m-514.779429 787.017143c-71.753143 71.68-192.804571 79.286857-272.018285 0-79.36-79.433143-71.68-200.411429 0-272.091429l129.170285-129.170286-68.096-68.169143-129.170285 129.243429a288.256 288.256 0 0 0 0 407.990857 288.329143 288.329143 0 0 0 407.917714 0l129.243429-129.243428-68.022858-68.022858-129.024 129.462858zM176.493714 108.544a10.678857 10.678857 0 0 0-15.067428 0l-52.809143 52.809143a10.678857 10.678857 0 0 0 0 15.067428L847.725714 915.675429c4.169143 4.169143 10.971429 4.169143 15.067429 0l52.809143-52.736a10.678857 10.678857 0 0 0 0-15.067429L176.493714 108.617143z"
									fill="#ff7171" p-id="1530"></path>
							</svg>
						</el-button>
					</el-tooltip>

					<el-tooltip v-if="data.conState" content="重载该数据库库表结构" placement="bottom" effect="light">
						<el-button :text="true" style="width: 27px;height: 27px;margin: 0;" @click.stop="refresh(data)">
							<el-icon size="18">
								<Refresh />
							</el-icon>
						</el-button>
					</el-tooltip>

					<el-tooltip :content="data.conState?'只允许断开连接时进行编辑':'编辑连接'" placement="bottom" effect="light">
						<div style="display: inline-block;">
							<el-button :disabled="data.conState" :text="true"
								style="width: 27px;height: 27px;margin: 0;" @click.stop="editDB(data)">
								<el-icon size="18">
									<Edit />
								</el-icon>
							</el-button>
						</div>
					</el-tooltip>

					<el-tooltip :content="data.conState?'只允许断开连接时删除配置':'删除连接配置'" placement="bottom" effect="light">
						<div style="display: inline-block;">
							<el-button :disabled="data.conState" :text="true"
								style="width: 27px;height: 27px;margin: 0;" @click.stop="deleteDB(data)">
								<el-icon size="18" color="#ff7171">
									<Delete />
								</el-icon>
							</el-button>
						</div>
					</el-tooltip>



					<!-- <a style="margin-left: 8px" @click="remove(node, data)"> Delete </a> -->
				</span>
			</span>
		</template>
	</el-tree>

</template>

<script setup>
	import {
		ref,
		// defineExpose,
		// defineEmits,
	} from 'vue';
	import {
		Coin,
		Delete,
		Edit,
		Refresh,
	} from '@element-plus/icons-vue'

	import {
		GetServerConfigList,
		DeleteServerConfig,
	} from '../../wailsjs/go/main/App'


	const emit = defineEmits(['openServerConfigEdit'])


	defineExpose({
		GetServerList
	})

	const treeForm = ref(null);
	let selectId = ref(0)
	let dataSource = ref()

	GetServerList();

	function GetServerList() {
		GetServerConfigList().then(result => {
			if (result.State == true) {
				dataSource.value = result.Data
			} else {
				ElMessage.error(result.Message)
			}
		})

	}

	const currentNodeChange = (data, node) => {
		// console.log("节点选中状态变化", data, node)
	}

	const handleNodeClick = (data, node, tn, e) => {
		//console.log(node)
	}

	const openDB = (obj) => {
		console.log("打开连接按钮点击事件:", obj)
		if(obj.has_record_pwd == false){
			ElMessageBox.prompt('请输入数据库服务器登录用户密码以继续操作:', '身份验证', {
			    confirmButtonText: '确认',
			    cancelButtonText: '取消',
				 inputPattern:/^.+$/,
			    inputErrorMessage: '密码不允许为空',
			  })
			    .then(({ value }) => {
					obj.children = [{
						children: null,
						conState: true,
						key: "dead4be5-110d-4fde-9e65-c52f85b480b9111",
						label: "base_basic",
						obj_type: "db",
					}]
			      // ElMessage({
			      //   type: 'success',
			      //   message: `Your email is:${value}`,
			      // })
			    })
			    .catch(() => {
			      ElMessage({
			        type: 'info',
			        message: '操作取消',
			      })
			    })
		}else{
			
		}
		obj.conState = true;
	}

	const closeDB = (obj) => {
		console.log("关闭按钮点击事件:", obj)
		obj.conState = false;
	}

	const refresh = (obj) => {
		console.log("刷新按钮点击事件:", obj)
	}

	const editDB = (obj) => {
		console.log("编辑按钮点击事件:", obj)
		emit('openServerConfigEdit', obj.key);
	}


	const deleteDB = (obj) => {
		// console.log("删除按钮点击事件:", obj)
		ElMessageBox.confirm(
				'删除服务器配置后将不可恢复,是否继续操作?',
				'确定删除?', {
					confirmButtonText: '确定',
					cancelButtonText: '取消',
					type: 'warning',
				}
			)
			.then(() => {
				DeleteServerConfig(obj.key).then(result => {
					if (result.State == true) {
						GetServerList();
						ElMessage({
							type: 'success',
							message: '删除成功',
						})
					} else {
						ElMessage.error(result.Message)
					}
				})
			})
			.catch(() => {
				ElMessage({
					type: 'info',
					message: '取消删除操作',
				})
			})
	}
</script>

<style>
	.el-tree--highlight-current .el-tree-node.is-current>.el-tree-node__content {
		background-color: #b3ebff;
	}

	.el-button.is-text:not(.is-disabled):hover,
	.el-button.is-text:not(.is-disabled):focus {
		background-color: #e7f0fd;
	}

	.custom-tree-node {
		flex: 1;
		display: flex;
		align-items: center;
		justify-content: space-between;
		font-size: 14px;
		/* padding-right: 8px; */
	}

	.tree-node-title {
		display: flex;
		justify-content: start;
		align-items: center;
	}
</style>
