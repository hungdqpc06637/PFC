package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCMarkingLocation struct {
	*BaseService
}

var PFCMarkingLoca = &PFCMarkingLocation{}

func (s *PFCMarkingLocation) GetAllPFCMarkingLocation(pfcModel *types.PFCModel) (*[]types.PFC_MarkingLocation, error) {
	var arrMarkingLocation *[]types.PFC_MarkingLocation
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(MarkingLocationID AS NVARCHAR(36)) AS MarkingLocationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_MarkingLocation
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
	).Scan(&arrMarkingLocation).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrMarkingLocation, nil
}

func (s *PFCMarkingLocation) InsertNewPFCMarkingLocation(req *types.PFC_MarkingLocation) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var MarkingLocationID string

	query := `
		INSERT INTO PFC_MarkingLocation(MarkingLocationID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.MarkingLocationID AS NVARCHAR(36)) AS MarkingLocationID
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
		&MarkingLocationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return MarkingLocationID, nil
}

func (s *PFCMarkingLocation) UpdatePFCMarkingLocation(req *types.PFC_MarkingLocation) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var MarkingLocationID string

	query := `
		UPDATE PFC_MarkingLocation
		SET Title = ?
		OUTPUT CAST(INSERTED.MarkingLocationID AS NVARCHAR(36))
		WHERE MarkingLocationID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.MarkingLocationID,
	).Scan(
		&MarkingLocationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return MarkingLocationID, nil
}

func (s *PFCMarkingLocation) DeletePFCMarkingLocation(req *types.PFC_MarkingLocation) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_MarkingLocation
		WHERE MarkingLocationID = ?
	`

	if err := tx.Exec(query,
		req.MarkingLocationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_MarkingLocation

func (s *PFCMarkingLocation) InsertNewPFCItemMarkingLocation(req *types.PFC_ItemMarkingLocation) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemMarkingLocationID string

	query := `
		INSERT INTO PFC_ItemMarkingLocation(ItemMarkingLocationID, MarkingLocationID, Component, ImageContent,TitleImage, Process, ItemIndex)
		OUTPUT CAST(INSERTED.ItemMarkingLocationID AS NVARCHAR(36)) AS ItemMarkingLocationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.MarkingLocationID,
		req.Component,
		req.ImageContent,
		req.TitleImage,
		req.Process,
		req.ItemIndex,
	).Scan(
		&ItemMarkingLocationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemMarkingLocationID, nil
}

func (s *PFCMarkingLocation) GetAllPFCItemMarkingLocation(pfcMarkingLocation *types.PFC_MarkingLocation) (*[]types.PFC_ItemMarkingLocation, error) {
	var arrItemMarkingLocation *[]types.PFC_ItemMarkingLocation
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemMarkingLocationID AS NVARCHAR(36)) AS ItemMarkingLocationID,
		CAST(MarkingLocationID AS NVARCHAR(36)) AS MarkingLocationID,
		Component, 
		ImageContent,
		TitleImage,
		Process,
		ItemIndex
	FROM PFC_ItemMarkingLocation
	WHERE MarkingLocationID = @MarkingLocationID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("MarkingLocationID", pfcMarkingLocation.MarkingLocationID),
	).Scan(&arrItemMarkingLocation).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemMarkingLocation, nil
}

func (s *PFCMarkingLocation) UpdatePFCItemMarkingLocation(req *types.PFC_ItemMarkingLocation) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemMarkingLocationID string

	query := `
		UPDATE PFC_ItemMarkingLocation
		SET Component = ?,
		ImageContent  = ?,
		TitleImage = ?,
		Process = ?,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemMarkingLocationID AS NVARCHAR(36)) AS ItemMarkingLocationID
		WHERE ItemMarkingLocationID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.TitleImage,
		req.Process,
		req.ItemIndex,
		req.ItemMarkingLocationID,
	).Scan(
		&ItemMarkingLocationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemMarkingLocationID, nil
}

func (s *PFCMarkingLocation) DeletePFCItemMarkingLocation(req *types.PFC_ItemMarkingLocation) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemMarkingLocation
		WHERE ItemMarkingLocationID = ?
	`

	if err := tx.Exec(query,
		req.ItemMarkingLocationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
