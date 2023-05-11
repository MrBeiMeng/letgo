package main

import (
	"fmt"
	"letgo_repo/system_file/utils"
	"strconv"
)

func main() {
	numArr := []string{"019", "024", "025", "029", "031", "034", "035", "036", "038", "039", "042", "043", "051", "052", "053", "061", "063", "067", "072", "076", "078", "081", "085", "093", "104", "105", "106", "107", "109", "126", "128", "129", "130", "138", "139", "142", "145", "146", "149", "150", "152", "159", "164", "165", "175", "176", "178", "179", "180", "182", "184", "187", "190", "192", "193", "201", "205", "206", "214", "218", "236", "238", "239", "247", "249", "250", "253", "254", "256", "261", "265", "270", "271", "273", "283", "285", "296", "297", "298", "305", "306", "307", "310", "314", "318", "320", "324", "327", "340", "346", "347", "349", "351", "354", "360", "361", "362", "365", "367", "369", "370", "374", "376", "379", "381", "389", "390", "391", "392", "394", "396", "397", "403", "407", "410", "413", "419", "423", "427", "432", "435", "439", "453", "456", "459", "462", "467", "471", "473", "476", "478", "479", "480", "482", "485", "486", "490", "495", "496", "497", "498", "502", "510", "513", "516", "517", "526", "527", "529", "531", "536", "537", "539", "540", "541", "542", "543", "548", "560", "564", "571", "572", "576", "583", "584", "587", "590", "594", "596", "601", "604", "609", "617", "619", "624", "631", "634", "637", "640", "641", "643", "651", "652", "657", "658", "670", "672", "674", "680", "684", "687", "690", "692", "697", "701", "705", "706", "709", "720", "721", "723", "724", "725", "726", "734", "738", "742", "749", "750", "758", "759", "765", "768", "781", "783", "784", "785", "791", "794", "795", "796", "801", "802", "803", "804", "807", "814", "817", "819", "825", "827", "830", "831", "845", "847", "850", "852", "857", "862", "863", "864", "869", "873", "874", "876", "890", "894", "903", "910", "912", "927", "938", "942", "943", "948", "950", "953", "957", "965", "967", "970", "971", "976", "981", "982", "983", "984", "985", "987"}

	tableArr := InitTableArr()

	for _, numStr := range numArr {

		firstNum, _ := strconv.Atoi(string(numStr[0]))

		secondNum, _ := strconv.Atoi(string(numStr[1]))

		tableArr[firstNum][secondNum] += fmt.Sprintf("%s,", numStr)
	}

	utils.TablePrint(tableArr, true)
}

func InitTableArr() [][]string {
	tableArr := make([][]string, 10)
	for i := range tableArr {
		tableArr[i] = make([]string, 10)
	}

	return tableArr
}

//"138\\548\\742\\850"