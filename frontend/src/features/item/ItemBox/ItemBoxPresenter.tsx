import { Item } from "@/grpc/generated/simple-drive_pb";
import {
	Box,
	Card,
	CardActionArea,
	CardContent,
	CardHeader,
	CardMedia,
	Grid,
	Paper,
	Typography,
} from "@mui/material";
import type { Props } from "./ItemBoxContainer";
import { useEffect, useRef } from "react";
import Image from "next/image";
import { getImageURL } from "@/utils/ImageURL";

export const ItemBoxPresenter = (props: Props) => {
	const ref = useRef();
	if (ref === null) return <></>;

	return (
		<Box ref={ref} sx={{ aspectRatio: "1/1" }}>
			<Card sx={{ height: "100%" }}>
				<CardActionArea onClick={() => props.handleOnClickItem(props.item)}>
					<CardHeader
						title={`${props.item.name} - ${props.item.type} - ${props.item.size?.value}`}
					></CardHeader>
					<CardMedia
						component="img"
						image={getImageURL(props.item.thumbnailId)}
					></CardMedia>
				</CardActionArea>
			</Card>
		</Box>
	);
};
