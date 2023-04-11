<template>
	<el-row justify="space-between">
		<el-col :span="5" style="text-align: left;padding-left: 30px;min-width: 300px;">
			<el-button plain @click="dialogFormVisible = true">
				<el-icon>
					<Plus />
				</el-icon>
				&nbsp;新建数据库连接
			</el-button>

			<el-tooltip content="导入连接配置" placement="bottom" effect="light">
				<el-button :text="true" style="width: 40px;">
					<el-image style="width: 20px; height: 20px" :src="img_import" fit="contain" />
				</el-button>
			</el-tooltip>

			<el-tooltip content="导出连接配置" placement="bottom" effect="light">
				<el-button :text="true" style="width: 40px;">
					<el-image style="width: 20px; height: 20px" :src="img_export" fit="contain" />
				</el-button>
			</el-tooltip>

		</el-col>
		<div>
			<div class="grid-content ep-bg-purple-light" />
		</div>
		<el-col :span="4" style="text-align: right;padding-right: 30px;">
			<el-button plain>
				<el-icon>
					<Tools />
				</el-icon>
				&nbsp;设置
			</el-button>
		</el-col>

	</el-row>

	<el-dialog v-model="dialogFormVisible" :show-close="false" class="addConfigWin">
		<template #header="{ close, titleId, titleClass }">

			<div class="my-header">
				<h4 :id="titleId" :class="titleClass">新连接配置</h4>
				<el-button text @click="close">
					<el-icon size="24">
						<Close />
					</el-icon>
				</el-button>
			</div>
		</template>
		<el-form :model="FormData" label-width="100px" ref="ruleFormRef" :rules="rules">
			<el-form-item label="连接名">
				<el-input v-model="FormData.local_name" placeholder="可以为空,为空时使用默认值: [DBType] _ [Host] : [Port]"
					autocomplete="off" />
			</el-form-item>
			<el-form-item label="数据库类型" prop="dbType">
				<el-select v-model="FormData.dbType" clearable placeholder="请选择数据库类型">
					<el-option label="MySql" value="MySql" />
				</el-select>
			</el-form-item>
			<el-form-item label="主机" prop="host">
				<el-col :span="18">
					<el-input v-model="FormData.host" placeholder="localhost" />
				</el-col>
				<el-col :span="2" class="text-center">
					<span class="text-gray-500">:</span>
				</el-col>
				<el-col :span="4">
					<el-input v-model="FormData.port" placeholder="3306" />
				</el-col>
			</el-form-item>
			<el-form-item label="用户名" prop="username">
				<el-input v-model="FormData.username" placeholder="root" />
			</el-form-item>
			<el-form-item label="密码" prop="password">
				<el-input type="password" show-password v-model="FormData.password" />
			</el-form-item>
			<el-form-item label="记住密码">
				<el-switch v-model="FormData.has_record_pwd" />
			</el-form-item>
			<el-form-item label="使用SSH通道">
				<el-switch v-model="FormData.hasUseSSH" />
			</el-form-item>

			<div v-if="FormData.hasUseSSH">
				<el-divider />
				<el-form-item label="SSH地址" prop="ssh_host">
					<el-col :span="18">
						<el-input v-model="FormData.ssh_host" placeholder="SSH 远程服务器" />
					</el-col>
					<el-col :span="2" class="text-center">
						<span class="text-gray-500">:</span>
					</el-col>
					<el-col :span="4">
						<el-input v-model="FormData.ssh_port" placeholder="22" />
					</el-col>
				</el-form-item>
				<el-form-item label="SSH用户名" prop="ssh_user">
					<el-input v-model="FormData.ssh_user" placeholder="验证SSH用户名" />
				</el-form-item>
				<el-form-item label="私钥" prop="ssh_keyfile">
					<el-col :span="2">
						<el-checkbox v-model="FormData.has_ssh_keyfile" />
					</el-col>
					<el-col :span="22">
						<input type="file" ref="fileSelect" hidden @change="fileChange">
						<el-input :disabled="!FormData.has_ssh_keyfile" readonly v-model="FormData.ssh_keyfile"
							placeholder="PEM格式私钥路径">
							<template #append>
								<el-tooltip content="选择本地文件" placement="bottom" effect="light">
									<el-button @click="btnChange" :disabled="!FormData.has_ssh_keyfile">
										...
									</el-button>
								</el-tooltip>
							</template>
						</el-input>
					</el-col>
				</el-form-item>
				<el-form-item label="密码" prop="ssh_password">
					<el-col :span="2">
						<el-checkbox v-model="FormData.has_ssh_pass" />
					</el-col>
					<el-col :span="22">
						<el-input :disabled="!FormData.has_ssh_pass" type="password" show-password
							v-model="FormData.ssh_password" placeholder="SSH用户密码" />
					</el-col>

				</el-form-item>
			</div>

		</el-form>

		<template #footer>
			<span class="dialog-footer">
				<el-button type="primary" @click="saveDBConfig">确定</el-button>
				<el-button @click="dialogFormVisible = false">取消</el-button>
			</span>
		</template>
	</el-dialog>

</template>


<script setup>
	import {
		Plus,
		Tools
	} from '@element-plus/icons-vue'
	import img_import from '../assets/images/import.png'
	import img_export from '../assets/images/export.png'
	import 'element-plus/theme-chalk/display.css'
	import {
		Close
	} from '@element-plus/icons-vue'
	import {
		SaveServerConfig,
		GetServerConfig
	} from '../../wailsjs/go/main/App'

	const fileSelect = ref();
	const ruleFormRef = ref();
	let dialogFormVisible = ref(false);
	let FormData = reactive({
		local_name: "",
		dbType: "MySql",
		host: "",
		port: "",
		username: "",
		password: "",
		has_record_pwd: false,
		hasUseSSH: false,
		ssh_host: "",
		ssh_port: "",
		ssh_user: "",
		has_ssh_keyfile: false,
		ssh_keyfile: "",
		has_ssh_pass: false,
		ssh_password: ""
	});

	const checkHostPort = (rule, value, callback) => {
		if (value == "") {
			return callback(new Error("请输入主机地址！"));
		}
		if (FormData.port == "") {
			return callback(new Error("请输入端口号！"));
		}
		return callback();
	};

	const checkSSHHostPort = (rule, value, callback) => {
		if (FormData.hasUseSSH) {
			if (value == "") {
				return callback(new Error("请输入SSH服务器地址！"));
			}
			if (FormData.ssh_port == "") {
				return callback(new Error("请输入SSH服务器端口号！"));
			}
		}
		return callback();
	};

	const checkSSHUser = (rule, value, callback) => {
		if (FormData.hasUseSSH) {
			if (value == "") {
				return callback(new Error("请输入SSH服务器用户名！"));
			}
		}
		return callback();
	};
	const checkSSHKeyFile = (rule, value, callback) => {
		if (FormData.hasUseSSH && FormData.has_ssh_keyfile) {
			if (value == "") {
				return callback(new Error("请选择SSH服务器密钥文件！"));
			}
			//验证文件是否存在
		}
		return callback();
	};
	const checkSSHPassWord = (rule, value, callback) => {
		if (FormData.hasUseSSH && FormData.has_ssh_pass) {
			if (value == "") {
				return callback(new Error("请输入SSH用户密码！"));
			}
			//验证文件是否存在
		}
		return callback();
	};



	const rules = reactive({
		local_name: [{
			required: true,
			message: '请输入连接名'
		}],
		dbType: [{
			required: true,
			message: '请选择数据库类型 ',
			trigger: 'change',
		}],
		host: [{
			required: true,
			validator: checkHostPort,
			trigger: 'blur'
		}],
		port: [{
			required: true,
			message: '请输入端口'
		}],
		username: [{
			required: true,
			message: '请输入登录用户名'
		}],
		password: [{
			required: true,
			message: '请输入登录密码'
		}],
		ssh_host: [{
			required: true,
			validator: checkSSHHostPort,
			trigger: 'blur'
		}],
		ssh_user: [{
			required: true,
			validator: checkSSHUser,
			trigger: 'blur'
		}],
		ssh_keyfile: [{
			required: true,
			validator: checkSSHKeyFile,
			trigger: 'blur'
		}],
		ssh_password: [{
			required: true,
			validator: checkSSHPassWord,
			trigger: 'blur'
		}]
	});



	function btnChange() {
		console.log(fileSelect)
		fileSelect.value.click();
	}

	function fileChange(e) {
		console.log(fileSelect.value.value)
		FormData.ssh_keyfile = fileSelect.value.value
	}

	function saveDBConfig() {
		console.log(ruleFormRef)
		ruleFormRef.value.validate((valid) => {
			if (valid) {
				SaveServerConfig(FormData).then(result => {
					console.log(result)
					
					GetServerConfig().then(result2 => {
						console.log(result2)
						
		
					})
					
				})

				//dialogFormVisible = false
			}
		})
	}
</script>

<style>
	.my-header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		max-height: 36px;
		align-items: center;
	}

	.addConfigWin {
		width: 600px
	}

	.addConfigWin .el-dialog__header {
		padding: 10px 20px;
	}

	.addConfigWin .el-dialog__body {
		/* height: 55vh; */
		max-height: 55vh;
		overflow: auto;
		padding: 0 20px;

		&::-webkit-scrollbar {
			width: 8px;
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

		/* 		&::-webkit-scrollbar-corner {
			background: #179a16;
		} */

	}

	.addConfigWin .el-dialog__footer {
		padding: 10px 20px;
	}
</style>
