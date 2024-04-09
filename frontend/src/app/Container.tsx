import React, { useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '../hooks/hooks';
import { fetchItems } from '../store/itemSlice';
import Presenter from './Presenter';

const ItemListContainer = () => {
    const dispatch = useAppDispatch();
    const items = useAppSelector((state) => state.item.itemList?.toObject());
    const loading = useAppSelector((state) => state.item.loading);

    useEffect(() => {
        dispatch(fetchItems());
    }, [dispatch]);

    return <Presenter items={items} loading={loading} />;
};

export default ItemListContainer;