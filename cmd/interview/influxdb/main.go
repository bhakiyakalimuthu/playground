package main

import "time"

func main(){



//Background: In InfluxDB we shard our data by time. E.g. we might have a 4 week retention period for some data, sharded by week.
	//So each week we create a new shard for the new week's data and drop the old shard.
//	This question is similar (but not identical) to a technical problem we had.
//
//		A 'Shard' is a grouping of data on disk. A Shard has a start and end time, and contains 'Points' (think rows of a table)
	//		where Shard.startTime <= Point.time < Shard.endTime
//
//	When adding a new point, we calculate whether there is an existing shard that can contain it or whether we need to add a new shard.
//
//		In general, we may have many shards with different start and end times, potentially overlapping, of different durations.


//	* For each point, how do we find out if it can go in any existing shard? For this question, we want a yes/no answer,
	//	we don't care which shard it belongs to, only if it belongs to one of the shards.
//
//Example:
//Shards: [200, 300], [250, 350], [0,100]
//	Points: [0 -> yes, 50 -> yes, 100 -> no, 150 -> no, 200 -> yes, 250 -> yes]

	// create map



}

func isPointPresent(shards []shard, in int64) bool{
	for _,k := range shards {
		t1 := time.Unix(k.t1,0)
		t2 := time.Unix(k.t2,0)
		point := time.Unix(in,0)
		if t1.Before(point) || t1 == point && t2.After(point){
			return true
		}

	}
	return false
}

type shard struct{
	t1,t2 int64
}
