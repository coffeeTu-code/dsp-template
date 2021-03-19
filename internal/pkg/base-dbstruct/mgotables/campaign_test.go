package mgotables

import (
	"fmt"
	"io/ioutil"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/mgo.v2/bson"
)

func TestCampaign_Parser(t *testing.T) {
	campaignDoc := new(Campaign)
	{
		data, err := ioutil.ReadFile("./testdata/campaign.json")
		if err != nil {
			t.Fatal(err)
		}
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, campaignDoc)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		bm, _ := bson.Marshal(campaignDoc)
		err := bson.Unmarshal(bm, campaignDoc)
		fmt.Println(err)
	}
}
