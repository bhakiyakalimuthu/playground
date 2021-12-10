//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
//
// Author Bhakiyaraj Kalimuthu
//

package cmd

var (
	Balance float32
)

func init()  {
	Balance = 50
}

func GetBalance() float32 {
	return Balance
}