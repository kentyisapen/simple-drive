"use client";
import React from "react";
import Container from "./Container";
import { store } from "@/store/store";
import { Provider } from "react-redux";

const Page = () => {
	return (
		<Provider store={store}>
			<div>
				<h1>Item List</h1>
				<Container />
			</div>
		</Provider>
	);
};

export default Page;
