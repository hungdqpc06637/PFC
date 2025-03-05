package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCOutsolePressingProcess struct {
	*BaseService
}

var PFCOutsolePressingProces = &PFCOutsolePressingProcess{}

func (s *PFCOutsolePressingProcess) GetAllPFCOutsolePressingProcess(pfcModel *types.PFCModel) (*[]types.PFC_OutsolePressingProcess, error) {
	var arrOutsolePressingProcess *[]types.PFC_OutsolePressingProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(OutsolePressingProcessID AS NVARCHAR(36)) AS OutsolePressingProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_OutsolePressingProcess
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
	).Scan(&arrOutsolePressingProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrOutsolePressingProcess, nil
}

func (s *PFCOutsolePressingProcess) InsertNewPFCOutsolePressingProcess(req *types.PFC_OutsolePressingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var OutsolePressingProcessID string

	query := `
		INSERT INTO PFC_OutsolePressingProcess
		(
		OutsolePressingProcessID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.OutsolePressingProcessID AS NVARCHAR(36)) AS OutsolePressingProcessID
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
		&OutsolePressingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return OutsolePressingProcessID, nil
}

func (s *PFCOutsolePressingProcess) UpdatePFCOutsolePressingProcess(req *types.PFC_OutsolePressingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var OutsolePressingProcessID string

	query := `
		UPDATE PFC_OutsolePressingProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.OutsolePressingProcessID AS NVARCHAR(36))
		WHERE OutsolePressingProcessID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.OutsolePressingProcessID,
	).Scan(
		&OutsolePressingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return OutsolePressingProcessID, nil
}

func (s *PFCOutsolePressingProcess) DeletePFCOutsolePressingProcess(req *types.PFC_OutsolePressingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_OutsolePressingProcess
		WHERE OutsolePressingProcessID = ?
	`

	if err := tx.Exec(query,
		req.OutsolePressingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemOutsolePressingProcess

func (s *PFCOutsolePressingProcess) InsertNewPFCItemOutsolePressingProcess(req *types.PFC_ItemOutsolePressingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemOutsolePressingProcessID string

	query := `
		INSERT INTO PFC_ItemOutsolePressingProcess
		(
				ItemOutsolePressingProcessID ,OutsolePressingProcessID, TableRow1,TableRow2
		)
		OUTPUT CAST(INSERTED.ItemOutsolePressingProcessID AS NVARCHAR(36)) AS ItemOutsolePressingProcessID
		VALUES (NEWID(), ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.OutsolePressingProcessID,
		req.TableRow1,
		req.TableRow2,
	).Scan(
		&ItemOutsolePressingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemOutsolePressingProcessID, nil
}

func (s *PFCOutsolePressingProcess) GetAllPFCItemOutsolePressingProcess(pfcOutsolePressingProcess *types.PFC_OutsolePressingProcess) (*[]types.PFC_ItemOutsolePressingProcess, error) {
	var arrItemOutsolePressingProcess *[]types.PFC_ItemOutsolePressingProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemOutsolePressingProcessID AS NVARCHAR(36)) AS ItemOutsolePressingProcessID,
		CAST(OutsolePressingProcessID AS NVARCHAR(36)) AS OutsolePressingProcessID, TableRow1, TableRow2
	FROM PFC_ItemOutsolePressingProcess
	WHERE OutsolePressingProcessID = @OutsolePressingProcessID 
	`
	err = db.Raw(query,
		sql.Named("OutsolePressingProcessID", pfcOutsolePressingProcess.OutsolePressingProcessID),
	).Scan(&arrItemOutsolePressingProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemOutsolePressingProcess, nil
}

func (s *PFCOutsolePressingProcess) UpdatePFCItemOutsolePressingProcess(req *types.PFC_ItemOutsolePressingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemOutsolePressingProcessID string

	query := `
		UPDATE PFC_ItemOutsolePressingProcess
		SET 
			TableRow1 = ?, TableRow2 = ?
		OUTPUT CAST(INSERTED.ItemOutsolePressingProcessID AS NVARCHAR(36)) AS ItemOutsolePressingProcessID
		WHERE ItemOutsolePressingProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.TableRow1,
		req.TableRow2,
		req.ItemOutsolePressingProcessID,
	).Scan(
		&ItemOutsolePressingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemOutsolePressingProcessID, nil
}

func (s *PFCOutsolePressingProcess) DeletePFCItemOutsolePressingProcess(req *types.PFC_ItemOutsolePressingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemOutsolePressingProcess
		WHERE ItemOutsolePressingProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemOutsolePressingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
