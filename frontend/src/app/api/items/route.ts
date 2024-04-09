// src/app/api/items/route.ts
import type { NextRequest } from "next/server";
import { NextResponse } from "next/server";
import { fileManagerClient } from "../../../grpc/client";
import { ListItemsRequest } from "../../../grpc/generated/simple-drive_pb";

export async function GET(req: NextRequest) {
	const request = new ListItemsRequest();
	// 必要に応じてrequestにパラメータを設定

	return new Promise((resolve, reject) => {
		fileManagerClient.listItems(request, (error, response) => {
			if (error) {
				return resolve(NextResponse.json({ error: error.message }));
			} else if (!response) {
				return resolve(
					NextResponse.json({ error: "No response received from the server" })
				);
			} else {
				return resolve(NextResponse.json(response.toObject()));
			}
		});
	});
}
