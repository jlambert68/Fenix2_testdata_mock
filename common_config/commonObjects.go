package common_config

const MaxBatchSizeSentToWorker = 100

type header_struct struct {
	headerId              int
	headerText            string
	headerType            int
	headerTypeName        string
	headerValueCount      int
	numberOfFormulas      int    //Number of formulas that are referencing valueSets for this header
	multiplicatorFactor   int64  //Multiplies the number of combination previous Headers generated
	numberOfRowsGenerated int64  // Hmmm, don't know what this is for
	valueComparator       []bool // Every value has True/False depending of if it is referenced or not in a formula
}

type valueSet_struct struct {
	valueSetId int
	HeaderId   int
	HeaderText string
	ValueText  string
}

type ruleIdentifier struct {
	ruleId      int
	headerId    int
	headerText  string
	headerIndex int
}

type ruleBaseBlock_struct struct {
	buildingBlockType       int // Describes the type, e.g. values: 'Header:HeaderValue', logical operator: 'AND', 'OR', 'NOT', separators: '(', ')'
	valueId                 int
	buildingBlockValue      string
	headerIndexIfappropiant int
}

type rule_struct struct {
	ruleIdentifier        ruleIdentifier
	requirement           string
	ruleBaseBlocks        []ruleBaseBlock_struct
	ruleText              string //This is a summary for text found in 'ruleBaseBlocks'
	ruleType              int
	ruleTypeName          string
	messageToUserOnTrue   string
	messageToUserOnFalse  string
	messageToFenixOnTrue  int
	messageToFenixOnFalse int
}

type combinationObject_struct struct {
	headers   []header_struct
	valueSets []valueSet_struct
	rules     []rule_struct
}

type messageToWorker_struct struct {
	startOfBatch      int32
	batchSize         int32
	combinationObject combinationObject_struct
}
