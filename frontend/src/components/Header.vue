<template>
	<el-row justify="space-between">
		<el-col :span="5" style="text-align: left;padding-left: 30px;min-width: 300px;">
			<el-button plain @click="openAddDialogForm">
				<el-icon>
					<Plus />
				</el-icon>
				&nbsp;新建数据库连接
			</el-button>

			<el-tooltip content="导入连接配置" placement="bottom" effect="light">
				<el-button :text="true" style="width: 40px;">
					<svg t="1681303368254" class="icon" viewBox="0 0 1032 1024" version="1.1"
						style="width: 20px; height: 20px" xmlns="http://www.w3.org/2000/svg" p-id="1064" width="200"
						height="200">
						<path
							d="M563.2 42.666667v640c0 23.466667-19.2 42.666667-42.666667 42.666666s-42.666667-19.2-42.666666-42.666666V42.666667c0-23.466667 19.2-42.666667 42.666666-42.666667s42.666667 19.2 42.666667 42.666667"
							fill="#707070" p-id="1065"></path>
						<path
							d="M806.1952 457.275733L550.5792 712.900267a42.7776 42.7776 0 0 1-60.330667 0c-16.597333-16.605867-16.597333-43.776 0-60.330667l255.616-255.624533a42.7776 42.7776 0 0 1 60.330667 0 42.7776 42.7776 0 0 1 0 60.330666"
							fill="#707070" p-id="1066"></path>
						<path
							d="M490.350933 712.8832L234.734933 457.275733a42.7776 42.7776 0 0 1 0-60.330666c16.597333-16.605867 43.776-16.605867 60.330667 0l255.616 255.616a42.7776 42.7776 0 0 1 0 60.330666 42.7776 42.7776 0 0 1-60.330667 0"
							fill="#707070" p-id="1067"></path>
						<path
							d="M904.533333 1024H136.533333c-70.570667 0-128-57.429333-128-128V640c0-23.594667 19.072-42.666667 42.666667-42.666667s42.666667 19.072 42.666667 42.666667v256c0 23.552 19.114667 42.666667 42.666666 42.666667h768c23.552 0 42.666667-19.114667 42.666667-42.666667V640c0-23.594667 19.072-42.666667 42.666667-42.666667s42.666667 19.072 42.666666 42.666667v256c0 70.570667-57.429333 128-128 128"
							fill="#707070" p-id="1068"></path>
					</svg>
				</el-button>
			</el-tooltip>

			<el-tooltip content="导出连接配置" placement="bottom" effect="light">
				<el-button :text="true" style="width: 40px;">
					<svg t="1681303432514" class="icon" viewBox="0 0 1024 1024" version="1.1"
						style="width: 20px; height: 20px" xmlns="http://www.w3.org/2000/svg" p-id="1365" width="200"
						height="200">
						<path
							d="M469.333333 682.666667V42.666667c0-23.466667 19.2-42.666667 42.666667-42.666667s42.666667 19.2 42.666667 42.666667v640c0 23.466667-19.2 42.666667-42.666667 42.666666s-42.666667-19.2-42.666667-42.666666"
							fill="#707070" p-id="1366"></path>
						<path
							d="M226.210133 268.0576L481.826133 12.433067a42.7776 42.7776 0 0 1 60.330667 0c16.597333 16.597333 16.597333 43.776 0 60.330666L286.5408 328.388267a42.7776 42.7776 0 0 1-60.330667 0 42.7776 42.7776 0 0 1 0-60.330667"
							fill="#707070" p-id="1367"></path>
						<path
							d="M542.0544 12.4416l255.616 255.616a42.7776 42.7776 0 0 1 0 60.330667c-16.597333 16.605867-43.776 16.605867-60.330667 0L481.723733 72.789333a42.7776 42.7776 0 0 1 0-60.330666 42.7776 42.7776 0 0 1 60.330667 0M896 1024H128c-70.570667 0-128-57.429333-128-128V640c0-23.594667 19.072-42.666667 42.666667-42.666667s42.666667 19.072 42.666666 42.666667v256c0 23.552 19.114667 42.666667 42.666667 42.666667h768c23.552 0 42.666667-19.114667 42.666667-42.666667V640c0-23.594667 19.072-42.666667 42.666666-42.666667s42.666667 19.072 42.666667 42.666667v256c0 70.570667-57.429333 128-128 128"
							fill="#707070" p-id="1368"></path>
					</svg>
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

	<el-dialog v-model="dialogFormVisible" :show-close="false" class="addConfigWin" :close-on-click-modal="false"
		destroy-on-close :close-on-press-escape="false" :draggable="true">
		<template #header="{ close, titleId, titleClass }">
			<div class="my-header">
				<h4 :id="titleId" :class="titleClass">{{dialogName}}</h4>
				<el-button text @click="AddConfigDialogClose">
					<el-icon size="24">
						<Close />
					</el-icon>
				</el-button>
			</div>
		</template>
		<el-form :model="FormData" label-width="100px" ref="ruleFormRef" :rules="rules">
			<el-form-item label="连接名">
				<el-input v-model="FormData.local_name" placeholder="可以为空,为空时使用默认值: [DBType]_[Host]:[Port]"
					autocomplete="off" />
			</el-form-item>
			<el-form-item label="数据库类型" prop="dbType">
				<el-select v-model="FormData.dbType" clearable placeholder="请选择数据库类型">
					<el-option label="MySql" value="MySql" />
				</el-select>
			</el-form-item>
			<el-form-item label="主机" required>
				<el-col :span="14">
					<el-form-item prop="host">
						<el-input v-model="FormData.host" placeholder="localhost" />
					</el-form-item>
				</el-col>
				<el-col :span="2" class="text-center">
					<span class="text-gray-500">:</span>
				</el-col>
				<el-col :span="8">
					<el-form-item prop="port">
						<el-input v-model="FormData.port" placeholder="3306" />
					</el-form-item>

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
				<el-form-item label="SSH地址" required>
					<el-col :span="14">
						<el-form-item prop="ssh_host">
							<el-input v-model="FormData.ssh_host" placeholder="SSH 远程服务器" />
						</el-form-item>
					</el-col>
					<el-col :span="2" class="text-center">
						<span class="text-gray-500">:</span>
					</el-col>
					<el-col :span="8">
						<el-form-item prop="ssh_port">
							<el-input v-model="FormData.ssh_port" placeholder="22" />
						</el-form-item>
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
			<el-row >
				<el-col :span="10" style="display: flex;justify-content: start;align-items: center;">
					<el-button plain>
						<el-icon><Promotion /></el-icon>
						&nbsp;测试连接
					</el-button>
					<div v-if="false" style="display: flex;justify-content: start;align-items: center;padding-left: 20px;">
						<el-icon><SuccessFilled /></el-icon>
						<span>连接成功</span>
					</div>
					
				</el-col>
				<el-col :span="14">
					<span class="dialog-footer">
						<el-button type="primary" @click="saveDBConfig">
							<span v-if="SaveType=='Add'">确定</span>
							<span v-else>保存</span>
						</el-button>
						<el-button @click="AddConfigDialogClose">取消</el-button>
					</span>
				</el-col>
			</el-row>

		</template>
	</el-dialog>

</template>


<script setup>
	import {
		Plus,
		Tools,
		Close,
		Promotion,
		SuccessFilled,
	} from '@element-plus/icons-vue'
	import 'element-plus/theme-chalk/display.css'
	// import {
	// 	defineExpose,
	// 	defineEmits
	// } from 'vue'
	import {
		AddServerConfig,
		GetServerConfigList,
		GetServerConfig,
		EditServerConfig,
	} from '../../wailsjs/go/main/App'


	const emit = defineEmits(['GetServerList'])
	const fileSelect = ref();
	const ruleFormRef = ref();
	let dialogFormVisible = ref(false);
	let FormData = reactive({
		key: "", //仅编辑时存在
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
	let SaveType = "Add"
	let dialogName = "新连接配置"

	defineExpose({
		OpenConfigEdit
	})

	const checkSSHHost = (rule, value, callback) => {
		if (FormData.hasUseSSH) {
			if (value == "") {
				return callback(new Error("请输入SSH服务器地址！"));
			}
		}
		return callback();
	};
	const checkSSHPort = (rule, value, callback) => {
		if (FormData.hasUseSSH) {
			if (value == "") {
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
			message: '请输入主机地址'
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
			validator: checkSSHHost,
			trigger: 'blur'
		}],
		ssh_port: [{
			required: true,
			validator: checkSSHPort,
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


	const openAddDialogForm = () => {
		SaveType = "Add"
		dialogName = "新建服务器配置";
		dialogFormVisible.value = true;
	}

	const btnChange = () => {
		//console.log(fileSelect)
		fileSelect.value.click();
	}

	const fileChange = (e) => {
		//console.log(fileSelect.value.value)
		FormData.ssh_keyfile = fileSelect.value.value
	}

	const saveDBConfig = () => {
		//console.log(ruleFormRef)
		ruleFormRef.value.validate((valid) => {
			if (valid) {
				if (SaveType == "Add") {
					AddServerConfig(FormData).then(result => {
						//console.log(result)
						if (result.State == true) {
							emit('GetServerList');
							ruleFormRef.value.resetFields();
							dialogFormVisible.value = false;
							ElMessage.success('新增成功！')
						} else {
							ElMessage.error(result.Message)
						}
					})
				} else {
					console.log(FormData)
					EditServerConfig(FormData).then(result => {
						console.log(result)
						if (result.State == true) {
							emit('GetServerList');
							ruleFormRef.value.resetFields();
							dialogFormVisible.value = false;
							ElMessage.success('保存成功！')
						} else {
							ElMessage.error(result.Message)
						}
					})
				}

			}
		})
	}

	const AddConfigDialogClose = () => {
		emptyFormData();
		//ruleFormRef.value.resetFields();
		dialogFormVisible.value = false;
	}

	function OpenConfigEdit(key) {
		//emptyFormData();
		GetServerConfig(key).then(result => {
			if (result.State == true) {
				FormData.key = result.Data.key
				FormData.local_name = result.Data.local_name
				FormData.dbType = result.Data.dbType
				FormData.host = result.Data.host
				FormData.port = result.Data.port
				FormData.username = result.Data.username
				FormData.password = result.Data.password
				FormData.has_record_pwd = result.Data.has_record_pwd
				FormData.hasUseSSH = result.Data.hasUseSSH
				FormData.ssh_host = result.Data.ssh_host
				FormData.ssh_port = result.Data.ssh_port
				FormData.ssh_user = result.Data.ssh_user
				FormData.has_ssh_keyfile = result.Data.has_ssh_keyfile
				FormData.ssh_keyfile = result.Data.ssh_keyfile
				FormData.has_ssh_pass = result.Data.has_ssh_pass
				FormData.ssh_password = result.Data.ssh_password

				SaveType = "Edit"
				dialogName = "编辑服务器配置";
				dialogFormVisible.value = true;
			} else {
				ElMessage.error(result.Message)
			}
		})
	}
	const emptyFormData = () => {
		FormData.key = ""
		FormData.local_name = ""
		FormData.dbType = ""
		FormData.host = ""
		FormData.port = ""
		FormData.username = ""
		FormData.password = ""
		FormData.has_record_pwd = false
		FormData.hasUseSSH = false
		FormData.ssh_host = ""
		FormData.ssh_port = ""
		FormData.ssh_user = ""
		FormData.has_ssh_keyfile = false
		FormData.ssh_keyfile = ""
		FormData.has_ssh_pass = false
		FormData.ssh_password = ""
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
