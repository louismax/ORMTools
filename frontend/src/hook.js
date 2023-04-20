import {
	v4 as uuidv4
} from 'uuid'

export const DBFieldOptions = () => {
	return ref([{
		key: "fieldNum",
		label: "数值类型",
		options: [{
				value: 'tinyint',
				label: 'TINYINT (tinyint)',
			},
			{
				value: 'smallint',
				label: 'SMALLINT (smallint)',
			},
			{
				value: 'mediumint',
				label: 'MEDIUMINT (mediumint)',
			},
			{
				value: 'integer',
				label: 'INTEGER (integer)',
			},
			{
				value: 'int',
				label: 'INT (int)',
			},
			{
				value: 'bigint',
				label: 'BIGINT (bigint)',
			},
			{
				value: 'float',
				label: 'FLOAT (float)',
			},
			{
				value: 'double',
				label: 'DOUBLE (double)',
			},
			{
				value: 'decimal',
				label: 'DECIMAL (decimal)',
			},
		],
	}, {
		key: "fieldstr",
		label: "字符串类型",
		options: [{
				value: 'char',
				label: 'CHAR (char)',
			},
			{
				value: 'varchar',
				label: 'VARCHAR (varchar)',
			},
			{
				value: 'tinyblob',
				label: 'TINYBLOB (tinyblob)',
			},
			{
				value: 'tinytext',
				label: 'TINYTEXT (tinytext)',
			},
			{
				value: 'blob',
				label: 'BLOB (blob)',
			},
			{
				value: 'text',
				label: 'TEXT (text)',
			},
			{
				value: 'mediumblob',
				label: 'MEDIUMBLOB (mediumblob)',
			},
			{
				value: 'mediumtext',
				label: 'MEDIUMTEXT (mediumtext)',
			},
			{
				value: 'longblob',
				label: 'LONGBLOB (longblob)',
			},
			{
				value: 'longtext',
				label: 'LONGTEXT (longtext)',
			},
		],
	}, {
		key: "fieldDate",
		label: "日期和时间类型",
		options: [{
				value: 'date',
				label: 'DATE (date)',
			},
			{
				value: 'time',
				label: 'TIME (time)',
			},
			{
				value: 'year',
				label: 'YEAR (year)',
			},
			{
				value: 'datetime',
				label: 'DATETIME (datetime)',
			},
			{
				value: 'timestamp',
				label: 'TIMESTAMP (timestamp)',
			}
		],
	}])
}


export const GolangFieldOptions = () => {
	return ref([{
			value: 'string',
			label: 'string',
		},
		{
			value: 'bool',
			label: 'bool',
		},
		{
			value: 'interface{}',
			label: 'interface{}',
		},
		{
			value: 'time.Time',
			label: 'time.Time',
		},
		{
			value: 'int',
			label: 'int',
		},
		{
			value: 'uint',
			label: 'uint',
		},
		{
			value: 'rune',
			label: 'rune',
		},
		{
			value: 'byte',
			label: 'byte',
		},
		{
			value: 'uintptr',
			label: 'uintptr',
		},
		{
			value: 'int8',
			label: 'int8',
		},
		{
			value: 'int16',
			label: 'int16',
		},
		{
			value: 'int32',
			label: 'int32',
		},
		{
			value: 'int64',
			label: 'int64',
		},
		{
			value: 'uint8',
			label: 'uint8',
		},
		{
			value: 'uint16',
			label: 'uint16',
		},
		{
			value: 'uint32',
			label: 'uint32',
		},
		{
			value: 'uint64',
			label: 'uint64',
		},
		{
			value: 'float32',
			label: 'float32',
		},
		{
			value: 'float64',
			label: 'float64',
		},
		{
			value: 'complex64',
			label: 'complex64',
		},
		{
			value: 'complex128',
			label: 'complex128',
		}
	])
}

export const DefaultFieldRule = () => {
	return ref([{
		id: uuidv4(),
		key: "tinyint",
		val: "int8"
	},{
		id: uuidv4(),
		key: "smallint",
		val: "int16"
	},{
		id: uuidv4(),
		key: "mediumint",
		val: "int32"
	},{
		id: uuidv4(),
		key: "integer",
		val: "int"
	},{
		id: uuidv4(),
		key: "int",
		val: "int"
	},{
		id: uuidv4(),
		key: "bigint",
		val: "int64"
	},{
		id: uuidv4(),
		key: "float",
		val: "float32"
	},{
		id: uuidv4(),
		key: "double",
		val: "float64"
	},{
		id: uuidv4(),
		key: "char",
		val: "string"
	},{
		id: uuidv4(),
		key: "varchar",
		val: "string"
	},{
		id: uuidv4(),
		key: "text",
		val: "string"
	},{
		id: uuidv4(),
		key: "longtext",
		val: "string"
	},{
		id: uuidv4(),
		key: "blob",
		val: "string"
	},{
		id: uuidv4(),
		key: "longblob",
		val: "string"
	},{
		id: uuidv4(),
		key: "datetime",
		val: "time.Time"
	},{
		id: uuidv4(),
		key: "timestamp",
		val: "time.Time"
	} ])
}
