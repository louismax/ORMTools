export namespace main {
	
	export class ServerConfig {
	    key: string;
	    local_name: string;
	    dbType: string;
	    host: string;
	    port: string;
	    username: string;
	    password: string;
	    has_record_pwd: boolean;
	    hasUseSSH: boolean;
	    ssh_host: string;
	    ssh_port: string;
	    ssh_user: string;
	    has_ssh_keyfile: boolean;
	    ssh_keyfile: string;
	    has_ssh_pass: boolean;
	    ssh_password: string;
	    create_date_unix: number;
	
	    static createFrom(source: any = {}) {
	        return new ServerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.local_name = source["local_name"];
	        this.dbType = source["dbType"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.has_record_pwd = source["has_record_pwd"];
	        this.hasUseSSH = source["hasUseSSH"];
	        this.ssh_host = source["ssh_host"];
	        this.ssh_port = source["ssh_port"];
	        this.ssh_user = source["ssh_user"];
	        this.has_ssh_keyfile = source["has_ssh_keyfile"];
	        this.ssh_keyfile = source["ssh_keyfile"];
	        this.has_ssh_pass = source["has_ssh_pass"];
	        this.ssh_password = source["ssh_password"];
	        this.create_date_unix = source["create_date_unix"];
	    }
	}

}

