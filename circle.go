package main

type Circle struct {
	List []struct {
		CircleCutImage struct {
			CreatedAt string `json:"createdAt"`
			FileSize  int64  `json:"fileSize"`
			Height    int64  `json:"height"`
			UpdatedAt string `json:"updatedAt"`
			URL       string `json:"url"`
			Width     int64  `json:"width"`
		} `json:"circleCutImage"`
		CreatedAt string `json:"createdAt"`
		Event     struct {
			AcceptCirclePaymentInfo bool   `json:"acceptCirclePaymentInfo"`
			CatalogImageDeadline    string `json:"catalogImageDeadline"`
			CreatedAt               string `json:"createdAt"`
			ElectionClosed          bool   `json:"electionClosed"`
			EndAt                   string `json:"endAt"`
			EventExhibitCourses     []struct {
				CreatedAt  string `json:"createdAt"`
				EndAt      string `json:"endAt"`
				ExhibitFee int64  `json:"exhibitFee"`
				ID         string `json:"id"`
				Name       string `json:"name"`
				Place      string `json:"place"`
				StartAt    string `json:"startAt"`
				UpdatedAt  string `json:"updatedAt"`
			} `json:"eventExhibitCourses"`
			ID             string `json:"id"`
			Name           string `json:"name"`
			OgpDescription string `json:"ogpDescription"`
			OgpImage       string `json:"ogpImage"`
			OgpTitle       string `json:"ogpTitle"`
			Place          string `json:"place"`
			RecruitEndAt   string `json:"recruitEndAt"`
			RecruitStartAt string `json:"recruitStartAt"`
			StartAt        string `json:"startAt"`
			UpdatedAt      string `json:"updatedAt"`
		} `json:"event"`
		EventExhibitCourse struct {
			CreatedAt  string `json:"createdAt"`
			EndAt      string `json:"endAt"`
			ExhibitFee int64  `json:"exhibitFee"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Place      string `json:"place"`
			StartAt    string `json:"startAt"`
			UpdatedAt  string `json:"updatedAt"`
		} `json:"eventExhibitCourse"`
		Genre                   string   `json:"genre"`
		GenreFreeFormat         string   `json:"genreFreeFormat"`
		ID                      string   `json:"id"`
		Name                    string   `json:"name"`
		NameRuby                string   `json:"nameRuby"`
		NextCircleExhibitInfoID string   `json:"nextCircleExhibitInfoID"`
		PenName                 string   `json:"penName"`
		PrevCircleExhibitInfoID string   `json:"prevCircleExhibitInfoID"`
		Spaces                  []string `json:"spaces"`
		UpdatedAt               string   `json:"updatedAt"`
		WebSiteURL              string   `json:"webSiteURL"`
	} `json:"list"`
}
