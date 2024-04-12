import { Item } from "@/grpc/generated/simple-drive_pb"
import { ItemBoxPresenter } from "./ItemBoxPresenter"

export type Props = {
    item: Item.AsObject
}

export const ItemBoxContainer = (props: Props) => {
    return ItemBoxPresenter(props)
}