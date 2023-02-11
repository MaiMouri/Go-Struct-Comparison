package main

import "fmt"

type Profile struct {
	ID       uint32  `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID uint32  `gorm:"column:member_id;type:int(11) unsigned;not null;index:idx01,priority:1;uniqueIndex:member_id,priority:1" json:"member_id"`
	Name     *string `gorm:"column:name;type:varchar(32)" json:"name"`
	NameKana *string `gorm:"column:name_kana;type:varchar(32)" json:"name_kana"`
	Nickname *string `gorm:"column:nickname;type:varchar(32)" json:"nickname"`
}

// 仮のプロフィール項目
var ProfileCompletedColumns = []string{
	// "ID",
	"Name",
	"NameKana",
	"Nickname",
}

func main() {
	var name = "John"
	var name_kana = "ジョン"
	var nickname = "ジェイ"
	profile1 := &Profile{
		ID:       1,
		MemberID: 1,
		Name:     &name,
		NameKana: &name_kana,
		Nickname: &nickname,
	}
	profile2 := &Profile{
		ID:       2,
		MemberID: 2,
		Name:     &name,
		NameKana: &name_kana,
	}
	profileC := &Profile{
		ID:       3,
		MemberID: 3,
		Name:     nil,
		NameKana: nil,
	}
	fmt.Println(profile1, profile2, profileC)

	// slice
	lists := ProfileCompletedColumns

	for _, list := range lists {
		fmt.Println(list)
		// profile1.listはできない
		// fmt.Printf("profile1: %v", profile1.list)
	}
}
