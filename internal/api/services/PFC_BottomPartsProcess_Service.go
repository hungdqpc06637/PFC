package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCBottomPartsProcess struct {
	*BaseService
}

var PFCBottomPartsProces = &PFCBottomPartsProcess{}

func (s *PFCBottomPartsProcess) GetAllPFCBottomPartsProcess(pfcModel *types.PFCModel) (*[]types.PFC_BottomPartsProcess, error) {
	var arrBottomPartsProcess *[]types.PFC_BottomPartsProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(BottomPartsProcessID AS NVARCHAR(36)) AS BottomPartsProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_BottomPartsProcess
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
	).Scan(&arrBottomPartsProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrBottomPartsProcess, nil
}

func (s *PFCBottomPartsProcess) InsertNewPFCBottomPartsProcess(req *types.PFC_BottomPartsProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var BottomPartsProcessID string

	query := `
		INSERT INTO PFC_BottomPartsProcess
		(
		BottomPartsProcessID , ModelType, ModelName, MaterialNumber, Title, ItemIndex
		)
		OUTPUT CAST(INSERTED.BottomPartsProcessID AS NVARCHAR(36)) AS BottomPartsProcessID
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
		&BottomPartsProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return BottomPartsProcessID, nil
}

func (s *PFCBottomPartsProcess) UpdatePFCBottomPartsProcess(req *types.PFC_BottomPartsProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var BottomPartsProcessID string

	query := `
		UPDATE PFC_BottomPartsProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.BottomPartsProcessID AS NVARCHAR(36))
		WHERE BottomPartsProcessID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.BottomPartsProcessID,
	).Scan(
		&BottomPartsProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return BottomPartsProcessID, nil
}

func (s *PFCBottomPartsProcess) DeletePFCBottomPartsProcess(req *types.PFC_BottomPartsProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_BottomPartsProcess
		WHERE BottomPartsProcessID = ?
	`

	if err := tx.Exec(query,
		req.BottomPartsProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_BottomPartsProcess

func (s *PFCBottomPartsProcess) InsertNewPFCItemBottomPartsProcess(req *types.PFC_ItemBottomPartsProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemBottomPartsProcessID string

	query := `
		INSERT INTO PFC_ItemBottomPartsProcess
		(
		ItemBottomPartsProcessID ,BottomPartsProcessID, Component, Material, Vendor, TableRow1, RemarksImages, RemarksSize
		)
		OUTPUT CAST(INSERTED.ItemBottomPartsProcessID AS NVARCHAR(36)) AS ItemBottomPartsProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.BottomPartsProcessID,
		req.Component,
		req.Material,
		req.Vendor,
		req.TableRow1,
		req.RemarksImages,
		req.RemarksSize,
	).Scan(
		&ItemBottomPartsProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemBottomPartsProcessID, nil
}

func (s *PFCBottomPartsProcess) GetAllPFCItemBottomPartsProcess(pfcBottomPartsProcess *types.PFC_BottomPartsProcess) (*[]types.PFC_ItemBottomPartsProcess, error) {
	var arrItemBottomPartsProcess *[]types.PFC_ItemBottomPartsProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemBottomPartsProcessID AS NVARCHAR(36)) AS ItemBottomPartsProcessID,
		CAST(BottomPartsProcessID AS NVARCHAR(36)) AS BottomPartsProcessID,
		 	Component, 
			Material, 
			Vendor, 
			TableRow1, 
			RemarksImages, 
			RemarksSize
	FROM PFC_ItemBottomPartsProcess
	WHERE BottomPartsProcessID = @BottomPartsProcessID
`
	err = db.Raw(query,
		sql.Named("BottomPartsProcessID", pfcBottomPartsProcess.BottomPartsProcessID),
	).Scan(&arrItemBottomPartsProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemBottomPartsProcess, nil
}

func (s *PFCBottomPartsProcess) UpdatePFCItemBottomPartsProcess(req *types.PFC_ItemBottomPartsProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemBottomPartsProcessID string

	query := `
		UPDATE PFC_ItemBottomPartsProcess
		SET 
		 	Component = ?, 
			Material = ?, 
			Vendor = ?, 
			TableRow1 = ?, 
			RemarksImages = ?, 
			RemarksSize = ?
		OUTPUT CAST(INSERTED.ItemBottomPartsProcessID AS NVARCHAR(36)) AS ItemBottomPartsProcessID
		WHERE ItemBottomPartsProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.Material,
		req.Vendor,
		req.TableRow1,
		req.RemarksImages,
		req.RemarksSize,
		req.ItemBottomPartsProcessID,
	).Scan(
		&ItemBottomPartsProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemBottomPartsProcessID, nil
}

func (s *PFCBottomPartsProcess) DeletePFCItemBottomPartsProcess(req *types.PFC_ItemBottomPartsProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemBottomPartsProcess
		WHERE ItemBottomPartsProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemBottomPartsProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
