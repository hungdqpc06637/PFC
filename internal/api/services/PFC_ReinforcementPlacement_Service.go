package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCReinforcementPlacement struct {
	*BaseService
}

var PFCReinforcementPlacemen = &PFCReinforcementPlacement{}

func (s *PFCReinforcementPlacement) GetAllPFCReinforcementPlacement(pfcModel *types.PFCModel) (*[]types.PFC_ReinforcementPlacement, error) {
	var arrReinforcementPlacement *[]types.PFC_ReinforcementPlacement
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(ReinforcementPlacementID AS NVARCHAR(36)) AS ReinforcementPlacementID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_ReinforcementPlacement
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
	).Scan(&arrReinforcementPlacement).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrReinforcementPlacement, nil
}

func (s *PFCReinforcementPlacement) InsertNewPFCReinforcementPlacement(req *types.PFC_ReinforcementPlacement) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ReinforcementPlacementID string

	query := `
		INSERT INTO PFC_ReinforcementPlacement(ReinforcementPlacementID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.ReinforcementPlacementID AS NVARCHAR(36)) AS ReinforcementPlacementID
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
		&ReinforcementPlacementID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ReinforcementPlacementID, nil
}

func (s *PFCReinforcementPlacement) UpdatePFCReinforcementPlacement(req *types.PFC_ReinforcementPlacement) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ReinforcementPlacementID string

	query := `
		UPDATE PFC_ReinforcementPlacement
		SET Title = ?
		OUTPUT CAST(INSERTED.ReinforcementPlacementID AS NVARCHAR(36))
		WHERE ReinforcementPlacementID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.ReinforcementPlacementID,
	).Scan(
		&ReinforcementPlacementID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ReinforcementPlacementID, nil
}

func (s *PFCReinforcementPlacement) DeletePFCReinforcementPlacement(req *types.PFC_ReinforcementPlacement) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ReinforcementPlacement
		WHERE ReinforcementPlacementID = ?
	`

	if err := tx.Exec(query,
		req.ReinforcementPlacementID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ReinforcementPlacement

func (s *PFCReinforcementPlacement) InsertNewPFCItemReinforcementPlacement(req *types.PFC_ItemReinforcementPlacement) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemReinforcementPlacementID string

	query := `
		INSERT INTO PFC_ItemReinforcementPlacement(ItemReinforcementPlacementID, ReinforcementPlacementID, Component, ImageContent, Material, Adhesive,AttachingMethod,Temp,Pressure,Time, ItemIndex)
		OUTPUT CAST(INSERTED.ItemReinforcementPlacementID AS NVARCHAR(36)) AS ItemReinforcementPlacementID
		VALUES (NEWID(), ?, ?, ?, ?, ?,?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.ReinforcementPlacementID,
		req.Component,
		req.ImageContent,
		req.Material,
		req.Adhesive,
		req.AttachingMethod,
		req.Temp,
		req.Pressure,
		req.Time,
		req.ItemIndex,
	).Scan(
		&ItemReinforcementPlacementID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemReinforcementPlacementID, nil
}

func (s *PFCReinforcementPlacement) GetAllPFCItemReinforcementPlacement(pfcReinforcementPlacement *types.PFC_ReinforcementPlacement) (*[]types.PFC_ItemReinforcementPlacement, error) {
	var arrItemReinforcementPlacement *[]types.PFC_ItemReinforcementPlacement
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemReinforcementPlacementID AS NVARCHAR(36)) AS ItemReinforcementPlacementID,
		CAST(ReinforcementPlacementID AS NVARCHAR(36)) AS ReinforcementPlacementID,
		Component, 
		ImageContent,
		Material ,
		Adhesive,
		AttachingMethod,
		Temp,
		Pressure,
		Time,
		ItemIndex
	FROM PFC_ItemReinforcementPlacement
	WHERE ReinforcementPlacementID = @ReinforcementPlacementID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("ReinforcementPlacementID", pfcReinforcementPlacement.ReinforcementPlacementID),
	).Scan(&arrItemReinforcementPlacement).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemReinforcementPlacement, nil
}

func (s *PFCReinforcementPlacement) UpdatePFCItemReinforcementPlacement(req *types.PFC_ItemReinforcementPlacement) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemReinforcementPlacementID string

	query := `
		UPDATE PFC_ItemReinforcementPlacement
		SET Component = ?,
		ImageContent  = ?,
		Material = ? ,
		Adhesive = ?,
		AttachingMethod= ?,
		Temp = ?,
		Pressure = ?,
		Time = ?,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemReinforcementPlacementID AS NVARCHAR(36)) AS ItemReinforcementPlacementID
		WHERE ItemReinforcementPlacementID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.Material,
		req.Adhesive,
		req.AttachingMethod,
		req.Temp,
		req.Pressure,
		req.Time,
		req.ItemIndex,
		req.ItemReinforcementPlacementID,
	).Scan(
		&ItemReinforcementPlacementID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemReinforcementPlacementID, nil
}

func (s *PFCReinforcementPlacement) DeletePFCItemReinforcementPlacement(req *types.PFC_ItemReinforcementPlacement) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemReinforcementPlacement
		WHERE ItemReinforcementPlacementID = ?
	`

	if err := tx.Exec(query,
		req.ItemReinforcementPlacementID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
