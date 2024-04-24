"use client";
import { Box, Modal } from "@mui/material";

import type { Props } from "./ItemModalContainer";

import { defaultLayoutPlugin } from "@react-pdf-viewer/default-layout";
import { Viewer, Worker } from "@react-pdf-viewer/core";

import "@react-pdf-viewer/core/lib/styles/index.css";
import "@react-pdf-viewer/default-layout/lib/styles/index.css";

type PresenterProps = Props & {
	blobUrl: string | undefined;
};

export const ItemModalPresenter: React.FC<PresenterProps> = ({
	isOpened,
	item,
	handleOnClose,
	blobUrl,
}) => {
	const defaultLayoutPluginInstance = defaultLayoutPlugin();
	if (!item) {
		return (
			<Modal open={isOpened} onClose={handleOnClose}>
				<div>loading...</div>
			</Modal>
		);
	}
	return (
		<Modal open={isOpened} onClose={handleOnClose}>
			<Box
				sx={{
					position: "absolute" as "absolute",
					top: "50%",
					left: "50%",
					transform: "translate(-50%, -50%)",
					width: "80%",
					height: "80%",
					objectFit: "contain ",
				}}
			>
				{blobUrl ? (
					<Worker workerUrl="https://unpkg.com/pdfjs-dist@3.11.174/build/pdf.worker.js">
						{" "}
						<Viewer fileUrl={blobUrl} plugins={[defaultLayoutPluginInstance]} />
					</Worker>
				) : (
					<>Loading...</>
				)}
			</Box>
		</Modal>
	);
};
