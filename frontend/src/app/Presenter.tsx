// src/app/Presenter.tsx
"use client";
import React from "react";
import { ListItemsResponse } from "../grpc/generated/simple-drive_pb";
import { Box, Container, Grid } from "@mui/material";
import { ItemBoxPresenter } from "@/features/item/ItemBox/ItemBoxPresenter";
import { ItemBoxContainer } from "@/features/item/ItemBox/ItemBoxContainer";

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
				sx={{ border: 1, borderRadius: "12px", borderColor: "#999999", boxShadow: "2px", minHeight: "1080px" }}
			>
				<Box sx={{ display: "grid", gap: "8px", gridTemplateColumns: "repeat(auto-fill, minmax(220px, 1fr))", paddingTop: "8px" }}>
					{items?.itemsList?.map((item, index) => (
						<ItemBoxContainer key={item.id} item={item} />
					))}
				</Box>
				<div>
					page: {items?.paging?.currentPage} / {items?.paging?.totalPages}
				</div>
			</Container>
		</div >
	);
};

export default Presenter;
