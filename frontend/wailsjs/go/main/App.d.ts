// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function AddServerConfig(arg1:main.ServerConfig):Promise<any>;

export function CloseDBConnect(arg1:string):Promise<any>;

export function DeleteServerConfig(arg1:string):Promise<any>;

export function EditServerConfig(arg1:main.ServerConfig):Promise<any>;

export function GetServerConfig(arg1:string):Promise<any>;

export function GetServerConfigList():Promise<any>;

export function OpenDBConnect(arg1:main.ServerConfig):Promise<any>;

export function QueryTableFieldList(arg1:string,arg2:string,arg3:string,arg4:string):Promise<any>;

export function QueryTableList(arg1:string,arg2:string):Promise<any>;

export function RefreshDBConnect(arg1:string):Promise<any>;

export function ReturnError(arg1:string):Promise<{[key: string]: any}>;

export function ReturnSuccess(arg1:any):Promise<{[key: string]: any}>;

export function TestDBConnect(arg1:main.ServerConfig):Promise<any>;
