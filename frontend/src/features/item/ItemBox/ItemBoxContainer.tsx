import { Item } from "@/grpc/generated/simple-drive_pb"
import { ItemBoxPresenter } from "./ItemBoxPresenter"

export type Props = {
    item: Item.AsObject
    handleOnClickItem: (item: Item.AsObject) => void
}

export const ItemBoxContainer = (props: Props) => {
    return ItemBoxPresenter(props)
}