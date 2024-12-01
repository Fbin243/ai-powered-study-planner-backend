package entity

import "ai-powered-study-planner-backend/pkg/db"

type Profile struct {
	*db.BaseModel `bson:",inline"`
	Name          string `json:"name" bson:"name"`
	Email         string `json:"email,omitempty" bson:"email"`
	FirebaseUID   string `json:"firebaseUID,omitempty" bson:"firebase_uid"`
	Picture       string `json:"picture,omitempty" bson:"picture"`
}
