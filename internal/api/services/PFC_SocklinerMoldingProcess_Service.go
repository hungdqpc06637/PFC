package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCSocklinerMoldingProcess struct {
	*BaseService
}

var PFCSocklinerMoldingProces = &PFCSocklinerMoldingProcess{}

func (s *PFCSocklinerMoldingProcess) GetAllPFCSocklinerMoldingProcess(pfcModel *types.PFCModel) (*[]types.PFC_SocklinerMoldingProcess, error) {
	var arrSocklinerMoldingProcess *[]types.PFC_SocklinerMoldingProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(SocklinerMoldingProcessID AS NVARCHAR(36)) AS SocklinerMoldingProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_SocklinerMoldingProcess
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
	).Scan(&arrSocklinerMoldingProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrSocklinerMoldingProcess, nil
}

func (s *PFCSocklinerMoldingProcess) InsertNewPFCSocklinerMoldingProcess(req *types.PFC_SocklinerMoldingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SocklinerMoldingProcessID string

	query := `
		INSERT INTO PFC_SocklinerMoldingProcess
		(
		SocklinerMoldingProcessID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.SocklinerMoldingProcessID AS NVARCHAR(36)) AS SocklinerMoldingProcessID
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
		&SocklinerMoldingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SocklinerMoldingProcessID, nil
}

func (s *PFCSocklinerMoldingProcess) UpdatePFCSocklinerMoldingProcess(req *types.PFC_SocklinerMoldingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SocklinerMoldingProcessID string

	query := `
		UPDATE PFC_SocklinerMoldingProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.SocklinerMoldingProcessID AS NVARCHAR(36))
		WHERE SocklinerMoldingProcessID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.SocklinerMoldingProcessID,
	).Scan(
		&SocklinerMoldingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SocklinerMoldingProcessID, nil
}

func (s *PFCSocklinerMoldingProcess) DeletePFCSocklinerMoldingProcess(req *types.PFC_SocklinerMoldingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_SocklinerMoldingProcess
		WHERE SocklinerMoldingProcessID = ?
	`

	if err := tx.Exec(query,
		req.SocklinerMoldingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemSocklinerMoldingProcess

func (s *PFCSocklinerMoldingProcess) InsertNewPFCItemSocklinerMoldingProcess(req *types.PFC_ItemSocklinerMoldingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSocklinerMoldingProcessID string

	query := `
		INSERT INTO PFC_ItemSocklinerMoldingProcess
		(
		ItemSocklinerMoldingProcessID ,SocklinerMoldingProcessID, RawMaterialName,Vendor,TableRow1,Remarks, RemarksImage, Model,Size 
		)
		OUTPUT CAST(INSERTED.ItemSocklinerMoldingProcessID AS NVARCHAR(36)) AS ItemSocklinerMoldingProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.SocklinerMoldingProcessID,
		req.RawMaterialName,
		req.Vendor,
		req.TableRow1,
		req.Remarks,
		req.RemarksImage,
		req.Model,
		req.Size,
	).Scan(
		&ItemSocklinerMoldingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSocklinerMoldingProcessID, nil
}

func (s *PFCSocklinerMoldingProcess) GetAllPFCItemSocklinerMoldingProcess(pfcSocklinerMoldingProcess *types.PFC_SocklinerMoldingProcess) (*[]types.PFC_ItemSocklinerMoldingProcess, error) {
	var arrItemSocklinerMoldingProcess *[]types.PFC_ItemSocklinerMoldingProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemSocklinerMoldingProcessID AS NVARCHAR(36)) AS ItemSocklinerMoldingProcessID,
		CAST(SocklinerMoldingProcessID AS NVARCHAR(36)) AS SocklinerMoldingProcessID,
		 	RawMaterialName,Vendor,TableRow1,Remarks, RemarksImage, Model,Size 
	FROM PFC_ItemSocklinerMoldingProcess
	WHERE SocklinerMoldingProcessID = @SocklinerMoldingProcessID
	`
	err = db.Raw(query,
		sql.Named("SocklinerMoldingProcessID", pfcSocklinerMoldingProcess.SocklinerMoldingProcessID),
	).Scan(&arrItemSocklinerMoldingProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemSocklinerMoldingProcess, nil
}

func (s *PFCSocklinerMoldingProcess) UpdatePFCItemSocklinerMoldingProcess(req *types.PFC_ItemSocklinerMoldingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSocklinerMoldingProcessID string

	query := `
		UPDATE PFC_ItemSocklinerMoldingProcess
		SET 
			RawMaterialName = ?,Vendor = ?,TableRow1 = ?,Remarks = ?, RemarksImage = ?, Model = ?,Size  = ?
		OUTPUT CAST(INSERTED.ItemSocklinerMoldingProcessID AS NVARCHAR(36)) AS ItemSocklinerMoldingProcessID
		WHERE ItemSocklinerMoldingProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.RawMaterialName,
		req.Vendor,
		req.TableRow1,
		req.Remarks,
		req.RemarksImage,
		req.Model,
		req.Size,
		req.ItemSocklinerMoldingProcessID,
	).Scan(
		&ItemSocklinerMoldingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSocklinerMoldingProcessID, nil
}

func (s *PFCSocklinerMoldingProcess) DeletePFCItemSocklinerMoldingProcess(req *types.PFC_ItemSocklinerMoldingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemSocklinerMoldingProcess
		WHERE ItemSocklinerMoldingProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemSocklinerMoldingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
