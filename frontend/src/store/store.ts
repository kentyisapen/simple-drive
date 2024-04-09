// src/store/store.ts
import { configureStore } from "@reduxjs/toolkit";
import itemReducer from "./itemSlice";

export const store = configureStore({
	reducer: {
		item: itemReducer, // itemSliceからexportされたreducerを使用
	},
});

// ルートステートの型をエクスポート
export type RootState = ReturnType<typeof store.getState>;

// アプリケーションのディスパッチ関数の型をエクスポート
export type AppDispatch = typeof store.dispatch;
