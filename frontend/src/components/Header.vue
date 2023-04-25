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
				<el-button :text="true" style="width: 40px;" @click="importConfig">
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
				<el-button :text="true" style="width: 40px;" @click="exportConfig">
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
			<el-button plain @click="openDialogConfigForm">
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
				<el-input type="password" :disabled="!FormData.has_record_pwd" show-password
					v-model="FormData.password" />
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
			<el-row>
				<el-col :span="10" style=" display: flex;justify-content: start;align-items: center;">
					<el-button plain @click="testServerConn">
						<el-icon>
							<Promotion />
						</el-icon>
						&nbsp;测试连接
					</el-button>
					<div ref="connSuccess"
						style="display: none;justify-content: start;align-items: center;padding-left: 20px;">
						<el-icon color="#67C23A">
							<SuccessFilled />
						</el-icon>
						<span class="el-form-item__label">连接成功</span>
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

	<el-dialog v-model="dialogConfigVisible" :show-close="false" class="addConfigWin" :close-on-click-modal="false"
		destroy-on-close :close-on-press-escape="false" :draggable="true" @close="CloseDialogConfig">
		<template #header="{ close, titleId, titleClass }">
			<div class="my-header">
				<h4 :id="titleId" :class="titleClass">设置</h4>
				<el-button text @click="close">
					<el-icon size="24">
						<Close />
					</el-icon>
				</el-button>
			</div>
		</template>

		<el-form :model="ConfigForm">
			<span class="form_item_Grade">通用</span>
			<el-form-item label="应用主题">
				<el-select v-model="ConfigForm.WindowTheme" value-key placeholder="请选择应用主题" @change="WindowThemeChange">
					<el-option label="跟随系统默认" value="SystemDefault" />
					<el-option label="浅色" value="Light" />
					<el-option label="深色" value="Dark" />
				</el-select>
			</el-form-item>
			<el-form-item label="代码主题">
				<el-select v-model="ConfigForm.CodeTheme" value-key placeholder="请选择代码主题" @change="CodeThemeChange">
					<el-option label="A 11 Y Dark" value="a11y-dark" />
					<el-option label="A 11 Y Light" value="a11y-light" />
					<el-option label="Agate" value="agate" />
					<el-option label="An Old Hope" value="an-old-hope" />
					<el-option label="Androidstudio" value="androidstudio" />
					<el-option label="Atom One Dark" value="atom-one-dark" />
					<el-option label="Atom One Light" value="atom-one-light" />
					<el-option label="Base16 / Atelier Cave" value="atelier-cave" />
					<el-option label="Base16 / Paraiso" value="paraiso" />
					<el-option label="Github" value="github" />
					<el-option label="Github Dark" value="github-dark" />
					<el-option label="Github Dark Dimmed" value="github-dark-dimmed" />
					<el-option label="Gradient Dark" value="gradient-dark" />
					<el-option label="Gradient Light" value="gradient-light" />
					<el-option label="Vs" value="vs" />
					<el-option label="Vs 2015" value="vs2015" />
				</el-select>
				<span>（PS:下方为代码主题演示）</span>
				<highlightjs class="exam_code_box" language='golang' :code="codeStr" />
			</el-form-item>
			<el-divider />
			<span class="form_item_Grade">连接树</span>

			<el-form-item label="忽略数据库" label-width="100px" class="custom_item custom_item_listX">
				<el-tag v-for="tag in ConfigForm.HideDBList" :key="tag" class="tag_custom" effect="plain" closable
					:disable-transitions="false" @close="handleDBListClose(tag)">
					{{ tag }}
				</el-tag>
				<el-input v-if="inputHideDBVisible" ref="InputHideDBRef" v-model="inputHideDBValue" style="width: 5rem;"
					size="small" @keyup.enter="hideDBInputConfirm" @blur="hideDBInputConfirm" />
				<el-button v-else class="button-new-tag" size="small" @click="showHideDBInput">
					+ 添加数据库名
				</el-button>
			</el-form-item>

			<el-form-item label="忽略表" label-width="100px" class="custom_item custom_item_listX">
				<el-tag v-for="tag in ConfigForm.HideTableList" :key="tag" class="tag_custom" effect="plain" closable
					:disable-transitions="false" @close="handleTableListClose(tag)">
					{{ tag }}
				</el-tag>
				<el-input v-if="inputHideTableVisible" ref="InputHideTableRef" v-model="inputHideTableValue"
					style="width: 5rem;" size="small" @keyup.enter="hideTableInputConfirm"
					@blur="hideTableInputConfirm" />
				<el-button v-else class="button-new-tag" size="small" @click="showHideTableInput">
					+ 添加表名
				</el-button>
			</el-form-item>

			<el-form-item label="显示表备注" label-width="100px" class="custom_item">
				<el-switch v-model="ConfigForm.HasTableComment" />
			</el-form-item>

			<el-divider />

			<span class="form_item_Grade">表结构</span>

			<el-form-item label="自定义表名" label-width="128px" class="custom_item">
				<el-switch v-model="ConfigForm.HasRewriteTableName" />
			</el-form-item>
			<el-form-item label="Json标签" label-width="128px" class="custom_item">
				<el-switch v-model="ConfigForm.HasJsonTag" />
			</el-form-item>
			<el-form-item label="Grom Column标签" label-width="128px" class="custom_item">
				<el-switch v-model="ConfigForm.HasGormColumnTag" />
			</el-form-item>
			<el-form-item label="忽略字段" label-width="128px" class="custom_item custom_item_listX">
				<el-tag v-for="tag in ConfigForm.HideTableColumnList" :key="tag" class="tag_custom" effect="plain"
					closable :disable-transitions="false" @close="handleTableColumnListClose(tag)">
					{{ tag }}
				</el-tag>
				<el-input v-if="inputHideTableColumnVisible" ref="InputHideTableColumnRef"
					v-model="inputHideTableColumnValue" style="width: 5rem;" size="small"
					@keyup.enter="hideTableColumnInputConfirm" @blur="hideTableColumnInputConfirm" />
				<el-button v-else class="button-new-tag" size="small" @click="showHideTableColumnInput">
					+ 添加字段名
				</el-button>
			</el-form-item>
			<el-form-item label="字段类型转换规则" label-width="128px" class="custom_item">
				<el-table :data="toFieldList" style="width: 100%" empty-text="暂无规则">
					<el-table-column label="数据库字段" align="center">
						<template #default="scope">
							<span style="margin-left: 10px">{{ scope.row.key }}</span>
						</template>
					</el-table-column>
					<el-table-column width="24px" align="center">
						<template #header>
							<el-icon>
								<DArrowRight />
							</el-icon>
						</template>

						<template #default="scope">
							<div>
								<el-icon>
									<DArrowRight />
								</el-icon>
							</div>
						</template>
					</el-table-column>
					<el-table-column label="Golang Struct 字段" align="center">
						<template #default="scope">
							<span style="margin-left: 10px">{{ scope.row.val }}</span>
						</template>
					</el-table-column>
					<el-table-column width="54px" align="center">
						<template #default="scope">
							<el-button v-if="!showAddField" @click="delFieldItem($event,scope.row.id,scope.row.key)"
								text :icon="Delete" style="width: 24px;" />
						</template>
					</el-table-column>
				</el-table>
				<div v-if="showAddField" style="width: 100vw;">
					<el-row>
						<el-col :span="8" style="text-align: right;">
							<el-select v-model="addFieldSelect.dbField" filterable allow-create default-first-option
								:reserve-keyword="false" size="small" style="width: 100px;" value-key
								placeholder="数据库字段">
								<el-option-group v-for="group in dbFieldOptions" :key="group.key" :label="group.label">
									<el-option v-for="item in group.options" :key="item.value" :label="item.label"
										:value="item.value"
										:disabled="toFieldList.findIndex((value)=>value.key==item.value) == 0" />
								</el-option-group>
							</el-select>
						</el-col>
						<el-col :span="4" style="text-align: right;padding-right: 3%;">
							<el-icon>
								<DArrowRight />
							</el-icon>
						</el-col>
						<el-col :span="7" style="text-align: right;">
							<el-select v-model="addFieldSelect.goField" filterable allow-create default-first-option
								:reserve-keyword="false" size="small" style="width: 100px;" value-key
								placeholder="Golang字段">
								<el-option v-for="item in golangFieldOptions" :key="item.value" :label="item.label"
									:value="item.value" />
							</el-select>
						</el-col>
						<el-col :span="5">
							<el-button text :icon="Select" style="width: 24px;" @click="addFieldList($event)"
								:disabled="addFieldSelect.dbField == '' || addFieldSelect.goField == ''" />
							<el-button text :icon="CloseBold" style="width: 24px;" @click="showAddField = false" />
						</el-col>

					</el-row>
				</div>
				<div v-if="!showAddField" style="width: 100vw;text-align: right;margin-top: 5px;">
					<el-button size="small" @click="showAddField = true">
						新增规则
					</el-button>
					<el-button v-if="toFieldList.length >0" size="small" @click="emptyFieldList($event)">
						全部清空
					</el-button>
					<el-button v-if="toFieldList.length ==0" size="small" @click="importDefaultFieldRule($event)">
						引入全部默认规则
					</el-button>
				</div>

			</el-form-item>

		</el-form>
	</el-dialog>

	<el-dialog v-model="showSvrCfgStr" class="eiSvrConfig" :title="SvrCfgType == 'import'?'导入连接配置':'导出连接配置'">
		<div style="text-align: left;" v-if="SvrCfgType != 'import'">
			<span style="color: #ff9292;">* 完整复制下方文本，在导入时粘贴此文本即可完成导入</span>
		</div>
		<div>
			<el-input v-model="SvrCfgStr" :autosize="{ minRows: 4, maxRows: 12 }" input-style="max-height: 50vh;"
				readonly type="textarea" />
		</div>

		<template #footer>
			<span class="dialog-footer"  v-if="SvrCfgType == 'import'">
				<el-button @click="readSvrStrFile">读取文件</el-button>
				<el-button @click="pasteSvrStr">从剪贴板粘贴</el-button>
				<el-button type="primary" @click="submitSvrStr">提交</el-button>
			</span>
			<span class="dialog-footer" v-else>
				<el-button @click="downloadSvrStrFile">保存为文件</el-button>
				<el-button @click="copySvrStr">复制到剪贴板</el-button>
			</span>
			
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
		DArrowRight,
		Delete,
		Select,
		CloseBold,
	} from '@element-plus/icons-vue'
	import 'element-plus/theme-chalk/display.css'
	import {
		AddServerConfig,
		GetServerConfigList,
		GetServerConfig,
		EditServerConfig,
		TestDBConnect,
		EditUserConfigItem,
		ExportServerConfigList,
		ImportServerConfigList,
		SaveServerConfigFile,
		ReadServerConfigFile,
	} from '../../wailsjs/go/main/App'
	import {
		useStore
	} from 'vuex'
	import {
		v4 as uuidv4
	} from 'uuid'

	import {
		DBFieldOptions,
		GolangFieldOptions,
		DefaultFieldRule
	} from '../hook.js'
	import{
		BtnTargetBlur
	} from '../tools.js'
	const dbFieldOptions = DBFieldOptions();
	const golangFieldOptions = GolangFieldOptions();

	const emit = defineEmits(['GetServerList'])
	const fileSelect = ref();
	const ruleFormRef = ref();
	const connSuccess = ref();
	const store = useStore();
	let dialogFormVisible = ref(false);
	let dialogConfigVisible = ref(false);
	let FormData = reactive({
		key: "", //仅编辑时存在
		local_name: "",
		dbType: "MySql",
		host: "",
		port: "",
		username: "",
		password: "",
		has_record_pwd: true,
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
	let ConfigForm = ref({})
	let codeStr = 'func (a *App) ApiTest(r interface{}) string {\n\treturn fmt.Sprintf("result:%s", r.(string))\n}'

	defineExpose({
		OpenConfigEdit
	})
	const checkPassword = (rule, value, callback) => {
		if (FormData.has_record_pwd) {
			if (value == "") {
				return callback(new Error("请输入登录密码！"));
			}
		}
		return callback();
	};
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
			validator: checkPassword,
			trigger: 'blur'
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


	const openAddDialogForm = (e) => {
		BtnTargetBlur(e);
		SaveType = "Add"
		dialogName = "新建服务器配置";
		dialogFormVisible.value = true;
	}

	const btnChange = () => {
		//console.log(fileSelect)
		fileSelect.value.click();
	}

	const fileChange = (e) => {
		FormData.ssh_keyfile = fileSelect.value.value
	}

	const saveDBConfig = () => {
		ruleFormRef.value.validate((valid) => {
			if (valid) {
				if (SaveType == "Add") {
					AddServerConfig(FormData).then(result => {
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

	const testServerConn = (e) => {
		BtnTargetBlur(e);
		ruleFormRef.value.validate((valid) => {
			if (valid) {
				const loading = ElLoading.service({
					//lock: true,
					text: 'Loading',
					//background: 'rgba(0, 0, 0, 0.7)',
				})
				TestDBConnect(FormData).then(result => {
					loading.close()
					//console.log(result)
					if (result.State == true) {
						ElMessage.success('测试连接成功！')
						connSuccess.value.style.display = "flex"
					} else {
						connSuccess.value.style.display = "none"
						ElMessageBox({
							title: "连接测试失败",
							message: result.Message,
							type: 'error'
						})
					}
				})
			}
		})
	}

	const openDialogConfigForm = (e) => {
		BtnTargetBlur(e);
		ConfigForm = store.state.userConfig;
		dialogConfigVisible.value = true;
	}


	const WindowThemeChange = (val) => {
		if (val == "SystemDefault") {
			window.runtime.WindowSetSystemDefaultTheme();
		} else if (val == "Light") {
			window.runtime.WindowSetLightTheme();
		} else if (val == "Dark") {
			window.runtime.WindowSetDarkTheme();
		}
		//console.log(val)
	}
	const CodeThemeChange = (val) => {
		//console.log(val)
		document.getElementsByTagName('html')[0].dataset.codeTheme = val
	}


	const inputHideDBVisible = ref(false)
	const InputHideDBRef = ref()
	const inputHideDBValue = ref('')
	const handleDBListClose = (tag) => {
		ConfigForm.HideDBList.splice(ConfigForm.HideDBList.indexOf(tag), 1)
	}
	const showHideDBInput = () => {
		inputHideDBVisible.value = true
		nextTick(() => {
			InputHideDBRef.value.input.focus()
		})
	}
	const hideDBInputConfirm = () => {
		if (inputHideDBValue.value) {
			ConfigForm.HideDBList.push(inputHideDBValue.value)
		}
		inputHideDBVisible.value = false
		inputHideDBValue.value = ''
	}

	const inputHideTableVisible = ref(false)
	const InputHideTableRef = ref()
	const inputHideTableValue = ref('')
	const handleTableListClose = (tag) => {
		ConfigForm.HideTableList.splice(ConfigForm.HideTableList.indexOf(tag), 1)
	}
	const showHideTableInput = () => {
		inputHideTableVisible.value = true
		nextTick(() => {
			InputHideTableRef.value.input.focus()
		})
	}
	const hideTableInputConfirm = () => {
		if (inputHideTableValue.value) {
			ConfigForm.HideTableList.push(inputHideTableValue.value)
		}
		inputHideTableVisible.value = false
		inputHideTableValue.value = ''
	}

	const inputHideTableColumnVisible = ref(false)
	const InputHideTableColumnRef = ref()
	const inputHideTableColumnValue = ref('')
	const handleTableColumnListClose = (tag) => {
		ConfigForm.HideTableColumnList.splice(ConfigForm.HideTableColumnList.indexOf(tag), 1)
	}
	const showHideTableColumnInput = () => {
		inputHideTableColumnVisible.value = true
		nextTick(() => {
			InputHideTableColumnRef.value.input.focus()
		})
	}
	const hideTableColumnInputConfirm = () => {
		if (inputHideTableColumnValue.value) {
			ConfigForm.HideTableColumnList.push(inputHideTableColumnValue.value)
		}
		inputHideTableColumnVisible.value = false
		inputHideTableColumnValue.value = ''
	}

	let toFieldList = ref([])
	let addFieldSelect = ref({
		dbField: "",
		goField: ""
	})

	let showAddField = ref(false)
	const delFieldItem = (e, id, key) => {
		BtnTargetBlur(e);
		delete ConfigForm.MySqlToStructFieldType[key];
		toFieldList.value.splice(toFieldList.value.findIndex(item => item.id === id), 1);
	}
	const addFieldList = (e) => {
		BtnTargetBlur(e);
		toFieldList.value.push({
			id: uuidv4(),
			key: addFieldSelect.value.dbField,
			val: addFieldSelect.value.goField
		})
		ConfigForm.MySqlToStructFieldType[addFieldSelect.value.dbField] = addFieldSelect.value.goField;

		addFieldSelect.value.dbField = "";
		addFieldSelect.value.goField = "";
		showAddField.value = false;
	}
	const emptyFieldList = (e) => {
		BtnTargetBlur(e);
		ConfigForm.MySqlToStructFieldType = {};
		toFieldList.value.splice(0)
	}
	const importDefaultFieldRule = (e) => {
		BtnTargetBlur(e);

		DefaultFieldRule().value.forEach(function(item) {
			ConfigForm.MySqlToStructFieldType[item.key] = item.val;
		})

		toFieldList.value = DefaultFieldRule().value
	}

	const CloseDialogConfig = () => {
		EditUserConfigItem(ConfigForm).then(result => {
			console.log(result)
			if (result.State == true) {
				store.commit('setUserConfig', result.Data)
			} else {
				connSuccess.value.style.display = "none"
				ElMessageBox({
					title: "保存配置失败",
					message: result.Message,
					type: 'error'
				})
			}
		})
	}

	let showSvrCfgStr = ref(false);
	let SvrCfgStr = ref('');
	let SvrCfgType = ref('')


	const exportConfig = (e) => {
		BtnTargetBlur(e);

		ExportServerConfigList().then(result => {
			if (result.State == true) {
				//console.log(result)
				showSvrCfgStr.value = true;
				SvrCfgStr.value = result.Data
				SvrCfgType.value = "export";
			} else {
				ElMessageBox({
					title: "导出配置失败",
					message: result.Message,
					type: 'error'
				})
			}
		})
	}
	
	const importConfig = (e) => {
		BtnTargetBlur(e);
		showSvrCfgStr.value = true;
		SvrCfgStr.value = "";
		SvrCfgType.value = "import";
	}
	
	const copySvrStr = (e) => {
		BtnTargetBlur(e);

		window.runtime.ClipboardSetText(SvrCfgStr.value).then(({
				value
			}) => {
				ElMessage({
					message: '复制成功!',
					type: 'success',
				})
			})
			.catch(() => {})
	}
	const pasteSvrStr = (e) => {
		BtnTargetBlur(e);
	
		window.runtime.ClipboardGetText().then((s) => {
				SvrCfgStr.value = s;
				ElMessage({
					message: '粘贴成功!',
					type: 'success',
				})
			})
			.catch(() => {})
	}
	
	const downloadSvrStrFile = (e) => {
		BtnTargetBlur(e);
		const loading = ElLoading.service({
			//lock: true,
			text: 'Loading',
			//background: 'rgba(0, 0, 0, 0.7)',
		})

		SaveServerConfigFile(SvrCfgStr.value).then(result => {
			loading.close();
			if (result.State == true) {
				if (result.Data == "cancel") {
					ElMessage('取消操作')
				} else {
					ElMessage({
						message: result.Data,
						type: 'success',
					})
				}
			} else {
				ElMessageBox({
					title: "导出配置失败",
					message: result.Message,
					type: 'error'
				})
			}
		})
	}
	
	const readSvrStrFile = (e) => {
		BtnTargetBlur(e);
		const loading = ElLoading.service({
			//lock: true,
			text: 'Loading',
			//background: 'rgba(0, 0, 0, 0.7)',
		})
		ReadServerConfigFile().then(result => {
			loading.close();
			console.log(result)
			if (result.State == true) {
				if (result.Data == "cancel") {
					ElMessage('取消操作')
				} else {
					SvrCfgStr.value = result.Data;
					ElMessage({
						message: '成功!',
						type: 'success',
					})
				}
			} else {
				ElMessageBox({
					title: "导出配置失败",
					message: result.Message,
					type: 'error'
				})
			}
		})
	}
	
	const submitSvrStr = (e) => {
		BtnTargetBlur(e);
		
		ImportServerConfigList(SvrCfgStr.value).then(result => {
			console.log(result)
			if (result.State == true) {
				emit('GetServerList');
				showSvrCfgStr.value = false;
				ElMessage({
					message: '导入成功!',
					type: 'success',
				})
				
			} else {
				ElMessageBox({
					title: "导出配置失败",
					message: result.Message,
					type: 'error'
				})
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
		padding: 10px 10px;
		margin-right: 0;
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

	.form_item_Grade {
		display: flex;
		font-weight: bold;
		font-size: 16px;
	}

	.tag_custom {
		margin-right: 5px;
		margin-bottom: 5px;
	}

	.custom_item .el-form-item__content {
		align-items: flex-start;

	}

	.custom_item_listX .el-form-item__content {
		padding-top: 3px;
	}

	.exam_code_box {
		width: 100vw;
		margin: 10px 0;
		border: 1px solid var(--el-border-color);

		& code {
			padding: 10px;
			text-align: left;
			line-height: 1.2rem;
		}
	}

	.el-table__empty-block {
		min-height: 50px !important;
	}

	.el-table__empty-text {
		line-height: 50px !important;
	}

	/* 	.eiSvrConfig{
		max-height: 50vh;
	} */
	.eiSvrConfig .el-dialog__body {
		padding-top: 0;
		padding-bottom: 0;

		& textarea {
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
		}

	}
</style>
