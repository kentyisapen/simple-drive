// src/app/Presenter.tsx
"use client";
import React from "react";
import { Item, ListItemsResponse } from "../grpc/generated/simple-drive_pb";
import { Box, Container, Grid } from "@mui/material";
import { ItemBoxPresenter } from "@/features/item/ItemBox/ItemBoxPresenter";
import { ItemBoxContainer } from "@/features/item/ItemBox/ItemBoxContainer";
import { ItemModalContainer } from "@/features/item/ItemModal/ItemModalContainer";

type PresenterProps = {
	items: ListItemsResponse.AsObject | null;
	loading: boolean;
	error: string | null;
	selectedItem: Item.AsObject | null
	isModalOpened: boolean
	setIsModalOpened: React.Dispatch<React.SetStateAction<boolean>>;
	handleOnClickItem: (item: Item.AsObject) => void;
	handleOnClose: () => void;
};

const Presenter: React.FC<PresenterProps> = ({ items, loading, error, handleOnClickItem, isModalOpened, selectedItem, handleOnClose }) => {
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
						<ItemBoxContainer key={item.id} item={item} handleOnClickItem={handleOnClickItem} />
					))}
				</Box>
				<div>
					page: {items?.paging?.currentPage} / {items?.paging?.totalPages}
				</div>
			</Container>
			<ItemModalContainer isOpened={isModalOpened} item={selectedItem} handleOnClose={handleOnClose}  />
		</div >
	);
};

export default Presenter;
