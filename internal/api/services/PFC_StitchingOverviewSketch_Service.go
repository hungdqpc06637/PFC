package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCStitchingOverviewSketch struct {
	*BaseService
}

var PFCStitchingOverviewSketc = &PFCStitchingOverviewSketch{}

func (s *PFCStitchingOverviewSketch) GetAllPFCStitchingOverviewSketch(pfcModel *types.PFCModel) (*[]types.PFC_StitchingOverviewSketch, error) {
	var arrStitchingOverviewSketch *[]types.PFC_StitchingOverviewSketch
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(StitchingOverviewSketchID AS NVARCHAR(36)) AS StitchingOverviewSketchID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_StitchingOverviewSketch
	WHERE
		ModelType = @ModelType
		AND ModelName = @ModelName
		AND MaterialNumber = @MaterialNumber
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("ModelType", pfcModel.ModelType),
		sql.Named("ModelName", pfcModel.ModelName),
		sql.Named("MaterialNumber", pfcModel.MaterialNumber),
	).Scan(&arrStitchingOverviewSketch).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrStitchingOverviewSketch, nil
}

func (s *PFCStitchingOverviewSketch) InsertNewPFCStitchingOverviewSketch(req *types.PFC_StitchingOverviewSketch) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var StitchingOverviewSketchID string

	query := `
		INSERT INTO PFC_StitchingOverviewSketch(StitchingOverviewSketchID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.StitchingOverviewSketchID AS NVARCHAR(36)) AS StitchingOverviewSketchID
		VALUES (NEWID(), ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
		req.Title,
		req.ItemIndex,
	).Scan(
		&StitchingOverviewSketchID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return StitchingOverviewSketchID, nil
}

func (s *PFCStitchingOverviewSketch) UpdatePFCStitchingOverviewSketch(req *types.PFC_StitchingOverviewSketch) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var StitchingOverviewSketchID string

	query := `
		UPDATE PFC_StitchingOverviewSketch
		SET Title = ?
		OUTPUT CAST(INSERTED.StitchingOverviewSketchID AS NVARCHAR(36))
		WHERE StitchingOverviewSketchID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.StitchingOverviewSketchID,
	).Scan(
		&StitchingOverviewSketchID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return StitchingOverviewSketchID, nil
}

func (s *PFCStitchingOverviewSketch) DeletePFCStitchingOverviewSketch(req *types.PFC_StitchingOverviewSketch) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_StitchingOverviewSketch
		WHERE StitchingOverviewSketchID = ?
	`

	if err := tx.Exec(query,
		req.StitchingOverviewSketchID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_StitchingOverviewSketch

func (s *PFCStitchingOverviewSketch) InsertNewPFCItemStitchingOverviewSketch(req *types.PFC_ItemStitchingOverviewSketch) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemStitchingOverviewSketchID string

	query := `
		INSERT INTO PFC_ItemStitchingOverviewSketch(ItemStitchingOverviewSketchID, StitchingOverviewSketchID, Component, ImageContent, RightFoot, ItemIndex)
		OUTPUT CAST(INSERTED.ItemStitchingOverviewSketchID AS NVARCHAR(36)) AS ItemStitchingOverviewSketchID
		VALUES (NEWID(), ?, ?, ?,?,?)

	`
	if err := tx.Raw(
		query,
		req.StitchingOverviewSketchID,
		req.Component,
		req.ImageContent,
		req.RightFoot,
		req.ItemIndex,
	).Scan(
		&ItemStitchingOverviewSketchID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemStitchingOverviewSketchID, nil
}

func (s *PFCStitchingOverviewSketch) GetAllPFCItemStitchingOverviewSketch(pfcStitchingOverviewSketch *types.PFC_StitchingOverviewSketch) (*[]types.PFC_ItemStitchingOverviewSketch, error) {
	var arrItemStitchingOverviewSketch *[]types.PFC_ItemStitchingOverviewSketch
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemStitchingOverviewSketchID AS NVARCHAR(36)) AS ItemStitchingOverviewSketchID,
		CAST(StitchingOverviewSketchID AS NVARCHAR(36)) AS StitchingOverviewSketchID,
		Component, 
		ImageContent,
		RightFoot ,
		ItemIndex
	FROM PFC_ItemStitchingOverviewSketch
	WHERE StitchingOverviewSketchID = @StitchingOverviewSketchID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("StitchingOverviewSketchID", pfcStitchingOverviewSketch.StitchingOverviewSketchID),
	).Scan(&arrItemStitchingOverviewSketch).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemStitchingOverviewSketch, nil
}

func (s *PFCStitchingOverviewSketch) UpdatePFCItemStitchingOverviewSketch(req *types.PFC_ItemStitchingOverviewSketch) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemStitchingOverviewSketchID string

	query := `
		UPDATE PFC_ItemStitchingOverviewSketch
		SET Component = ?,
		ImageContent  = ?,
		RightFoot = ?  ,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemStitchingOverviewSketchID AS NVARCHAR(36)) AS ItemStitchingOverviewSketchID
		WHERE ItemStitchingOverviewSketchID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.RightFoot,
		req.ItemIndex,
		req.ItemStitchingOverviewSketchID,
	).Scan(
		&ItemStitchingOverviewSketchID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemStitchingOverviewSketchID, nil
}

func (s *PFCStitchingOverviewSketch) DeletePFCItemStitchingOverviewSketch(req *types.PFC_ItemStitchingOverviewSketch) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemStitchingOverviewSketch
		WHERE ItemStitchingOverviewSketchID = ?
	`

	if err := tx.Exec(query,
		req.ItemStitchingOverviewSketchID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
