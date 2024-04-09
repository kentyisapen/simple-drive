// src/store/itemSlice.ts
import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import { ListItemsResponse } from "../grpc/generated/simple-drive_pb";

// ステートの初期値を定義
interface ItemState {
	items: ListItemsResponse.AsObject | null;
	loading: boolean;
	error: string | null;
}

const initialState: ItemState = {
	items: null,
	loading: false,
	error: null,
};

// アイテム取得の非同期関数
export const fetchItems = createAsyncThunk("items/fetchItems", async () => {
	const response = await fetch("/api/items");
	return (await response.json()) as ListItemsResponse.AsObject;
});

const itemSlice = createSlice({
	name: "item",
	initialState,
	reducers: {},
	extraReducers: (builder) => {
		builder.addCase(fetchItems.pending, (state) => {
			state.loading = true;
		});
		builder.addCase(
			fetchItems.fulfilled,
			(state, action: PayloadAction<ListItemsResponse.AsObject>) => {
				state.items = action.payload;
				state.loading = false;
			}
		);
		builder.addCase(fetchItems.rejected, (state, action) => {
			state.loading = false;
			state.error = action.error.message || "Something went wrong";
		});
	},
});

export default itemSlice.reducer;
