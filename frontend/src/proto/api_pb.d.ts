// package: api
// file: api.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class LoginRequest extends jspb.Message { 
    getLogintoken(): string;
    setLogintoken(value: string): LoginRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LoginRequest.AsObject;
    static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LoginRequest;
    static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
    export type AsObject = {
        logintoken: string,
    }
}

export class RepeatNewsRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): RepeatNewsRequest;
    getCount(): number;
    setCount(value: number): RepeatNewsRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RepeatNewsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RepeatNewsRequest): RepeatNewsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RepeatNewsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RepeatNewsRequest;
    static deserializeBinaryFromReader(message: RepeatNewsRequest, reader: jspb.BinaryReader): RepeatNewsRequest;
}

export namespace RepeatNewsRequest {
    export type AsObject = {
        token: string,
        count: number,
    }
}

export class LoginReply extends jspb.Message { 
    getStatus(): number;
    setStatus(value: number): LoginReply;
    getToken(): string;
    setToken(value: string): LoginReply;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LoginReply.AsObject;
    static toObject(includeInstance: boolean, msg: LoginReply): LoginReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LoginReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LoginReply;
    static deserializeBinaryFromReader(message: LoginReply, reader: jspb.BinaryReader): LoginReply;
}

export namespace LoginReply {
    export type AsObject = {
        status: number,
        token: string,
    }
}

export class NewsReply extends jspb.Message { 
    getId(): number;
    setId(value: number): NewsReply;
    getMessage(): string;
    setMessage(value: string): NewsReply;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewsReply.AsObject;
    static toObject(includeInstance: boolean, msg: NewsReply): NewsReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewsReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewsReply;
    static deserializeBinaryFromReader(message: NewsReply, reader: jspb.BinaryReader): NewsReply;
}

export namespace NewsReply {
    export type AsObject = {
        id: number,
        message: string,
    }
}