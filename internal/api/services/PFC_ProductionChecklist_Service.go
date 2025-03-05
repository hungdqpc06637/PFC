package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCProductionChecklist struct {
	*BaseService
}

var PFCProductionChecklis = &PFCProductionChecklist{}

func (s *PFCProductionChecklist) GetAllPFCProductionChecklist(pfcModel *types.PFCModel) (*[]types.PFC_ProductionChecklist, error) {
	var arrProductionChecklist *[]types.PFC_ProductionChecklist
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(ProductionChecklistID AS NVARCHAR(36)) AS ProductionChecklistID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_ProductionChecklist
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
	).Scan(&arrProductionChecklist).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrProductionChecklist, nil
}

func (s *PFCProductionChecklist) InsertNewPFCProductionChecklist(req *types.PFC_ProductionChecklist) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ProductionChecklistID string

	query := `
		INSERT INTO PFC_ProductionChecklist
		(
		ProductionChecklistID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.ProductionChecklistID AS NVARCHAR(36)) AS ProductionChecklistID
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
		&ProductionChecklistID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ProductionChecklistID, nil
}

func (s *PFCProductionChecklist) UpdatePFCProductionChecklist(req *types.PFC_ProductionChecklist) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ProductionChecklistID string

	query := `
		UPDATE PFC_ProductionChecklist
		SET Title = ?
		OUTPUT CAST(INSERTED.ProductionChecklistID AS NVARCHAR(36))
		WHERE ProductionChecklistID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.ProductionChecklistID,
	).Scan(
		&ProductionChecklistID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ProductionChecklistID, nil
}

func (s *PFCProductionChecklist) DeletePFCProductionChecklist(req *types.PFC_ProductionChecklist) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ProductionChecklist
		WHERE ProductionChecklistID = ?
	`

	if err := tx.Exec(query,
		req.ProductionChecklistID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemProductionChecklist

func (s *PFCProductionChecklist) InsertNewPFCItemProductionChecklist(req *types.PFC_ItemProductionChecklist) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemProductionChecklistID string

	query := `
		INSERT INTO PFC_ItemProductionChecklist
		(
		ItemProductionChecklistID ,ProductionChecklistID, ShoeLacingInstruction,OtherAccessoryInfo,IDSMeasurement,InnerBoxSchedule,LaceSchedule, 
		 WrappingPaperSchedule, ArchCookieSchedule,CardboardFootForm, MoldedFootForm ,StuffingPaper
		)
		OUTPUT CAST(INSERTED.ItemProductionChecklistID AS NVARCHAR(36)) AS ItemProductionChecklistID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.ProductionChecklistID,
		req.ShoeLacingInstruction,
		req.OtherAccessoryInfo,
		req.IDSMeasurement,
		req.InnerBoxSchedule,
		req.LaceSchedule,
		req.WrappingPaperSchedule,
		req.ArchCookieSchedule,
		req.CardboardFootForm,
		req.MoldedFootForm,
		req.StuffingPaper,
	).Scan(
		&ItemProductionChecklistID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemProductionChecklistID, nil
}

func (s *PFCProductionChecklist) GetAllPFCItemProductionChecklist(pfcProductionChecklist *types.PFC_ProductionChecklist) (*[]types.PFC_ItemProductionChecklist, error) {
	var arrItemProductionChecklist *[]types.PFC_ItemProductionChecklist
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemProductionChecklistID AS NVARCHAR(36)) AS ItemProductionChecklistID,
		CAST(ProductionChecklistID AS NVARCHAR(36)) AS ProductionChecklistID,
		 	ShoeLacingInstruction,OtherAccessoryInfo,IDSMeasurement,InnerBoxSchedule,LaceSchedule,  WrappingPaperSchedule, ArchCookieSchedule,CardboardFootForm, MoldedFootForm ,StuffingPaper
	FROM PFC_ItemProductionChecklist
	WHERE ProductionChecklistID = @ProductionChecklistID
	`
	err = db.Raw(query,
		sql.Named("ProductionChecklistID", pfcProductionChecklist.ProductionChecklistID),
	).Scan(&arrItemProductionChecklist).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemProductionChecklist, nil
}

func (s *PFCProductionChecklist) UpdatePFCItemProductionChecklist(req *types.PFC_ItemProductionChecklist) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemProductionChecklistID string

	query := `
		UPDATE PFC_ItemProductionChecklist
		SET 
			ShoeLacingInstruction = ? ,OtherAccessoryInfo = ? ,IDSMeasurement = ? ,InnerBoxSchedule = ? ,LaceSchedule = ? , 
			WrappingPaperSchedule = ? , ArchCookieSchedule = ? ,CardboardFootForm = ? , MoldedFootForm = ?  ,StuffingPaper = ? 
		OUTPUT CAST(INSERTED.ItemProductionChecklistID AS NVARCHAR(36)) AS ItemProductionChecklistID
		WHERE ItemProductionChecklistID = ?
	`
	if err := tx.Raw(
		query,
		req.ShoeLacingInstruction,
		req.OtherAccessoryInfo,
		req.IDSMeasurement,
		req.InnerBoxSchedule,
		req.LaceSchedule,
		req.WrappingPaperSchedule,
		req.ArchCookieSchedule,
		req.CardboardFootForm,
		req.MoldedFootForm,
		req.StuffingPaper,
		req.ItemProductionChecklistID,
	).Scan(
		&ItemProductionChecklistID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemProductionChecklistID, nil
}

func (s *PFCProductionChecklist) DeletePFCItemProductionChecklist(req *types.PFC_ItemProductionChecklist) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemProductionChecklist
		WHERE ItemProductionChecklistID = ?
	`

	if err := tx.Exec(query,
		req.ItemProductionChecklistID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
