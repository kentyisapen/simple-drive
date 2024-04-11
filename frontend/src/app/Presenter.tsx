// src/app/Presenter.tsx
"use client";
import React from "react";
import { ListItemsResponse } from "../grpc/generated/simple-drive_pb";
import { Container, Grid } from "@mui/material";
import { ItemBoxPresenter } from "@/features/item/ItemBox/ItemBoxPresenter";

type PresenterProps = {
	items: ListItemsResponse.AsObject | null;
	loading: boolean;
	error: string | null;
};

const Presenter: React.FC<PresenterProps> = ({ items, loading, error }) => {
	if (loading) return <div>Loading...</div>;
	if (error) return <div>Error: {error}</div>;

	return (
		<div>
			<h1>Item List</h1>
			<Container
				sx={{ border: 1, borderRadius: "12px", borderColor: "#999999" }}
			>
				<Grid container spacing={2} wrap="wrap">
					{items?.itemsList?.map((item, index) => (
						<ItemBoxPresenter />

						// <div key={index}>
						// 	{item.name} - {item.type} - {item.size?.value}
						// </div>
					))}
					{[1, 2, 3, 4, 5, 6].map(() => (
						<ItemBoxPresenter />
					))}
				</Grid>

				<div>
					page: {items?.paging?.currentPage} / {items?.paging?.totalPages}
				</div>
			</Container>
		</div>
	);
};

export default Presenter;
