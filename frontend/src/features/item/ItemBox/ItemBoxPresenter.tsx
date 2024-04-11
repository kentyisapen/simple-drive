import { Item } from "@/grpc/generated/simple-drive_pb";
import { Box, Grid, Paper } from "@mui/material";

type Props = {
	item: Item.AsObject;
};

export const ItemBoxPresenter = () => {
	return (
		<Grid
			item
			sx={{
				minWidth: 160, // 最小幅
				maxWidth: 240, // 最大幅
				width: "calc(20vw - 16px)", // ブラウザの幅に応じて調整
				flexGrow: 1,
			}}
		>
			<Paper
				sx={{
					width: "100%", // 親要素の幅に合わせる
					paddingTop: "100%", // 高さを幅と同じにする
					position: "relative",
				}}
			>
				<Box
					sx={{
						position: "absolute",
						top: 0,
						left: 0,
						bottom: 0,
						right: 0,
						display: "flex",
						justifyContent: "center",
						alignItems: "center",
					}}
				>
					{`Item`}
				</Box>
			</Paper>
		</Grid>
	);
};
