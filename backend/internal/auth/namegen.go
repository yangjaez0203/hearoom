package auth

import (
	"fmt"
	"math/rand"
)

var adjectives = []string{
	"즐거운", "행복한", "신나는", "용감한", "멋진",
	"귀여운", "따뜻한", "빛나는", "활발한", "느긋한",
	"다정한", "씩씩한", "재빠른", "조용한", "명랑한",
	"상냥한", "든든한", "솔직한", "유쾌한", "차분한",
}

var nouns = []string{
	"고양이", "강아지", "토끼", "여우", "펭귄",
	"코알라", "판다", "수달", "다람쥐", "부엉이",
	"돌고래", "기린", "햄스터", "고슴도치", "알파카",
	"북극곰", "레서판다", "카피바라", "치타", "미어캣",
}

func GenerateName() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	num := rand.Intn(100)
	return fmt.Sprintf("%s%s%d", adj, noun, num)
}
