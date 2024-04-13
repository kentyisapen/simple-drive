// internal/domain/model/Item.go
package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ItemType はアイテムのタイプを定義します。
type ItemType string

// ItemType の定数です。
const (
	FileType      ItemType = "FILE"
	DirectoryType ItemType = "DIRECTORY"
)

// Item はアイテムを表すドメインモデルです。
type Item struct {
	ID             uuid.UUID `db:"id"`
	ParentID       uuid.UUID `db:"parent_id"` // NULLを許容するためにポインタ型を使用
	ThumbnailId    uuid.UUID `db:"thumbnail_id"`
	Name           string    `db:"name"`
	Type           ItemType  `db:"type"`
	Size           *int64    `db:"size"` // NULLを許容するためにポインタ型を使用
	CreatedAt      time.Time `db:"created_at"`
	LastModifiedAt time.Time `db:"last_modified_at"`
}

// ToPB はItem構造体をpb.Itemメッセージに変換します。
func (i *Item) ToPB() *pb.Item {
	return &pb.Item{
		Id:             i.ID.String(),
		ParentId:       i.ParentID.String(),
		ThumbnailId:    i.ThumbnailId.String(),
		Name:           i.Name,
		Type:           pb.ItemType(pb.ItemType_value[string(i.Type)]), // 型変換が必要になる場合があります
		Size:           wrapperspb.Int64(wrapperspb.Int64Value{Value: *i.Size}.Value),
		CreatedAt:      timestamppb.New(i.CreatedAt),
		LastModifiedAt: timestamppb.New(i.LastModifiedAt),
	}
}
