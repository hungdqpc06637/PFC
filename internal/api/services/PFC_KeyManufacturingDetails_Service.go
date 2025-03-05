package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCKeyManufacturingDetails struct {
	*BaseService
}

var PFCKeyManufacturingDetail = &PFCKeyManufacturingDetails{}

func (s *PFCKeyManufacturingDetails) GetAllPFCKeyManufacturingDetails(pfcModel *types.PFCModel) (*[]types.PFC_KeyManufacturingDetails, error) {
	var arrKeyManufacturingDetails *[]types.PFC_KeyManufacturingDetails
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(KeyManufacturingDetailsID AS NVARCHAR(36)) AS KeyManufacturingDetailsID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_KeyManufacturingDetails
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
	).Scan(&arrKeyManufacturingDetails).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrKeyManufacturingDetails, nil
}

func (s *PFCKeyManufacturingDetails) InsertNewPFCKeyManufacturingDetails(req *types.PFC_KeyManufacturingDetails) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var KeyManufacturingDetailsID string

	query := `
		INSERT INTO PFC_KeyManufacturingDetails
		(
		KeyManufacturingDetailsID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.KeyManufacturingDetailsID AS NVARCHAR(36)) AS KeyManufacturingDetailsID
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
		&KeyManufacturingDetailsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return KeyManufacturingDetailsID, nil
}

func (s *PFCKeyManufacturingDetails) UpdatePFCKeyManufacturingDetails(req *types.PFC_KeyManufacturingDetails) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var KeyManufacturingDetailsID string

	query := `
		UPDATE PFC_KeyManufacturingDetails
		SET Title = ?
		OUTPUT CAST(INSERTED.KeyManufacturingDetailsID AS NVARCHAR(36))
		WHERE KeyManufacturingDetailsID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.KeyManufacturingDetailsID,
	).Scan(
		&KeyManufacturingDetailsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return KeyManufacturingDetailsID, nil
}

func (s *PFCKeyManufacturingDetails) DeletePFCKeyManufacturingDetails(req *types.PFC_KeyManufacturingDetails) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_KeyManufacturingDetails
		WHERE KeyManufacturingDetailsID = ?
	`

	if err := tx.Exec(query,
		req.KeyManufacturingDetailsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemKeyManufacturingDetails

func (s *PFCKeyManufacturingDetails) InsertNewPFCItemKeyManufacturingDetails(req *types.PFC_ItemKeyManufacturingDetails) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemKeyManufacturingDetailsID string

	query := `
		INSERT INTO PFC_ItemKeyManufacturingDetails
		(
		ItemKeyManufacturingDetailsID ,KeyManufacturingDetailsID, KeyManufacturingProcess,DetailPicture, ProcessDetailKeyCheckPoint
		)
		OUTPUT CAST(INSERTED.ItemKeyManufacturingDetailsID AS NVARCHAR(36)) AS ItemKeyManufacturingDetailsID
		VALUES (NEWID(), ?, ? ,?,?)

	`
	if err := tx.Raw(
		query,
		req.KeyManufacturingDetailsID,
		req.KeyManufacturingProcess,
		req.DetailPicture,
		req.ProcessDetailKeyCheckPoint,
	).Scan(
		&ItemKeyManufacturingDetailsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemKeyManufacturingDetailsID, nil
}

func (s *PFCKeyManufacturingDetails) GetAllPFCItemKeyManufacturingDetails(pfcKeyManufacturingDetails *types.PFC_KeyManufacturingDetails) (*[]types.PFC_ItemKeyManufacturingDetails, error) {
	var arrItemKeyManufacturingDetails *[]types.PFC_ItemKeyManufacturingDetails
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemKeyManufacturingDetailsID AS NVARCHAR(36)) AS ItemKeyManufacturingDetailsID,
		CAST(KeyManufacturingDetailsID AS NVARCHAR(36)) AS KeyManufacturingDetailsID,
		 	 KeyManufacturingProcess,DetailPicture, ProcessDetailKeyCheckPoint
	FROM PFC_ItemKeyManufacturingDetails
	WHERE KeyManufacturingDetailsID = @KeyManufacturingDetailsID
	`
	err = db.Raw(query,
		sql.Named("KeyManufacturingDetailsID", pfcKeyManufacturingDetails.KeyManufacturingDetailsID),
	).Scan(&arrItemKeyManufacturingDetails).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemKeyManufacturingDetails, nil
}

func (s *PFCKeyManufacturingDetails) UpdatePFCItemKeyManufacturingDetails(req *types.PFC_ItemKeyManufacturingDetails) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemKeyManufacturingDetailsID string

	query := `
		UPDATE PFC_ItemKeyManufacturingDetails
		SET   KeyManufacturingProcess = ?,DetailPicture = ?, ProcessDetailKeyCheckPoint = ?
		OUTPUT CAST(INSERTED.ItemKeyManufacturingDetailsID AS NVARCHAR(36)) AS ItemKeyManufacturingDetailsID
		WHERE ItemKeyManufacturingDetailsID = ?
	`
	if err := tx.Raw(
		query,
		req.KeyManufacturingProcess,
		req.DetailPicture,
		req.ProcessDetailKeyCheckPoint,
		req.ItemKeyManufacturingDetailsID,
	).Scan(
		&ItemKeyManufacturingDetailsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemKeyManufacturingDetailsID, nil
}

func (s *PFCKeyManufacturingDetails) DeletePFCItemKeyManufacturingDetails(req *types.PFC_ItemKeyManufacturingDetails) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemKeyManufacturingDetails
		WHERE ItemKeyManufacturingDetailsID = ?
	`

	if err := tx.Exec(query,
		req.ItemKeyManufacturingDetailsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
