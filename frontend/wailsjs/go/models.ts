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
	    // Go type: time
	    create_date: any;
	
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
	        this.create_date = this.convertValues(source["create_date"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

