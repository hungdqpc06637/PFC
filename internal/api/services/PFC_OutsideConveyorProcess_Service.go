package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCOutsideConveyorProcess struct {
	*BaseService
}

var PFCOutsideConveyorProces = &PFCOutsideConveyorProcess{}

func (s *PFCOutsideConveyorProcess) GetAllPFCOutsideConveyorProcess(pfcModel *types.PFCModel) (*[]types.PFC_OutsideConveyorProcess, error) {
	var arrOutsideConveyorProcess *[]types.PFC_OutsideConveyorProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(OutsideConveyorProcessID AS NVARCHAR(36)) AS OutsideConveyorProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_OutsideConveyorProcess
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
	).Scan(&arrOutsideConveyorProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrOutsideConveyorProcess, nil
}

func (s *PFCOutsideConveyorProcess) InsertNewPFCOutsideConveyorProcess(req *types.PFC_OutsideConveyorProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var OutsideConveyorProcessID string

	query := `
		INSERT INTO PFC_OutsideConveyorProcess
		(
		OutsideConveyorProcessID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.OutsideConveyorProcessID AS NVARCHAR(36)) AS OutsideConveyorProcessID
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
		&OutsideConveyorProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return OutsideConveyorProcessID, nil
}

func (s *PFCOutsideConveyorProcess) UpdatePFCOutsideConveyorProcess(req *types.PFC_OutsideConveyorProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var OutsideConveyorProcessID string

	query := `
		UPDATE PFC_OutsideConveyorProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.OutsideConveyorProcessID AS NVARCHAR(36))
		WHERE OutsideConveyorProcessID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.OutsideConveyorProcessID,
	).Scan(
		&OutsideConveyorProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return OutsideConveyorProcessID, nil
}

func (s *PFCOutsideConveyorProcess) DeletePFCOutsideConveyorProcess(req *types.PFC_OutsideConveyorProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_OutsideConveyorProcess
		WHERE OutsideConveyorProcessID = ?
	`

	if err := tx.Exec(query,
		req.OutsideConveyorProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemOutsideConveyorProcess

func (s *PFCOutsideConveyorProcess) InsertNewPFCItemOutsideConveyorProcess(req *types.PFC_ItemOutsideConveyorProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemOutsideConveyorProcessID string

	query := `
		INSERT INTO PFC_ItemOutsideConveyorProcess
		(
		ItemOutsideConveyorProcessID ,OutsideConveyorProcessID, ComponentName,Material,Vendor,TableRow1 
		)
		OUTPUT CAST(INSERTED.ItemOutsideConveyorProcessID AS NVARCHAR(36)) AS ItemOutsideConveyorProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.OutsideConveyorProcessID,
		req.ComponentName,
		req.Material,
		req.Vendor,
		req.TableRow1,
	).Scan(
		&ItemOutsideConveyorProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemOutsideConveyorProcessID, nil
}

func (s *PFCOutsideConveyorProcess) GetAllPFCItemOutsideConveyorProcess(pfcOutsideConveyorProcess *types.PFC_OutsideConveyorProcess) (*[]types.PFC_ItemOutsideConveyorProcess, error) {
	var arrItemOutsideConveyorProcess *[]types.PFC_ItemOutsideConveyorProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemOutsideConveyorProcessID AS NVARCHAR(36)) AS ItemOutsideConveyorProcessID,
		CAST(OutsideConveyorProcessID AS NVARCHAR(36)) AS OutsideConveyorProcessID,
		 	ComponentName,Material,Vendor,TableRow1
	FROM PFC_ItemOutsideConveyorProcess
	WHERE OutsideConveyorProcessID = @OutsideConveyorProcessID
	`
	err = db.Raw(query,
		sql.Named("OutsideConveyorProcessID", pfcOutsideConveyorProcess.OutsideConveyorProcessID),
	).Scan(&arrItemOutsideConveyorProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemOutsideConveyorProcess, nil
}

func (s *PFCOutsideConveyorProcess) UpdatePFCItemOutsideConveyorProcess(req *types.PFC_ItemOutsideConveyorProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemOutsideConveyorProcessID string

	query := `
		UPDATE PFC_ItemOutsideConveyorProcess
		SET 
			ComponentName= ?,Material= ?,Vendor= ?,TableRow1= ? 
		OUTPUT CAST(INSERTED.ItemOutsideConveyorProcessID AS NVARCHAR(36)) AS ItemOutsideConveyorProcessID
		WHERE ItemOutsideConveyorProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.ComponentName,
		req.Material,
		req.Vendor,
		req.TableRow1,
		req.ItemOutsideConveyorProcessID,
	).Scan(
		&ItemOutsideConveyorProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemOutsideConveyorProcessID, nil
}

func (s *PFCOutsideConveyorProcess) DeletePFCItemOutsideConveyorProcess(req *types.PFC_ItemOutsideConveyorProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemOutsideConveyorProcess
		WHERE ItemOutsideConveyorProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemOutsideConveyorProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
