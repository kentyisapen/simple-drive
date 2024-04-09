// src/grpc/client.ts
import { FileManagerClient } from "./generated/simple-drive_grpc_pb";
import { credentials } from "@grpc/grpc-js";

// gRPCサーバーのエンドポイントを指定
const GRPC_SERVER_ENDPOINT = "backend:8021"; // 適宜変更してください

export const fileManagerClient = new FileManagerClient(
	GRPC_SERVER_ENDPOINT,
	credentials.createInsecure()
);
