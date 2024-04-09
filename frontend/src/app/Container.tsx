// src/app/Container.tsx
"use client";
import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "../hooks/hooks";
import { fetchItems } from "../store/itemSlice";
import Presenter from "./Presenter";

const Container = () => {
	const dispatch = useAppDispatch();
	const { items, loading, error } = useAppSelector((state) => state.item);

	useEffect(() => {
		dispatch(fetchItems());
	}, [dispatch]);

	return <Presenter items={items} loading={loading} error={error} />;
};

export default Container;
