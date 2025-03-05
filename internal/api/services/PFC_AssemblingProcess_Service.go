package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCAssemblingProcess struct {
	*BaseService
}

var PFCAssemblingProces = &PFCAssemblingProcess{}

func (s *PFCAssemblingProcess) GetAllPFCAssemblingProcess(pfcModel *types.PFCModel) (*[]types.PFC_AssemblingProcess, error) {
	var arrAssemblingProcess *[]types.PFC_AssemblingProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(AssemblingProcessID AS NVARCHAR(36)) AS AssemblingProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_AssemblingProcess
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
	).Scan(&arrAssemblingProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrAssemblingProcess, nil
}

func (s *PFCAssemblingProcess) InsertNewPFCAssemblingProcess(req *types.PFC_AssemblingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var AssemblingProcessID string

	query := `
		INSERT INTO PFC_AssemblingProcess
		(
		AssemblingProcessID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.AssemblingProcessID AS NVARCHAR(36)) AS AssemblingProcessID
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
		&AssemblingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return AssemblingProcessID, nil
}

func (s *PFCAssemblingProcess) UpdatePFCAssemblingProcess(req *types.PFC_AssemblingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var AssemblingProcessID string

	query := `
		UPDATE PFC_AssemblingProcess
		SET Title = ?
		OUTPUT CAST(INSERTED.AssemblingProcessID AS NVARCHAR(36))
		WHERE AssemblingProcessID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.AssemblingProcessID,
	).Scan(
		&AssemblingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return AssemblingProcessID, nil
}

func (s *PFCAssemblingProcess) DeletePFCAssemblingProcess(req *types.PFC_AssemblingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_AssemblingProcess
		WHERE AssemblingProcessID = ?
	`

	if err := tx.Exec(query,
		req.AssemblingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemAssemblingProcess

func (s *PFCAssemblingProcess) InsertNewPFCItemAssemblingProcess(req *types.PFC_ItemAssemblingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemAssemblingProcessID string

	query := `
		INSERT INTO PFC_ItemAssemblingProcess
		(
		ItemAssemblingProcessID ,AssemblingProcessID, TableRow1 ,ImagesContent, LastingLaceType,LastingClipSchedule,ChillerMoldSpecification,TotalWBSB
		)
		OUTPUT CAST(INSERTED.ItemAssemblingProcessID AS NVARCHAR(36)) AS ItemAssemblingProcessID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.AssemblingProcessID,
		req.TableRow1,
		req.ImagesContent,
		req.LastingLaceType,
		req.LastingClipSchedule,
		req.ChillerMoldSpecification,
		req.TotalWBSB,
	).Scan(
		&ItemAssemblingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemAssemblingProcessID, nil
}

func (s *PFCAssemblingProcess) GetAllPFCItemAssemblingProcess(pfcAssemblingProcess *types.PFC_AssemblingProcess) (*[]types.PFC_ItemAssemblingProcess, error) {
	var arrItemAssemblingProcess *[]types.PFC_ItemAssemblingProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemAssemblingProcessID AS NVARCHAR(36)) AS ItemAssemblingProcessID,
		CAST(AssemblingProcessID AS NVARCHAR(36)) AS AssemblingProcessID,
		 	TableRow1 ,ImagesContent, LastingLaceType,LastingClipSchedule,ChillerMoldSpecification,TotalWBSB
	FROM PFC_ItemAssemblingProcess
	WHERE AssemblingProcessID = @AssemblingProcessID
	`
	err = db.Raw(query,
		sql.Named("AssemblingProcessID", pfcAssemblingProcess.AssemblingProcessID),
	).Scan(&arrItemAssemblingProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemAssemblingProcess, nil
}

func (s *PFCAssemblingProcess) UpdatePFCItemAssemblingProcess(req *types.PFC_ItemAssemblingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemAssemblingProcessID string

	query := `
		UPDATE PFC_ItemAssemblingProcess
		SET 
			TableRow1 = ?  ,ImagesContent = ? , LastingLaceType = ? ,LastingClipSchedule = ? ,ChillerMoldSpecification = ? ,TotalWBSB = ? 
		OUTPUT CAST(INSERTED.ItemAssemblingProcessID AS NVARCHAR(36)) AS ItemAssemblingProcessID
		WHERE ItemAssemblingProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.TableRow1,
		req.ImagesContent,
		req.LastingLaceType,
		req.LastingClipSchedule,
		req.ChillerMoldSpecification,
		req.TotalWBSB,
		req.ItemAssemblingProcessID,
	).Scan(
		&ItemAssemblingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemAssemblingProcessID, nil
}

func (s *PFCAssemblingProcess) DeletePFCItemAssemblingProcess(req *types.PFC_ItemAssemblingProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemAssemblingProcess
		WHERE ItemAssemblingProcessID = ?
	`

	if err := tx.Exec(query,
		req.ItemAssemblingProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
