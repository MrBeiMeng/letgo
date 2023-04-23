package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm/clause"
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/models"
	"letgo_repo/system_file/generate"
	utils2 "letgo_repo/system_file/utils"
	"time"
)

func main() {

	skip := 0
	headerMap := make(map[string]string)

	cookies := "_gid=GA1.2.1633828563.1672150479; gr_user_id=fdfd41b7-6fc9-4c54-a5c0-b39e42bb9cd0; _bl_uid=CmlpqcaU6U2b2m79XzRFmy9qq9ma; a2873925c34ecbd2_gr_last_sent_cs1=beimengclub; csrftoken=Qg2oPO181zv8CQdSl68cyiD2pXyVFVEkgtEvMXvY4NKvqU31Xjj9mk7HlKO9YnPC; _ga_PDVPZYN3CW=GS1.1.1672223327.2.1.1672223365.0.0.0; _ga=GA1.2.1932331919.1672150479; NEW_QUESTION_DETAIL_PAGE_V2=1; aliyungf_tc=9104dc1f720e8887754329aa4765af05fb52108a4d5d57cec2691c437bef2f46; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1672150058,1672223220,1672223328,1672302020; NEW_PROBLEMLIST_PAGE=1; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiNDA3MDI5OCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiZDVjNWJhYmUzYjZmODQxYzYwN2ViNGIxMGZhNjY0N2Y1MmEwMGZmZmRhNDhlZTlmNDdhZDhkOGMwNGJmNDY3NiIsImlkIjo0MDcwMjk4LCJlbWFpbCI6IjExOTIzODQ3MjJAcXEuY29tIiwidXNlcm5hbWUiOiJiZWltZW5nY2x1YiIsInVzZXJfc2x1ZyI6ImJlaW1lbmdjbHViIiwiYXZhdGFyIjoiaHR0cHM6Ly9hc3NldHMubGVldGNvZGUuY24vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9iZWltZW5nY2x1Yi9hdmF0YXJfMTY2NzUzNTc0OC5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTY3MjIyMzM2NC43NDk1MjAzLCJleHBpcmVkX3RpbWVfIjoxNjc0NzU5NjAwLCJ2ZXJzaW9uX2tleV8iOjAsImxhdGVzdF90aW1lc3RhbXBfIjoxNjcyMzE3MDk4fQ.G5OZMva64jdi73u1bbrYu1K5kRTfaq1QXtw6lwjItn4; Hm_lpvt_f0faad39bcf8471e3ab3ef70125152c3=1672317101; a2873925c34ecbd2_gr_session_id=28d7810b-e93b-4933-97ee-ad2d688b4e10; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=28d7810b-e93b-4933-97ee-ad2d688b4e10; a2873925c34ecbd2_gr_cs1=beimengclub; a2873925c34ecbd2_gr_session_id_28d7810b-e93b-4933-97ee-ad2d688b4e10=true; _gat=1"

	headerMap["random-uuid"] = "ee0085a4-4dcd-1b24-d939-a573dd5c93ff"
	headerMap["content-type"] = "application/json"
	headerMap["origin"] = "https://leetcode.cn"
	headerMap["x-csrftoken"] = "Qg2oPO181zv8CQdSl68cyiD2pXyVFVEkgtEvMXvY4NKvqU31Xjj9mk7HlKO9YnPC"
	headerMap["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

	for i := 0; i < 29; i++ {
		skip = 100 * i
		startedTime := time.Now()
		println(fmt.Sprintf("%s\t|正在归档：[%d/%d]", startedTime.Format("2006-01-02:15:04:13"), skip, 28*100))
		initQuestionList(cookies, headerMap, skip)
		spend := time.Now().Sub(startedTime) / time.Microsecond
		println(fmt.Sprintf("完成[%d]阶段,用时[%s]毫秒", i, spend))

		println("延时10秒")
		time.Sleep(time.Second * 10)

		if spend/1000 < 30 {
			time.Sleep(time.Second * 30)
			panic("间隔不到30秒，已补时30秒")
		}

	}

	println(utils2.GetColorGreen("都已完成"))

}

func initQuestionList(cookies string, headerMap map[string]string, skip int) {
	// 获取列表
	postBody := "{\"query\":\"\\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\\n  problemsetQuestionList(\\n    categorySlug: $categorySlug\\n    limit: $limit\\n    skip: $skip\\n    filters: $filters\\n  ) {\\n    hasMore\\n    total\\n    questions {\\n      acRate\\n      difficulty\\n      freqBar\\n      frontendQuestionId\\n      isFavor\\n      paidOnly\\n      solutionNum\\n      status\\n      title\\n      titleCn\\n      titleSlug\\n      topicTags {\\n        name\\n        nameTranslated\\n        id\\n        slug\\n      }\\n      extra {\\n        hasVideoSolution\\n        topCompanyTags {\\n          imgUrl\\n          slug\\n          numSubscribed\\n        }\\n      }\\n    }\\n  }\\n}\\n    \",\"variables\":{\"categorySlug\":\"\",\"skip\":%d,\"limit\":100,\"filters\":{}}}"
	questionListBytes := utils2.HttpPost(`https://leetcode.cn/graphql/`, cookies, headerMap, fmt.Sprintf(postBody, skip))

	var questionList generate.QuestionList
	err := json.Unmarshal(questionListBytes, &questionList)
	if err != nil {
		panic(err)
	}

	for index, question := range questionList.Data.ProblemSetQuestionList.Questions {

		fmt.Printf("[%d]", index)

		mQuestion := convToModelQuestion(question)

		if SaveData(mQuestion) {
			return
		}

	}
}

func SaveData(mQuestion models.Question) bool {
	db := data_access.MysqlDB

	err := db.Save(&mQuestion.TopicTags).Error
	if err != nil {
		println(utils2.GetColorRed(err.Error()))
		//return true
	}

	err = db.Save(&mQuestion.TopCompanyTags).Error
	if err != nil {
		println(utils2.GetColorRed(err.Error()))
		//return true
	}

	err = db.Omit(clause.Associations).Clauses(clause.OnConflict{UpdateAll: true}).Create(&mQuestion).Error
	if err != nil {
		println(utils2.GetColorRed(err.Error()))
		//return true
	}

	for _, topicTag := range mQuestion.TopicTags {
		topicTag.ID = 0
		err := db.Where("name = ?", topicTag.Name).First(&topicTag).Error
		if err != nil {
			println(err.Error())
			//return true
		}

		insertSql := `replace into questions_topic_tags (question_id, topic_tags_id) values (?,?);`
		err = db.Exec(insertSql, &mQuestion.ID, &topicTag.ID).Error
		if err != nil {
			println(utils2.GetColorRed(err.Error()))
			//return true
		}
	}

	for _, topCompanyTag := range mQuestion.TopCompanyTags {
		topCompanyTag.ID = 0
		err := db.Where("slug = ?", topCompanyTag.Slug).First(&topCompanyTag).Error
		if err != nil {
			println(err.Error())
			return true
		}

		insertSql := `replace into questions_top_company_tags (question_id, top_company_tags_id) values (?,?);`
		err = db.Exec(insertSql, &mQuestion.ID, &topCompanyTag.ID).Error
		if err != nil {
			println(utils2.GetColorRed(err.Error()))
			return true
		}
	}
	return false
}

func convToModelQuestion(question generate.Questions) (result models.Question) {
	for _, topTag := range question.TopicTags {
		var mTopTag models.TopicTags

		_ = utils2.SimpleCopyProperties(&mTopTag, &topTag)
		result.TopicTags = append(result.TopicTags, mTopTag)
	}

	for _, topCompanyTag := range question.Extra.TopCompanyTags {
		var mTopCompanyTag models.TopCompanyTags

		_ = utils2.SimpleCopyProperties(&mTopCompanyTag, &topCompanyTag)
		result.TopCompanyTags = append(result.TopCompanyTags, mTopCompanyTag)
	}

	result.CompanyTagNum = question.Extra.CompanyTagNum
	result.HasVideoSolution = question.Extra.HasVideoSolution

	_ = utils2.SimpleCopyProperties(&result, &question)
	return result
}

func InitQuestionDetail(cookies string, headerMap map[string]string, question models.Question) []byte {
	reqBody := "{\"query\":\"\\n    query questionEditorData($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    codeSnippets {\\n      lang\\n      langSlug\\n      code\\n    }\\n    envInfo\\n    enableRunCode\\n  }\\n}\\n    \",\"variables\":{\"titleSlug\":\"%s\"}}"
	questionDetail := utils2.HttpPost(`https://leetcode.cn/graphql/`, cookies, headerMap, fmt.Sprintf(reqBody, question.TitleSlug))
	return questionDetail
}
