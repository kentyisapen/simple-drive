import React from 'react';
import { Item, ListItemsResponse } from '../grpc/generated/simple-drive_pb';

type ItemListPresenterProps = {
    items: ListItemsResponse.AsObject;
    loading: boolean;
};

const ItemListPresenter: React.FC<ItemListPresenterProps> = ({ items, loading }) => {
    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <ul>
            {items.itemsList.map((item) => (
                <li key={item.id}>
                    {item.name} - {item.type} - Created At: {item.createdAt?.seconds}
                </li>
            ))}
        </ul>
    );
};

export default ItemListPresenter;