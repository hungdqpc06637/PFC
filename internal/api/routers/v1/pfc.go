package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPFCModelRouter(router *gin.RouterGroup) {
	router.POST("/uploadFilePFCModel", controllers.UploadFile.UploadFilePFCModel)
	router.POST("/uploadFilePFCModelFromFolderPFCModel", controllers.UploadFile.UploadFilePFCModelFromFolderPFCModel)
	router.GET("/downloadFilePFCModel", controllers.UploadFile.DownloadFilePFCModel)
	router.GET("/downloadFilePFCModelFromFolderPFCModel", controllers.UploadFile.DownloadFilePFCModelFromFolderPFCModel)
	router.DELETE("/deleteFilePFCModel", controllers.UploadFile.DeleteFilePFCModel)
	router.DELETE("/deleteFilePFCModelFromFolderPFCModel", controllers.UploadFile.DeleteFilePFCModelFromFolderPFCModel)
	router.DELETE("/deleteModelDirectory", controllers.UploadFile.DeleteModelDirectory)
	//PFC Model
	router.GET("/pfcModel", controllers.PFCModel.GetAllPFCModel)
	router.POST("/pfcModel", controllers.PFCModel.InsertNewPFCModel)
	router.PUT("/pfcModel", controllers.PFCModel.UpdatePFCModel)
	router.DELETE("/pfcModel", controllers.PFCModel.DeletePFCModel)

	//PFC Page Setup
	router.POST("/getPFCPageSetup", controllers.PFCPageSetup.GetPFCPageSetup) //The payload object should be used with the POST method for the GET feature
	router.POST("/pfcPageSetup", controllers.PFCPageSetup.InsertNewPFCPageSetup)
	router.PUT("/pfcPageSetup", controllers.PFCPageSetup.UpdatePFCPageSetup)
	router.DELETE("/pfcPageSetupByModelID", controllers.PFCPageSetup.DeletePFCPageSetupByModelID)

	//PFC Lamination Process
	router.POST("/getPFCLaminationProcess", controllers.PFCLaminationProcess.GetAllPFCLaminationProcess) //The payload object should be used with the POST method for the GET feature
	router.POST("/pfcLaminationProcess", controllers.PFCLaminationProcess.InsertNewPFCLaminationProcess)
	router.PUT("/pfcLaminationProcess", controllers.PFCLaminationProcess.UpdatePFCLaminationProcess)
	router.DELETE("/pfcLaminationProcess", controllers.PFCLaminationProcess.DeletePFCLaminationProcess)
	router.DELETE("/pfcLaminationProcessByModelID", controllers.PFCLaminationProcess.DeletePFCLaminationProcessByModelID)

	//PFC Material Description Service
	router.POST("/pfcMaterialDescription", controllers.PFCMaterialDescription.InsertPFCMaterialDescription)
	router.PUT("/pfcMaterialDescription", controllers.PFCMaterialDescription.UpdatePFCMaterialDescription)
	router.DELETE("/pfcMaterialDescription", controllers.PFCMaterialDescription.DeletePFCMaterialDescription)
	router.DELETE("/pfcMaterialDescriptionByModelID", controllers.PFCMaterialDescription.DeletePFCMaterialDescriptionByModelID)

	//PFC Adhesive Type
	router.POST("/pfcAdhesiveType", controllers.PFCAdhesiveType.InsertPFCAdhesiveType)
	router.PUT("/pfcAdhesiveType", controllers.PFCAdhesiveType.UpdatePFCAdhesiveType)
	router.DELETE("/pfcAdhesiveType", controllers.PFCAdhesiveType.DeletePFCAdhesiveType)
	router.DELETE("/pfcAdhesiveTypeByModelID", controllers.PFCAdhesiveType.DeletePFCAdhesiveTypeByModelID)

	//PFC Adhesive Other Type
	router.POST("/pfcAdhesiveOtherType", controllers.PFCAdhesiveOtherType.InsertPFCAdhesiveOtherType)
	router.PUT("/pfcAdhesiveOtherType", controllers.PFCAdhesiveOtherType.UpdatePFCAdhesiveOtherType)
	router.DELETE("/pfcAdhesiveOtherType", controllers.PFCAdhesiveOtherType.DeletePFCAdhesiveOtherType)
	router.DELETE("/pfcAdhesiveOtherTypeByModelID", controllers.PFCAdhesiveOtherType.DeletePFCAdhesiveOtherTypeByModelID)

	//PFC Roll
	router.POST("/pfcRoll", controllers.PFCRoll.InsertPFCRoll)
	router.PUT("/pfcRoll", controllers.PFCRoll.UpdatePFCRoll)
	router.DELETE("/pfcRoll", controllers.PFCRoll.DeletePFCRoll)
	router.DELETE("/pfcRollByModelID", controllers.PFCRoll.DeletePFCRollByModelID)

	//PFC Upper Cutting Die Schedule
	router.POST("/getPFCUpperCuttingDieSchedule", controllers.PFCUpperCuttingDieSchedule.GetAllPFCUpperCuttingDieSchedule)
	router.POST("/pfcUpperCuttingDieSchedule", controllers.PFCUpperCuttingDieSchedule.InsertPFCUpperCuttingDieSchedule)
	router.PUT("/pfcUpperCuttingDieSchedule", controllers.PFCUpperCuttingDieSchedule.UpdatePFCUpperCuttingDieSchedule)
	router.DELETE("/pfcUpperCuttingDieSchedule", controllers.PFCUpperCuttingDieSchedule.DeletePFCUpperCuttingDieSchedule)
	router.DELETE("/pfcUpperCuttingDieScheduleByModelID", controllers.PFCUpperCuttingDieSchedule.DeletePFCUpperCuttingDieScheduleByModelID)

	//PFC Item Upper Cutting Die Schedule
	router.POST("/getPFCItemUpperCuttingDieScheduleByModelID", controllers.PFCItemUpperCuttingDieSchedule.GetAllPFCItemUpperCuttingDieScheduleByModelID)
	router.POST("/getPFCItemUpperCuttingDieSchedule", controllers.PFCItemUpperCuttingDieSchedule.GetAllPFCItemUpperCuttingDieSchedule)
	router.POST("/pfcItemUpperCuttingDieSchedule", controllers.PFCItemUpperCuttingDieSchedule.InsertPFCItemUpperCuttingDieSchedule)
	router.PUT("/pfcItemUpperCuttingDieSchedule", controllers.PFCItemUpperCuttingDieSchedule.UpdatePFCItemUpperCuttingDieSchedule)
	router.DELETE("/pfcItemUpperCuttingDieSchedule", controllers.PFCItemUpperCuttingDieSchedule.DeletePFCItemUpperCuttingDieSchedule)
	router.DELETE("/pfcItemUpperCuttingDieScheduleByModelID", controllers.PFCItemUpperCuttingDieSchedule.DeletePFCItemUpperCuttingDieScheduleByModelID)

	//PFC Perforation Specification
	router.POST("/getPFCPerforationSpecificationByModelID", controllers.PFCPerforationSpecification.GetAllPFCPerforationSpecification)
	router.POST("/pfcPerforationSpecification", controllers.PFCPerforationSpecification.InsertPFCPerforationSpecification)
	router.PUT("/pfcPerforationSpecification", controllers.PFCPerforationSpecification.UpdatePFCPerforationSpecification)
	router.DELETE("/pfcPerforationSpecification", controllers.PFCPerforationSpecification.DeletePFCPerforationSpecification)

	router.POST("/getPFCItemPerforationSpecification", controllers.PFCPerforationSpecification.GetAllPFCItemPerforationSpecification)
	router.POST("/pfcItemPerforationSpecification", controllers.PFCPerforationSpecification.InsertPFCItemPerforationSpecification)
	router.PUT("/pfcItemPerforationSpecification", controllers.PFCPerforationSpecification.UpdatePFCItemPerforationSpecification)
	router.DELETE("/pfcItemPerforationSpecification", controllers.PFCPerforationSpecification.DeletePFCItemPerforationSpecification)

	//PFC Upper Logo Specification
	router.POST("/getPFCUpperLogoSpecificationByModelID", controllers.PFCUpperLogoSpecification.GetAllPFCUpperLogoSpecification)
	router.POST("/pfcUpperLogoSpecification", controllers.PFCUpperLogoSpecification.InsertPFCUpperLogoSpecification)
	router.PUT("/pfcUpperLogoSpecification", controllers.PFCUpperLogoSpecification.UpdatePFCUpperLogoSpecification)
	router.DELETE("/pfcUpperLogoSpecification", controllers.PFCUpperLogoSpecification.DeletePFCUpperLogoSpecification)

	router.POST("/getPFCItemUpperLogoSpecification", controllers.PFCUpperLogoSpecification.GetAllPFCItemUpperLogoSpecification)
	router.POST("/pfcItemUpperLogoSpecification", controllers.PFCUpperLogoSpecification.InsertPFCItemUpperLogoSpecification)
	router.PUT("/pfcItemUpperLogoSpecification", controllers.PFCUpperLogoSpecification.UpdatePFCItemUpperLogoSpecification)
	router.DELETE("/pfcItemUpperLogoSpecification", controllers.PFCUpperLogoSpecification.DeletePFCItemUpperLogoSpecification)

	//PFC Elastic Gore Specification
	router.POST("/getPFCElasticGoreSpecificationByModelID", controllers.PFCElasticGoreSpecification.GetAllPFCElasticGoreSpecification)
	router.POST("/pfcElasticGoreSpecification", controllers.PFCElasticGoreSpecification.InsertPFCElasticGoreSpecification)
	router.PUT("/pfcElasticGoreSpecification", controllers.PFCElasticGoreSpecification.UpdatePFCElasticGoreSpecification)
	router.DELETE("/pfcElasticGoreSpecification", controllers.PFCElasticGoreSpecification.DeletePFCElasticGoreSpecification)

	router.POST("/getPFCItemElasticGoreSpecification", controllers.PFCElasticGoreSpecification.GetAllPFCItemElasticGoreSpecification)
	router.POST("/pfcItemElasticGoreSpecification", controllers.PFCElasticGoreSpecification.InsertPFCItemElasticGoreSpecification)
	router.PUT("/pfcItemElasticGoreSpecification", controllers.PFCElasticGoreSpecification.UpdatePFCItemElasticGoreSpecification)
	router.DELETE("/pfcItemElasticGoreSpecification", controllers.PFCElasticGoreSpecification.DeletePFCItemElasticGoreSpecification)

	//PFC Skiving Instructions
	router.POST("/getPFCSkivingInstructionsByModelID", controllers.PFCSkivingInstruction.GetAllPFCSkivingInstructions)
	router.POST("/pfcSkivingInstruction", controllers.PFCSkivingInstruction.InsertPFCSkivingInstructions)
	router.PUT("/pfcSkivingInstruction", controllers.PFCSkivingInstruction.UpdatePFCSkivingInstructions)
	router.DELETE("/pfcSkivingInstruction", controllers.PFCSkivingInstruction.DeletePFCSkivingInstructions)

	router.POST("/getPFCItemSkivingInstruction", controllers.PFCSkivingInstruction.GetAllPFCItemSkivingInstructions) //item
	router.POST("/pfcItemSkivingInstruction", controllers.PFCSkivingInstruction.InsertPFCItemSkivingInstructions)    //item
	router.PUT("/pfcItemSkivingInstruction", controllers.PFCSkivingInstruction.UpdatePFCItemSkivingInstructions)     //item
	router.DELETE("/pfcItemSkivingInstruction", controllers.PFCSkivingInstruction.DeletePFCItemSkivingInstructions)  //item

	//PFC Marking Location
	router.POST("/getPFCMarkingLocationByModelID", controllers.PFCMarkingLoca.GetAllPFCMarkingLocation)
	router.POST("/pfcMarkingLocation", controllers.PFCMarkingLoca.InsertPFCMarkingLocation)
	router.PUT("/pfcMarkingLocation", controllers.PFCMarkingLoca.UpdatePFCMarkingLocation)
	router.DELETE("/pfcMarkingLocation", controllers.PFCMarkingLoca.DeletePFCMarkingLocation)

	router.POST("/getPFCItemMarkingLocation", controllers.PFCMarkingLoca.GetAllPFCItemMarkingLocation) //item
	router.POST("/pfcItemMarkingLocation", controllers.PFCMarkingLoca.InsertPFCItemMarkingLocation)    //item
	router.PUT("/pfcItemMarkingLocation", controllers.PFCMarkingLoca.UpdatePFCItemMarkingLocation)     //item
	router.DELETE("/pfcItemMarkingLocation", controllers.PFCMarkingLoca.DeletePFCItemMarkingLocation)  //item

	//PFC Reinforcement Placement
	router.POST("/getPFCReinforcementPlacementByModelID", controllers.PFCReinforcementPlacemen.GetAllPFCReinforcementPlacement)
	router.POST("/pfcReinforcementPlacement", controllers.PFCReinforcementPlacemen.InsertPFCReinforcementPlacement)
	router.PUT("/pfcReinforcementPlacement", controllers.PFCReinforcementPlacemen.UpdatePFCReinforcementPlacement)
	router.DELETE("/pfcReinforcementPlacement", controllers.PFCReinforcementPlacemen.DeletePFCReinforcementPlacement)

	router.POST("/getPFCItemReinforcementPlacement", controllers.PFCReinforcementPlacemen.GetAllPFCItemReinforcementPlacement) //item
	router.POST("/pfcItemReinforcementPlacement", controllers.PFCReinforcementPlacemen.InsertPFCItemReinforcementPlacement)    //item
	router.PUT("/pfcItemReinforcementPlacement", controllers.PFCReinforcementPlacemen.UpdatePFCItemReinforcementPlacement)     //item
	router.DELETE("/pfcItemReinforcementPlacement", controllers.PFCReinforcementPlacemen.DeletePFCItemReinforcementPlacement)  //item

	//PFC Second Process
	router.POST("/getPFCSecondProcessByModelID", controllers.PFCSecondProces.GetAllPFCSecondProcess)
	router.POST("/pfcSecondProcess", controllers.PFCSecondProces.InsertPFCSecondProcess)
	router.PUT("/pfcSecondProcess", controllers.PFCSecondProces.UpdatePFCSecondProcess)
	router.DELETE("/pfcSecondProcess", controllers.PFCSecondProces.DeletePFCSecondProcess)

	router.POST("/getPFCItemSecondProcess", controllers.PFCSecondProces.GetAllPFCItemSecondProcess) //item
	router.POST("/pfcItemSecondProcess", controllers.PFCSecondProces.InsertPFCItemSecondProcess)    //item
	router.PUT("/pfcItemSecondProcess", controllers.PFCSecondProces.UpdatePFCItemSecondProcess)     //item
	router.DELETE("/pfcItemSecondProcess", controllers.PFCSecondProces.DeletePFCItemSecondProcess)  //item

	//PFC Computer Stitching Schedule
	router.POST("/getPFCComputerStitchingScheduleByModelID", controllers.PFCComputerStitchingSchedu.GetAllPFCComputerStitchingSchedule)
	router.POST("/pfcComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.InsertPFCComputerStitchingSchedule)
	router.PUT("/pfcComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.UpdatePFCComputerStitchingSchedule)
	router.DELETE("/pfcComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.DeletePFCComputerStitchingSchedule)

	router.POST("/getPFCItemComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.GetAllPFCItemComputerStitchingSchedule) //item
	router.POST("/pfcItemComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.InsertPFCItemComputerStitchingSchedule)    //item
	router.PUT("/pfcItemComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.UpdatePFCItemComputerStitchingSchedule)     //item
	router.DELETE("/pfcItemComputerStitchingSchedule", controllers.PFCComputerStitchingSchedu.DeletePFCItemComputerStitchingSchedule)  //item

	//PFC Stitching Overview Sketch
	router.POST("/getPFCStitchingOverviewSketchByModelID", controllers.PFCStitchingOverviewSketc.GetAllPFCStitchingOverviewSketch)
	router.POST("/pfcStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.InsertPFCStitchingOverviewSketch)
	router.PUT("/pfcStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.UpdatePFCStitchingOverviewSketch)
	router.DELETE("/pfcStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.DeletePFCStitchingOverviewSketch)

	router.POST("/getPFCItemStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.GetAllPFCItemStitchingOverviewSketch) //item
	router.POST("/pfcItemStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.InsertPFCItemStitchingOverviewSketch)    //item
	router.PUT("/pfcItemStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.UpdatePFCItemStitchingOverviewSketch)     //item
	router.DELETE("/pfcItemStitchingOverviewSketch", controllers.PFCStitchingOverviewSketc.DeletePFCItemStitchingOverviewSketch)  //item

	//PFC Stitching Instruction
	router.POST("/getPFCStitchingInstructionByModelID", controllers.PFCStitchingInstructio.GetAllPFCStitchingInstruction)
	router.POST("/pfcStitchingInstruction", controllers.PFCStitchingInstructio.InsertPFCStitchingInstruction)
	router.PUT("/pfcStitchingInstruction", controllers.PFCStitchingInstructio.UpdatePFCStitchingInstruction)
	router.DELETE("/pfcStitchingInstruction", controllers.PFCStitchingInstructio.DeletePFCStitchingInstruction)

	router.POST("/getPFCItemStitchingInstruction", controllers.PFCStitchingInstructio.GetAllPFCItemStitchingInstruction) //item
	router.POST("/pfcItemStitchingInstruction", controllers.PFCStitchingInstructio.InsertPFCItemStitchingInstruction)    //item
	router.PUT("/pfcItemStitchingInstruction", controllers.PFCStitchingInstructio.UpdatePFCItemStitchingInstruction)     //item
	router.DELETE("/pfcItemStitchingInstruction", controllers.PFCStitchingInstructio.DeletePFCItemStitchingInstruction)  //item

	//PFC Bottom Cutting DieSchedule
	router.POST("/getPFCBottomCuttingDieScheduleByModelID", controllers.PFCBottomCuttingDieSchedul.GetAllPFCBottomCuttingDieSchedule)
	router.POST("/pfcBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.InsertPFCBottomCuttingDieSchedule)
	router.PUT("/pfcBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.UpdatePFCBottomCuttingDieSchedule)
	router.DELETE("/pfcBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.DeletePFCBottomCuttingDieSchedule)

	router.POST("/getPFCItemBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.GetAllPFCItemBottomCuttingDieSchedule) //item
	router.POST("/pfcItemBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.InsertPFCItemBottomCuttingDieSchedule)    //item
	router.PUT("/pfcItemBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.UpdatePFCItemBottomCuttingDieSchedule)     //item
	router.DELETE("/pfcItemBottomCuttingDieSchedule", controllers.PFCBottomCuttingDieSchedul.DeletePFCItemBottomCuttingDieSchedule)  //item

	//PFC Bottom Logo Specification
	router.POST("/getPFCBottomLogoSpecificationByModelID", controllers.PFCBottomLogoSpecificatio.GetAllPFCBottomLogoSpecification)
	router.POST("/pfcBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.InsertPFCBottomLogoSpecification)
	router.PUT("/pfcBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.UpdatePFCBottomLogoSpecification)
	router.DELETE("/pfcBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.DeletePFCBottomLogoSpecification)

	router.POST("/getPFCItemBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.GetAllPFCItemBottomLogoSpecification) //item
	router.POST("/pfcItemBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.InsertPFCItemBottomLogoSpecification)    //item
	router.PUT("/pfcItemBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.UpdatePFCItemBottomLogoSpecification)     //item
	router.DELETE("/pfcItemBottomLogoSpecification", controllers.PFCBottomLogoSpecificatio.DeletePFCItemBottomLogoSpecification)  //item

	//PFC Bottom Parts Process
	router.POST("/getPFCBottomPartsProcessByModelID", controllers.PFCBottomPartsProces.GetAllPFCBottomPartsProcess)
	router.POST("/pfcBottomPartsProcess", controllers.PFCBottomPartsProces.InsertPFCBottomPartsProcess)
	router.PUT("/pfcBottomPartsProcess", controllers.PFCBottomPartsProces.UpdatePFCBottomPartsProcess)
	router.DELETE("/pfcBottomPartsProcess", controllers.PFCBottomPartsProces.DeletePFCBottomPartsProcess)

	router.POST("/getPFCItemBottomPartsProcess", controllers.PFCBottomPartsProces.GetAllPFCItemBottomPartsProcess) //item
	router.POST("/pfcItemBottomPartsProcess", controllers.PFCBottomPartsProces.InsertPFCItemBottomPartsProcess)    //item
	router.PUT("/pfcItemBottomPartsProcess", controllers.PFCBottomPartsProces.UpdatePFCItemBottomPartsProcess)     //item
	router.DELETE("/pfcItemBottomPartsProcess", controllers.PFCBottomPartsProces.DeletePFCItemBottomPartsProcess)  //item

	//PFC Bottom Silk Screen Process
	router.POST("/getPFCBottomSilkScreenProcessByModelID", controllers.PFCBottomSilkScreenProces.GetAllPFCBottomSilkScreenProcess)
	router.POST("/pfcBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.InsertPFCBottomSilkScreenProcess)
	router.PUT("/pfcBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.UpdatePFCBottomSilkScreenProcess)
	router.DELETE("/pfcBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.DeletePFCBottomSilkScreenProcess)

	router.POST("/getPFCItemBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.GetAllPFCItemBottomSilkScreenProcess) //item
	router.POST("/pfcItemBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.InsertPFCItemBottomSilkScreenProcess)    //item
	router.PUT("/pfcItemBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.UpdatePFCItemBottomSilkScreenProcess)     //item
	router.DELETE("/pfcItemBottomSilkScreenProcess", controllers.PFCBottomSilkScreenProces.DeletePFCItemBottomSilkScreenProcess)  //item

	//PFC Outsole Specification
	router.POST("/getPFCOutsoleSpecificationByModelID", controllers.PFCOutsoleSpecificatio.GetAllPFCOutsoleSpecification)
	router.POST("/pfcOutsoleSpecification", controllers.PFCOutsoleSpecificatio.InsertPFCOutsoleSpecification)
	router.PUT("/pfcOutsoleSpecification", controllers.PFCOutsoleSpecificatio.UpdatePFCOutsoleSpecification)
	router.DELETE("/pfcOutsoleSpecification", controllers.PFCOutsoleSpecificatio.DeletePFCOutsoleSpecification)

	router.POST("/getPFCItemOutsoleSpecification", controllers.PFCOutsoleSpecificatio.GetAllPFCItemOutsoleSpecification) //item
	router.POST("/pfcItemOutsoleSpecification", controllers.PFCOutsoleSpecificatio.InsertPFCItemOutsoleSpecification)    //item
	router.PUT("/pfcItemOutsoleSpecification", controllers.PFCOutsoleSpecificatio.UpdatePFCItemOutsoleSpecification)     //item
	router.DELETE("/pfcItemOutsoleSpecification", controllers.PFCOutsoleSpecificatio.DeletePFCItemOutsoleSpecification)  //item

	//PFC Outsole Pressing Process
	router.POST("/getPFCOutsolePressingProcessByModelID", controllers.PFCOutsolePressingProces.GetAllPFCOutsolePressingProcess)
	router.POST("/pfcOutsolePressingProcess", controllers.PFCOutsolePressingProces.InsertPFCOutsolePressingProcess)
	router.PUT("/pfcOutsolePressingProcess", controllers.PFCOutsolePressingProces.UpdatePFCOutsolePressingProcess)
	router.DELETE("/pfcOutsolePressingProcess", controllers.PFCOutsolePressingProces.DeletePFCOutsolePressingProcess)

	router.POST("/getPFCItemOutsolePressingProcess", controllers.PFCOutsolePressingProces.GetAllPFCItemOutsolePressingProcess) //item
	router.POST("/pfcItemOutsolePressingProcess", controllers.PFCOutsolePressingProces.InsertPFCItemOutsolePressingProcess)    //item
	router.PUT("/pfcItemOutsolePressingProcess", controllers.PFCOutsolePressingProces.UpdatePFCItemOutsolePressingProcess)     //item
	router.DELETE("/pfcItemOutsolePressingProcess", controllers.PFCOutsolePressingProces.DeletePFCItemOutsolePressingProcess)  //item

	//PFC Rubber Component Specification
	router.POST("/getPFCRubberComponentSpecificationByModelID", controllers.PFCRubberComponentSpecificatio.GetAllPFCRubberComponentSpecification)
	router.POST("/pfcRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.InsertPFCRubberComponentSpecification)
	router.PUT("/pfcRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.UpdatePFCRubberComponentSpecification)
	router.DELETE("/pfcRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.DeletePFCRubberComponentSpecification)

	router.POST("/getPFCItemRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.GetAllPFCItemRubberComponentSpecification) //item
	router.POST("/pfcItemRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.InsertPFCItemRubberComponentSpecification)    //item
	router.PUT("/pfcItemRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.UpdatePFCItemRubberComponentSpecification)     //item
	router.DELETE("/pfcItemRubberComponentSpecification", controllers.PFCRubberComponentSpecificatio.DeletePFCItemRubberComponentSpecification)  //item

	//PFC Sockliner Specification
	router.POST("/getPFCSocklinerSpecificationByModelID", controllers.PFCSocklinerSpecificatio.GetAllPFCSocklinerSpecification)
	router.POST("/pfcSocklinerSpecification", controllers.PFCSocklinerSpecificatio.InsertPFCSocklinerSpecification)
	router.PUT("/pfcSocklinerSpecification", controllers.PFCSocklinerSpecificatio.UpdatePFCSocklinerSpecification)
	router.DELETE("/pfcSocklinerSpecification", controllers.PFCSocklinerSpecificatio.DeletePFCSocklinerSpecification)

	router.POST("/getPFCItemSocklinerSpecification", controllers.PFCSocklinerSpecificatio.GetAllPFCItemSocklinerSpecification) //item
	router.POST("/pfcItemSocklinerSpecification", controllers.PFCSocklinerSpecificatio.InsertPFCItemSocklinerSpecification)    //item
	router.PUT("/pfcItemSocklinerSpecification", controllers.PFCSocklinerSpecificatio.UpdatePFCItemSocklinerSpecification)     //item
	router.DELETE("/pfcItemSocklinerSpecification", controllers.PFCSocklinerSpecificatio.DeletePFCItemSocklinerSpecification)  //item

	//PFC Heel Wedge Specification
	router.POST("/getPFCHeelWedgeSpecificationByModelID", controllers.PFCHeelWedgeSpecificatio.GetAllPFCHeelWedgeSpecification)
	router.POST("/pfcHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.InsertPFCHeelWedgeSpecification)
	router.PUT("/pfcHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.UpdatePFCHeelWedgeSpecification)
	router.DELETE("/pfcHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.DeletePFCHeelWedgeSpecification)

	router.POST("/getPFCItemHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.GetAllPFCItemHeelWedgeSpecification) //item
	router.POST("/pfcItemHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.InsertPFCItemHeelWedgeSpecification)    //item
	router.PUT("/pfcItemHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.UpdatePFCItemHeelWedgeSpecification)     //item
	router.DELETE("/pfcItemHeelWedgeSpecification", controllers.PFCHeelWedgeSpecificatio.DeletePFCItemHeelWedgeSpecification)  //item

	//PFC Logo Application Process
	router.POST("/getPFCLogoApplicationProcessByModelID", controllers.PFCLogoApplicationProces.GetAllPFCLogoApplicationProcess)
	router.POST("/pfcLogoApplicationProcess", controllers.PFCLogoApplicationProces.InsertPFCLogoApplicationProcess)
	router.PUT("/pfcLogoApplicationProcess", controllers.PFCLogoApplicationProces.UpdatePFCLogoApplicationProcess)
	router.DELETE("/pfcLogoApplicationProcess", controllers.PFCLogoApplicationProces.DeletePFCLogoApplicationProcess)

	router.POST("/getPFCItemLogoApplicationProcess", controllers.PFCLogoApplicationProces.GetAllPFCItemLogoApplicationProcess) //item
	router.POST("/pfcItemLogoApplicationProcess", controllers.PFCLogoApplicationProces.InsertPFCItemLogoApplicationProcess)    //item
	router.PUT("/pfcItemLogoApplicationProcess", controllers.PFCLogoApplicationProces.UpdatePFCItemLogoApplicationProcess)     //item
	router.DELETE("/pfcItemLogoApplicationProcess", controllers.PFCLogoApplicationProces.DeletePFCItemLogoApplicationProcess)  //item

	//PFC Sockliner Molding Process
	router.POST("/getPFCSocklinerMoldingProcessByModelID", controllers.PFCSocklinerMoldingProces.GetAllPFCSocklinerMoldingProcess)
	router.POST("/pfcSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.InsertPFCSocklinerMoldingProcess)
	router.PUT("/pfcSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.UpdatePFCSocklinerMoldingProcess)
	router.DELETE("/pfcSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.DeletePFCSocklinerMoldingProcess)

	router.POST("/getPFCItemSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.GetAllPFCItemSocklinerMoldingProcess) //item
	router.POST("/pfcItemSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.InsertPFCItemSocklinerMoldingProcess)    //item
	router.PUT("/pfcItemSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.UpdatePFCItemSocklinerMoldingProcess)     //item
	router.DELETE("/pfcItemSocklinerMoldingProcess", controllers.PFCSocklinerMoldingProces.DeletePFCItemSocklinerMoldingProcess)  //item

	//PFC Sockliner Graphic Process
	router.POST("/getPFCSocklinerGraphicProcessByModelID", controllers.PFCSocklinerGraphicProces.GetAllPFCSocklinerGraphicProcess)
	router.POST("/pfcSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.InsertPFCSocklinerGraphicProcess)
	router.PUT("/pfcSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.UpdatePFCSocklinerGraphicProcess)
	router.DELETE("/pfcSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.DeletePFCSocklinerGraphicProcess)

	router.POST("/getPFCItemSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.GetAllPFCItemSocklinerGraphicProcess) //item
	router.POST("/pfcItemSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.InsertPFCItemSocklinerGraphicProcess)    //item
	router.PUT("/pfcItemSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.UpdatePFCItemSocklinerGraphicProcess)     //item
	router.DELETE("/pfcItemSocklinerGraphicProcess", controllers.PFCSocklinerGraphicProces.DeletePFCItemSocklinerGraphicProcess)  //item

	//PFC Outside Conveyor Process
	router.POST("/getPFCOutsideConveyorProcessByModelID", controllers.PFCOutsideConveyorProces.GetAllPFCOutsideConveyorProcess)
	router.POST("/pfcOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.InsertPFCOutsideConveyorProcess)
	router.PUT("/pfcOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.UpdatePFCOutsideConveyorProcess)
	router.DELETE("/pfcOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.DeletePFCOutsideConveyorProcess)

	router.POST("/getPFCItemOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.GetAllPFCItemOutsideConveyorProcess) //item
	router.POST("/pfcItemOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.InsertPFCItemOutsideConveyorProcess)    //item
	router.PUT("/pfcItemOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.UpdatePFCItemOutsideConveyorProcess)     //item
	router.DELETE("/pfcItemOutsideConveyorProcess", controllers.PFCOutsideConveyorProces.DeletePFCItemOutsideConveyorProcess)  //item

	//PFC Assembling Process
	router.POST("/getPFCAssemblingProcessByModelID", controllers.PFCAssemblingProces.GetAllPFCAssemblingProcess)
	router.POST("/pfcAssemblingProcess", controllers.PFCAssemblingProces.InsertPFCAssemblingProcess)
	router.PUT("/pfcAssemblingProcess", controllers.PFCAssemblingProces.UpdatePFCAssemblingProcess)
	router.DELETE("/pfcAssemblingProcess", controllers.PFCAssemblingProces.DeletePFCAssemblingProcess)

	router.POST("/getPFCItemAssemblingProcess", controllers.PFCAssemblingProces.GetAllPFCItemAssemblingProcess) //item
	router.POST("/pfcItemAssemblingProcess", controllers.PFCAssemblingProces.InsertPFCItemAssemblingProcess)    //item
	router.PUT("/pfcItemAssemblingProcess", controllers.PFCAssemblingProces.UpdatePFCItemAssemblingProcess)     //item
	router.DELETE("/pfcItemAssemblingProcess", controllers.PFCAssemblingProces.DeletePFCItemAssemblingProcess)  //item

	//PFC Pressing Pad Specification
	router.POST("/getPFCPressingPadSpecificationByModelID", controllers.PFCPressingPadSpecificatio.GetAllPFCPressingPadSpecification)
	router.POST("/pfcPressingPadSpecification", controllers.PFCPressingPadSpecificatio.InsertPFCPressingPadSpecification)
	router.PUT("/pfcPressingPadSpecification", controllers.PFCPressingPadSpecificatio.UpdatePFCPressingPadSpecification)
	router.DELETE("/pfcPressingPadSpecification", controllers.PFCPressingPadSpecificatio.DeletePFCPressingPadSpecification)

	router.POST("/getPFCItemPressingPadSpecification", controllers.PFCPressingPadSpecificatio.GetAllPFCItemPressingPadSpecification) //item
	router.POST("/pfcItemPressingPadSpecification", controllers.PFCPressingPadSpecificatio.InsertPFCItemPressingPadSpecification)    //item
	router.PUT("/pfcItemPressingPadSpecification", controllers.PFCPressingPadSpecificatio.UpdatePFCItemPressingPadSpecification)     //item
	router.DELETE("/pfcItemPressingPadSpecification", controllers.PFCPressingPadSpecificatio.DeletePFCItemPressingPadSpecification)  //item

	//PFC Production Checklist
	router.POST("/getPFCProductionChecklistByModelID", controllers.PFCProductionChecklis.GetAllPFCProductionChecklist)
	router.POST("/pfcProductionChecklist", controllers.PFCProductionChecklis.InsertPFCProductionChecklist)
	router.PUT("/pfcProductionChecklist", controllers.PFCProductionChecklis.UpdatePFCProductionChecklist)
	router.DELETE("/pfcProductionChecklist", controllers.PFCProductionChecklis.DeletePFCProductionChecklist)

	router.POST("/getPFCItemProductionChecklist", controllers.PFCProductionChecklis.GetAllPFCItemProductionChecklist) //item
	router.POST("/pfcItemProductionChecklist", controllers.PFCProductionChecklis.InsertPFCItemProductionChecklist)    //item
	router.PUT("/pfcItemProductionChecklist", controllers.PFCProductionChecklis.UpdatePFCItemProductionChecklist)     //item
	router.DELETE("/pfcItemProductionChecklist", controllers.PFCProductionChecklis.DeletePFCItemProductionChecklist)  //item

	//PFC Upper Measurement Spec
	router.POST("/getPFCUpperMeasurementSpecByModelID", controllers.PFCUpperMeasurementSpe.GetAllPFCUpperMeasurementSpec)
	router.POST("/pfcUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.InsertPFCUpperMeasurementSpec)
	router.PUT("/pfcUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.UpdatePFCUpperMeasurementSpec)
	router.DELETE("/pfcUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.DeletePFCUpperMeasurementSpec)

	router.POST("/getPFCItemUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.GetAllPFCItemUpperMeasurementSpec) //item
	router.POST("/pfcItemUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.InsertPFCItemUpperMeasurementSpec)    //item
	router.PUT("/pfcItemUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.UpdatePFCItemUpperMeasurementSpec)     //item
	router.DELETE("/pfcItemUpperMeasurementSpec", controllers.PFCUpperMeasurementSpe.DeletePFCItemUpperMeasurementSpec)  //item

	//PFC Key Manufacturing Details
	router.POST("/getPFCKeyManufacturingDetailsByModelID", controllers.PFCKeyManufacturingDetail.GetAllPFCKeyManufacturingDetails)
	router.POST("/pfcKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.InsertPFCKeyManufacturingDetails)
	router.PUT("/pfcKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.UpdatePFCKeyManufacturingDetails)
	router.DELETE("/pfcKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.DeletePFCKeyManufacturingDetails)

	router.POST("/getPFCItemKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.GetAllPFCItemKeyManufacturingDetails) //item
	router.POST("/pfcItemKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.InsertPFCItemKeyManufacturingDetails)    //item
	router.PUT("/pfcItemKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.UpdatePFCItemKeyManufacturingDetails)     //item
	router.DELETE("/pfcItemKeyManufacturingDetails", controllers.PFCKeyManufacturingDetail.DeletePFCItemKeyManufacturingDetails)  //item

}
