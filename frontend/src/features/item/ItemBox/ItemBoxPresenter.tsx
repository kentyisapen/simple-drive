import { Item } from "@/grpc/generated/simple-drive_pb";
import { Box, Card, CardContent, CardHeader, Grid, Paper, Typography } from "@mui/material";
import type { Props } from "./ItemBoxContainer";

export const ItemBoxPresenter = (props: Props) => {
	return (
		<Box sx={{ aspectRatio: "1/1" }} >
			<Card sx={{ height: "100%" }}>
				<CardContent >
					<Typography variant="h5" component="div">
						{props.item.name} - {props.item.type} - {props.item.size?.value}
					</Typography>
					{props.item.id}
				</CardContent>
			</Card>
		</Box >);
};
