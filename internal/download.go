package internal

import (
	"fmt"
	"strconv"

	"github.com/idorocodes/warpget/pkg"
)

func DownloadFile(url string,chunks string) (bool, error) {


	

	fullPath,err := pkg.HandleFileCreation(url,)


	fmt.Printf("Splitted into %v chunks", chunks);


	IntChunks,err := strconv.Atoi(chunks)

	if err != nil{
		return false,fmt.Errorf("error while converting string into int : %w",err)
	}


	
	return true, nil
}
