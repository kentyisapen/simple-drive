import { FileManagerClient } from './generated/simple-drive_grpc_pb';
import { credentials } from '@grpc/grpc-js';

// gRPCサーバーのエンドポイントを指定
const GRPC_SERVER_ENDPOINT = 'http://localhost:8021';

// FileManagerClient のインスタンスを作成
const fileManagerClient = new FileManagerClient(GRPC_SERVER_ENDPOINT, credentials.createInsecure());

export default fileManagerClient;