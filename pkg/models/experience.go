package models

import "time"

// Experience is struct of experience
type Experience struct {
	CompanyName    string
	EmploymentType string
	StartDate      time.Time
	EndDate        time.Time
	Projects       []Project
}

const (
	// FullTimeEmployment is "正社員"
	FullTimeEmployment = "正社員"
	// PartTimeEmployment is "パートタイム"
	PartTimeEmployment = "パートタイム"
	// Freelance is "フリーランス"
	Freelance = "フリーランス"
	// Internship is "インターン"
	Internship = "インターン"
)

// GetAllExperiences returns all expericnes
func GetAllExperiences() []Experience {
	return []Experience{
		{
			CompanyName:    "株式会社マネーフォワード",
			EmploymentType: FullTimeEmployment,
			StartDate:      time.Date(2019, 12, 1, 0, 0, 0, 0, time.Local),
			Projects: []Project{
				{
					Description: "ID基盤サービスを開発。認証基盤サービスとしてファーストパーティ向けのAPIを設計・開発。",
					Technologies: []Technology{
						{
							Name:     "Ruby on Rails",
							Versions: []string{"5"},
						},
						{
							Name:     "Vue.js",
							Versions: []string{"2"},
						},
						{
							Name: "Docker",
						},
					},
				},
			},
		},
		{
			CompanyName:    "株式会社ティンク",
			EmploymentType: FullTimeEmployment,
			StartDate:      time.Date(2019, 6, 1, 0, 0, 0, 0, time.Local),
			EndDate:        time.Date(2019, 8, 1, 0, 0, 0, 0, time.Local),
			Projects: []Project{
				{
					Description: "web、ios、androidから利用するAPIのレスポンス速度とサーバー負荷を改善。サーバーのスループットを1.8倍に、レスポンスタイムを1/2に改善。",
					Technologies: []Technology{
						{
							Name:     "Ruby on Rails",
							Versions: []string{"5"},
						},
					},
				},
				{
					Description: "Webフロントアプリ画面をリニューアル。",
					Technologies: []Technology{
						{
							Name: "Nust.js",
						},
						{
							Name: "Storybook",
						},
					},
				},
			},
		},
		{
			CompanyName:    "株式会社ポケットペア",
			EmploymentType: PartTimeEmployment,
			StartDate:      time.Date(2019, 5, 1, 0, 0, 0, 0, time.Local),
			EndDate:        time.Date(2019, 10, 1, 0, 0, 0, 0, time.Local),
			Projects: []Project{
				{
					Description: "ウェブサイトを独自にカテゴライズし、ウェブサイトを分析するツールを開発。",
					Technologies: []Technology{
						{
							Name: "Ruby on Rails",
						},
						{
							Name: "TypeScript",
						},
						{
							Name: "Puppeteer",
						},
						{
							Name: "Docker",
						},
						{
							Name: "MySQL",
						},
						{
							Name: "Redis",
						},
						{
							Name: "Sidekiq",
						},
						{
							Name: "EC2",
						},
						{
							Name: "S3",
						},
						{
							Name: "RDS",
						},
					},
				},
				{
					Description: "100人同時接続ゲームを作る前に100人同時接続が可能かどうかを検証するための検証環境を構築。",
					Technologies: []Technology{
						{
							Name: "ansible",
						},
						{
							Name: "EC2",
						},
					},
				},
			},
		},
		{
			CompanyName:    "株式会社プレスブログ",
			StartDate:      time.Date(2017, 8, 1, 0, 0, 0, 0, time.Local),
			EndDate:        time.Date(2019, 5, 1, 0, 0, 0, 0, time.Local),
			EmploymentType: Freelance,
			Projects: []Project{
				{
					Description: "ブログサービスを開発。",
					Technologies: []Technology{
						{
							Name:     "Ruby on Rails",
							Versions: []string{"4"},
						},
						{
							Name:     "Vue.js",
							Versions: []string{"2"},
						},
						{
							Name: "Nginx",
						},
						{
							Name: "Cloudflare",
						},
						{
							Name:     "MySQL",
							Versions: []string{"5.6"},
						},
						{
							Name: "Redis",
						},
						{
							Name: "EC2",
						},
						{
							Name: "S3",
						},
						{
							Name: "RDS",
						},
					},
				},
			},
		},
		{
			CompanyName:    "株式会社RUC",
			StartDate:      time.Date(2016, 12, 1, 0, 0, 0, 0, time.Local),
			EndDate:        time.Date(2017, 3, 1, 0, 0, 0, 0, time.Local),
			EmploymentType: FullTimeEmployment,
			Projects: []Project{
				{
					Description: "Wordpressを使ったWebサイト制作の受託開発",
					Technologies: []Technology{
						{
							Name: "WordPress",
						},
					},
				},
				{
					Description: "Google Drive内のファイルを読み取り、最適なフォルダ構成を提案する士業向けのChromeブラウザの拡張機能を開発",
					Technologies: []Technology{
						{
							Name: "Chrome拡張",
						},
						{
							Name:     "Ruby on Rails",
							Versions: []string{"4"},
						},
					},
				},
			},
		},
		{
			CompanyName:    "株式会社NVT",
			StartDate:      time.Date(2016, 6, 1, 0, 0, 0, 0, time.Local),
			EndDate:        time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
			EmploymentType: Freelance,
			Projects: []Project{
				{
					Description: "中小製造業の属人的な見積もりを解決するソフトのwebフロントとAPI開発",
					Technologies: []Technology{
						{
							Name:     "Ruby on Rails",
							Versions: []string{"4"},
						},
						{
							Name:     "AngularJS",
							Versions: []string{"1.8"},
						},
						{
							Name:     "MySQL",
							Versions: []string{"5.6"},
						},
					},
				},
			},
		},
		{
			CompanyName:    "株式会社div",
			StartDate:      time.Date(2015, 7, 1, 0, 0, 0, 0, time.Local),
			EndDate:        time.Date(2017, 3, 1, 0, 0, 0, 0, time.Local),
			EmploymentType: Internship,
			Projects: []Project{
				{
					Description: "プログラミング学習教室におけるメンター業務",
					Technologies: []Technology{
						{
							Name:     "Ruby on Rails",
							Versions: []string{"4"},
						},
					},
				},
				{
					Description: "社内システムを開発",
					Technologies: []Technology{
						{
							Name:     "Ruby on Rails",
							Versions: []string{"4"},
						},
					},
				},
			},
		},
	}
}
