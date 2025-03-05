package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCLogoApplicationProcess struct {
	*BaseService
}

var PFCLogoApplicationProces = &PFCLogoApplicationProcess{}

func (s *PFCLogoApplicationProcess) GetAllPFCLogoApplicationProcess(pfcModel *types.PFCModel) (*[]types.PFC_LogoApplicationProcess, error) {
	var arrLogoApplicationProcess *[]types.PFC_LogoApplicationProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(LogoApplicationProcessID AS NVARCHAR(36)) AS LogoApplicationProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_LogoApplicationProcess
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
	).Scan(&arrLogoApplicationProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrLogoApplicationProcess, nil
}

func (s *PFCLogoApplicationProcess) InsertNewPFCLogoApplicationProcess(req *types.PFC_LogoApplicationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var LogoApplicationProcessID string

	query := `
		INSERT INTO PFC_LogoApplicationProcess
		(
		LogoApplicationProcessID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.LogoApplicationProcessID AS NVARCHAR(36)) AS LogoApplicationProcessID
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
		&LogoApplicationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return LogoApplicationProcessID, nil
}

func (s *PFCLogoApplicationProcess) UpdatePFCLogoApplicationProcess(req *types.PFC_LogoApplicationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var LogoApplicationProcessID string

	query := `
		UPDATE PFC_LogoApplicationProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.LogoApplicationProcessID AS NVARCHAR(36))
		WHERE LogoApplicationProcessID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.LogoApplicationProcessID,
	).Scan(
		&LogoApplicationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return LogoApplicationProcessID, nil
}

func (s *PFCLogoApplicationProcess) DeletePFCLogoApplicationProcess(req *types.PFC_LogoApplicationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_LogoApplicationProcess
		WHERE LogoApplicationProcessID = ?
	`

	if err := tx.Exec(query,
		req.LogoApplicationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemLogoApplicationProcess

func (s *PFCLogoApplicationProcess) InsertNewPFCItemLogoApplicationProcess(req *types.PFC_ItemLogoApplicationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemLogoApplicationProcessID string

	query := `
		INSERT INTO PFC_ItemLogoApplicationProcess
		(
		ItemLogoApplicationProcessID ,LogoApplicationProcessID, ComponentName, LogoSockliner,Vendor,TableRow1,Remarks, RemarksImage, Size 
		)
		OUTPUT CAST(INSERTED.ItemLogoApplicationProcessID AS NVARCHAR(36)) AS ItemLogoApplicationProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.LogoApplicationProcessID,
		req.ComponentName,
		req.LogoSockliner,
		req.Vendor,
		req.TableRow1,
		req.Remarks,
		req.RemarksImage,
		req.Size,
	).Scan(
		&ItemLogoApplicationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemLogoApplicationProcessID, nil
}

func (s *PFCLogoApplicationProcess) GetAllPFCItemLogoApplicationProcess(pfcLogoApplicationProcess *types.PFC_LogoApplicationProcess) (*[]types.PFC_ItemLogoApplicationProcess, error) {
	var arrItemLogoApplicationProcess *[]types.PFC_ItemLogoApplicationProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemLogoApplicationProcessID AS NVARCHAR(36)) AS ItemLogoApplicationProcessID,
		CAST(LogoApplicationProcessID AS NVARCHAR(36)) AS LogoApplicationProcessID,
		 	ComponentName, LogoSockliner,Vendor,TableRow1,Remarks, RemarksImage, Size 
	FROM PFC_ItemLogoApplicationProcess
	WHERE LogoApplicationProcessID = @LogoApplicationProcessID
	`
	err = db.Raw(query,
		sql.Named("LogoApplicationProcessID", pfcLogoApplicationProcess.LogoApplicationProcessID),
	).Scan(&arrItemLogoApplicationProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemLogoApplicationProcess, nil
}

func (s *PFCLogoApplicationProcess) UpdatePFCItemLogoApplicationProcess(req *types.PFC_ItemLogoApplicationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemLogoApplicationProcessID string

	query := `
		UPDATE PFC_ItemLogoApplicationProcess
		SET 
			ComponentName  = ?, LogoSockliner = ?,Vendor = ?,TableRow1 = ?,Remarks = ?, RemarksImage = ?, Size  = ?
		OUTPUT CAST(INSERTED.ItemLogoApplicationProcessID AS NVARCHAR(36)) AS ItemLogoApplicationProcessID
		WHERE ItemLogoApplicationProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.ComponentName,
		req.LogoSockliner,
		req.Vendor,
		req.TableRow1,
		req.Remarks,
		req.RemarksImage,
		req.Size,
		req.ItemLogoApplicationProcessID,
	).Scan(
		&ItemLogoApplicationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemLogoApplicationProcessID, nil
}

func (s *PFCLogoApplicationProcess) DeletePFCItemLogoApplicationProcess(req *types.PFC_ItemLogoApplicationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemLogoApplicationProcess
		WHERE ItemLogoApplicationProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemLogoApplicationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
