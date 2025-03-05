package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCSocklinerGraphicProcess struct {
	*BaseService
}

var PFCSocklinerGraphicProces = &PFCSocklinerGraphicProcess{}

func (s *PFCSocklinerGraphicProcess) GetAllPFCSocklinerGraphicProcess(pfcModel *types.PFCModel) (*[]types.PFC_SocklinerGraphicProcess, error) {
	var arrSocklinerGraphicProcess *[]types.PFC_SocklinerGraphicProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(SocklinerGraphicProcessID AS NVARCHAR(36)) AS SocklinerGraphicProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_SocklinerGraphicProcess
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
	).Scan(&arrSocklinerGraphicProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrSocklinerGraphicProcess, nil
}

func (s *PFCSocklinerGraphicProcess) InsertNewPFCSocklinerGraphicProcess(req *types.PFC_SocklinerGraphicProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SocklinerGraphicProcessID string

	query := `
		INSERT INTO PFC_SocklinerGraphicProcess
		(
		SocklinerGraphicProcessID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.SocklinerGraphicProcessID AS NVARCHAR(36)) AS SocklinerGraphicProcessID
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
		&SocklinerGraphicProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SocklinerGraphicProcessID, nil
}

func (s *PFCSocklinerGraphicProcess) UpdatePFCSocklinerGraphicProcess(req *types.PFC_SocklinerGraphicProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SocklinerGraphicProcessID string

	query := `
		UPDATE PFC_SocklinerGraphicProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.SocklinerGraphicProcessID AS NVARCHAR(36))
		WHERE SocklinerGraphicProcessID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.SocklinerGraphicProcessID,
	).Scan(
		&SocklinerGraphicProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SocklinerGraphicProcessID, nil
}

func (s *PFCSocklinerGraphicProcess) DeletePFCSocklinerGraphicProcess(req *types.PFC_SocklinerGraphicProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_SocklinerGraphicProcess
		WHERE SocklinerGraphicProcessID = ?
	`

	if err := tx.Exec(query,
		req.SocklinerGraphicProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemSocklinerGraphicProcess

func (s *PFCSocklinerGraphicProcess) InsertNewPFCItemSocklinerGraphicProcess(req *types.PFC_ItemSocklinerGraphicProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSocklinerGraphicProcessID string

	query := `
		INSERT INTO PFC_ItemSocklinerGraphicProcess
		(
		ItemSocklinerGraphicProcessID ,SocklinerGraphicProcessID, ComponentName,SocklinerLogo,Vendor,TableRow1,Remarks, RemarksImage, ModelSize 
		)
		OUTPUT CAST(INSERTED.ItemSocklinerGraphicProcessID AS NVARCHAR(36)) AS ItemSocklinerGraphicProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.SocklinerGraphicProcessID,
		req.ComponentName,
		req.SocklinerLogo,
		req.Vendor,
		req.TableRow1,
		req.Remarks,
		req.RemarksImage,
		req.ModelSize,
	).Scan(
		&ItemSocklinerGraphicProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSocklinerGraphicProcessID, nil
}

func (s *PFCSocklinerGraphicProcess) GetAllPFCItemSocklinerGraphicProcess(pfcSocklinerGraphicProcess *types.PFC_SocklinerGraphicProcess) (*[]types.PFC_ItemSocklinerGraphicProcess, error) {
	var arrItemSocklinerGraphicProcess *[]types.PFC_ItemSocklinerGraphicProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemSocklinerGraphicProcessID AS NVARCHAR(36)) AS ItemSocklinerGraphicProcessID,
		CAST(SocklinerGraphicProcessID AS NVARCHAR(36)) AS SocklinerGraphicProcessID,
		 	ComponentName,SocklinerLogo,Vendor,TableRow1,Remarks, RemarksImage, ModelSize 
	FROM PFC_ItemSocklinerGraphicProcess
	WHERE SocklinerGraphicProcessID = @SocklinerGraphicProcessID
	`
	err = db.Raw(query,
		sql.Named("SocklinerGraphicProcessID", pfcSocklinerGraphicProcess.SocklinerGraphicProcessID),
	).Scan(&arrItemSocklinerGraphicProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemSocklinerGraphicProcess, nil
}

func (s *PFCSocklinerGraphicProcess) UpdatePFCItemSocklinerGraphicProcess(req *types.PFC_ItemSocklinerGraphicProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSocklinerGraphicProcessID string

	query := `
		UPDATE PFC_ItemSocklinerGraphicProcess
		SET 
			ComponentName= ?,SocklinerLogo= ?,Vendor= ?,TableRow1= ?,Remarks= ?, RemarksImage= ?, ModelSize = ?
		OUTPUT CAST(INSERTED.ItemSocklinerGraphicProcessID AS NVARCHAR(36)) AS ItemSocklinerGraphicProcessID
		WHERE ItemSocklinerGraphicProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.ComponentName,
		req.SocklinerLogo,
		req.Vendor,
		req.TableRow1,
		req.Remarks,
		req.RemarksImage,
		req.ModelSize,
		req.ItemSocklinerGraphicProcessID,
	).Scan(
		&ItemSocklinerGraphicProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSocklinerGraphicProcessID, nil
}

func (s *PFCSocklinerGraphicProcess) DeletePFCItemSocklinerGraphicProcess(req *types.PFC_ItemSocklinerGraphicProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemSocklinerGraphicProcess
		WHERE ItemSocklinerGraphicProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemSocklinerGraphicProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
