// GENERATED CODE -- DO NOT EDIT!

// package: simpledrive
// file: simple-drive.proto

import * as simple_drive_pb from "./simple-drive_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "@grpc/grpc-js";

interface IFileManagerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  healthCheck: grpc.MethodDefinition<simple_drive_pb.HealthCheckRequest, simple_drive_pb.HealthCheckResponse>;
  listItems: grpc.MethodDefinition<simple_drive_pb.ListItemsRequest, simple_drive_pb.ListItemsResponse>;
  createItem: grpc.MethodDefinition<simple_drive_pb.ItemCreateRequest, simple_drive_pb.Item>;
  getItem: grpc.MethodDefinition<google_protobuf_wrappers_pb.StringValue, simple_drive_pb.Item>;
  updateItem: grpc.MethodDefinition<simple_drive_pb.ItemUpdateRequest, simple_drive_pb.Item>;
  deleteItem: grpc.MethodDefinition<google_protobuf_wrappers_pb.StringValue, google_protobuf_empty_pb.Empty>;
}

export const FileManagerService: IFileManagerService;

export interface IFileManagerServer extends grpc.UntypedServiceImplementation {
  healthCheck: grpc.handleUnaryCall<simple_drive_pb.HealthCheckRequest, simple_drive_pb.HealthCheckResponse>;
  listItems: grpc.handleUnaryCall<simple_drive_pb.ListItemsRequest, simple_drive_pb.ListItemsResponse>;
  createItem: grpc.handleUnaryCall<simple_drive_pb.ItemCreateRequest, simple_drive_pb.Item>;
  getItem: grpc.handleUnaryCall<google_protobuf_wrappers_pb.StringValue, simple_drive_pb.Item>;
  updateItem: grpc.handleUnaryCall<simple_drive_pb.ItemUpdateRequest, simple_drive_pb.Item>;
  deleteItem: grpc.handleUnaryCall<google_protobuf_wrappers_pb.StringValue, google_protobuf_empty_pb.Empty>;
}

export class FileManagerClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  healthCheck(argument: simple_drive_pb.HealthCheckRequest, callback: grpc.requestCallback<simple_drive_pb.HealthCheckResponse>): grpc.ClientUnaryCall;
  healthCheck(argument: simple_drive_pb.HealthCheckRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.HealthCheckResponse>): grpc.ClientUnaryCall;
  healthCheck(argument: simple_drive_pb.HealthCheckRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.HealthCheckResponse>): grpc.ClientUnaryCall;
  listItems(argument: simple_drive_pb.ListItemsRequest, callback: grpc.requestCallback<simple_drive_pb.ListItemsResponse>): grpc.ClientUnaryCall;
  listItems(argument: simple_drive_pb.ListItemsRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.ListItemsResponse>): grpc.ClientUnaryCall;
  listItems(argument: simple_drive_pb.ListItemsRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.ListItemsResponse>): grpc.ClientUnaryCall;
  createItem(argument: simple_drive_pb.ItemCreateRequest, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  createItem(argument: simple_drive_pb.ItemCreateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  createItem(argument: simple_drive_pb.ItemCreateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  getItem(argument: google_protobuf_wrappers_pb.StringValue, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  getItem(argument: google_protobuf_wrappers_pb.StringValue, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  getItem(argument: google_protobuf_wrappers_pb.StringValue, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  updateItem(argument: simple_drive_pb.ItemUpdateRequest, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  updateItem(argument: simple_drive_pb.ItemUpdateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  updateItem(argument: simple_drive_pb.ItemUpdateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<simple_drive_pb.Item>): grpc.ClientUnaryCall;
  deleteItem(argument: google_protobuf_wrappers_pb.StringValue, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  deleteItem(argument: google_protobuf_wrappers_pb.StringValue, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  deleteItem(argument: google_protobuf_wrappers_pb.StringValue, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
}
