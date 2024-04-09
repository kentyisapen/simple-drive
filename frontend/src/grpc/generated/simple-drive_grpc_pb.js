// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// proto/simple-drive.proto
'use strict';
var grpc = require('@grpc/grpc-js');
var simple$drive_pb = require('./simple-drive_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_StringValue(arg) {
  if (!(arg instanceof google_protobuf_wrappers_pb.StringValue)) {
    throw new Error('Expected argument of type google.protobuf.StringValue');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_StringValue(buffer_arg) {
  return google_protobuf_wrappers_pb.StringValue.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_HealthCheckRequest(arg) {
  if (!(arg instanceof simple$drive_pb.HealthCheckRequest)) {
    throw new Error('Expected argument of type simpledrive.HealthCheckRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_HealthCheckRequest(buffer_arg) {
  return simple$drive_pb.HealthCheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_HealthCheckResponse(arg) {
  if (!(arg instanceof simple$drive_pb.HealthCheckResponse)) {
    throw new Error('Expected argument of type simpledrive.HealthCheckResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_HealthCheckResponse(buffer_arg) {
  return simple$drive_pb.HealthCheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_Item(arg) {
  if (!(arg instanceof simple$drive_pb.Item)) {
    throw new Error('Expected argument of type simpledrive.Item');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_Item(buffer_arg) {
  return simple$drive_pb.Item.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_ItemCreateRequest(arg) {
  if (!(arg instanceof simple$drive_pb.ItemCreateRequest)) {
    throw new Error('Expected argument of type simpledrive.ItemCreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_ItemCreateRequest(buffer_arg) {
  return simple$drive_pb.ItemCreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_ItemUpdateRequest(arg) {
  if (!(arg instanceof simple$drive_pb.ItemUpdateRequest)) {
    throw new Error('Expected argument of type simpledrive.ItemUpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_ItemUpdateRequest(buffer_arg) {
  return simple$drive_pb.ItemUpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_ListItemsRequest(arg) {
  if (!(arg instanceof simple$drive_pb.ListItemsRequest)) {
    throw new Error('Expected argument of type simpledrive.ListItemsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_ListItemsRequest(buffer_arg) {
  return simple$drive_pb.ListItemsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_simpledrive_ListItemsResponse(arg) {
  if (!(arg instanceof simple$drive_pb.ListItemsResponse)) {
    throw new Error('Expected argument of type simpledrive.ListItemsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_simpledrive_ListItemsResponse(buffer_arg) {
  return simple$drive_pb.ListItemsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// サービス定義
var FileManagerService = exports.FileManagerService = {
  healthCheck: {
    path: '/simpledrive.FileManager/HealthCheck',
    requestStream: false,
    responseStream: false,
    requestType: simple$drive_pb.HealthCheckRequest,
    responseType: simple$drive_pb.HealthCheckResponse,
    requestSerialize: serialize_simpledrive_HealthCheckRequest,
    requestDeserialize: deserialize_simpledrive_HealthCheckRequest,
    responseSerialize: serialize_simpledrive_HealthCheckResponse,
    responseDeserialize: deserialize_simpledrive_HealthCheckResponse,
  },
  listItems: {
    path: '/simpledrive.FileManager/ListItems',
    requestStream: false,
    responseStream: false,
    requestType: simple$drive_pb.ListItemsRequest,
    responseType: simple$drive_pb.ListItemsResponse,
    requestSerialize: serialize_simpledrive_ListItemsRequest,
    requestDeserialize: deserialize_simpledrive_ListItemsRequest,
    responseSerialize: serialize_simpledrive_ListItemsResponse,
    responseDeserialize: deserialize_simpledrive_ListItemsResponse,
  },
  createItem: {
    path: '/simpledrive.FileManager/CreateItem',
    requestStream: false,
    responseStream: false,
    requestType: simple$drive_pb.ItemCreateRequest,
    responseType: simple$drive_pb.Item,
    requestSerialize: serialize_simpledrive_ItemCreateRequest,
    requestDeserialize: deserialize_simpledrive_ItemCreateRequest,
    responseSerialize: serialize_simpledrive_Item,
    responseDeserialize: deserialize_simpledrive_Item,
  },
  getItem: {
    path: '/simpledrive.FileManager/GetItem',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_wrappers_pb.StringValue,
    responseType: simple$drive_pb.Item,
    requestSerialize: serialize_google_protobuf_StringValue,
    requestDeserialize: deserialize_google_protobuf_StringValue,
    responseSerialize: serialize_simpledrive_Item,
    responseDeserialize: deserialize_simpledrive_Item,
  },
  updateItem: {
    path: '/simpledrive.FileManager/UpdateItem',
    requestStream: false,
    responseStream: false,
    requestType: simple$drive_pb.ItemUpdateRequest,
    responseType: simple$drive_pb.Item,
    requestSerialize: serialize_simpledrive_ItemUpdateRequest,
    requestDeserialize: deserialize_simpledrive_ItemUpdateRequest,
    responseSerialize: serialize_simpledrive_Item,
    responseDeserialize: deserialize_simpledrive_Item,
  },
  deleteItem: {
    path: '/simpledrive.FileManager/DeleteItem',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_wrappers_pb.StringValue,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_google_protobuf_StringValue,
    requestDeserialize: deserialize_google_protobuf_StringValue,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.FileManagerClient = grpc.makeGenericClientConstructor(FileManagerService);
