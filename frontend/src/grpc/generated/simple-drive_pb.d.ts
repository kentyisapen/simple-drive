// package: simpledrive
// file: simple-drive.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class Item extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasParentId(): boolean;
  clearParentId(): void;
  getParentId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setParentId(value?: google_protobuf_wrappers_pb.StringValue): void;

  getName(): string;
  setName(value: string): void;

  getType(): ItemTypeMap[keyof ItemTypeMap];
  setType(value: ItemTypeMap[keyof ItemTypeMap]): void;

  hasSize(): boolean;
  clearSize(): void;
  getSize(): google_protobuf_wrappers_pb.Int64Value | undefined;
  setSize(value?: google_protobuf_wrappers_pb.Int64Value): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasLastModifiedAt(): boolean;
  clearLastModifiedAt(): void;
  getLastModifiedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastModifiedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Item.AsObject;
  static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Item, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Item;
  static deserializeBinaryFromReader(message: Item, reader: jspb.BinaryReader): Item;
}

export namespace Item {
  export type AsObject = {
    id: string,
    parentId?: google_protobuf_wrappers_pb.StringValue.AsObject,
    name: string,
    type: ItemTypeMap[keyof ItemTypeMap],
    size?: google_protobuf_wrappers_pb.Int64Value.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastModifiedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class ItemCreateRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getType(): ItemTypeMap[keyof ItemTypeMap];
  setType(value: ItemTypeMap[keyof ItemTypeMap]): void;

  hasParentId(): boolean;
  clearParentId(): void;
  getParentId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setParentId(value?: google_protobuf_wrappers_pb.StringValue): void;

  getFile(): Uint8Array | string;
  getFile_asU8(): Uint8Array;
  getFile_asB64(): string;
  setFile(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemCreateRequest): ItemCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemCreateRequest;
  static deserializeBinaryFromReader(message: ItemCreateRequest, reader: jspb.BinaryReader): ItemCreateRequest;
}

export namespace ItemCreateRequest {
  export type AsObject = {
    name: string,
    type: ItemTypeMap[keyof ItemTypeMap],
    parentId?: google_protobuf_wrappers_pb.StringValue.AsObject,
    file: Uint8Array | string,
  }
}

export class ItemUpdateRequest extends jspb.Message {
  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasParentId(): boolean;
  clearParentId(): void;
  getParentId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setParentId(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemUpdateRequest): ItemUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemUpdateRequest;
  static deserializeBinaryFromReader(message: ItemUpdateRequest, reader: jspb.BinaryReader): ItemUpdateRequest;
}

export namespace ItemUpdateRequest {
  export type AsObject = {
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
    parentId?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class ListItemsRequest extends jspb.Message {
  hasParentId(): boolean;
  clearParentId(): void;
  getParentId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setParentId(value?: google_protobuf_wrappers_pb.StringValue): void;

  getPage(): number;
  setPage(value: number): void;

  getSize(): number;
  setSize(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListItemsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListItemsRequest): ListItemsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListItemsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListItemsRequest;
  static deserializeBinaryFromReader(message: ListItemsRequest, reader: jspb.BinaryReader): ListItemsRequest;
}

export namespace ListItemsRequest {
  export type AsObject = {
    parentId?: google_protobuf_wrappers_pb.StringValue.AsObject,
    page: number,
    size: number,
  }
}

export class ListItemsResponse extends jspb.Message {
  clearItemsList(): void;
  getItemsList(): Array<Item>;
  setItemsList(value: Array<Item>): void;
  addItems(value?: Item, index?: number): Item;

  hasPaging(): boolean;
  clearPaging(): void;
  getPaging(): Paging | undefined;
  setPaging(value?: Paging): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListItemsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListItemsResponse): ListItemsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListItemsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListItemsResponse;
  static deserializeBinaryFromReader(message: ListItemsResponse, reader: jspb.BinaryReader): ListItemsResponse;
}

export namespace ListItemsResponse {
  export type AsObject = {
    itemsList: Array<Item.AsObject>,
    paging?: Paging.AsObject,
  }
}

export class HealthCheckRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HealthCheckRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HealthCheckRequest): HealthCheckRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HealthCheckRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HealthCheckRequest;
  static deserializeBinaryFromReader(message: HealthCheckRequest, reader: jspb.BinaryReader): HealthCheckRequest;
}

export namespace HealthCheckRequest {
  export type AsObject = {
  }
}

export class HealthCheckResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HealthCheckResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HealthCheckResponse): HealthCheckResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HealthCheckResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HealthCheckResponse;
  static deserializeBinaryFromReader(message: HealthCheckResponse, reader: jspb.BinaryReader): HealthCheckResponse;
}

export namespace HealthCheckResponse {
  export type AsObject = {
    status: string,
  }
}

export class Paging extends jspb.Message {
  getCurrentPage(): number;
  setCurrentPage(value: number): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  getTotalItems(): number;
  setTotalItems(value: number): void;

  getTotalPages(): number;
  setTotalPages(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Paging.AsObject;
  static toObject(includeInstance: boolean, msg: Paging): Paging.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Paging, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Paging;
  static deserializeBinaryFromReader(message: Paging, reader: jspb.BinaryReader): Paging;
}

export namespace Paging {
  export type AsObject = {
    currentPage: number,
    pageSize: number,
    totalItems: number,
    totalPages: number,
  }
}

export interface ItemTypeMap {
  FILE: 0;
  DIRECTORY: 1;
}

export const ItemType: ItemTypeMap;

