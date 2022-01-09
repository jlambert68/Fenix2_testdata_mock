package common_config

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-gota/gota/series"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

func HashValues(valuesToHash []string) string {

	hash_string := ""
	sha256_hash := ""

	// Allways sort values before hash them

	sort.Strings(valuesToHash)
	//Hash all values
	for _, valueToHash := range valuesToHash {
		hash_string = hash_string + valueToHash

		hash := sha256.New()
		hash.Write([]byte(hash_string))
		sha256_hash = hex.EncodeToString(hash.Sum(nil))
		hash_string = sha256_hash

	}

	return sha256_hash

}

func setFromList(list []string) (set []string) {
	ks := make(map[string]bool) // map to keep track of repeats

	for _, e := range list {
		if _, v := ks[e]; !v {
			ks[e] = true
			set = append(set, e)
		}
	}
	return
}

func uniqueGotaSeries(s series.Series) series.Series {
	return series.New(setFromList(s.Records()), s.Type(), s.Name)
}

func uniqueGotaSeriesAsStringArray(s series.Series) []string {
	return uniqueGotaSeries(s).Records()
}

func hashChildrenAndWriteToDataStore(level int, currentMerklePath string, valuesToHash []string, isEndLeafNode bool) string {

	hash_string := ""
	sha256_hash := ""

	sort.Strings(valuesToHash)

	// Hash all leaves for rowHashValue in valuesToHash
	for _, valueToHash := range valuesToHash {
		hash_string = hash_string + valueToHash

		hash := sha256.New()
		hash.Write([]byte(hash_string))
		sha256_hash = hex.EncodeToString(hash.Sum(nil))
		hash_string = sha256_hash

	}

	MerkleHash := sha256_hash

	// # Add MerkleHash and sub leaf nodes to table if not end node. If End node then only save ref to it self
	if isEndLeafNode == true {
		// Add row
		//merkleTreeToUse.loc[merkleTreeToUse.shape[0]] = [level, currentMerklePath, MerkleHash, MerkleHash]
		newRowDataFrame := dataframe.New(
			series.New([]int{level}, series.Int, "MerkleLevel"),
			series.New([]string{currentMerklePath}, series.String, "MerklePath"),
			series.New([]string{MerkleHash}, series.String, "MerkleHash"),
			series.New([]string{MerkleHash}, series.String, "MerkleChildHash"),
		)

		tempDF := merkleTreeDataFrame.RBind(newRowDataFrame)
		merkleTreeDataFrame = tempDF

	} else {
		for _, rowHashValue := range valuesToHash {
			// Add row
			//merkleTreeToUse.loc[merkleTreeToUse.shape[0]] = [level, currentMerklePath, MerkleHash, rowHashValue ]
			newRowDataFrame := dataframe.New(
				series.New([]int{level}, series.Int, "MerkleLevel"),
				series.New([]string{currentMerklePath}, series.String, "MerklePath"),
				series.New([]string{MerkleHash}, series.String, "MerkleHash"),
				series.New([]string{rowHashValue}, series.String, "MerkleChildHash"),
			)
			tempDF := merkleTreeDataFrame.RBind(newRowDataFrame)
			merkleTreeDataFrame = tempDF

		}
	}

	return MerkleHash

}

func recursiveTreeCreator(level int, currentMerkleFilterPath string, dataFrameToWorkOn dataframe.DataFrame, currentMerklePath string) string {
	level = level + 1
	startPosition := 0
	endPosition := strings.Index(currentMerkleFilterPath, "/")

	// Check if we found end of Tree
	if endPosition == -1 {
		// Leaf node, process rows

		// Sort on row - hashes
		dataFrameToWorkOn = dataFrameToWorkOn.Arrange(dataframe.Sort("TestDataHash"))

		// Hash all row - hashes into one hash
		valuesToHash := uniqueGotaSeriesAsStringArray(dataFrameToWorkOn.Col("TestDataHash"))

		// Hash and store
		MerkleHash := hashChildrenAndWriteToDataStore(level, currentMerklePath, valuesToHash, true)

		return MerkleHash

	} else {
		// Get merklePathlabel
		merklePathLabel := currentMerkleFilterPath[startPosition:endPosition]
		currentMerkleFilterPath := currentMerkleFilterPath[endPosition+1:]

		// Get Unique values
		uniqueValuesForSpecifiedColumn := uniqueGotaSeriesAsStringArray(dataFrameToWorkOn.Col(merklePathLabel))

		valuesToHash := []string{}

		// Loop over all unique values in column
		for _, uniqueValue := range uniqueValuesForSpecifiedColumn {
			//newFilteredDataFrame := dataFrameToWorkOn[dataFrameToWorkOn[merklePathLabel] == uniqueValue]
			newFilteredDataFrame := dataFrameToWorkOn.Filter(
				dataframe.F{
					Colname:    merklePathLabel,
					Comparator: series.Eq,
					Comparando: uniqueValue,
				})

			// Recursive call to get next level, if there is one
			localMerkleHash := recursiveTreeCreator(level, currentMerkleFilterPath, newFilteredDataFrame, currentMerklePath+uniqueValue+"/")

			if len(localMerkleHash) != 0 {
				valuesToHash = append(valuesToHash, localMerkleHash)
			} else {
				log.Fatalln("We are at the end node - **** Should never happened ****")
			}
		}

		// Add MerkleHash and nodes to table
		merkleHash := hashChildrenAndWriteToDataStore(level, currentMerklePath, valuesToHash, false)
		return merkleHash

	}
	return ""
}

// Dataframe holding original File's MerkleTree
var merkleTreeDataFrame dataframe.DataFrame

// Dataframe holding changed File's MerkleTree
var changedFilesMerkleTreeDataFrame dataframe.DataFrame

// Process incoming csv file and create MerkleRootHash and MerkleTree
func LoadAndProcessFile(fileToprocess string) (string, dataframe.DataFrame) {

	irisCsv, err := os.Open(fileToprocess)
	if err != nil {
		log.Fatal(err)
	}
	defer irisCsv.Close()

	df := dataframe.ReadCSV(irisCsv,
		dataframe.WithDelimiter(';'),
		dataframe.HasHeader(true))

	df = df.Arrange(dataframe.Sort("TestDataId"))

	numberOfRows := df.Nrow()
	df = df.Mutate(
		series.New(make([]string, numberOfRows), series.String, "TestDataHash"))

	number_of_columns_to_process := df.Ncol() - 1 // Don't process Hash column

	for rowCounter := 0; rowCounter < numberOfRows; rowCounter++ {
		var valuesToHash []string
		for columnCounter := 0; columnCounter < number_of_columns_to_process; columnCounter++ {
			valueToHash := df.Elem(rowCounter, columnCounter).String()
			valuesToHash = append(valuesToHash, valueToHash)
		}

		// Hash all values for row
		hashedRow := HashValues(valuesToHash)
		df.Elem(rowCounter, number_of_columns_to_process).Set(hashedRow)

	}

	// Columns for MerkleTree DataFrame
	merkleTreeDataFrame = dataframe.New(
		series.New([]int{}, series.Int, "MerkleLevel"),
		series.New([]string{}, series.String, "MerklePath"),
		series.New([]string{}, series.String, "MerkleHash"),
		series.New([]string{}, series.String, "MerkleChildHash"),
	)

	merkleFilterPath := "AccountEnvironment/ClientJuristictionCountryCode/MarketSubType/MarketName/" //SecurityType/"

	merkleHash := recursiveTreeCreator(0, merkleFilterPath, df, "MerkleRoot/")

	return merkleHash, merkleTreeDataFrame
}
