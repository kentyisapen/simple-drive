// src/app/Presenter.tsx
"use client";
import React from "react";
import { ListItemsResponse } from "../grpc/generated/simple-drive_pb";

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
			{items?.itemsList?.map((item, index) => (
				<div key={index}>
					{item.name} - {item.type} - {item.size.value}
				</div>
			))}
		</div>
	);
};

export default Presenter;
