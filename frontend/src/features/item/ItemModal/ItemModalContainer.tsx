"use client";
import { Modal } from "@mui/material";
import { ItemModalPresenter } from "./ItemModalPresenter";
import { Item } from "@/grpc/generated/simple-drive_pb";
import { useEffect, useState } from "react";
import { fetchItemBinary } from "@/utils/fetchItemBinary";

export type Props = {
	isOpened: boolean;
	item: Item.AsObject | null;
	handleOnClose: () => void;
};

export const ItemModalContainer: React.FC<Props> = ({
	isOpened,
	item,
	handleOnClose,
}) => {
	const [blobUrl, setBlobUrl] = useState<string | undefined>();

	useEffect(() => {
		(async () => {
			if (!item) {
				return;
			}
			const res = await fetchItemBinary(item.id);
			if (res.error) {
				return;
			}
			setBlobUrl(URL.createObjectURL(new Blob([res.data])));
		})();
	}, [item]);
	return (
		<ItemModalPresenter
			isOpened={isOpened}
			item={item}
			handleOnClose={handleOnClose}
			blobUrl={blobUrl}
		/>
	);
};
