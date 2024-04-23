// src/app/Container.tsx
"use client";
import React, { useEffect, useState } from "react";
import { useAppDispatch, useAppSelector } from "../hooks/hooks";
import { fetchItems } from "../store/itemSlice";
import Presenter from "./Presenter";
import { Item } from "@/grpc/generated/simple-drive_pb";

const Container = () => {
	const dispatch = useAppDispatch();
	const { items, loading, error } = useAppSelector((state) => state.item);

	const [selectedItem, setSelectedItem] = useState<null | Item.AsObject>(null)
	const [modalIsOpened, setIsModalOpened] = useState(false)

	const handleOnClickItem = (item: Item.AsObject) => {
		setSelectedItem(item)
		setIsModalOpened(true)
	}

	const handleOnClose = () => {
		setIsModalOpened(false)
	}

	useEffect(() => {
		dispatch(fetchItems());
	}, [dispatch]);

	return <Presenter items={items} loading={loading} error={error} selectedItem={selectedItem} isModalOpened={modalIsOpened} setIsModalOpened={setIsModalOpened} handleOnClickItem={handleOnClickItem} handleOnClose={handleOnClose} />;
};

export default Container;
