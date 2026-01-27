package rows

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/models"
)

func InsertDeleted(ctx context.Context, row any, sourceTable string, rowID int64) error {
	jsonRow, err := json.Marshal(row)
	if err != nil {
		return huma.Error500InternalServerError(fmt.Sprintf("%s deleted, but not backed up in database.", sourceTable))
	}

	err = db.Queries.InsertDeletedRow(ctx, queries.InsertDeletedRowParams{
		SourceTable: sourceTable,
		SourceID:    strconv.FormatInt(rowID, 10),
		Data:        json.RawMessage(jsonRow),
	})
	if err != nil {
		return models.WrapDBErr(err)
	}

	return nil
}
