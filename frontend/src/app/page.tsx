"use client";
import React from "react";
import Container from "./Container";
import { store } from "@/store/store";
import { Provider } from "react-redux";

const Page = () => {
	return (
		<Provider store={store}>
			<div>
				<Container />
			</div>
		</Provider>
	);
};

export default Page;
