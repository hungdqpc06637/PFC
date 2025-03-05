package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCSecondProcess struct {
	*BaseService
}

var PFCSecondProces = &PFCSecondProcess{}

func (s *PFCSecondProcess) GetAllPFCSecondProcess(pfcModel *types.PFCModel) (*[]types.PFC_SecondProcess, error) {
	var arrSecondProcess *[]types.PFC_SecondProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(SecondProcessID AS NVARCHAR(36)) AS SecondProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_SecondProcess
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
	).Scan(&arrSecondProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrSecondProcess, nil
}

func (s *PFCSecondProcess) InsertNewPFCSecondProcess(req *types.PFC_SecondProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SecondProcessID string

	query := `
		INSERT INTO PFC_SecondProcess(SecondProcessID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.SecondProcessID AS NVARCHAR(36)) AS SecondProcessID
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
		&SecondProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SecondProcessID, nil
}

func (s *PFCSecondProcess) UpdatePFCSecondProcess(req *types.PFC_SecondProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SecondProcessID string

	query := `
		UPDATE PFC_SecondProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.SecondProcessID AS NVARCHAR(36))
		WHERE SecondProcessID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.SecondProcessID,
	).Scan(
		&SecondProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SecondProcessID, nil
}

func (s *PFCSecondProcess) DeletePFCSecondProcess(req *types.PFC_SecondProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_SecondProcess
		WHERE SecondProcessID = ?
	`

	if err := tx.Exec(query,
		req.SecondProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_SecondProcess

func (s *PFCSecondProcess) InsertNewPFCItemSecondProcess(req *types.PFC_ItemSecondProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSecondProcessID string

	query := `
		INSERT INTO PFC_ItemSecondProcess(ItemSecondProcessID, SecondProcessID, Component, ImageContent, Material,Method, ItemIndex)
		OUTPUT CAST(INSERTED.ItemSecondProcessID AS NVARCHAR(36)) AS ItemSecondProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?,?)

	`
	if err := tx.Raw(
		query,
		req.SecondProcessID,
		req.Component,
		req.ImageContent,
		req.Material,
		req.Method,
		req.ItemIndex,
	).Scan(
		&ItemSecondProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSecondProcessID, nil
}

func (s *PFCSecondProcess) GetAllPFCItemSecondProcess(pfcSecondProcess *types.PFC_SecondProcess) (*[]types.PFC_ItemSecondProcess, error) {
	var arrItemSecondProcess *[]types.PFC_ItemSecondProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemSecondProcessID AS NVARCHAR(36)) AS ItemSecondProcessID,
		CAST(SecondProcessID AS NVARCHAR(36)) AS SecondProcessID,
		Component, 
		ImageContent,
		Material ,
		Method,
		ItemIndex
	FROM PFC_ItemSecondProcess
	WHERE SecondProcessID = @SecondProcessID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("SecondProcessID", pfcSecondProcess.SecondProcessID),
	).Scan(&arrItemSecondProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemSecondProcess, nil
}

func (s *PFCSecondProcess) UpdatePFCItemSecondProcess(req *types.PFC_ItemSecondProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSecondProcessID string

	query := `
		UPDATE PFC_ItemSecondProcess
		SET Component = ?,
		ImageContent  = ?,
		Material = ? ,
		Method = ?,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemSecondProcessID AS NVARCHAR(36)) AS ItemSecondProcessID
		WHERE ItemSecondProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.Material,
		req.Method,
		req.ItemIndex,
		req.ItemSecondProcessID,
	).Scan(
		&ItemSecondProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSecondProcessID, nil
}

func (s *PFCSecondProcess) DeletePFCItemSecondProcess(req *types.PFC_ItemSecondProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemSecondProcess
		WHERE ItemSecondProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemSecondProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
