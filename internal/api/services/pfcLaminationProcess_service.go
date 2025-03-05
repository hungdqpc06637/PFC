package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCLaminationProcessService struct {
	*BaseService
}

var PFCLaminationProcess = &PFCLaminationProcessService{}

func (s *PFCLaminationProcessService) GetAllPFCLaminationProcess(requestParams *types.PFCModel) (*types.PFC_LaminationProcess, error) {
	var laminationProcess *types.PFC_LaminationProcess
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID,
		ModelType,
		ModelName,
		MaterialNumber,
		ComponentName,
		Vendor,
		Winding,
		Cooling,
		EndStep
	FROM PFC_LaminationProcess
	WHERE
		ModelType = @ModelType
		AND ModelName = @ModelName
		AND MaterialNumber = @MaterialNumber
`
	err = db.Raw(query,
		sql.Named("ModelType", requestParams.ModelType),
		sql.Named("ModelName", requestParams.ModelName),
		sql.Named("MaterialNumber", requestParams.MaterialNumber),
	).Scan(&laminationProcess).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()
	resLaminationProcess, _ := PFCMaterialDescription.GetAllPFCMaterialDescription(laminationProcess.LaminationProcessID)
	resAdhesiveType, _ := PFCAdhesiveType.GetAllPFCAdhesiveType(laminationProcess.LaminationProcessID)
	resAdhesiveOtherType, _ := PFCAdhesiveOtherType.GetAllPFCAdhesiveOtherType(laminationProcess.LaminationProcessID)
	resRolls, _ := PFCRoll.GetAllPFCRolls(laminationProcess.LaminationProcessID)
	laminationProcess.MaterialDescriptions = resLaminationProcess
	laminationProcess.AdhesiveType = resAdhesiveType
	laminationProcess.AdhesiveOtherType = resAdhesiveOtherType
	laminationProcess.Rolls = resRolls

	return laminationProcess, nil
}

func (s *PFCLaminationProcessService) InsertNewPFCLaminationProcess(req *types.PFC_LaminationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var laminationProcessID string

	query := `
		INSERT INTO PFC_LaminationProcess(LaminationProcessID, ModelType, ModelName, MaterialNumber, ComponentName, Vendor, Winding, Cooling, EndStep)
		OUTPUT CAST(INSERTED.LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID
		VALUES
		(NEWID(), ?, ?, ?, ?, ?, ?, ?, ?)
	`
	if err := tx.Raw(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
		req.ComponentName,
		req.Vendor,
		req.Winding,
		req.Cooling,
		req.EndStep,
	).Scan(
		&laminationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return laminationProcessID, nil
}

func (s *PFCLaminationProcessService) UpdatePFCLaminationProcess(req *types.PFC_LaminationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var laminationProcessID string

	query := `
		UPDATE PFC_LaminationProcess
		SET 
			ComponentName = ?, 
			Vendor = ?, 
			Winding = ?, 
			Cooling = ?, 
			EndStep = ?
		OUTPUT CAST(INSERTED.LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID
		WHERE 
			LaminationProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.ComponentName,
		req.Vendor,
		req.Winding,
		req.Cooling,
		req.EndStep,
		req.LaminationProcessID,
	).Scan(
		&laminationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return laminationProcessID, nil
}

func (s *PFCLaminationProcessService) DeletePFCLaminationProcess(req *types.PFC_LaminationProcess) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var laminationProcessID string

	query := `
		DELETE PFC_LaminationProcess
		OUTPUT CAST(DELETED.LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID
		WHERE LaminationProcessID = ?
	`
	if err := tx.Raw(
		query,
		req.LaminationProcessID,
	).Scan(
		&laminationProcessID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return laminationProcessID, nil
}

func (s *PFCLaminationProcessService) DeletePFCLaminationProcessByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_LaminationProcess
		WHERE LaminationProcessID = (SELECT TOP(1) LaminationProcessID FROM PFC_LaminationProcess WHERE ModelType = ? AND ModelName = ? AND MaterialNumber = ?)
	`

	if err := tx.Exec(query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
