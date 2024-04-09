import { ListItemsRequest, ListItemsResponse } from '@/grpc/generated/simple-drive_pb';
import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';

import client from "../grpc/client"

// ステートの初期値を定義
interface ItemState {
    itemList: ListItemsResponse | null;
    loading: boolean;
    error: string | null;
}

const initialState: ItemState = {
    itemList: null,
    loading: false,
    error: null,
};

// スライスの定義
const itemSlice = createSlice({
    name: 'item',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchItems.pending, (state) => {
            state.loading = true;
            state.error = null;
        });
        builder.addCase(fetchItems.fulfilled, (state, action: PayloadAction<ListItemsResponse>) => {
            state.itemList = action.payload;
            state.loading = false;
        });
        builder.addCase(fetchItems.rejected, (state, action) => {
            state.loading = false;
            state.error = action.error.message || 'Something went wrong';
        });
    },
});

export const fetchItems = createAsyncThunk<ListItemsResponse>(
    'items/fetchItems',
    async (_, { rejectWithValue }) => {
        const request = new ListItemsRequest();

        return new Promise<ListItemsResponse>((resolve, reject) => {
            client.listItems(request, (error, response) => {
                if (error) {
                    reject(rejectWithValue(error));
                } else if (!response) {
                    reject(rejectWithValue(new Error('Response is null')));
                } else {
                    resolve(response); // response は ListItemsResponse 型であることが保証されている
                }
            });
        });
    }
);

export default itemSlice.reducer