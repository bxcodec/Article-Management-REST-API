package models



type ErrorModel struct{

 ErrorCode    	int 	`boil:"error_code" json:"error_code" toml:"error_code" yaml:"error_code"`
 Hints			string	`boil:"hints" json:"hints" toml:"hints" yaml:"hints"`
 Info			string	`boil:"info" json:"info" toml:"info" yaml:"info"`
}
