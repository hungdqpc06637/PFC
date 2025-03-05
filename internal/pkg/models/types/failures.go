package types

import "database/sql"

// TYPE - Failures
type Failures struct {
	Failure_id            string `json:"Failure_id"`
	Stage_id              string `json:"Stage_id"`
	Category_id           string `json:"Category_id"`
	SubCategory_id        string `json:"Sub_category_id"`
	Function_Requirements string `json:"Function_Requirements"`
	PotentialFailure      string `json:"PotentialFailure"`
	PotentialCause        string `json:"Potential_Cause"`
	PotentialEffec        string `json:"Potential_effec"`
	Sev                   string `json:"Sev"`
	Occ                   string `json:"Occ"`
	Det                   string `json:"Det"`
	Prevention            string `json:"Prevention"`
	Detection             string `json:"Detection"`
}

// type - Failure Images
type FailureImages struct {
	ImageId     string `json:"Image_id"`
	FailureId   string `json:"Failure_id"`
	Image       string `json:"Image"`
	Description string `json:"Description"`
}

// type - Failure SKU
type FailureSKU struct {
	FailureId string `json:"Failure_id"`
	SKUId     string `json:"SKU_id"`
}

// type - GetFailuresList
type GetFailuresList struct {
	Failure_id       string `json:"Failure_id"`
	Stage_name       string `json:"Stage_name"`
	SKU_code         string `json:"SKU_code"`
	DEVCODE			 string `json:"SR"`
	PotentialFailure string `json:"PotentialFailure"`
	Potential_effec  string `json:"Potential_effec"`
	Sev              string `json:"Sev"`
	Occ              string `json:"Occ"`
	Det              string `json:"Det"`
	RPN              string `json:"RPN"`
}

// type - GetFailureDetail
type GetFailureDetail struct {
	Failure_id            string         `json:"Failure_id"`
	Stage_id              string         `json:"Stage_id"`
	Stage_name            string         `json:"Stage_name"`
	SKU_code              string         `json:"SKU_code"`
	Category_name         string         `json:"Category_name"`
	SubCategory_name      string         `json:"SubCategory_name"`
	Image                 string         `json:"Image"`
	Function_Requirements string         `json:"Function_Requirements"`
	PotentialFailure      string         `json:"PotentialFailure"`
	Potential_effec       string         `json:"Potential_effec"`
	Potential_Cause       string         `json:"Potential_Cause"`
	Sev                   string         `json:"Sev"`
	Occ                   string         `json:"Occ"`
	Det                   string         `json:"Det"`
	RPN                   string         `json:"RPN"`
	Recommended_Actions   sql.NullString `json:"Recommended_Actions"`
}

// type - FailureSOD
type GetActionResultSOD struct {
	Failure_Id   string `json:"Failure_id"`
	Stage_Id     string `json:"Stage_id"`
	Stage_name   string `json:"Stage_name"`
	SEV_Name     string `json:"SEV_Name"`
	SEV          string `json:"SEV"`
	OCC_Name     string `json:"OCC_Name"`
	OCC          string `json:"OCC"`
	DET_Name     string `json:"DET_Name"`
	DET          string `json:"DET"`
	UpdateDate   string `json:"UpdateDate"`
	Action_Taken 	string `json:"Action_Taken"`
	Process_Name 	string `json:"Process_Name"`
	MFG_Resources 	string `json:"MFG_Resources"`
	Characteristics string `json:"Characteristics"`
	Spec_Tolerance	string `json:"Spec_Tolerance"`
	Eval_Tech 		string `json:"Eval_Tech"`
	Sample_Size_Freq string `json:"Sample_Size_Freq"`
	Resp_Depts 		string `json:"Resp_Depts"`
	Ctrl_Method_Form string `json:"Ctrl_Method_Form"` 
	Reaction_Plan 	string `json:"Reaction_Plan"`
}

// type - LatestStageSOD
type LatestStageSOD struct {
	Stage_Id   string `json:"Stage_Id"`
	Stage_name string `json:"Stage_Name"`
}

type GetIMGSKU struct {
	IMGName   string `json:"Image"`
}
type GetSkuList struct {
	SKU   string `json:"SKU"`
}

