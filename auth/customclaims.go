package auth

import (
	"encoding/json"
	"fmt"

	"firebase.google.com/go/auth"
)

func (f *Firebase) InitCustomClaims(user *auth.UserRecord, voto_id int64) error {
	var customClaims UserCustomClaims
	customClaims.IsAdmin = false
	customClaims.Creator = []int{}
	customClaims.Trustperson = []int{}
	customClaims.Candidate = []int{}
	customClaims.Media = []int{}
	customClaims.VOTO_id = voto_id

	err := f.SetCustomUserClaims(user, customClaims)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) GetCustomUserClaims(user *auth.UserRecord) (UserCustomClaims, error) {

	var customClaims UserCustomClaims
	cC := user.CustomClaims
	jsonStr, err := json.Marshal(cC)
	if err != nil {
		return customClaims, err
	}
	if err := json.Unmarshal(jsonStr, &customClaims); err != nil {
		return customClaims, err
	}
	return customClaims, nil
}

func (f *Firebase) SetCustomUserClaims(user *auth.UserRecord, claims UserCustomClaims) error {

	jsonStr, err := json.Marshal(claims)
	if err != nil {
		fmt.Println(err)
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal(jsonStr, &mapData); err != nil {
		fmt.Println(err)
	}
	err = f.Auth.SetCustomUserClaims(f.Ctx, user.UID, mapData)
	return err
}

func (f *Firebase) AddAdminRole(user *auth.UserRecord) error {

	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	claims.IsAdmin = true
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) RemoveAdminRole(user *auth.UserRecord) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	claims.IsAdmin = false
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) AddCreatorRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	claims.Creator = append(claims.Creator, instanceID)
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) RemoveCreatorRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	for i := 0; i < len(claims.Creator); i++ {
		if claims.Creator[i] == instanceID {
			claims.Creator = append(claims.Creator[:i], claims.Creator[i+1:]...)
		}
	}
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) AddTrustPersonRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	claims.Trustperson = append(claims.Trustperson, instanceID)
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}
func (f *Firebase) RemoveTrustPersonRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	for i := 0; i < len(claims.Trustperson); i++ {
		if claims.Trustperson[i] == instanceID {
			claims.Trustperson = append(claims.Trustperson[:i], claims.Trustperson[i+1:]...)
		}
	}
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}
func (f *Firebase) AddCandidateRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	claims.Candidate = append(claims.Candidate, instanceID)
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}
func (f *Firebase) RemoveCandidateRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	for i := 0; i < len(claims.Candidate); i++ {
		if claims.Candidate[i] == instanceID {
			claims.Candidate = append(claims.Candidate[:i], claims.Candidate[i+1:]...)
		}
	}
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}
func (f *Firebase) AddMediaRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	claims.Media = append(claims.Media, instanceID)
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}
func (f *Firebase) RemoveMediaRole(user *auth.UserRecord, instanceID int) error {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return err
	}
	for i := 0; i < len(claims.Media); i++ {
		if claims.Media[i] == instanceID {
			claims.Media = append(claims.Media[:i], claims.Media[i+1:]...)
		}
	}
	err = f.SetCustomUserClaims(user, claims)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) GetUserIdFromCustomClaims(user *auth.UserRecord) (int64, error) {
	claims, err := f.GetCustomUserClaims(user)
	if err != nil {
		return 0, err
	}
	return claims.VOTO_id, nil
}
